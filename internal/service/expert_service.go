// Package service — 专家检查 + 专家库 业务层。
//
// expert_service.go 实现两组 CRUD：
//
//	ExpertInspection / ExpertCompany / ExpertInspectionUser  - 检查计划三层级
//	BlocExpertUser                                           - 专家库
//
// 跟 Laravel ExpertInspectionsController + BlocExpertsController 对应。子表
// 不走「整体替换 kids」语义（前端是按 row 单独增删的），所以这里每个 entity
// 都给完整 CRUD。
package service

import (
	"archive/zip"
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gorm.io/gorm"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/model"
)

type ExpertService struct {
	db *gorm.DB
	// inspectionNoticeTemplate 是 Laravel public/enterprise_export.docx 在 Go 项目下的镜像。
	// configs/templates/enterprise_export.docx；Newhandler 时由 router 注入路径。
	inspectionNoticeTemplate string
	// noticeStorageDir 是生成的 docx 落盘根目录，路径段 inspection-notices/YYYY/MM/<filename>。
	// router 启动时按 `<项目根>/storage/app` 传入，跟 Laravel storage_path('app/...') 对齐。
	noticeStorageDir string
}

func NewExpertService(db *gorm.DB, templatePath, storageDir string) *ExpertService {
	return &ExpertService{
		db:                       db,
		inspectionNoticeTemplate: templatePath,
		noticeStorageDir:         storageDir,
	}
}

// ── ExpertInspection（检查计划） ─────────────────────────────────────

// ExpertInspectionListParams — Laravel 的 (company_id, expert_user_id, title) 三态过滤。
//
//   - 仅 company_id：本企业作为 (a) 计划方  或  (b) 受检企业。
//   - 仅 expert_user_id：本人作为已指派专家。
//   - 同时传：上述任一命中。
//   - title：模糊搜索。
type ExpertInspectionListParams struct {
	Page         int
	PerPage      int
	CompanyID    int64
	ExpertUserID int64
	Title        string
}

func (s *ExpertService) ListExpertInspection(ctx context.Context, p ExpertInspectionListParams) ([]model.ExpertInspection, int64, error) {
	page, per := normalizePage(p.Page, p.PerPage)
	q := s.db.WithContext(ctx).Model(&model.ExpertInspection{})

	hasCompany := p.CompanyID > 0
	hasExpert := p.ExpertUserID > 0
	switch {
	case hasCompany && hasExpert:
		// company_id matches plan creator
		// OR plan contains an expert_company with check_company_id = company_id
		// OR plan contains an expert_inspection_user with expert_user_id = userId
		q = q.Where(
			s.db.Where("company_id = ?", p.CompanyID).
				Or("id IN (?)",
					s.db.Model(&model.ExpertCompany{}).
						Select("expert_inspection_id").
						Where("check_company_id = ?", p.CompanyID),
				).
				Or("id IN (?)",
					s.db.Model(&model.ExpertCompany{}).
						Select("expert_inspection_id").
						Where("id IN (?)",
							s.db.Model(&model.ExpertInspectionUser{}).
								Select("expert_company_id").
								Where("expert_user_id = ?", p.ExpertUserID),
						),
				),
		)
	case hasCompany:
		q = q.Where(
			s.db.Where("company_id = ?", p.CompanyID).
				Or("id IN (?)",
					s.db.Model(&model.ExpertCompany{}).
						Select("expert_inspection_id").
						Where("check_company_id = ?", p.CompanyID),
				),
		)
	case hasExpert:
		q = q.Where("id IN (?)",
			s.db.Model(&model.ExpertCompany{}).
				Select("expert_inspection_id").
				Where("id IN (?)",
					s.db.Model(&model.ExpertInspectionUser{}).
						Select("expert_company_id").
						Where("expert_user_id = ?", p.ExpertUserID),
				),
		)
	}

	if v := strings.TrimSpace(p.Title); v != "" {
		q = q.Where("title ILIKE ?", "%"+v+"%")
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.ExpertInspection
	if err := q.Preload("ExpertCompanies").
		Order("id DESC").Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *ExpertService) GetExpertInspection(ctx context.Context, id int64) (*model.ExpertInspection, error) {
	var m model.ExpertInspection
	err := s.db.WithContext(ctx).
		Preload("ExpertCompanies.CheckCompany", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id, title")
		}).
		Preload("ExpertCompanies.ExpertInspectionUsers.ExpertUser", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id, nickname, phone")
		}).
		First(&m, id).Error
	if err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

func (s *ExpertService) CreateExpertInspection(ctx context.Context, m *model.ExpertInspection) error {
	if m.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return wrapWrite(s.db.WithContext(ctx).Create(m).Error)
}

func (s *ExpertService) UpdateExpertInspection(ctx context.Context, id int64, updates map[string]any) error {
	return updateRow(s.db.WithContext(ctx), &model.ExpertInspection{}, id, updates)
}

// DeleteExpertInspection — 软删除主表；ExpertCompany / ExpertInspectionUser 由 ON DELETE CASCADE
// 在硬删除时清理；soft delete 场景下子表跟着 deleted_at IS NULL 过滤即可。
func (s *ExpertService) DeleteExpertInspection(ctx context.Context, id int64) error {
	res := s.db.WithContext(ctx).Delete(&model.ExpertInspection{}, id)
	if res.Error != nil {
		return wrapWrite(res.Error)
	}
	if res.RowsAffected == 0 {
		return errs.New(errs.CodeNotFound, "专家检查不存在")
	}
	return nil
}

// ── ExpertCompany（受检企业） ───────────────────────────────────────

type ExpertCompanyListParams struct {
	Page               int
	PerPage            int
	ExpertInspectionID int64
}

func (s *ExpertService) ListExpertCompany(ctx context.Context, p ExpertCompanyListParams) ([]model.ExpertCompany, int64, error) {
	page, per := normalizePage(p.Page, p.PerPage)
	q := s.db.WithContext(ctx).Model(&model.ExpertCompany{})
	if p.ExpertInspectionID > 0 {
		q = q.Where("expert_inspection_id = ?", p.ExpertInspectionID)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.ExpertCompany
	if err := q.
		Preload("CheckCompany", func(tx *gorm.DB) *gorm.DB { return tx.Select("id, title") }).
		Preload("ExpertInspectionUsers.ExpertUser", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id, nickname, phone")
		}).
		Order("id DESC").Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *ExpertService) GetExpertCompany(ctx context.Context, id int64) (*model.ExpertCompany, error) {
	var m model.ExpertCompany
	err := s.db.WithContext(ctx).
		Preload("ExpertInspection").
		Preload("CheckCompany").
		Preload("ExpertInspectionUsers.ExpertUser", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id, nickname, phone")
		}).
		First(&m, id).Error
	if err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

func (s *ExpertService) CreateExpertCompany(ctx context.Context, m *model.ExpertCompany) error {
	if m.CompanyID == 0 || m.ExpertInspectionID == 0 || m.CheckCompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少必填字段")
	}
	if err := s.db.WithContext(ctx).Create(m).Error; err != nil {
		return wrapWrite(err)
	}
	// 重新加载 checkCompany 关联，handler 直接 response.Created(record)。
	return s.db.WithContext(ctx).
		Preload("CheckCompany", func(tx *gorm.DB) *gorm.DB { return tx.Select("id, title") }).
		First(m, m.ID).Error
}

func (s *ExpertService) UpdateExpertCompany(ctx context.Context, id int64, updates map[string]any) error {
	return updateRow(s.db.WithContext(ctx), &model.ExpertCompany{}, id, updates)
}

func (s *ExpertService) DeleteExpertCompany(ctx context.Context, id int64) error {
	res := s.db.WithContext(ctx).Delete(&model.ExpertCompany{}, id)
	if res.Error != nil {
		return wrapWrite(res.Error)
	}
	if res.RowsAffected == 0 {
		return errs.New(errs.CodeNotFound, "专家检查企业不存在")
	}
	return nil
}

// ConfirmExpertCompany — 受检企业确认收到检查通知；写 confirm_time = now 并把 status 推到 2。
func (s *ExpertService) ConfirmExpertCompany(ctx context.Context, id int64, uid int64) (*model.ExpertCompany, error) {
	var ec model.ExpertCompany
	if err := s.db.WithContext(ctx).First(&ec, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	now := model.DateTime(time.Now())
	updates := map[string]any{
		"confirm_time": now,
		"status":       2,
		"last_user_id": uid,
	}
	if err := s.db.WithContext(ctx).Model(&ec).Updates(updates).Error; err != nil {
		return nil, wrapWrite(err)
	}
	if err := s.db.WithContext(ctx).First(&ec, id).Error; err != nil {
		return nil, err
	}
	return &ec, nil
}

// ── ExpertInspectionUser（专家指派） ────────────────────────────────

type ExpertInspectionUserListParams struct {
	Page            int
	PerPage         int
	ExpertCompanyID int64
}

func (s *ExpertService) ListExpertInspectionUser(ctx context.Context, p ExpertInspectionUserListParams) ([]model.ExpertInspectionUser, int64, error) {
	page, per := normalizePage(p.Page, p.PerPage)
	q := s.db.WithContext(ctx).Model(&model.ExpertInspectionUser{}).
		Where("expert_company_id = ?", p.ExpertCompanyID)
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.ExpertInspectionUser
	if err := q.Preload("ExpertUser", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("id, nickname, phone")
	}).
		Order("id DESC").Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *ExpertService) CreateExpertInspectionUser(ctx context.Context, m *model.ExpertInspectionUser) error {
	if m.CompanyID == 0 || m.ExpertCompanyID == 0 || m.ExpertUserID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少必填字段")
	}
	if err := s.db.WithContext(ctx).Create(m).Error; err != nil {
		return wrapWrite(err)
	}
	return s.db.WithContext(ctx).
		Preload("ExpertUser", func(tx *gorm.DB) *gorm.DB { return tx.Select("id, nickname, phone") }).
		First(m, m.ID).Error
}

func (s *ExpertService) UpdateExpertInspectionUser(ctx context.Context, id int64, updates map[string]any) error {
	return updateRow(s.db.WithContext(ctx), &model.ExpertInspectionUser{}, id, updates)
}

func (s *ExpertService) DeleteExpertInspectionUser(ctx context.Context, id int64) error {
	res := s.db.WithContext(ctx).Delete(&model.ExpertInspectionUser{}, id)
	if res.Error != nil {
		return wrapWrite(res.Error)
	}
	if res.RowsAffected == 0 {
		return errs.New(errs.CodeNotFound, "专家用户不存在")
	}
	return nil
}

// ConfirmExpertInspectionUser — 专家点击「我确认接受指派」，写 confirm_time = now，status = 2。
func (s *ExpertService) ConfirmExpertInspectionUser(ctx context.Context, id int64, uid int64) (*model.ExpertInspectionUser, error) {
	var eu model.ExpertInspectionUser
	if err := s.db.WithContext(ctx).First(&eu, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	now := model.DateTime(time.Now())
	updates := map[string]any{
		"confirm_time": now,
		"status":       2,
		"last_user_id": uid,
	}
	if err := s.db.WithContext(ctx).Model(&eu).Updates(updates).Error; err != nil {
		return nil, wrapWrite(err)
	}
	if err := s.db.WithContext(ctx).First(&eu, id).Error; err != nil {
		return nil, err
	}
	return &eu, nil
}

// ── 仪表板通知与「我的检查/指派」 ───────────────────────────────────

// CompanyNotifications — 受检企业仪表板通知：notify_company_time <= now 且尚未确认的记录。
//
// 跟 Laravel 一致：notify_company_time 一旦到达即出现在仪表板，无需额外触发。
func (s *ExpertService) CompanyNotifications(ctx context.Context, companyID int64) ([]model.ExpertCompany, error) {
	var rows []model.ExpertCompany
	err := s.db.WithContext(ctx).
		Where("check_company_id = ?", companyID).
		Where("confirm_time IS NULL").
		Where("notify_company_time IS NOT NULL").
		Where("notify_company_time <= ?", time.Now()).
		Preload("ExpertInspection", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id, title, description, start_date, end_date")
		}).
		Preload("ExpertInspectionUsers.ExpertUser", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id, nickname, phone")
		}).
		Order("notify_company_time DESC").
		Find(&rows).Error
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// MyCheckInspections —「我的检查」分页：合并
//
//	(a) 本企业被检查（check_company_id = companyID 且 notify_company_time <= now）
//	(b) 本人作为已指派专家（expert_inspection_users.expert_user_id = userID，立即可见）
type MyCheckInspectionsParams struct {
	Page         int
	PerPage      int
	CompanyID    int64
	ExpertUserID int64
}

func (s *ExpertService) MyCheckInspections(ctx context.Context, p MyCheckInspectionsParams) ([]model.ExpertCompany, int64, error) {
	page, per := normalizePage(p.Page, p.PerPage)
	q := s.db.WithContext(ctx).Model(&model.ExpertCompany{})

	companyClause := s.db.
		Where("check_company_id = ?", p.CompanyID).
		Where("notify_company_time IS NOT NULL").
		Where("notify_company_time <= ?", time.Now())
	if p.ExpertUserID > 0 {
		expertClause := s.db.Where("id IN (?)",
			s.db.Model(&model.ExpertInspectionUser{}).
				Select("expert_company_id").
				Where("expert_user_id = ?", p.ExpertUserID),
		)
		q = q.Where(s.db.Where(companyClause).Or(expertClause))
	} else {
		q = q.Where(companyClause)
	}

	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.ExpertCompany
	if err := q.
		Preload("CheckCompany", func(tx *gorm.DB) *gorm.DB { return tx.Select("id, title") }).
		Preload("ExpertInspection", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id, title, description, start_date, end_date, company_id")
		}).
		Preload("ExpertInspectionUsers.ExpertUser", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id, nickname, phone")
		}).
		Order("id DESC").Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

// ExpertNotifications — 专家端仪表板：本专家所有未确认（confirm_time IS NULL）的指派。
func (s *ExpertService) ExpertNotifications(ctx context.Context, userID int64) ([]model.ExpertInspectionUser, error) {
	if userID <= 0 {
		return []model.ExpertInspectionUser{}, nil
	}
	var rows []model.ExpertInspectionUser
	err := s.db.WithContext(ctx).
		Where("expert_user_id = ?", userID).
		Where("confirm_time IS NULL").
		Preload("ExpertCompany.ExpertInspection", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id, title, description, start_date, end_date")
		}).
		Preload("ExpertCompany.CheckCompany", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id, title")
		}).
		Order("id DESC").
		Find(&rows).Error
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// MyExpertAssignments — 专家端：本专家被指派的全部检查记录（指派即可见，无需 notify）。
type MyExpertAssignmentsParams struct {
	Page    int
	PerPage int
	UserID  int64
}

func (s *ExpertService) MyExpertAssignments(ctx context.Context, p MyExpertAssignmentsParams) ([]model.ExpertInspectionUser, int64, error) {
	if p.UserID <= 0 {
		return []model.ExpertInspectionUser{}, 0, nil
	}
	page, per := normalizePage(p.Page, p.PerPage)
	q := s.db.WithContext(ctx).Model(&model.ExpertInspectionUser{}).Where("expert_user_id = ?", p.UserID)
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.ExpertInspectionUser
	if err := q.
		Preload("ExpertCompany.ExpertInspection", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id, title, description, start_date, end_date, company_id")
		}).
		Preload("ExpertCompany.CheckCompany", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id, title")
		}).
		Order("id DESC").Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

// ── 检查通知单 ──────────────────────────────────────────────────────

// GenerateInspectionNotice — 生成「现场检查通知单」docx，落到 storage/app/inspection-notices/YYYY/MM。
//
// 权限：仅检查计划创建企业可生成（handler 层校验 ExpertCompany.ExpertInspection.CompanyID == userCompanyID）。
//
// docx 模板使用 PHP PhpWord 的 ${...} 语法，本实现不引入额外依赖：直接把 .docx
// 当成 zip，逐个 word/document.xml + word/header*.xml + word/footer*.xml 做字符串替换，
// 重新打包写出。模板里只有简单字段，没有跨 run 切片的复杂场景，足够用。
type InspectionNoticeContext struct {
	Title        string // 「安全检查通知单」固定值
	CompanyTitle string // 受检企业名
	UserName     string // 指派专家姓名拼接
	Content      string // 检查计划主题
	Time         string // check_company_time 字符串
	CreatedTime  string // expert_companies.created_at
}

func (s *ExpertService) GenerateInspectionNotice(ctx context.Context, expertCompanyID int64, userCompanyID int64) (*model.ExpertCompany, error) {
	var ec model.ExpertCompany
	err := s.db.WithContext(ctx).
		Preload("ExpertInspection").
		Preload("CheckCompany").
		Preload("ExpertInspectionUsers.ExpertUser").
		First(&ec, expertCompanyID).Error
	if err != nil {
		return nil, mapNotFound(err)
	}
	if ec.ExpertInspection == nil || ec.ExpertInspection.CompanyID != userCompanyID {
		return nil, errs.New(errs.CodeForbidden, "仅检查计划创建企业可生成通知单")
	}

	if s.inspectionNoticeTemplate == "" {
		return nil, errs.New(errs.CodeInternal, "通知单模板路径未配置")
	}
	if _, err := os.Stat(s.inspectionNoticeTemplate); err != nil {
		return nil, errs.New(errs.CodeInternal, "通知单模板缺失：enterprise_export.docx")
	}

	checkCompanyName := ""
	if ec.CheckCompany != nil {
		checkCompanyName = ec.CheckCompany.Title
	}
	expertNames := joinExpertNames(ec.ExpertInspectionUsers)
	planTitle := ""
	if ec.ExpertInspection != nil {
		planTitle = ec.ExpertInspection.Title
	}
	checkTime := ""
	if ec.CheckCompanyTime != nil {
		checkTime = time.Time(*ec.CheckCompanyTime).Format("2006-01-02 15:04:05")
	}
	createdTime := ec.CreatedAt.Format("2006-01-02 15:04:05")

	replacements := map[string]string{
		"${title}":         "安全检查通知单",
		"${company_title}": checkCompanyName,
		"${user_name}":     expertNames,
		"${content}":       planTitle,
		"${time}":          checkTime,
		"${created_time}":  createdTime,
	}

	now := time.Now()
	relDir := filepath.Join("inspection-notices", now.Format("2006"), now.Format("01"))
	absDir := filepath.Join(s.noticeStorageDir, relDir)
	if err := os.MkdirAll(absDir, 0o755); err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "创建通知单目录失败", err)
	}

	suffix, err := randomHex(3)
	if err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "生成随机后缀失败", err)
	}
	filename := fmt.Sprintf("检查通知单-%s-%s-%s.docx", checkCompanyName, now.Format("20060102"), suffix)
	relPath := filepath.Join(relDir, filename)
	absPath := filepath.Join(s.noticeStorageDir, relPath)

	if err := substituteDocxTemplate(s.inspectionNoticeTemplate, absPath, replacements); err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "生成通知单失败", err)
	}

	rel := relPath
	if err := s.db.WithContext(ctx).Model(&ec).
		Update("inspection_notice_url", rel).Error; err != nil {
		return nil, wrapWrite(err)
	}
	if err := s.db.WithContext(ctx).First(&ec, expertCompanyID).Error; err != nil {
		return nil, err
	}
	return &ec, nil
}

// InspectionNoticeFile — 用户下载通知单时调用，返回绝对路径供 handler c.File 使用。
//
// 权限：计划创建方 / 已指派专家可下载；受检企业本身没有访问权（保持 Laravel 现状）。
func (s *ExpertService) InspectionNoticeFile(ctx context.Context, expertCompanyID, userID, userCompanyID int64) (string, string, error) {
	var ec model.ExpertCompany
	err := s.db.WithContext(ctx).
		Preload("ExpertInspection").
		Preload("ExpertInspectionUsers").
		First(&ec, expertCompanyID).Error
	if err != nil {
		return "", "", mapNotFound(err)
	}

	planCompanyID := int64(0)
	if ec.ExpertInspection != nil {
		planCompanyID = ec.ExpertInspection.CompanyID
	}
	isOwner := planCompanyID != 0 && planCompanyID == userCompanyID
	isAssignedExpert := false
	for _, u := range ec.ExpertInspectionUsers {
		if u.ExpertUserID == userID {
			isAssignedExpert = true
			break
		}
	}
	if !isOwner && !isAssignedExpert {
		return "", "", errs.New(errs.CodeForbidden, "无权限下载通知单")
	}
	if ec.InspectionNoticeURL == nil || *ec.InspectionNoticeURL == "" {
		return "", "", errs.New(errs.CodeUnprocessable, "通知单尚未生成")
	}
	abs := filepath.Join(s.noticeStorageDir, *ec.InspectionNoticeURL)
	if _, err := os.Stat(abs); err != nil {
		return "", "", errs.New(errs.CodeNotFound, "通知单文件不存在")
	}
	return abs, filepath.Base(abs), nil
}

func joinExpertNames(users []model.ExpertInspectionUser) string {
	out := make([]string, 0, len(users))
	for _, u := range users {
		if u.ExpertUser != nil && u.ExpertUser.Nickname != "" {
			out = append(out, u.ExpertUser.Nickname)
		}
	}
	return strings.Join(out, "、")
}

func randomHex(n int) (string, error) {
	buf := make([]byte, n)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	return hex.EncodeToString(buf), nil
}

// substituteDocxTemplate — 把模板 docx 中所有 word/*.xml 文件里的 ${var} 占位符替换后另存。
//
// .docx 是 zip 格式，里面 word/document.xml / word/header*.xml / word/footer*.xml 是
// 主要文本载体。简单字符串替换在 PhpWord 风格模板上够用（占位符不会被切到不同的 <w:t>
// run 里——前提是模板里就是直接打的 `${title}`，没有中间格式变化）。
func substituteDocxTemplate(srcPath, dstPath string, replacements map[string]string) error {
	r, err := zip.OpenReader(srcPath)
	if err != nil {
		return fmt.Errorf("open template: %w", err)
	}
	defer r.Close()

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return fmt.Errorf("create output: %w", err)
	}
	defer dstFile.Close()
	w := zip.NewWriter(dstFile)
	defer w.Close()

	for _, f := range r.File {
		fw, err := w.CreateHeader(&zip.FileHeader{
			Name:     f.Name,
			Method:   f.Method,
			Modified: f.Modified,
		})
		if err != nil {
			return fmt.Errorf("zip header %s: %w", f.Name, err)
		}
		fr, err := f.Open()
		if err != nil {
			return fmt.Errorf("open entry %s: %w", f.Name, err)
		}
		// Only run substitution on the textual XML parts — body / header / footer / endnotes / footnotes.
		if isDocxTextPart(f.Name) {
			buf := new(bytes.Buffer)
			if _, err := io.Copy(buf, fr); err != nil {
				fr.Close()
				return fmt.Errorf("read entry %s: %w", f.Name, err)
			}
			fr.Close()
			content := buf.String()
			for k, v := range replacements {
				content = strings.ReplaceAll(content, k, escapeXML(v))
			}
			if _, err := fw.Write([]byte(content)); err != nil {
				return fmt.Errorf("write entry %s: %w", f.Name, err)
			}
			continue
		}
		// Binary parts copied verbatim.
		if _, err := io.Copy(fw, fr); err != nil {
			fr.Close()
			return fmt.Errorf("copy entry %s: %w", f.Name, err)
		}
		fr.Close()
	}
	return nil
}

func isDocxTextPart(name string) bool {
	if !strings.HasPrefix(name, "word/") || !strings.HasSuffix(name, ".xml") {
		return false
	}
	base := strings.TrimPrefix(name, "word/")
	switch {
	case base == "document.xml":
		return true
	case strings.HasPrefix(base, "header") && strings.HasSuffix(base, ".xml"):
		return true
	case strings.HasPrefix(base, "footer") && strings.HasSuffix(base, ".xml"):
		return true
	case base == "footnotes.xml" || base == "endnotes.xml":
		return true
	}
	return false
}

// escapeXML — 替换值需要转义 XML 特殊字符；模板正文是 XML，user_name / company_title 里出现
// 「&、<、>」都会破坏文档结构。
func escapeXML(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for _, r := range s {
		switch r {
		case '&':
			b.WriteString("&amp;")
		case '<':
			b.WriteString("&lt;")
		case '>':
			b.WriteString("&gt;")
		case '"':
			b.WriteString("&quot;")
		case '\'':
			b.WriteString("&apos;")
		default:
			b.WriteRune(r)
		}
	}
	return b.String()
}

// ── BlocExpertUser（专家库） ────────────────────────────────────────

type BlocExpertListParams struct {
	Page      int
	PerPage   int
	CompanyID int64
	Nickname  string
}

// ListBlocExpert — 集团专家库列表。Nickname 时通过 join users 模糊搜索。
func (s *ExpertService) ListBlocExpert(ctx context.Context, p BlocExpertListParams) ([]model.BlocExpertUser, int64, error) {
	page, per := normalizePage(p.Page, p.PerPage)
	q := s.db.WithContext(ctx).Model(&model.BlocExpertUser{}).Where("company_id = ?", p.CompanyID)
	if v := strings.TrimSpace(p.Nickname); v != "" {
		q = q.Where("expert_user_id IN (?)",
			s.db.Model(&model.User{}).Select("id").Where("nickname ILIKE ?", "%"+v+"%"),
		)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.BlocExpertUser
	if err := q.
		Preload("Company", func(tx *gorm.DB) *gorm.DB { return tx.Select("id, title") }).
		Preload("User", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id, nickname, phone, department_name, email")
		}).
		Order("id DESC").Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

// GetBlocExpertUserByID — Laravel show 接口实际是 User::query()->with('company')->find($id)，
// 返回的是 User 而非 BlocExpertUser；frontend `专家详情` 里取 user.company.title。
func (s *ExpertService) GetBlocExpertUserByID(ctx context.Context, id int64) (*model.User, error) {
	var u model.User
	if err := s.db.WithContext(ctx).First(&u, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &u, nil
}

func (s *ExpertService) CreateBlocExpert(ctx context.Context, m *model.BlocExpertUser) error {
	if m.CompanyID == 0 || m.ExpertUserID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id 或 expert_user_id")
	}
	// 唯一索引拦截，但提前给一个清晰的错误信息。
	var dup int64
	if err := s.db.WithContext(ctx).Model(&model.BlocExpertUser{}).
		Where("company_id = ? AND expert_user_id = ?", m.CompanyID, m.ExpertUserID).
		Count(&dup).Error; err != nil {
		return err
	}
	if dup > 0 {
		return errs.New(errs.CodeConflict, "该用户已是本企业专家")
	}
	return wrapWrite(s.db.WithContext(ctx).Create(m).Error)
}

func (s *ExpertService) DeleteBlocExpert(ctx context.Context, id int64) error {
	res := s.db.WithContext(ctx).Delete(&model.BlocExpertUser{}, id)
	if res.Error != nil {
		return wrapWrite(res.Error)
	}
	if res.RowsAffected == 0 {
		return errs.New(errs.CodeNotFound, "该专家不存在")
	}
	return nil
}

// ListAvailableExperts —「添加专家」面板：列出 status=1 且尚未在该公司专家库中的用户。
func (s *ExpertService) ListAvailableExperts(ctx context.Context, companyID int64, nickname string) ([]model.User, error) {
	subQuery := s.db.Model(&model.BlocExpertUser{}).
		Select("expert_user_id").Where("company_id = ?", companyID)
	q := s.db.WithContext(ctx).Model(&model.User{}).
		Where("status = ?", 1).
		Where("id NOT IN (?)", subQuery)
	if v := strings.TrimSpace(nickname); v != "" {
		q = q.Where("nickname ILIKE ?", "%"+v+"%")
	}
	var rows []model.User
	if err := q.Order("id ASC").Find(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}
