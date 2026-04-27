package service

import (
	"context"
	"strings"

	"gorm.io/gorm"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/model"
	"ai_hs_go/internal/repository"
)

// AccidentEmergencyService 聚合事故应急 4 个生产视图模块 + 1 个集团视图模块：
//
//   - accident_bulletins / kids        生产事故快报（含子表整体替换）
//   - bloc-accident-bulletins          集团事故快报（按 company subtree 跨企业聚合）
//   - emergency_drills                 应急演练
//   - emergency_supplies               应急物资（无子表）
//   - emergency_teams / kids           应急队伍（含子表整体替换）
//
// 子表 nil 切片 = 不动；非 nil（含空切片）= 整体替换。Laravel 写法 `$kids = $request->input('kids')`
// 同语义：传 `kids: []` 清空，不传该 key 留旧数据。
type AccidentEmergencyService struct {
	db          *gorm.DB
	companyRepo *repository.CompanyRepo
}

func NewAccidentEmergencyService(db *gorm.DB, companyRepo *repository.CompanyRepo) *AccidentEmergencyService {
	return &AccidentEmergencyService{db: db, companyRepo: companyRepo}
}

// ── 事故快报（生产视图） ──────────────────────────────────────────────

type AccidentBulletinListParams struct {
	Page          int
	PerPage       int
	CompanyID     int64
	Type          string
	AccidentLevel string
}

func (s *AccidentEmergencyService) ListAccidentBulletin(ctx context.Context, p AccidentBulletinListParams) ([]model.AccidentBulletin, int64, error) {
	page, per := normalizePage(p.Page, p.PerPage)
	q := s.db.WithContext(ctx).Model(&model.AccidentBulletin{}).Where("company_id = ?", p.CompanyID)
	if v := strings.TrimSpace(p.Type); v != "" {
		q = q.Where("type ILIKE ?", "%"+v+"%")
	}
	if v := strings.TrimSpace(p.AccidentLevel); v != "" {
		q = q.Where("accident_level ILIKE ?", "%"+v+"%")
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.AccidentBulletin
	if err := q.Preload("Kids").Order("id DESC").
		Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *AccidentEmergencyService) GetAccidentBulletin(ctx context.Context, id int64) (*model.AccidentBulletin, error) {
	var m model.AccidentBulletin
	if err := s.db.WithContext(ctx).Preload("Kids").First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

type AccidentBulletinInput struct {
	Bulletin *model.AccidentBulletin
	Kids     []model.AccidentBulletinKid
}

func (s *AccidentEmergencyService) CreateAccidentBulletin(ctx context.Context, in AccidentBulletinInput) error {
	if in.Bulletin.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(in.Bulletin).Error; err != nil {
			return wrapWrite(err)
		}
		if len(in.Kids) > 0 {
			uid := in.Bulletin.UserID
			cid := in.Bulletin.CompanyID
			for i := range in.Kids {
				in.Kids[i].AccidentBulletinID = in.Bulletin.ID
				in.Kids[i].CompanyID = cid
				in.Kids[i].UserID = uid
				in.Kids[i].LastUserID = uid
			}
			if err := tx.Create(&in.Kids).Error; err != nil {
				return wrapWrite(err)
			}
		}
		return nil
	})
}

// UpdateAccidentBulletin — kids 为 nil 表示不动；非 nil（含空切片）表示整体替换。
func (s *AccidentEmergencyService) UpdateAccidentBulletin(
	ctx context.Context,
	id int64,
	updates map[string]any,
	kids []model.AccidentBulletinKid,
	replaceKids bool,
	uid int64,
) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		res := tx.Model(&model.AccidentBulletin{}).Where("id = ?", id).Updates(updates)
		if res.Error != nil {
			return wrapWrite(res.Error)
		}
		if res.RowsAffected == 0 {
			return errs.New(errs.CodeNotFound, "记录不存在")
		}
		if !replaceKids {
			return nil
		}
		var b model.AccidentBulletin
		if err := tx.First(&b, id).Error; err != nil {
			return err
		}
		if err := tx.Where("accident_bulletin_id = ?", id).Delete(&model.AccidentBulletinKid{}).Error; err != nil {
			return err
		}
		if len(kids) == 0 {
			return nil
		}
		for i := range kids {
			kids[i].AccidentBulletinID = id
			kids[i].CompanyID = b.CompanyID
			kids[i].UserID = uid
			kids[i].LastUserID = uid
		}
		if err := tx.Create(&kids).Error; err != nil {
			return wrapWrite(err)
		}
		return nil
	})
}

func (s *AccidentEmergencyService) DeleteAccidentBulletin(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("accident_bulletin_id = ?", id).Delete(&model.AccidentBulletinKid{}).Error; err != nil {
			return err
		}
		return tx.Delete(&model.AccidentBulletin{}, id).Error
	})
}

// ── 集团事故快报（跨企业 list + 年度聚合） ────────────────────────────

type BlocAccidentBulletinListParams struct {
	Page          int
	PerPage       int
	CompanyID     int64
	Type          string
	AccidentLevel string
	HappenYear    string
}

func (s *AccidentEmergencyService) companySubtreeIDs(ctx context.Context, companyID int64) ([]int64, error) {
	ids, err := s.companyRepo.DescendantIDs(ctx, companyID)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return []int64{companyID}, nil
	}
	return ids, nil
}

func (s *AccidentEmergencyService) ListBlocAccidentBulletin(ctx context.Context, p BlocAccidentBulletinListParams) ([]model.AccidentBulletin, int64, error) {
	page, per := normalizePage(p.Page, p.PerPage)
	ids, err := s.companySubtreeIDs(ctx, p.CompanyID)
	if err != nil {
		return nil, 0, err
	}
	q := s.db.WithContext(ctx).Model(&model.AccidentBulletin{}).Where("company_id IN ?", ids)
	if v := strings.TrimSpace(p.Type); v != "" {
		q = q.Where("type = ?", v)
	}
	if v := strings.TrimSpace(p.AccidentLevel); v != "" {
		q = q.Where("accident_level = ?", v)
	}
	if v := strings.TrimSpace(p.HappenYear); v != "" {
		q = q.Where("EXTRACT(YEAR FROM happen_time) = ?", v)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.AccidentBulletin
	if err := q.Preload("Company").
		Order("happen_time DESC NULLS LAST").Order("id DESC").
		Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

// AnnualStat — 跟前端 AnnualStat 接口字段对齐。
type AnnualStat struct {
	Year          string `json:"year"`
	AccidentCount int64  `json:"accident_count"`
	HurtCount     int64  `json:"hurt_count"`
	DeathCount    int64  `json:"death_count"`
}

func (s *AccidentEmergencyService) BlocAccidentAnnualStats(ctx context.Context, companyID int64) ([]AnnualStat, error) {
	ids, err := s.companySubtreeIDs(ctx, companyID)
	if err != nil {
		return nil, err
	}
	var rows []AnnualStat
	err = s.db.WithContext(ctx).
		Table("accident_bulletins").
		Where("company_id IN ?", ids).
		Where("deleted_at IS NULL").
		Where("happen_time IS NOT NULL").
		Select(
			"to_char(happen_time, 'YYYY') AS year, "+
				"COUNT(*) AS accident_count, "+
				"COALESCE(SUM(staff_hurt), 0) + COALESCE(SUM(not_staff_hurt), 0) AS hurt_count, "+
				"COALESCE(SUM(staff_die), 0) + COALESCE(SUM(not_staff_die), 0) AS death_count",
		).
		Group("to_char(happen_time, 'YYYY')").
		Order("year").
		Scan(&rows).Error
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// ── 应急演练 ──────────────────────────────────────────────────────────

type EmergencyDrillListParams struct {
	Page      int
	PerPage   int
	CompanyID int64
	Name      string
}

func (s *AccidentEmergencyService) ListEmergencyDrill(ctx context.Context, p EmergencyDrillListParams) ([]model.EmergencyDrill, int64, error) {
	page, per := normalizePage(p.Page, p.PerPage)
	q := s.db.WithContext(ctx).Model(&model.EmergencyDrill{}).Where("company_id = ?", p.CompanyID)
	if kw := strings.TrimSpace(p.Name); kw != "" {
		q = q.Where("name ILIKE ?", "%"+kw+"%")
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.EmergencyDrill
	if err := q.Order("id DESC").Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *AccidentEmergencyService) GetEmergencyDrill(ctx context.Context, id int64) (*model.EmergencyDrill, error) {
	var m model.EmergencyDrill
	if err := s.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

func (s *AccidentEmergencyService) CreateEmergencyDrill(ctx context.Context, m *model.EmergencyDrill) error {
	if m.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return wrapWrite(s.db.WithContext(ctx).Create(m).Error)
}

func (s *AccidentEmergencyService) UpdateEmergencyDrill(ctx context.Context, id int64, updates map[string]any) error {
	return updateRow(s.db.WithContext(ctx), &model.EmergencyDrill{}, id, updates)
}

func (s *AccidentEmergencyService) DeleteEmergencyDrill(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.EmergencyDrill{}, id).Error
}

// ── 应急物资 ──────────────────────────────────────────────────────────

type EmergencySuppliesListParams struct {
	Page      int
	PerPage   int
	CompanyID int64
	Name      string
}

func (s *AccidentEmergencyService) ListEmergencySupplies(ctx context.Context, p EmergencySuppliesListParams) ([]model.EmergencySupplies, int64, error) {
	page, per := normalizePage(p.Page, p.PerPage)
	q := s.db.WithContext(ctx).Model(&model.EmergencySupplies{}).Where("company_id = ?", p.CompanyID)
	if kw := strings.TrimSpace(p.Name); kw != "" {
		q = q.Where("name ILIKE ?", "%"+kw+"%")
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.EmergencySupplies
	if err := q.Order("id DESC").Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *AccidentEmergencyService) GetEmergencySupplies(ctx context.Context, id int64) (*model.EmergencySupplies, error) {
	var m model.EmergencySupplies
	if err := s.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

func (s *AccidentEmergencyService) CreateEmergencySupplies(ctx context.Context, m *model.EmergencySupplies) error {
	if m.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return wrapWrite(s.db.WithContext(ctx).Create(m).Error)
}

func (s *AccidentEmergencyService) UpdateEmergencySupplies(ctx context.Context, id int64, updates map[string]any) error {
	return updateRow(s.db.WithContext(ctx), &model.EmergencySupplies{}, id, updates)
}

func (s *AccidentEmergencyService) DeleteEmergencySupplies(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.EmergencySupplies{}, id).Error
}

// ── 应急队伍（含 kids 整体替换） ──────────────────────────────────────

type EmergencyTeamListParams struct {
	Page      int
	PerPage   int
	CompanyID int64
	Name      string
}

func (s *AccidentEmergencyService) ListEmergencyTeam(ctx context.Context, p EmergencyTeamListParams) ([]model.EmergencyTeam, int64, error) {
	page, per := normalizePage(p.Page, p.PerPage)
	q := s.db.WithContext(ctx).Model(&model.EmergencyTeam{}).Where("company_id = ?", p.CompanyID)
	if kw := strings.TrimSpace(p.Name); kw != "" {
		q = q.Where("name ILIKE ?", "%"+kw+"%")
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.EmergencyTeam
	if err := q.Preload("Kids").Order("id DESC").
		Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *AccidentEmergencyService) GetEmergencyTeam(ctx context.Context, id int64) (*model.EmergencyTeam, error) {
	var m model.EmergencyTeam
	if err := s.db.WithContext(ctx).Preload("Kids").First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

type EmergencyTeamInput struct {
	Team *model.EmergencyTeam
	Kids []model.EmergencyTeamKid
}

func (s *AccidentEmergencyService) CreateEmergencyTeam(ctx context.Context, in EmergencyTeamInput) error {
	if in.Team.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(in.Team).Error; err != nil {
			return wrapWrite(err)
		}
		if len(in.Kids) > 0 {
			uid := in.Team.UserID
			cid := in.Team.CompanyID
			for i := range in.Kids {
				in.Kids[i].EmergencyTeamID = in.Team.ID
				in.Kids[i].CompanyID = cid
				in.Kids[i].UserID = uid
				in.Kids[i].LastUserID = uid
			}
			if err := tx.Create(&in.Kids).Error; err != nil {
				return wrapWrite(err)
			}
		}
		return nil
	})
}

func (s *AccidentEmergencyService) UpdateEmergencyTeam(
	ctx context.Context,
	id int64,
	updates map[string]any,
	kids []model.EmergencyTeamKid,
	replaceKids bool,
	uid int64,
) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		res := tx.Model(&model.EmergencyTeam{}).Where("id = ?", id).Updates(updates)
		if res.Error != nil {
			return wrapWrite(res.Error)
		}
		if res.RowsAffected == 0 {
			return errs.New(errs.CodeNotFound, "记录不存在")
		}
		if !replaceKids {
			return nil
		}
		var t model.EmergencyTeam
		if err := tx.First(&t, id).Error; err != nil {
			return err
		}
		if err := tx.Where("emergency_team_id = ?", id).Delete(&model.EmergencyTeamKid{}).Error; err != nil {
			return err
		}
		if len(kids) == 0 {
			return nil
		}
		for i := range kids {
			kids[i].EmergencyTeamID = id
			kids[i].CompanyID = t.CompanyID
			kids[i].UserID = uid
			kids[i].LastUserID = uid
		}
		if err := tx.Create(&kids).Error; err != nil {
			return wrapWrite(err)
		}
		return nil
	})
}

func (s *AccidentEmergencyService) DeleteEmergencyTeam(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("emergency_team_id = ?", id).Delete(&model.EmergencyTeamKid{}).Error; err != nil {
			return err
		}
		return tx.Delete(&model.EmergencyTeam{}, id).Error
	})
}
