package service

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/model"
	"ai_hs_go/internal/repository"
)

const (
	CompanyStatusPending  int16 = 0
	CompanyStatusDisabled int16 = 2
)

type CompanyService struct {
	db   *gorm.DB
	repo *repository.CompanyRepo
	// rbac 用于在 Create / CreateRoot 后给自动派生的 c{id}_admin 用户分配默认角色：
	//   CreateRoot → tenant_root_admin（一级企业管理员，全部权限）
	//   Create     → tenant_admin     （子公司管理员，业务模块权限子集）
	// 可为 nil — 单元测试 / 命令行工具不需要 RBAC 时使用 NewCompanyService。
	rbac *RBACService
}

func NewCompanyService(db *gorm.DB, repo *repository.CompanyRepo) *CompanyService {
	return &CompanyService{db: db, repo: repo}
}

// WithRBAC 注入 RBAC service 让 Create / CreateRoot 自动给派生 admin 用户分配角色。
// router 启动时调用一次。链式 API 让既有调用方（无 RBAC 需求）的构造保持不变。
func (s *CompanyService) WithRBAC(rbac *RBACService) *CompanyService {
	s.rbac = rbac
	return s
}

type CompanyListParams struct {
	Page              int
	PerPage           int
	ParentID          *int64
	RootIDFilter      *int64 // applied for non-platform-admin to scope view
	Keyword           string
	IncludeDisabled   bool
}

func (s *CompanyService) List(ctx context.Context, p CompanyListParams) ([]model.Company, int64, error) {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PerPage <= 0 || p.PerPage > 100 {
		p.PerPage = 20
	}
	q := s.db.WithContext(ctx).Model(&model.Company{}).
		Preload("EnterpriseType").
		Preload("OperationalType")
	if p.ParentID != nil {
		q = q.Where("parent_id = ?", *p.ParentID)
	}
	if p.RootIDFilter != nil {
		q = q.Where("root_id = ?", *p.RootIDFilter)
	}
	if kw := strings.TrimSpace(p.Keyword); kw != "" {
		like := "%" + kw + "%"
		q = q.Where("title ILIKE ? OR credit_code ILIKE ?", like, like)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.Company
	if err := q.Order("id DESC").Offset((p.Page - 1) * p.PerPage).Limit(p.PerPage).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

// ReviewListParams filters the pending-review queue. ReviewerCompanyID limits
// pending companies to direct children + grandchildren of the reviewer; pass
// 0 (and IsPlatform=true) for platform admins to see every pending row.
type ReviewListParams struct {
	Page              int
	PerPage           int
	Title             string
	ReviewerCompanyID int64
	IsPlatform        bool
}

func (s *CompanyService) ReviewList(ctx context.Context, p ReviewListParams) ([]model.Company, int64, error) {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PerPage <= 0 || p.PerPage > 100 {
		p.PerPage = 15
	}
	q := s.db.WithContext(ctx).Model(&model.Company{}).
		Where("status = ?", CompanyStatusPending).
		Preload("EnterpriseType").
		Preload("OperationalType").
		Preload("Parent")
	if t := strings.TrimSpace(p.Title); t != "" {
		q = q.Where("title ILIKE ?", "%"+t+"%")
	}
	if !p.IsPlatform {
		// Mirror Laravel: a non-platform reviewer can only see pending
		// companies whose parent or grandparent is their company.
		if p.ReviewerCompanyID == 0 {
			return nil, 0, nil
		}
		q = q.Where(
			"parent_id = ? OR parent_id IN (SELECT id FROM companies WHERE parent_id = ? AND deleted_at IS NULL)",
			p.ReviewerCompanyID, p.ReviewerCompanyID,
		)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.Company
	if err := q.Order("id DESC").
		Offset((p.Page - 1) * p.PerPage).Limit(p.PerPage).
		Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *CompanyService) PublicList(ctx context.Context) ([]model.Company, error) {
	var rows []model.Company
	if err := s.db.WithContext(ctx).
		Where("status = ?", CompanyStatusEnable).
		Order("id").
		Find(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}

type CreateCompanyInput struct {
	Title              string
	EnterpriseTypeID   *int64
	OperationalTypeID  int64
	AreaID             *int64
	Address            *string
	CreditCode         *string
	BusinessScope      *string
	ParentID           *int64 // nil ⇒ root tenant; only platform_admin can create
	IsProduce          int16
	IsBloc             int16
	CreatedBy          int64
}

// AdminCredentials are the auto-generated initial credentials for the
// per-company `c{id}_admin` user. Mirrors Laravel's CompanyController::store
// response shape under the `admin_credentials` key.
type AdminCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Create creates a sub-company AND auto-provisions its initial
// `c{id}_admin` user with `must_change_credentials=true`. Mirrors Laravel's
// CompanyController::store, which always pairs company creation with admin
// creation in a single transaction. Use CreateRoot for tenant roots.
func (s *CompanyService) Create(ctx context.Context, in CreateCompanyInput) (*model.Company, *AdminCredentials, error) {
	if strings.TrimSpace(in.Title) == "" {
		return nil, nil, errs.New(errs.CodeUnprocessable, "企业名称不能为空")
	}
	if in.OperationalTypeID == 0 {
		in.OperationalTypeID = 1
	}

	var company *model.Company
	var creds *AdminCredentials
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		c := &model.Company{
			Title:             in.Title,
			EnterpriseTypeID:  in.EnterpriseTypeID,
			OperationalTypeID: in.OperationalTypeID,
			AreaID:            in.AreaID,
			Address:           in.Address,
			CreditCode:        in.CreditCode,
			BusinessScope:     in.BusinessScope,
			ParentID:          in.ParentID,
			IsProduce:         in.IsProduce,
			IsBloc:            in.IsBloc,
			Status:            CompanyStatusPending,
		}
		c.UserID = in.CreatedBy
		c.LastUserID = in.CreatedBy
		if err := tx.Create(c).Error; err != nil {
			return err
		}
		company = c

		username := fmt.Sprintf("c%d_admin", c.ID)
		plainPassword, err := generateInitialPassword(12)
		if err != nil {
			return err
		}
		hashed, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		nickname := truncateRunes(c.Title, 20) + "管理员"
		u := &model.User{
			Username:              username,
			Password:              string(hashed),
			Nickname:              nickname,
			CompanyID:             &c.ID,
			Status:                UserStatusEnable,
			MustChangeCredentials: true,
		}
		if err := tx.Create(u).Error; err != nil {
			return err
		}
		creds = &AdminCredentials{Username: username, Password: plainPassword}
		return nil
	})
	if err != nil {
		return nil, nil, errs.Wrap(errs.CodeInternal, "创建企业失败", err)
	}
	// 给派生的 c{id}_admin 分配 tenant_admin 角色（子公司管理员）。
	// RBAC 写入失败只记录但不回退企业创建 — 角色可后续在「角色管理」页面手动补救。
	if s.rbac != nil {
		var newAdmin model.User
		if err := s.db.WithContext(ctx).Where("username = ?", creds.Username).First(&newAdmin).Error; err == nil {
			_ = s.rbac.AssignSingleRoleToUser(ctx, newAdmin.ID, "tenant_admin")
		}
	}
	return company, creds, nil
}

// CreateRoot creates a root tenant + initial admin user (`c{id}_admin`).
// Caller must be platform admin. Distinct from Create because tenant roots
// have parent_id = NULL and skip the pending-review gate (they're auto-
// enabled by the platform admin's own privilege).
func (s *CompanyService) CreateRoot(ctx context.Context, in CreateCompanyInput, initialPassword string) (*model.Company, *model.User, error) {
	in.ParentID = nil
	var company *model.Company
	var admin *model.User
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		c := &model.Company{
			Title:             in.Title,
			EnterpriseTypeID:  in.EnterpriseTypeID,
			OperationalTypeID: in.OperationalTypeID,
			AreaID:            in.AreaID,
			Address:           in.Address,
			CreditCode:        in.CreditCode,
			BusinessScope:     in.BusinessScope,
			ParentID:          nil,
			IsProduce:         in.IsProduce,
			IsBloc:            in.IsBloc,
			Status:            CompanyStatusPending,
		}
		if c.OperationalTypeID == 0 {
			c.OperationalTypeID = 1
		}
		c.UserID = in.CreatedBy
		c.LastUserID = in.CreatedBy
		if err := tx.Create(c).Error; err != nil {
			return err
		}
		company = c

		username := fmt.Sprintf("c%d_admin", c.ID)
		if initialPassword == "" {
			initialPassword = "Aifix@2026"
		}
		hashed, err := bcrypt.GenerateFromPassword([]byte(initialPassword), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u := &model.User{
			Username:               username,
			Password:               string(hashed),
			Nickname:               c.Title + "管理员",
			CompanyID:              &c.ID,
			Status:                 UserStatusEnable,
			MustChangeCredentials:  true,
		}
		if err := tx.Create(u).Error; err != nil {
			return err
		}
		admin = u

		// auto-enable the tenant on creation by platform admin
		if err := tx.Model(c).Update("status", CompanyStatusEnable).Error; err != nil {
			return err
		}
		// reload to pick up root_id set by the AFTER INSERT trigger
		if err := tx.First(c, c.ID).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, nil, errs.Wrap(errs.CodeInternal, "创建租户失败", err)
	}
	// 给一级企业的 c{id}_admin 分配 tenant_root_admin 角色（全部业务权限）。
	// 失败不回退（同 Create）— 平台管理员可在角色管理页面补救。
	if s.rbac != nil && admin != nil {
		_ = s.rbac.AssignSingleRoleToUser(ctx, admin.ID, "tenant_root_admin")
	}
	return company, admin, nil
}

// UpdateCompanyInput is the union of every editable field on the current-info
// page. Each pointer encodes "did the client send it?" — nil means "leave
// alone", non-nil (including empty string) means "write this value".
type UpdateCompanyInput struct {
	Title               *string
	EnterpriseTypeID    *int64
	OperationalTypeID   *int64
	AreaID              *int64
	Address             *string
	Established         *string // ISO date; service parses
	CreditCode          *string
	BusinessScope       *string
	CompanyIntroduction *string
	Longitude           *float64
	Dimension           *float64
	GeographicRange     *string
	Acreage             *float64
	SuperviseName       *string
	SuperviseTel        *string
	SuperviseFix        *string
	SuperviseEmail      *string
	SupervisePosition   *string
	IsProduce           *int16
	IsBloc              *int16
	UpdatedBy           int64
}

func (s *CompanyService) Update(ctx context.Context, id int64, in UpdateCompanyInput) (*model.Company, error) {
	var c model.Company
	if err := s.db.WithContext(ctx).First(&c, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.New(errs.CodeNotFound, "企业不存在")
		}
		return nil, err
	}
	updates := map[string]any{"last_user_id": in.UpdatedBy}
	if in.Title != nil {
		updates["title"] = *in.Title
	}
	if in.EnterpriseTypeID != nil {
		updates["enterprise_type_id"] = *in.EnterpriseTypeID
	}
	if in.OperationalTypeID != nil {
		updates["operational_type_id"] = *in.OperationalTypeID
	}
	if in.AreaID != nil {
		updates["area_id"] = *in.AreaID
	}
	if in.Address != nil {
		updates["address"] = in.Address
	}
	if in.Established != nil {
		// `established` is a DATE column; GORM accepts time.Time. Parse the
		// ISO date the frontend sends; on parse failure we leave the column
		// alone instead of failing the whole save.
		if t := parseDateOnly(*in.Established); t != nil {
			updates["established"] = t
		} else if *in.Established == "" {
			updates["established"] = nil
		}
	}
	if in.CreditCode != nil {
		updates["credit_code"] = in.CreditCode
	}
	if in.BusinessScope != nil {
		updates["business_scope"] = in.BusinessScope
	}
	if in.CompanyIntroduction != nil {
		updates["company_introduction"] = in.CompanyIntroduction
	}
	if in.Longitude != nil {
		updates["longitude"] = *in.Longitude
	}
	if in.Dimension != nil {
		updates["dimension"] = *in.Dimension
	}
	if in.GeographicRange != nil {
		updates["geographic_range"] = in.GeographicRange
	}
	if in.Acreage != nil {
		updates["acreage"] = *in.Acreage
	}
	if in.SuperviseName != nil {
		updates["supervise_name"] = in.SuperviseName
	}
	if in.SuperviseTel != nil {
		updates["supervise_tel"] = in.SuperviseTel
	}
	if in.SuperviseFix != nil {
		updates["supervise_fix"] = in.SuperviseFix
	}
	if in.SuperviseEmail != nil {
		updates["supervise_email"] = in.SuperviseEmail
	}
	if in.SupervisePosition != nil {
		updates["supervise_position"] = in.SupervisePosition
	}
	if in.IsProduce != nil {
		updates["is_produce"] = *in.IsProduce
	}
	if in.IsBloc != nil {
		updates["is_bloc"] = *in.IsBloc
	}
	if err := s.db.WithContext(ctx).Model(&c).Updates(updates).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "更新失败", err)
	}
	return &c, nil
}

// parseDateOnly returns a *model.Date suitable for assigning to date-only
// columns. nil on empty / unrecognised input.
func parseDateOnly(s string) *model.Date {
	if s == "" {
		return nil
	}
	for _, f := range []string{"2006-01-02", "2006/01/02", "2006/1/2"} {
		if t, err := time.Parse(f, s); err == nil {
			d := model.Date(t)
			return &d
		}
	}
	return nil
}

func (s *CompanyService) Delete(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.Company{}, id).Error
}

// ToggleStatus flips status between Enable and Disabled.
func (s *CompanyService) ToggleStatus(ctx context.Context, id int64, by int64) (*model.Company, error) {
	c, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, errs.New(errs.CodeNotFound, "企业不存在")
	}
	target := CompanyStatusEnable
	if c.Status == CompanyStatusEnable {
		target = CompanyStatusDisabled
	}
	if err := s.db.WithContext(ctx).Model(c).
		Updates(map[string]any{"status": target, "last_user_id": by}).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "切换状态失败", err)
	}
	c.Status = target
	return c, nil
}

// Review writes the reviewer's decision to status. Mirrors Laravel which
// passes the literal target status (1=enable / approve, 2=disabled / reject)
// rather than a boolean — the frontend approval page sends `{status: 1|2}`.
func (s *CompanyService) Review(ctx context.Context, id int64, status int16, by int64) (*model.Company, error) {
	if status != CompanyStatusEnable && status != CompanyStatusDisabled {
		return nil, errs.New(errs.CodeUnprocessable, "无效的审核状态")
	}
	c, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, errs.New(errs.CodeNotFound, "企业不存在")
	}
	if err := s.db.WithContext(ctx).Model(c).
		Updates(map[string]any{"status": status, "last_user_id": by}).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "审核失败", err)
	}
	c.Status = status
	return c, nil
}

// TreeNode is the recursive form returned by GET /companies-tree etc.
//
// Both `children` (Go-side convention, used by /platform/forest) and
// `companies` (Laravel relation name, used by /enterprise/grouping) point at
// the same slice — handlers populate one and we mirror to the other in
// MarshalJSON so neither page has to special-case its read path.
type TreeNode struct {
	ID       int64       `json:"id"`
	Title    string      `json:"title"`
	ParentID *int64      `json:"parent_id"`
	Pid      *int64      `json:"pid"`
	RootID   int64       `json:"root_id"`
	Status   int16       `json:"status"`
	IsBloc   int16       `json:"is_bloc"`
	IsProduce int16      `json:"is_produce"`
	Children []*TreeNode `json:"children,omitempty"`
	// Mirrors Children for Laravel-style consumers; populated alongside.
	Companies []*TreeNode `json:"companies,omitempty"`
}

// Tree returns the subtree rooted at root. Pass 0 for "no specific root" — but
// callers should always supply one (forest endpoint loops).
func (s *CompanyService) Tree(ctx context.Context, rootID int64) (*TreeNode, error) {
	rows, err := s.repo.Subtree(ctx, rootID)
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, errs.New(errs.CodeNotFound, "企业不存在")
	}
	byID := make(map[int64]*TreeNode, len(rows))
	for i := range rows {
		row := rows[i]
		byID[row.ID] = &TreeNode{
			ID:        row.ID,
			Title:     row.Title,
			ParentID:  row.ParentID,
			Pid:       row.ParentID,
			RootID:    row.RootID,
			Status:    row.Status,
			IsBloc:    row.IsBloc,
			IsProduce: row.IsProduce,
		}
	}
	var root *TreeNode
	for i := range rows {
		row := rows[i]
		node := byID[row.ID]
		if row.ID == rootID {
			root = node
			continue
		}
		if row.ParentID != nil {
			if parent, ok := byID[*row.ParentID]; ok {
				parent.Children = append(parent.Children, node)
				parent.Companies = parent.Children // keep both keys in sync
			}
		}
	}
	if root == nil {
		root = byID[rootID]
	}
	return root, nil
}

// Forest returns the full forest of trees. Platform admin only.
func (s *CompanyService) Forest(ctx context.Context) ([]*TreeNode, error) {
	roots, err := s.repo.Forest(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]*TreeNode, 0, len(roots))
	for _, r := range roots {
		t, err := s.Tree(ctx, r.ID)
		if err == nil {
			out = append(out, t)
		}
	}
	return out, nil
}

func (s *CompanyService) Repo() *repository.CompanyRepo { return s.repo }

// generateInitialPassword returns a printable random password of the given
// length. The character set excludes ambiguous chars (0/O, 1/l) to make
// hand-typing the credentials less error prone for new admins on first
// login. No symbols — Laravel uses `Str::password(12, symbols: false)` and
// we mirror that.
func generateInitialPassword(length int) (string, error) {
	const charset = "ABCDEFGHJKMNPQRSTUVWXYZabcdefghjkmnpqrstuvwxyz23456789"
	b := make([]byte, length)
	max := big.NewInt(int64(len(charset)))
	for i := range b {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		b[i] = charset[n.Int64()]
	}
	return string(b), nil
}

// truncateRunes truncates a Chinese string to at most n runes (not bytes).
func truncateRunes(s string, n int) string {
	r := []rune(s)
	if len(r) <= n {
		return s
	}
	return string(r[:n])
}
