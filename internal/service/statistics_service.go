// Package service — 数据统计模块（集团视图 + 生产视图共享）业务层。
//
// StatisticsService 实现 Laravel MonthStatisticsController / YearStatisticsController /
// NoticeStatsController 的所有接口：
//
//   月度数据上报：CRUD + stuff (聚合本企业 + 直接下级) + statistics (跨直接下级汇总)
//   年度数据上报：CRUD + stuff (聚合) + statistics (跨直接下级汇总)
//   通知公告反馈统计：summary (每个下级 office/new 通知的反馈/阅读情况) + pending (尚未反馈/阅读明细)
//
// 子表（month_statistic_controls / month_statistic_accidents / year_statistic_kids）
// 用整体替换语义：service 收到 nil 切片 = 不动，非 nil（含空切片）= 全替换。
package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/model"
)

type StatisticsService struct {
	db *gorm.DB
}

func NewStatisticsService(db *gorm.DB) *StatisticsService {
	return &StatisticsService{db: db}
}

// ════════════════════════════════════════════════════════════════════
//   月度数据上报
// ════════════════════════════════════════════════════════════════════

// MonthStatisticListParams — index 入参。CompanyID 跟 Pid 至少给一个；
// Pid 走集团视图（直接下级），include_pid 把父企业也算进去；CompanyID 走生产视图。
type MonthStatisticListParams struct {
	Page       int
	PerPage    int
	CompanyID  int64
	Pid        int64
	IncludePid bool
}

func (s *StatisticsService) ListMonthStatistic(ctx context.Context, p MonthStatisticListParams) ([]model.MonthStatistic, int64, error) {
	page, per := normalizePage(p.Page, p.PerPage)
	if p.CompanyID == 0 && p.Pid == 0 {
		return nil, 0, errs.New(errs.CodeUnprocessable, "company_id 或 pid 至少提供一个")
	}
	q := s.db.WithContext(ctx).Model(&model.MonthStatistic{})
	if p.Pid > 0 {
		ids, err := s.subordinateIDs(ctx, p.Pid)
		if err != nil {
			return nil, 0, err
		}
		if p.IncludePid {
			ids = append(ids, p.Pid)
		}
		if len(ids) == 0 {
			return []model.MonthStatistic{}, 0, nil
		}
		q = q.Where("company_id IN ?", ids)
	} else {
		q = q.Where("company_id = ?", p.CompanyID)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.MonthStatistic
	if err := q.
		Preload("MonthStatisticControl").
		Preload("MonthStatisticAccident").
		Preload("Company", func(tx *gorm.DB) *gorm.DB { return tx.Select("id, title") }).
		Order("created_at DESC").
		Offset((page - 1) * per).Limit(per).
		Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *StatisticsService) GetMonthStatistic(ctx context.Context, id int64) (*model.MonthStatistic, error) {
	var m model.MonthStatistic
	if err := s.db.WithContext(ctx).
		Preload("MonthStatisticControl").
		Preload("MonthStatisticAccident").
		First(&m, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.New(errs.CodeNotFound, "月报不存在")
		}
		return nil, err
	}
	return &m, nil
}

// MonthStatisticInput — Create / Update 的入参。Controls / Accidents 为 nil
// 时不动子表；非 nil（含空切片）= 整体替换。
type MonthStatisticInput struct {
	Record    *model.MonthStatistic
	Controls  *[]model.MonthStatisticControl
	Accidents *[]model.MonthStatisticAccident
	UserID    int64
}

func (s *StatisticsService) CreateMonthStatistic(ctx context.Context, in MonthStatisticInput) (*model.MonthStatistic, error) {
	if in.Record == nil || in.Record.CompanyID == 0 {
		return nil, errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	// Laravel 行为：当月本企业已有月报时拒绝重复提交。
	now := time.Now()
	var existing int64
	if err := s.db.WithContext(ctx).Model(&model.MonthStatistic{}).
		Where("company_id = ? AND EXTRACT(YEAR FROM created_at) = ? AND EXTRACT(MONTH FROM created_at) = ?",
			in.Record.CompanyID, now.Year(), int(now.Month())).
		Count(&existing).Error; err != nil {
		return nil, err
	}
	if existing > 0 {
		return nil, errs.New(errs.CodeUnprocessable, "你已经提交过当月月报了，如填报错误请前往修改")
	}

	in.Record.UserID = in.UserID
	in.Record.LastUserID = in.UserID
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(in.Record).Error; err != nil {
			return wrapWrite(err)
		}
		if in.Controls != nil {
			if err := s.replaceControls(tx, in.Record.ID, in.Record.CompanyID, in.UserID, *in.Controls); err != nil {
				return err
			}
		}
		if in.Accidents != nil {
			if err := s.replaceAccidents(tx, in.Record.ID, in.Record.CompanyID, in.UserID, *in.Accidents); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return s.GetMonthStatistic(ctx, in.Record.ID)
}

func (s *StatisticsService) UpdateMonthStatistic(ctx context.Context, id int64, in MonthStatisticInput) (*model.MonthStatistic, error) {
	var m model.MonthStatistic
	if err := s.db.WithContext(ctx).First(&m, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.New(errs.CodeNotFound, "月报不存在")
		}
		return nil, err
	}
	companyID := m.CompanyID
	if in.Record != nil && in.Record.CompanyID > 0 {
		companyID = in.Record.CompanyID
	}
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		updates := monthStatisticUpdates(in.Record, in.UserID)
		if len(updates) > 0 {
			if err := tx.Model(&m).Updates(updates).Error; err != nil {
				return wrapWrite(err)
			}
		}
		if in.Controls != nil {
			if err := tx.Where("month_statistic_id = ?", id).Delete(&model.MonthStatisticControl{}).Error; err != nil {
				return wrapWrite(err)
			}
			if err := s.replaceControls(tx, id, companyID, in.UserID, *in.Controls); err != nil {
				return err
			}
		}
		if in.Accidents != nil {
			if err := tx.Where("month_statistic_id = ?", id).Delete(&model.MonthStatisticAccident{}).Error; err != nil {
				return wrapWrite(err)
			}
			if err := s.replaceAccidents(tx, id, companyID, in.UserID, *in.Accidents); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return s.GetMonthStatistic(ctx, id)
}

func (s *StatisticsService) DeleteMonthStatistic(ctx context.Context, id int64) error {
	res := s.db.WithContext(ctx).Delete(&model.MonthStatistic{}, id)
	if res.Error != nil {
		return wrapWrite(res.Error)
	}
	if res.RowsAffected == 0 {
		return errs.New(errs.CodeNotFound, "月报不存在")
	}
	return nil
}

// monthStatisticUpdates — 拷贝可变字段到 updates；patch 语义。
// 这里采用宽松策略：所有列直接整列覆盖（Laravel 行为一致），同时强制 last_user_id。
func monthStatisticUpdates(rec *model.MonthStatistic, userID int64) map[string]any {
	updates := map[string]any{"last_user_id": userID}
	if rec == nil {
		return updates
	}
	if rec.CompanyID > 0 {
		updates["company_id"] = rec.CompanyID
	}
	updates["average_employee"] = rec.AverageEmployee
	updates["safety_meeting_count"] = rec.SafetyMeetingCount
	updates["fire_meeting_count"] = rec.FireMeetingCount
	updates["safety_check_count"] = rec.SafetyCheckCount
	updates["fire_check_count"] = rec.FireCheckCount
	updates["safety_funding"] = rec.SafetyFunding
	updates["fire_funding"] = rec.FireFunding
	updates["hydrating_seedlings"] = rec.HydratingSeedlings
	updates["retrofit_electrical"] = rec.RetrofitElectrical
	updates["fire_extinguisher"] = rec.FireExtinguisher
	updates["smart_fire"] = rec.SmartFire
	updates["safety_training_count"] = rec.SafetyTrainingCount
	updates["safety_training_number"] = rec.SafetyTrainingNumber
	updates["new_people_training_number"] = rec.NewPeopleTrainingNumber
	updates["fire_training_count"] = rec.FireTrainingCount
	updates["fire_training_number"] = rec.FireTrainingNumber
	updates["safety_emergency_count"] = rec.SafetyEmergencyCount
	updates["safety_emergency_number"] = rec.SafetyEmergencyNumber
	updates["fire_emergency_count"] = rec.FireEmergencyCount
	updates["fire_emergency_number"] = rec.FireEmergencyNumber
	updates["some_hidden_danger"] = rec.SomeHiddenDanger
	updates["some_rectify_hidden_danger"] = rec.SomeRectifyHiddenDanger
	updates["no_rectify_info"] = rec.NoRectifyInfo
	updates["big_hidden_danger"] = rec.BigHiddenDanger
	updates["big_rectify_hidden_danger"] = rec.BigRectifyHiddenDanger
	updates["govern_hidden_danger_count"] = rec.GovernHiddenDangerCount
	updates["fire_fighting_danger"] = rec.FireFightingDanger
	updates["fire_fighting_hidden_danger"] = rec.FireFightingHiddenDanger
	updates["fire_fighting_hidden_danger_rectify"] = rec.FireFightingHiddenDangerRectify
	updates["accident_bulletin_count"] = rec.AccidentBulletinCount
	return updates
}

func (s *StatisticsService) replaceControls(tx *gorm.DB, monthID, companyID, userID int64, items []model.MonthStatisticControl) error {
	if len(items) == 0 {
		return nil
	}
	rows := make([]model.MonthStatisticControl, 0, len(items))
	for _, v := range items {
		v.MonthStatisticID = monthID
		v.CompanyID = companyID
		v.UserID = userID
		v.LastUserID = userID
		v.ID = 0
		rows = append(rows, v)
	}
	if err := tx.Create(&rows).Error; err != nil {
		return wrapWrite(err)
	}
	return nil
}

func (s *StatisticsService) replaceAccidents(tx *gorm.DB, monthID, companyID, userID int64, items []model.MonthStatisticAccident) error {
	if len(items) == 0 {
		return nil
	}
	rows := make([]model.MonthStatisticAccident, 0, len(items))
	for _, v := range items {
		v.MonthStatisticID = monthID
		v.CompanyID = companyID
		v.UserID = userID
		v.LastUserID = userID
		v.ID = 0
		rows = append(rows, v)
	}
	if err := tx.Create(&rows).Error; err != nil {
		return wrapWrite(err)
	}
	return nil
}

// ── stuff（预填充） ─────────────────────────────────────────────────

// MonthStuffResponse — 月报预填充：30 个数值字段 + 治理计划/事故详情明细 + 企业上下文。
// 字段名跟 Laravel monthCountKeys() 输出对齐，前端 Index.tsx 直接消费。
type MonthStuffResponse struct {
	AverageEmployee                 int                                 `json:"average_employee"`
	SafetyMeetingCount              int                                 `json:"safety_meeting_count"`
	FireMeetingCount                int                                 `json:"fire_meeting_count"`
	SafetyCheckCount                int                                 `json:"safety_check_count"`
	FireCheckCount                  int                                 `json:"fire_check_count"`
	SafetyFunding                   int                                 `json:"safety_funding"`
	FireFunding                     int                                 `json:"fire_funding"`
	HydratingSeedlings              int                                 `json:"hydrating_seedlings"`
	RetrofitElectrical              int                                 `json:"retrofit_electrical"`
	FireExtinguisher                int                                 `json:"fire_extinguisher"`
	SmartFire                       int                                 `json:"smart_fire"`
	SafetyTrainingCount             int                                 `json:"safety_training_count"`
	SafetyTrainingNumber            int                                 `json:"safety_training_number"`
	NewPeopleTrainingNumber         int                                 `json:"new_people_training_number"`
	FireTrainingCount               int                                 `json:"fire_training_count"`
	FireTrainingNumber              int                                 `json:"fire_training_number"`
	SafetyEmergencyCount            int                                 `json:"safety_emergency_count"`
	SafetyEmergencyNumber           int                                 `json:"safety_emergency_number"`
	FireEmergencyCount              int                                 `json:"fire_emergency_count"`
	FireEmergencyNumber             int                                 `json:"fire_emergency_number"`
	SomeHiddenDanger                int                                 `json:"some_hidden_danger"`
	SomeRectifyHiddenDanger         int                                 `json:"some_rectify_hidden_danger"`
	BigHiddenDanger                 int                                 `json:"big_hidden_danger"`
	BigRectifyHiddenDanger          int                                 `json:"big_rectify_hidden_danger"`
	GovernHiddenDangerCount         int                                 `json:"govern_hidden_danger_count"`
	FireFightingDanger              int                                 `json:"fire_fighting_danger"`
	FireFightingHiddenDanger        int                                 `json:"fire_fighting_hidden_danger"`
	FireFightingHiddenDangerRectify int                                 `json:"fire_fighting_hidden_danger_rectify"`
	AccidentBulletinCount           int                                 `json:"accident_bulletin_count"`
	NoRectifyInfo                   *string                             `json:"no_rectify_info"`
	GovernHiddenDanger              []model.MonthlyInvestigation        `json:"govern_hidden_danger"`
	AccidentBulletin                []MonthAccidentDetail               `json:"accident_bulletin"`
	CompanyID                       int64                               `json:"company_id"`
	CompanyName                     string                              `json:"company_name"`
}

// MonthAccidentDetail — accident 子表预填项（合并 Kid + 父 AccidentBulletin 字段）。
type MonthAccidentDetail struct {
	Name               *string `json:"name"`
	Department         *string `json:"department"`
	Sex                *string `json:"sex"`
	Age                *int32  `json:"age"`
	WorkType           *string `json:"work_type"`
	WorkAge            *int32  `json:"work_age"`
	EducationLevel     *string `json:"education_level"`
	DegreeInjury       *string `json:"degree_injury"`
	DaysOff            *int32  `json:"days_off"`
	HappenTime         *model.Date `json:"happen_time"`
	Type               *string `json:"type"`
	DirectCause        *string `json:"direct_cause"`
	Content            *string `json:"content"`
	DirectLossProperty *string `json:"direct_loss_property"`
}

func (s *StatisticsService) MonthStuff(ctx context.Context, companyID int64) (*MonthStuffResponse, error) {
	var company model.Company
	if err := s.db.WithContext(ctx).First(&company, companyID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.New(errs.CodeNotFound, "企业不存在")
		}
		return nil, err
	}
	now := time.Now()
	year := now.Year()
	month := int(now.Month())

	var selfCounts monthCounts
	if company.IsProduce == 1 {
		c, err := s.aggregateMonth(ctx, []int64{company.ID}, year, month)
		if err != nil {
			return nil, err
		}
		selfCounts = c
	}

	var subIDs []int64
	if company.IsBloc == 1 {
		ids, err := s.subordinateIDs(ctx, company.ID)
		if err != nil {
			return nil, err
		}
		subIDs = ids
	}
	subCounts, err := s.aggregateMonth(ctx, subIDs, year, month)
	if err != nil {
		return nil, err
	}

	// 起始职工数量：本企业 Personnel.total_people（仅生产端） + 下级聚合。
	if company.IsProduce == 1 {
		var personnel model.Personnel
		err := s.db.WithContext(ctx).Where("company_id = ?", company.ID).First(&personnel).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		selfCounts.AverageEmployee = personnel.TotalPeople + subCounts.AverageEmployee
	} else {
		selfCounts.AverageEmployee = subCounts.AverageEmployee
	}

	combined := selfCounts.add(subCounts)
	combined.AverageEmployee = selfCounts.AverageEmployee // 已包含子聚合，避免重复加

	// 治理计划详情 + 事故详情：本企业 + 直接下级合并去重。
	detailIDs := uniqueIDs(append([]int64{company.ID}, subIDs...))

	var governs []model.MonthlyInvestigation
	if err := s.db.WithContext(ctx).
		Where("company_id IN ?", detailIDs).
		Where("EXTRACT(YEAR FROM inv_time) = ?", year).
		Where("EXTRACT(MONTH FROM inv_time) = ?", month).
		Where("level = ?", "重大隐患").
		Where("is_govern = 1").
		Find(&governs).Error; err != nil {
		return nil, err
	}

	var kids []model.AccidentBulletinKid
	if err := s.db.WithContext(ctx).
		Where("company_id IN ?", detailIDs).
		Where("EXTRACT(YEAR FROM created_at) = ?", year).
		Where("EXTRACT(MONTH FROM created_at) = ?", month).
		Find(&kids).Error; err != nil {
		return nil, err
	}
	bulletinIDs := make([]int64, 0, len(kids))
	for _, k := range kids {
		bulletinIDs = append(bulletinIDs, k.AccidentBulletinID)
	}
	bulletins := map[int64]*model.AccidentBulletin{}
	if len(bulletinIDs) > 0 {
		var bs []model.AccidentBulletin
		if err := s.db.WithContext(ctx).Where("id IN ?", uniqueIDs(bulletinIDs)).Find(&bs).Error; err != nil {
			return nil, err
		}
		for i := range bs {
			bulletins[bs[i].ID] = &bs[i]
		}
	}
	accidents := make([]MonthAccidentDetail, 0, len(kids))
	for _, k := range kids {
		bulletin := bulletins[k.AccidentBulletinID]
		d := MonthAccidentDetail{
			Name:           k.Name,
			Department:     k.Department,
			Sex:            k.Sex,
			Age:            k.Age,
			WorkType:       k.WorkType,
			WorkAge:        k.WorkAge,
			EducationLevel: k.EducationLevel,
			DegreeInjury:   k.DegreeInjury,
			DaysOff:        k.DaysOff,
		}
		if d.Department == nil && bulletin != nil {
			d.Department = bulletin.Department
		}
		if d.DegreeInjury == nil && bulletin != nil {
			d.DegreeInjury = bulletin.DegreeInjury
		}
		if bulletin != nil {
			d.HappenTime = bulletin.HappenTime
			d.Type = bulletin.Type
			d.DirectCause = bulletin.DirectCause
			d.Content = bulletin.Content
			d.DirectLossProperty = bulletin.DirectLossProperty
		}
		accidents = append(accidents, d)
	}

	return &MonthStuffResponse{
		AverageEmployee:                 combined.AverageEmployee,
		SafetyMeetingCount:              combined.SafetyMeetingCount,
		FireMeetingCount:                combined.FireMeetingCount,
		SafetyCheckCount:                combined.SafetyCheckCount,
		FireCheckCount:                  combined.FireCheckCount,
		SafetyFunding:                   combined.SafetyFunding,
		FireFunding:                     combined.FireFunding,
		HydratingSeedlings:              combined.HydratingSeedlings,
		RetrofitElectrical:              combined.RetrofitElectrical,
		FireExtinguisher:                combined.FireExtinguisher,
		SmartFire:                       combined.SmartFire,
		SafetyTrainingCount:             combined.SafetyTrainingCount,
		SafetyTrainingNumber:            combined.SafetyTrainingNumber,
		NewPeopleTrainingNumber:         combined.NewPeopleTrainingNumber,
		FireTrainingCount:               combined.FireTrainingCount,
		FireTrainingNumber:              combined.FireTrainingNumber,
		SafetyEmergencyCount:            combined.SafetyEmergencyCount,
		SafetyEmergencyNumber:           combined.SafetyEmergencyNumber,
		FireEmergencyCount:              combined.FireEmergencyCount,
		FireEmergencyNumber:             combined.FireEmergencyNumber,
		SomeHiddenDanger:                combined.SomeHiddenDanger,
		SomeRectifyHiddenDanger:         combined.SomeRectifyHiddenDanger,
		BigHiddenDanger:                 combined.BigHiddenDanger,
		BigRectifyHiddenDanger:          combined.BigRectifyHiddenDanger,
		GovernHiddenDangerCount:         combined.GovernHiddenDangerCount,
		FireFightingDanger:              combined.FireFightingDanger,
		FireFightingHiddenDanger:        combined.FireFightingHiddenDanger,
		FireFightingHiddenDangerRectify: combined.FireFightingHiddenDangerRectify,
		AccidentBulletinCount:           combined.AccidentBulletinCount,
		NoRectifyInfo:                   nil,
		GovernHiddenDanger:              governs,
		AccidentBulletin:                accidents,
		CompanyID:                       company.ID,
		CompanyName:                     company.Title,
	}, nil
}

// monthCounts — aggregateMonth 内部表示；同时被 MonthStatisticsResponse 嵌入做汇总输出。
// 字段 json 标签直接对齐 Laravel monthCountKeys()。
type monthCounts struct {
	AverageEmployee                 int `json:"average_employee"`
	SafetyMeetingCount              int `json:"safety_meeting_count"`
	FireMeetingCount                int `json:"fire_meeting_count"`
	SafetyCheckCount                int `json:"safety_check_count"`
	FireCheckCount                  int `json:"fire_check_count"`
	SafetyFunding                   int `json:"safety_funding"`
	FireFunding                     int `json:"fire_funding"`
	HydratingSeedlings              int `json:"hydrating_seedlings"`
	RetrofitElectrical              int `json:"retrofit_electrical"`
	FireExtinguisher                int `json:"fire_extinguisher"`
	SmartFire                       int `json:"smart_fire"`
	SafetyTrainingCount             int `json:"safety_training_count"`
	SafetyTrainingNumber            int `json:"safety_training_number"`
	NewPeopleTrainingNumber         int `json:"new_people_training_number"`
	FireTrainingCount               int `json:"fire_training_count"`
	FireTrainingNumber              int `json:"fire_training_number"`
	SafetyEmergencyCount            int `json:"safety_emergency_count"`
	SafetyEmergencyNumber           int `json:"safety_emergency_number"`
	FireEmergencyCount              int `json:"fire_emergency_count"`
	FireEmergencyNumber             int `json:"fire_emergency_number"`
	SomeHiddenDanger                int `json:"some_hidden_danger"`
	SomeRectifyHiddenDanger         int `json:"some_rectify_hidden_danger"`
	BigHiddenDanger                 int `json:"big_hidden_danger"`
	BigRectifyHiddenDanger          int `json:"big_rectify_hidden_danger"`
	GovernHiddenDangerCount         int `json:"govern_hidden_danger_count"`
	FireFightingDanger              int `json:"fire_fighting_danger"`
	FireFightingHiddenDanger        int `json:"fire_fighting_hidden_danger"`
	FireFightingHiddenDangerRectify int `json:"fire_fighting_hidden_danger_rectify"`
	AccidentBulletinCount           int `json:"accident_bulletin_count"`
}

func (a monthCounts) add(b monthCounts) monthCounts {
	return monthCounts{
		AverageEmployee:                 a.AverageEmployee + b.AverageEmployee,
		SafetyMeetingCount:              a.SafetyMeetingCount + b.SafetyMeetingCount,
		FireMeetingCount:                a.FireMeetingCount + b.FireMeetingCount,
		SafetyCheckCount:                a.SafetyCheckCount + b.SafetyCheckCount,
		FireCheckCount:                  a.FireCheckCount + b.FireCheckCount,
		SafetyFunding:                   a.SafetyFunding + b.SafetyFunding,
		FireFunding:                     a.FireFunding + b.FireFunding,
		HydratingSeedlings:              a.HydratingSeedlings + b.HydratingSeedlings,
		RetrofitElectrical:              a.RetrofitElectrical + b.RetrofitElectrical,
		FireExtinguisher:                a.FireExtinguisher + b.FireExtinguisher,
		SmartFire:                       a.SmartFire + b.SmartFire,
		SafetyTrainingCount:             a.SafetyTrainingCount + b.SafetyTrainingCount,
		SafetyTrainingNumber:            a.SafetyTrainingNumber + b.SafetyTrainingNumber,
		NewPeopleTrainingNumber:         a.NewPeopleTrainingNumber + b.NewPeopleTrainingNumber,
		FireTrainingCount:               a.FireTrainingCount + b.FireTrainingCount,
		FireTrainingNumber:              a.FireTrainingNumber + b.FireTrainingNumber,
		SafetyEmergencyCount:            a.SafetyEmergencyCount + b.SafetyEmergencyCount,
		SafetyEmergencyNumber:           a.SafetyEmergencyNumber + b.SafetyEmergencyNumber,
		FireEmergencyCount:              a.FireEmergencyCount + b.FireEmergencyCount,
		FireEmergencyNumber:             a.FireEmergencyNumber + b.FireEmergencyNumber,
		SomeHiddenDanger:                a.SomeHiddenDanger + b.SomeHiddenDanger,
		SomeRectifyHiddenDanger:         a.SomeRectifyHiddenDanger + b.SomeRectifyHiddenDanger,
		BigHiddenDanger:                 a.BigHiddenDanger + b.BigHiddenDanger,
		BigRectifyHiddenDanger:          a.BigRectifyHiddenDanger + b.BigRectifyHiddenDanger,
		GovernHiddenDangerCount:         a.GovernHiddenDangerCount + b.GovernHiddenDangerCount,
		FireFightingDanger:              a.FireFightingDanger + b.FireFightingDanger,
		FireFightingHiddenDanger:        a.FireFightingHiddenDanger + b.FireFightingHiddenDanger,
		FireFightingHiddenDangerRectify: a.FireFightingHiddenDangerRectify + b.FireFightingHiddenDangerRectify,
		AccidentBulletinCount:           a.AccidentBulletinCount + b.AccidentBulletinCount,
	}
}

// aggregateMonth — 聚合给定企业集合在 year/month 的当月生产数据。
// 跟 Laravel MonthStatisticsController::aggregateMonth 一一对应；空集合直接返回零值。
func (s *StatisticsService) aggregateMonth(ctx context.Context, companyIDs []int64, year, month int) (monthCounts, error) {
	if len(companyIDs) == 0 {
		return monthCounts{}, nil
	}
	db := s.db.WithContext(ctx)
	var c monthCounts

	// Personnel 起始职工：sum(total_people)。
	var avgRow struct {
		Total int
	}
	if err := db.Model(&model.Personnel{}).
		Where("company_id IN ?", companyIDs).
		Select("COALESCE(SUM(total_people), 0) AS total").Scan(&avgRow).Error; err != nil {
		return c, err
	}
	c.AverageEmployee = avgRow.Total

	// InternalMeeting：CASE WHEN type 拆 安全/消防 例会次数（按 created_at 月份）。
	var meetingRow struct {
		SafetyMeeting int
		FireMeeting   int
	}
	if err := db.Model(&model.InternalMeeting{}).
		Where("company_id IN ?", companyIDs).
		Where("EXTRACT(YEAR FROM created_at) = ?", year).
		Where("EXTRACT(MONTH FROM created_at) = ?", month).
		Select(`COALESCE(SUM(CASE WHEN type = '安全例会' THEN 1 ELSE 0 END), 0) AS safety_meeting,
		         COALESCE(SUM(CASE WHEN type = '消防例会' THEN 1 ELSE 0 END), 0) AS fire_meeting`).
		Scan(&meetingRow).Error; err != nil {
		return c, err
	}
	c.SafetyMeetingCount = meetingRow.SafetyMeeting
	c.FireMeetingCount = meetingRow.FireMeeting

	// MonthlyInvestigation（按 created_at 月份）：safety_check / fire_check。
	var checkRow struct {
		SafetyCheck int
		FireCheck   int
	}
	if err := db.Model(&model.MonthlyInvestigation{}).
		Where("company_id IN ?", companyIDs).
		Where("EXTRACT(YEAR FROM created_at) = ?", year).
		Where("EXTRACT(MONTH FROM created_at) = ?", month).
		Select(`COALESCE(SUM(CASE WHEN type <> '消防' THEN 1 ELSE 0 END), 0) AS safety_check,
		         COALESCE(SUM(CASE WHEN type = '消防' THEN 1 ELSE 0 END), 0) AS fire_check`).
		Scan(&checkRow).Error; err != nil {
		return c, err
	}
	c.SafetyCheckCount = checkRow.SafetyCheck
	c.FireCheckCount = checkRow.FireCheck

	// FundingExecution：sum(budget) JOIN security_investments.type。
	var safetyFunding float64
	if err := db.Model(&model.FundingExecution{}).
		Joins("LEFT JOIN security_investments si ON si.id = funding_executions.security_investment_id").
		Where("funding_executions.company_id IN ?", companyIDs).
		Where("EXTRACT(YEAR FROM funding_executions.created_at) = ?", year).
		Where("EXTRACT(MONTH FROM funding_executions.created_at) = ?", month).
		Where("si.type = ?", "安全投入").
		Select("COALESCE(SUM(funding_executions.budget), 0)").Scan(&safetyFunding).Error; err != nil {
		return c, err
	}
	c.SafetyFunding = int(safetyFunding)
	var fireFunding float64
	if err := db.Model(&model.FundingExecution{}).
		Joins("LEFT JOIN security_investments si ON si.id = funding_executions.security_investment_id").
		Where("funding_executions.company_id IN ?", companyIDs).
		Where("EXTRACT(YEAR FROM funding_executions.created_at) = ?", year).
		Where("EXTRACT(MONTH FROM funding_executions.created_at) = ?", month).
		Where("si.type = ?", "消防投入").
		Select("COALESCE(SUM(funding_executions.budget), 0)").Scan(&fireFunding).Error; err != nil {
		return c, err
	}
	c.FireFunding = int(fireFunding)

	// FireEquipment：按 special_type 分桶计数。
	type kv struct {
		SpecialType string
		Cnt         int
	}
	var equipRows []kv
	if err := db.Model(&model.FireEquipment{}).
		Where("company_id IN ?", companyIDs).
		Where("EXTRACT(YEAR FROM created_at) = ?", year).
		Where("EXTRACT(MONTH FROM created_at) = ?", month).
		Select("special_type, COUNT(*) AS cnt").
		Group("special_type").
		Scan(&equipRows).Error; err != nil {
		return c, err
	}
	for _, r := range equipRows {
		switch r.SpecialType {
		case "水带苗子":
			c.HydratingSeedlings = r.Cnt
		case "改造电气线路":
			c.RetrofitElectrical = r.Cnt
		case "灭火器":
			c.FireExtinguisher = r.Cnt
		}
	}

	// FireFacility：单一 special_type=智慧消防系统 计数。
	var smartFire int64
	if err := db.Model(&model.FireFacility{}).
		Where("company_id IN ?", companyIDs).
		Where("EXTRACT(YEAR FROM created_at) = ?", year).
		Where("EXTRACT(MONTH FROM created_at) = ?", month).
		Where("special_type = ?", "智慧消防系统").
		Count(&smartFire).Error; err != nil {
		return c, err
	}
	c.SmartFire = int(smartFire)

	// InternalTraining：按 training_date 月份，training_type 分桶 cnt + sum(number)。
	type trainingKV struct {
		TrainingType string
		Cnt          int
		Num          int
	}
	var trainRows []trainingKV
	if err := db.Model(&model.InternalTraining{}).
		Where("company_id IN ?", companyIDs).
		Where("EXTRACT(YEAR FROM training_date) = ?", year).
		Where("EXTRACT(MONTH FROM training_date) = ?", month).
		Select("training_type, COUNT(*) AS cnt, COALESCE(SUM(number), 0) AS num").
		Group("training_type").
		Scan(&trainRows).Error; err != nil {
		return c, err
	}
	for _, r := range trainRows {
		switch r.TrainingType {
		case "安全培训":
			c.SafetyTrainingCount = r.Cnt
			c.SafetyTrainingNumber = r.Num
		case "新员工培训":
			c.NewPeopleTrainingNumber = r.Num
		case "消防培训":
			c.FireTrainingCount = r.Cnt
			c.FireTrainingNumber = r.Num
		}
	}

	// EmergencyDrill：按 drill_time 月份，drill_type 分桶 cnt + sum(drill_number)。
	type drillKV struct {
		DrillType string
		Cnt       int
		Num       int
	}
	var drillRows []drillKV
	if err := db.Model(&model.EmergencyDrill{}).
		Where("company_id IN ?", companyIDs).
		Where("EXTRACT(YEAR FROM drill_time) = ?", year).
		Where("EXTRACT(MONTH FROM drill_time) = ?", month).
		// drill_number 是 VARCHAR — Laravel 按字符串 SUM 隐式转，PG 必须显式 cast。
		Select("drill_type, COUNT(*) AS cnt, COALESCE(SUM(NULLIF(drill_number, '')::numeric), 0) AS num").
		Group("drill_type").
		Scan(&drillRows).Error; err != nil {
		return c, err
	}
	for _, r := range drillRows {
		switch r.DrillType {
		case "安全生产":
			c.SafetyEmergencyCount = r.Cnt
			c.SafetyEmergencyNumber = r.Num
		case "消防安全":
			c.FireEmergencyCount = r.Cnt
			c.FireEmergencyNumber = r.Num
		}
	}

	// MonthlyInvestigation（按 inv_time 月份）：8 个 CASE WHEN 维度。
	var invRow struct {
		SomeHidden     int
		SomeRectify    int
		BigHidden      int
		BigRectify     int
		GovernCount    int
		FireDanger     int
		FireLaw        int
		FireLawRectify int
	}
	if err := db.Model(&model.MonthlyInvestigation{}).
		Where("company_id IN ?", companyIDs).
		Where("EXTRACT(YEAR FROM inv_time) = ?", year).
		Where("EXTRACT(MONTH FROM inv_time) = ?", month).
		Select(`COALESCE(SUM(CASE WHEN level = '一般隐患' THEN 1 ELSE 0 END), 0) AS some_hidden,
		         COALESCE(SUM(CASE WHEN level = '一般隐患' AND correct_result <> '未整改' THEN 1 ELSE 0 END), 0) AS some_rectify,
		         COALESCE(SUM(CASE WHEN level = '重大隐患' THEN 1 ELSE 0 END), 0) AS big_hidden,
		         COALESCE(SUM(CASE WHEN level = '重大隐患' AND correct_result <> '未整改' THEN 1 ELSE 0 END), 0) AS big_rectify,
		         COALESCE(SUM(CASE WHEN level = '重大隐患' AND is_govern = 1 THEN 1 ELSE 0 END), 0) AS govern_count,
		         COALESCE(SUM(CASE WHEN type = '消防' THEN 1 ELSE 0 END), 0) AS fire_danger,
		         COALESCE(SUM(CASE WHEN is_fire_fighting = 1 THEN 1 ELSE 0 END), 0) AS fire_law,
		         COALESCE(SUM(CASE WHEN is_fire_fighting = 1 AND correct_result <> '未整改' THEN 1 ELSE 0 END), 0) AS fire_law_rectify`).
		Scan(&invRow).Error; err != nil {
		return c, err
	}
	c.SomeHiddenDanger = invRow.SomeHidden
	c.SomeRectifyHiddenDanger = invRow.SomeRectify
	c.BigHiddenDanger = invRow.BigHidden
	c.BigRectifyHiddenDanger = invRow.BigRectify
	c.GovernHiddenDangerCount = invRow.GovernCount
	c.FireFightingDanger = invRow.FireDanger
	c.FireFightingHiddenDanger = invRow.FireLaw
	c.FireFightingHiddenDangerRectify = invRow.FireLawRectify

	// AccidentBulletin：按 created_at 月份计数。
	var accidentCount int64
	if err := db.Model(&model.AccidentBulletin{}).
		Where("company_id IN ?", companyIDs).
		Where("EXTRACT(YEAR FROM created_at) = ?", year).
		Where("EXTRACT(MONTH FROM created_at) = ?", month).
		Count(&accidentCount).Error; err != nil {
		return c, err
	}
	c.AccidentBulletinCount = int(accidentCount)

	return c, nil
}

// ── statistics（汇总） ─────────────────────────────────────────────

// MonthStatisticsResponse — 30 列 SUM + 数据条数 + date + 下级数。embedded
// monthCounts 字段已带 json tag，直接平铺到顶层跟 Laravel 输出对齐。
type MonthStatisticsResponse struct {
	monthCounts
	CountNumber      int    `json:"count_number"`
	Date             string `json:"date"`
	SubordinateCount int    `json:"subordinate_count"`
}

func (s *StatisticsService) MonthStatisticStatistics(ctx context.Context, companyID int64, dateTime string) (*MonthStatisticsResponse, error) {
	if companyID <= 0 {
		return nil, errs.New(errs.CodeUnprocessable, "company_id 必填")
	}
	year, month, err := parseYearMonth(dateTime)
	if err != nil {
		return nil, err
	}
	subIDs, err := s.subordinateIDs(ctx, companyID)
	if err != nil {
		return nil, err
	}
	resp := &MonthStatisticsResponse{
		Date:             fmt.Sprintf("%04d-%02d", year, month),
		SubordinateCount: len(subIDs),
	}
	if len(subIDs) == 0 {
		return resp, nil
	}

	type sumRow struct {
		AverageEmployee                 int
		SafetyMeetingCount              int
		FireMeetingCount                int
		SafetyCheckCount                int
		FireCheckCount                  int
		SafetyFunding                   int
		FireFunding                     int
		HydratingSeedlings              int
		RetrofitElectrical              int
		FireExtinguisher                int
		SmartFire                       int
		SafetyTrainingCount             int
		SafetyTrainingNumber            int
		NewPeopleTrainingNumber         int
		FireTrainingCount               int
		FireTrainingNumber              int
		SafetyEmergencyCount            int
		SafetyEmergencyNumber           int
		FireEmergencyCount              int
		FireEmergencyNumber             int
		SomeHiddenDanger                int
		SomeRectifyHiddenDanger         int
		BigHiddenDanger                 int
		BigRectifyHiddenDanger          int
		GovernHiddenDangerCount         int
		FireFightingDanger              int
		FireFightingHiddenDanger        int
		FireFightingHiddenDangerRectify int
		AccidentBulletinCount           int
		CountNumber                     int
	}
	var row sumRow
	expr := `COALESCE(SUM(average_employee), 0) AS average_employee,
		COALESCE(SUM(safety_meeting_count), 0) AS safety_meeting_count,
		COALESCE(SUM(fire_meeting_count), 0) AS fire_meeting_count,
		COALESCE(SUM(safety_check_count), 0) AS safety_check_count,
		COALESCE(SUM(fire_check_count), 0) AS fire_check_count,
		COALESCE(SUM(safety_funding), 0) AS safety_funding,
		COALESCE(SUM(fire_funding), 0) AS fire_funding,
		COALESCE(SUM(hydrating_seedlings), 0) AS hydrating_seedlings,
		COALESCE(SUM(retrofit_electrical), 0) AS retrofit_electrical,
		COALESCE(SUM(fire_extinguisher), 0) AS fire_extinguisher,
		COALESCE(SUM(smart_fire), 0) AS smart_fire,
		COALESCE(SUM(safety_training_count), 0) AS safety_training_count,
		COALESCE(SUM(safety_training_number), 0) AS safety_training_number,
		COALESCE(SUM(new_people_training_number), 0) AS new_people_training_number,
		COALESCE(SUM(fire_training_count), 0) AS fire_training_count,
		COALESCE(SUM(fire_training_number), 0) AS fire_training_number,
		COALESCE(SUM(safety_emergency_count), 0) AS safety_emergency_count,
		COALESCE(SUM(safety_emergency_number), 0) AS safety_emergency_number,
		COALESCE(SUM(fire_emergency_count), 0) AS fire_emergency_count,
		COALESCE(SUM(fire_emergency_number), 0) AS fire_emergency_number,
		COALESCE(SUM(some_hidden_danger), 0) AS some_hidden_danger,
		COALESCE(SUM(some_rectify_hidden_danger), 0) AS some_rectify_hidden_danger,
		COALESCE(SUM(big_hidden_danger), 0) AS big_hidden_danger,
		COALESCE(SUM(big_rectify_hidden_danger), 0) AS big_rectify_hidden_danger,
		COALESCE(SUM(govern_hidden_danger_count), 0) AS govern_hidden_danger_count,
		COALESCE(SUM(fire_fighting_danger), 0) AS fire_fighting_danger,
		COALESCE(SUM(fire_fighting_hidden_danger), 0) AS fire_fighting_hidden_danger,
		COALESCE(SUM(fire_fighting_hidden_danger_rectify), 0) AS fire_fighting_hidden_danger_rectify,
		COALESCE(SUM(accident_bulletin_count), 0) AS accident_bulletin_count,
		COUNT(*) AS count_number`
	if err := s.db.WithContext(ctx).Model(&model.MonthStatistic{}).
		Where("company_id IN ?", subIDs).
		Where("EXTRACT(YEAR FROM created_at) = ?", year).
		Where("EXTRACT(MONTH FROM created_at) = ?", month).
		Select(expr).Scan(&row).Error; err != nil {
		return nil, err
	}
	resp.AverageEmployee = row.AverageEmployee
	resp.SafetyMeetingCount = row.SafetyMeetingCount
	resp.FireMeetingCount = row.FireMeetingCount
	resp.SafetyCheckCount = row.SafetyCheckCount
	resp.FireCheckCount = row.FireCheckCount
	resp.SafetyFunding = row.SafetyFunding
	resp.FireFunding = row.FireFunding
	resp.HydratingSeedlings = row.HydratingSeedlings
	resp.RetrofitElectrical = row.RetrofitElectrical
	resp.FireExtinguisher = row.FireExtinguisher
	resp.SmartFire = row.SmartFire
	resp.SafetyTrainingCount = row.SafetyTrainingCount
	resp.SafetyTrainingNumber = row.SafetyTrainingNumber
	resp.NewPeopleTrainingNumber = row.NewPeopleTrainingNumber
	resp.FireTrainingCount = row.FireTrainingCount
	resp.FireTrainingNumber = row.FireTrainingNumber
	resp.SafetyEmergencyCount = row.SafetyEmergencyCount
	resp.SafetyEmergencyNumber = row.SafetyEmergencyNumber
	resp.FireEmergencyCount = row.FireEmergencyCount
	resp.FireEmergencyNumber = row.FireEmergencyNumber
	resp.SomeHiddenDanger = row.SomeHiddenDanger
	resp.SomeRectifyHiddenDanger = row.SomeRectifyHiddenDanger
	resp.BigHiddenDanger = row.BigHiddenDanger
	resp.BigRectifyHiddenDanger = row.BigRectifyHiddenDanger
	resp.GovernHiddenDangerCount = row.GovernHiddenDangerCount
	resp.FireFightingDanger = row.FireFightingDanger
	resp.FireFightingHiddenDanger = row.FireFightingHiddenDanger
	resp.FireFightingHiddenDangerRectify = row.FireFightingHiddenDangerRectify
	resp.AccidentBulletinCount = row.AccidentBulletinCount
	resp.CountNumber = row.CountNumber
	return resp, nil
}

// ════════════════════════════════════════════════════════════════════
//   年度数据上报
// ════════════════════════════════════════════════════════════════════

type YearStatisticListParams struct {
	Page       int
	PerPage    int
	CompanyID  int64
	Pid        int64
	IncludePid bool
}

func (s *StatisticsService) ListYearStatistic(ctx context.Context, p YearStatisticListParams) ([]model.YearStatistic, int64, error) {
	page, per := normalizePage(p.Page, p.PerPage)
	if p.CompanyID == 0 && p.Pid == 0 {
		return nil, 0, errs.New(errs.CodeUnprocessable, "company_id 或 pid 至少提供一个")
	}
	q := s.db.WithContext(ctx).Model(&model.YearStatistic{})
	if p.Pid > 0 {
		ids, err := s.subordinateIDs(ctx, p.Pid)
		if err != nil {
			return nil, 0, err
		}
		if p.IncludePid {
			ids = append(ids, p.Pid)
		}
		if len(ids) == 0 {
			return []model.YearStatistic{}, 0, nil
		}
		q = q.Where("company_id IN ?", ids)
	} else {
		q = q.Where("company_id = ?", p.CompanyID)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.YearStatistic
	if err := q.
		Preload("YearStatisticKid").
		Preload("Company", func(tx *gorm.DB) *gorm.DB { return tx.Select("id, title") }).
		Order("id DESC").
		Offset((page - 1) * per).Limit(per).
		Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *StatisticsService) GetYearStatistic(ctx context.Context, id int64) (*model.YearStatistic, error) {
	var m model.YearStatistic
	if err := s.db.WithContext(ctx).
		Preload("YearStatisticKid").
		First(&m, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.New(errs.CodeNotFound, "年度统计不存在")
		}
		return nil, err
	}
	return &m, nil
}

type YearStatisticInput struct {
	Record *model.YearStatistic
	Kids   *[]model.YearStatisticKid
	UserID int64
}

func (s *StatisticsService) CreateYearStatistic(ctx context.Context, in YearStatisticInput) (*model.YearStatistic, error) {
	if in.Record == nil || in.Record.CompanyID == 0 {
		return nil, errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	now := time.Now()
	var existing int64
	if err := s.db.WithContext(ctx).Model(&model.YearStatistic{}).
		Where("company_id = ? AND EXTRACT(YEAR FROM created_at) = ?", in.Record.CompanyID, now.Year()).
		Count(&existing).Error; err != nil {
		return nil, err
	}
	if existing > 0 {
		return nil, errs.New(errs.CodeUnprocessable, "你已经提交过当年年报了，如填报错误请前往修改")
	}
	in.Record.UserID = in.UserID
	in.Record.LastUserID = in.UserID
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(in.Record).Error; err != nil {
			return wrapWrite(err)
		}
		if in.Kids != nil {
			if err := s.replaceYearKids(tx, in.Record.ID, in.Record.CompanyID, in.UserID, *in.Kids); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return s.GetYearStatistic(ctx, in.Record.ID)
}

func (s *StatisticsService) UpdateYearStatistic(ctx context.Context, id int64, in YearStatisticInput) (*model.YearStatistic, error) {
	var m model.YearStatistic
	if err := s.db.WithContext(ctx).First(&m, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.New(errs.CodeNotFound, "年度统计不存在")
		}
		return nil, err
	}
	companyID := m.CompanyID
	if in.Record != nil && in.Record.CompanyID > 0 {
		companyID = in.Record.CompanyID
	}
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		updates := yearStatisticUpdates(in.Record, in.UserID)
		if len(updates) > 0 {
			if err := tx.Model(&m).Updates(updates).Error; err != nil {
				return wrapWrite(err)
			}
		}
		if in.Kids != nil {
			if err := tx.Where("year_statistic_id = ?", id).Delete(&model.YearStatisticKid{}).Error; err != nil {
				return wrapWrite(err)
			}
			if err := s.replaceYearKids(tx, id, companyID, in.UserID, *in.Kids); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return s.GetYearStatistic(ctx, id)
}

func (s *StatisticsService) DeleteYearStatistic(ctx context.Context, id int64) error {
	res := s.db.WithContext(ctx).Delete(&model.YearStatistic{}, id)
	if res.Error != nil {
		return wrapWrite(res.Error)
	}
	if res.RowsAffected == 0 {
		return errs.New(errs.CodeNotFound, "年度统计不存在")
	}
	return nil
}

func yearStatisticUpdates(rec *model.YearStatistic, userID int64) map[string]any {
	updates := map[string]any{"last_user_id": userID}
	if rec == nil {
		return updates
	}
	if rec.CompanyID > 0 {
		updates["company_id"] = rec.CompanyID
	}
	updates["company_name"] = rec.CompanyName
	updates["base_area"] = rec.BaseArea
	updates["enterprise_type"] = rec.EnterpriseType
	updates["safety_funding"] = rec.SafetyFunding
	updates["fire_fighting"] = rec.FireFighting
	updates["compliance_time"] = rec.ComplianceTime
	updates["revaluation_time"] = rec.RevaluationTime
	return updates
}

func (s *StatisticsService) replaceYearKids(tx *gorm.DB, yearID, companyID, userID int64, items []model.YearStatisticKid) error {
	if len(items) == 0 {
		return nil
	}
	rows := make([]model.YearStatisticKid, 0, len(items))
	for _, v := range items {
		v.YearStatisticID = yearID
		v.CompanyID = companyID
		v.UserID = userID
		v.LastUserID = userID
		v.ID = 0
		rows = append(rows, v)
	}
	if err := tx.Create(&rows).Error; err != nil {
		return wrapWrite(err)
	}
	return nil
}

// YearStuffResponse — 年报预填充。kids 自动按 type=1（三类人员） / type=2（消防控制室）
// 分组拼接，跟前端表单分两个 tab 渲染对齐。
type YearStuffResponse struct {
	CompanyID       int64                 `json:"company_id"`
	CompanyName     string                `json:"company_name"`
	EnterpriseType  string                `json:"enterprise_type"`
	BaseArea        string                `json:"base_area"`
	ComplianceTime  *model.Date           `json:"compliance_time"`
	RevaluationTime *model.Date           `json:"revaluation_time"`
	SafetyFunding   int                   `json:"safety_funding"`
	FireFighting    int                   `json:"fire_fighting"`
	Kids            []YearStuffKid        `json:"kids"`
}

type YearStuffKid struct {
	Name                *string `json:"name"`
	Type                int     `json:"type"`
	ShouldObtainPeople  *int    `json:"should_obtain_people"`
	AlreadyObtainPeople *int    `json:"already_obtain_people"`
	NoObtainPeople      *string `json:"no_obtain_people"`
	TrainPeople         *int    `json:"train_people"`
}

func (s *StatisticsService) YearStuff(ctx context.Context, companyID int64) (*YearStuffResponse, error) {
	var company model.Company
	if err := s.db.WithContext(ctx).
		Preload("EnterpriseType").
		First(&company, companyID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.New(errs.CodeNotFound, "企业不存在")
		}
		return nil, err
	}
	resp := &YearStuffResponse{CompanyID: company.ID, CompanyName: company.Title}
	if company.EnterpriseType != nil {
		resp.EnterpriseType = company.EnterpriseType.Title
	}
	// Company 没有 BaseArea relation——按 area_id 单独查 allname。
	baseArea := ""
	if company.AreaID != nil {
		var area model.BaseArea
		err := s.db.WithContext(ctx).Where("area_id = ?", *company.AreaID).First(&area).Error
		if err == nil && area.AllName != nil {
			baseArea = *area.AllName
		} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}
	if company.Address != nil && *company.Address != "" {
		if baseArea != "" {
			baseArea = baseArea + "/" + *company.Address
		} else {
			baseArea = *company.Address
		}
	}
	resp.BaseArea = baseArea

	var ps model.ProductionStandardization
	err := s.db.WithContext(ctx).Where("company_id = ?", company.ID).First(&ps).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if err == nil {
		resp.ComplianceTime = ps.GetTime
		resp.RevaluationTime = ps.ValidPeriod
	}

	year := time.Now().Year()
	var ownSafety, ownFire float64
	if company.IsProduce == 1 {
		if err := s.db.WithContext(ctx).Model(&model.SecurityInvestment{}).
			Where("company_id = ?", company.ID).
			Where("year = ?", strconv.Itoa(year)).
			Where("type = ?", "安全投入").
			Select("COALESCE(SUM(budget), 0)").Scan(&ownSafety).Error; err != nil {
			return nil, err
		}
		if err := s.db.WithContext(ctx).Model(&model.SecurityInvestment{}).
			Where("company_id = ?", company.ID).
			Where("year = ?", strconv.Itoa(year)).
			Where("type = ?", "消防投入").
			Select("COALESCE(SUM(budget), 0)").Scan(&ownFire).Error; err != nil {
			return nil, err
		}
	}

	var subSafety, subFire float64
	if company.IsBloc == 1 {
		subIDs, err := s.subordinateIDs(ctx, company.ID)
		if err != nil {
			return nil, err
		}
		if len(subIDs) > 0 {
			if err := s.db.WithContext(ctx).Model(&model.YearStatistic{}).
				Where("company_id IN ?", subIDs).
				Where("EXTRACT(YEAR FROM created_at) = ?", year).
				Select("COALESCE(SUM(NULLIF(safety_funding, '')::numeric), 0)").Scan(&subSafety).Error; err != nil {
				return nil, err
			}
			if err := s.db.WithContext(ctx).Model(&model.YearStatistic{}).
				Where("company_id IN ?", subIDs).
				Where("EXTRACT(YEAR FROM created_at) = ?", year).
				Select("COALESCE(SUM(NULLIF(fire_fighting, '')::numeric), 0)").Scan(&subFire).Error; err != nil {
				return nil, err
			}
		}
	}
	resp.SafetyFunding = int(ownSafety + subSafety)
	resp.FireFighting = int(ownFire + subFire)

	resp.Kids = []YearStuffKid{}
	if company.IsProduce == 1 {
		var slStatuses []model.CertificateStatus
		if err := s.db.WithContext(ctx).
			Where("company_id = ?", company.ID).
			Where("type IN ?", []int16{1, 2, 3}).Find(&slStatuses).Error; err != nil {
			return nil, err
		}
		for _, st := range slStatuses {
			resp.Kids = append(resp.Kids, YearStuffKid{
				Name:                st.Name,
				Type:                1,
				ShouldObtainPeople:  st.ShouldObtainPeople,
				AlreadyObtainPeople: st.AlreadyObtainPeople,
				NoObtainPeople:      st.NoObtainPeople,
				TrainPeople:         st.TrainPeople,
			})
		}
		var xfStatuses []model.CertificateStatus
		if err := s.db.WithContext(ctx).
			Where("company_id = ?", company.ID).
			Where("type IN ?", []int16{4, 5, 6}).Find(&xfStatuses).Error; err != nil {
			return nil, err
		}
		for _, st := range xfStatuses {
			resp.Kids = append(resp.Kids, YearStuffKid{
				Name:                st.Name,
				Type:                2,
				ShouldObtainPeople:  st.ShouldObtainPeople,
				AlreadyObtainPeople: st.AlreadyObtainPeople,
				NoObtainPeople:      st.NoObtainPeople,
				TrainPeople:         st.TrainPeople,
			})
		}
	}
	return resp, nil
}

// YearStatisticsResponse — 年报跨直接下级汇总。sl_data / xf_data 按 name 分组求和。
type YearStatisticsResponse struct {
	SafetyFunding    int                  `json:"safety_funding"`
	FireFighting     int                  `json:"fire_fighting"`
	CountNumber      int                  `json:"count_number"`
	SLData           []YearStatGrouped    `json:"sl_data"`
	XFData           []YearStatGrouped    `json:"xf_data"`
	Date             string               `json:"date"`
	SubordinateCount int                  `json:"subordinate_count"`
}

type YearStatGrouped struct {
	Name                string `json:"name"`
	ShouldObtainPeople  int    `json:"should_obtain_people"`
	AlreadyObtainPeople int    `json:"already_obtain_people"`
	TrainPeople         int    `json:"train_people"`
}

func (s *StatisticsService) YearStatisticStatistics(ctx context.Context, companyID int64, dateTime string) (*YearStatisticsResponse, error) {
	if companyID <= 0 {
		return nil, errs.New(errs.CodeUnprocessable, "company_id 必填")
	}
	year, err := parseYear(dateTime)
	if err != nil {
		return nil, err
	}
	subIDs, err := s.subordinateIDs(ctx, companyID)
	if err != nil {
		return nil, err
	}
	resp := &YearStatisticsResponse{
		Date:             strconv.Itoa(year),
		SubordinateCount: len(subIDs),
		SLData:           []YearStatGrouped{},
		XFData:           []YearStatGrouped{},
	}
	if len(subIDs) == 0 {
		return resp, nil
	}

	var totals struct {
		SafetyFunding float64
		FireFighting  float64
		CountNumber   int
	}
	if err := s.db.WithContext(ctx).Model(&model.YearStatistic{}).
		Where("company_id IN ?", subIDs).
		Where("EXTRACT(YEAR FROM created_at) = ?", year).
		Select(`COALESCE(SUM(NULLIF(safety_funding, '')::numeric), 0) AS safety_funding,
		         COALESCE(SUM(NULLIF(fire_fighting, '')::numeric), 0) AS fire_fighting,
		         COUNT(*) AS count_number`).Scan(&totals).Error; err != nil {
		return nil, err
	}
	resp.SafetyFunding = int(totals.SafetyFunding)
	resp.FireFighting = int(totals.FireFighting)
	resp.CountNumber = totals.CountNumber

	groupQuery := func(kidType int) ([]YearStatGrouped, error) {
		var rows []YearStatGrouped
		if err := s.db.WithContext(ctx).Model(&model.YearStatisticKid{}).
			Where("company_id IN ?", subIDs).
			Where("EXTRACT(YEAR FROM created_at) = ?", year).
			Where("type = ?", kidType).
			Select(`name,
			         COALESCE(SUM(should_obtain_people), 0) AS should_obtain_people,
			         COALESCE(SUM(already_obtain_people), 0) AS already_obtain_people,
			         COALESCE(SUM(train_people), 0) AS train_people`).
			Group("name").Scan(&rows).Error; err != nil {
			return nil, err
		}
		return rows, nil
	}
	if rows, err := groupQuery(1); err != nil {
		return nil, err
	} else if rows != nil {
		resp.SLData = rows
	}
	if rows, err := groupQuery(2); err != nil {
		return nil, err
	} else if rows != nil {
		resp.XFData = rows
	}
	return resp, nil
}

// ════════════════════════════════════════════════════════════════════
//   通知公告反馈统计
// ════════════════════════════════════════════════════════════════════

// NoticeStatsRow — summary 行。office = 任务型（按 feedback_status），new = 通知型（按 read_status）。
type NoticeStatsRow struct {
	CompanyID     int64  `json:"company_id"`
	CompanyName   string `json:"company_name"`
	OfficeTotal   int64  `json:"office_total"`
	OfficeFedBack int64  `json:"office_fed_back"`
	OfficePending int64  `json:"office_pending"`
	NewTotal      int64  `json:"new_total"`
	NewRead       int64  `json:"new_read"`
	NewPending    int64  `json:"new_pending"`
	PendingTotal  int64  `json:"pending_total"`
}

func (s *StatisticsService) NoticeStatsSummary(ctx context.Context, companyID int64) ([]NoticeStatsRow, error) {
	if companyID <= 0 {
		return nil, errs.New(errs.CodeUnprocessable, "company_id 必填")
	}
	var subordinates []model.Company
	if err := s.db.WithContext(ctx).Model(&model.Company{}).
		Select("id, title").
		Where("parent_id = ?", companyID).
		Order("id").Find(&subordinates).Error; err != nil {
		return nil, err
	}
	if len(subordinates) == 0 {
		return []NoticeStatsRow{}, nil
	}
	ids := make([]int64, 0, len(subordinates))
	for _, c := range subordinates {
		ids = append(ids, c.ID)
	}

	type officeKV struct {
		CompanyID int64
		Total     int64
		FedBack   int64
	}
	var officeRows []officeKV
	if err := s.db.WithContext(ctx).Model(&model.OfficeNotifyScope{}).
		Where("company_id IN ?", ids).
		Select(`company_id, COUNT(*) AS total,
		         SUM(CASE WHEN feedback_status = '已反馈' THEN 1 ELSE 0 END) AS fed_back`).
		Group("company_id").Scan(&officeRows).Error; err != nil {
		return nil, err
	}
	officeMap := map[int64]officeKV{}
	for _, r := range officeRows {
		officeMap[r.CompanyID] = r
	}

	type newKV struct {
		CompanyID int64
		Total     int64
		ReadCount int64
	}
	var newRows []newKV
	if err := s.db.WithContext(ctx).Model(&model.NewNotifyScope{}).
		Where("company_id IN ?", ids).
		Select(`company_id, COUNT(*) AS total,
		         SUM(CASE WHEN read_status = '已读' THEN 1 ELSE 0 END) AS read_count`).
		Group("company_id").Scan(&newRows).Error; err != nil {
		return nil, err
	}
	newMap := map[int64]newKV{}
	for _, r := range newRows {
		newMap[r.CompanyID] = r
	}

	out := make([]NoticeStatsRow, 0, len(subordinates))
	for _, c := range subordinates {
		office := officeMap[c.ID]
		nw := newMap[c.ID]
		officePending := office.Total - office.FedBack
		if officePending < 0 {
			officePending = 0
		}
		newPending := nw.Total - nw.ReadCount
		if newPending < 0 {
			newPending = 0
		}
		out = append(out, NoticeStatsRow{
			CompanyID:     c.ID,
			CompanyName:   c.Title,
			OfficeTotal:   office.Total,
			OfficeFedBack: office.FedBack,
			OfficePending: officePending,
			NewTotal:      nw.Total,
			NewRead:       nw.ReadCount,
			NewPending:    newPending,
			PendingTotal:  officePending + newPending,
		})
	}
	return out, nil
}

// NoticePendingResponse — pending 接口返回。office / new 各自一份明细列表。
type NoticePendingResponse struct {
	Office []NoticePendingOffice `json:"office"`
	New    []NoticePendingNew    `json:"new"`
}

type NoticePendingOffice struct {
	ScopeID          int64       `json:"scope_id"`
	NoticeID         int64       `json:"notice_id"`
	NoticeName       *string     `json:"notice_name"`
	Comment          *string     `json:"comment"`
	AlertOthers      *string     `json:"alert_others"`
	FeedbackDeadline *model.Date `json:"feedback_deadline"`
	NoticeCreatedAt  *time.Time  `json:"notice_created_at"`
}

type NoticePendingNew struct {
	ScopeID         int64      `json:"scope_id"`
	NoticeID        int64      `json:"notice_id"`
	NoticeName      *string    `json:"notice_name"`
	Comment         *string    `json:"comment"`
	AlertOthers     *string    `json:"alert_others"`
	NoticeCreatedAt *time.Time `json:"notice_created_at"`
}

func (s *StatisticsService) NoticeStatsPending(ctx context.Context, subordinateID int64) (*NoticePendingResponse, error) {
	if subordinateID <= 0 {
		return nil, errs.New(errs.CodeUnprocessable, "subordinate_id 必填")
	}
	var officeScopes []model.OfficeNotifyScope
	if err := s.db.WithContext(ctx).
		Preload("OfficeNotify").
		Where("company_id = ?", subordinateID).
		Where("(feedback_status IS NULL OR feedback_status <> ?)", "已反馈").
		Order("id DESC").Find(&officeScopes).Error; err != nil {
		return nil, err
	}
	office := make([]NoticePendingOffice, 0, len(officeScopes))
	for _, sc := range officeScopes {
		row := NoticePendingOffice{ScopeID: sc.ID, NoticeID: sc.OfficeNotifyID}
		if sc.OfficeNotify != nil {
			row.NoticeName = sc.OfficeNotify.Name
			row.Comment = sc.OfficeNotify.Comment
			row.AlertOthers = sc.OfficeNotify.AlertOthers
			row.FeedbackDeadline = sc.OfficeNotify.FeedbackTime
			ca := sc.OfficeNotify.CreatedAt
			row.NoticeCreatedAt = &ca
		}
		office = append(office, row)
	}

	var newScopes []model.NewNotifyScope
	if err := s.db.WithContext(ctx).
		Preload("NewNotify").
		Where("company_id = ?", subordinateID).
		Where("(read_status IS NULL OR read_status <> ?)", "已读").
		Order("id DESC").Find(&newScopes).Error; err != nil {
		return nil, err
	}
	newOut := make([]NoticePendingNew, 0, len(newScopes))
	for _, sc := range newScopes {
		row := NoticePendingNew{ScopeID: sc.ID, NoticeID: sc.NewNotifyID}
		if sc.NewNotify != nil {
			row.NoticeName = sc.NewNotify.Name
			row.Comment = sc.NewNotify.Comment
			row.AlertOthers = sc.NewNotify.AlertOthers
			ca := sc.NewNotify.CreatedAt
			row.NoticeCreatedAt = &ca
		}
		newOut = append(newOut, row)
	}
	return &NoticePendingResponse{Office: office, New: newOut}, nil
}

// ════════════════════════════════════════════════════════════════════
//   helpers
// ════════════════════════════════════════════════════════════════════

func (s *StatisticsService) subordinateIDs(ctx context.Context, parentID int64) ([]int64, error) {
	var ids []int64
	if err := s.db.WithContext(ctx).Model(&model.Company{}).
		Where("parent_id = ?", parentID).Pluck("id", &ids).Error; err != nil {
		return nil, err
	}
	return ids, nil
}

func uniqueIDs(in []int64) []int64 {
	seen := make(map[int64]struct{}, len(in))
	out := make([]int64, 0, len(in))
	for _, v := range in {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		out = append(out, v)
	}
	return out
}

// parseYearMonth — 解析 "YYYY-MM"。
func parseYearMonth(s string) (int, int, error) {
	t, err := time.Parse("2006-01", s)
	if err != nil {
		return 0, 0, errs.New(errs.CodeUnprocessable, "date_time 必须是 YYYY-MM")
	}
	return t.Year(), int(t.Month()), nil
}

// parseYear — 解析 "YYYY"。
func parseYear(s string) (int, error) {
	if len(s) != 4 {
		return 0, errs.New(errs.CodeUnprocessable, "date_time 必须是 4 位年份")
	}
	y, err := strconv.Atoi(s)
	if err != nil {
		return 0, errs.New(errs.CodeUnprocessable, "date_time 必须是 4 位年份")
	}
	return y, nil
}
