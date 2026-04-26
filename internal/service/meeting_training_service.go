package service

import (
	"context"
	"strings"

	"gorm.io/gorm"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/model"
)

// MeetingTrainingService 聚合会议培训四个模块的 CRUD：
//   - internal_meetings (+ internal_meeting_kids) 内部会议（带参会人员子表）
//   - internal_trainings  内部培训
//   - external_trainings  外部培训
//   - other_trainings     相关方培训
type MeetingTrainingService struct {
	db *gorm.DB
}

func NewMeetingTrainingService(db *gorm.DB) *MeetingTrainingService {
	return &MeetingTrainingService{db: db}
}

// ── 内部会议（含子表） ──────────────────────────────────────────────────

type InternalMeetingListParams struct {
	Page      int
	PerPage   int
	CompanyID int64
	Name      string
}

func (s *MeetingTrainingService) ListInternalMeeting(ctx context.Context, p InternalMeetingListParams) ([]model.InternalMeeting, int64, error) {
	page, per := normalizePage(p.Page, p.PerPage)
	q := s.db.WithContext(ctx).Model(&model.InternalMeeting{}).Where("company_id = ?", p.CompanyID)
	if kw := strings.TrimSpace(p.Name); kw != "" {
		q = q.Where("name ILIKE ?", "%"+kw+"%")
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.InternalMeeting
	if err := q.Preload("Kids.JoinUser").Order("id DESC").Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *MeetingTrainingService) GetInternalMeeting(ctx context.Context, id int64) (*model.InternalMeeting, error) {
	var m model.InternalMeeting
	if err := s.db.WithContext(ctx).Preload("Kids.JoinUser").First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

type CreateInternalMeetingInput struct {
	Meeting *model.InternalMeeting
	Kids    []model.InternalMeetingKid
}

func (s *MeetingTrainingService) CreateInternalMeeting(ctx context.Context, in CreateInternalMeetingInput) error {
	if in.Meeting.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(in.Meeting).Error; err != nil {
			return wrapWrite(err)
		}
		for i := range in.Kids {
			in.Kids[i].InternalMeetingID = in.Meeting.ID
			in.Kids[i].CompanyID = in.Meeting.CompanyID
		}
		if len(in.Kids) > 0 {
			if err := tx.Create(&in.Kids).Error; err != nil {
				return wrapWrite(err)
			}
		}
		return nil
	})
}

func (s *MeetingTrainingService) UpdateInternalMeeting(ctx context.Context, id int64, updates map[string]any, kids []model.InternalMeetingKid) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		res := tx.Model(&model.InternalMeeting{}).Where("id = ?", id).Updates(updates)
		if res.Error != nil {
			return wrapWrite(res.Error)
		}
		if res.RowsAffected == 0 {
			return errs.New(errs.CodeNotFound, "记录不存在")
		}
		if kids != nil {
			if err := tx.Where("internal_meeting_id = ?", id).Delete(&model.InternalMeetingKid{}).Error; err != nil {
				return err
			}
			if len(kids) > 0 {
				var meeting model.InternalMeeting
				if err := tx.First(&meeting, id).Error; err != nil {
					return err
				}
				for i := range kids {
					kids[i].InternalMeetingID = id
					kids[i].CompanyID = meeting.CompanyID
				}
				if err := tx.Create(&kids).Error; err != nil {
					return wrapWrite(err)
				}
			}
		}
		return nil
	})
}

func (s *MeetingTrainingService) DeleteInternalMeeting(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("internal_meeting_id = ?", id).Delete(&model.InternalMeetingKid{}).Error; err != nil {
			return err
		}
		return tx.Delete(&model.InternalMeeting{}, id).Error
	})
}

// ── 内部培训 / 外部培训 / 相关方培训 ────────────────────────────────────
//
// 三类培训表结构相近、均无子表，写入语义和 design_units 完全一致。

type TrainingListParams struct {
	Page      int
	PerPage   int
	CompanyID int64
	Name      string
}

func (s *MeetingTrainingService) ListInternalTraining(ctx context.Context, p TrainingListParams) ([]model.InternalTraining, int64, error) {
	page, per := normalizePage(p.Page, p.PerPage)
	q := s.db.WithContext(ctx).Model(&model.InternalTraining{}).Where("company_id = ?", p.CompanyID)
	if kw := strings.TrimSpace(p.Name); kw != "" {
		q = q.Where("name ILIKE ?", "%"+kw+"%")
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.InternalTraining
	if err := q.Order("id DESC").Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *MeetingTrainingService) GetInternalTraining(ctx context.Context, id int64) (*model.InternalTraining, error) {
	var m model.InternalTraining
	if err := s.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

func (s *MeetingTrainingService) CreateInternalTraining(ctx context.Context, m *model.InternalTraining) error {
	if m.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return wrapWrite(s.db.WithContext(ctx).Create(m).Error)
}

func (s *MeetingTrainingService) UpdateInternalTraining(ctx context.Context, id int64, updates map[string]any) error {
	return updateRow(s.db.WithContext(ctx), &model.InternalTraining{}, id, updates)
}

func (s *MeetingTrainingService) DeleteInternalTraining(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.InternalTraining{}, id).Error
}

func (s *MeetingTrainingService) ListExternalTraining(ctx context.Context, p TrainingListParams) ([]model.ExternalTraining, int64, error) {
	page, per := normalizePage(p.Page, p.PerPage)
	q := s.db.WithContext(ctx).Model(&model.ExternalTraining{}).Where("company_id = ?", p.CompanyID)
	if kw := strings.TrimSpace(p.Name); kw != "" {
		q = q.Where("name ILIKE ?", "%"+kw+"%")
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.ExternalTraining
	if err := q.Order("id DESC").Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *MeetingTrainingService) GetExternalTraining(ctx context.Context, id int64) (*model.ExternalTraining, error) {
	var m model.ExternalTraining
	if err := s.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

func (s *MeetingTrainingService) CreateExternalTraining(ctx context.Context, m *model.ExternalTraining) error {
	if m.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return wrapWrite(s.db.WithContext(ctx).Create(m).Error)
}

func (s *MeetingTrainingService) UpdateExternalTraining(ctx context.Context, id int64, updates map[string]any) error {
	return updateRow(s.db.WithContext(ctx), &model.ExternalTraining{}, id, updates)
}

func (s *MeetingTrainingService) DeleteExternalTraining(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.ExternalTraining{}, id).Error
}

func (s *MeetingTrainingService) ListOtherTraining(ctx context.Context, p TrainingListParams) ([]model.OtherTraining, int64, error) {
	page, per := normalizePage(p.Page, p.PerPage)
	q := s.db.WithContext(ctx).Model(&model.OtherTraining{}).Where("company_id = ?", p.CompanyID)
	if kw := strings.TrimSpace(p.Name); kw != "" {
		q = q.Where("name ILIKE ?", "%"+kw+"%")
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.OtherTraining
	if err := q.Order("id DESC").Offset((page - 1) * per).Limit(per).Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func (s *MeetingTrainingService) GetOtherTraining(ctx context.Context, id int64) (*model.OtherTraining, error) {
	var m model.OtherTraining
	if err := s.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return nil, mapNotFound(err)
	}
	return &m, nil
}

func (s *MeetingTrainingService) CreateOtherTraining(ctx context.Context, m *model.OtherTraining) error {
	if m.CompanyID == 0 {
		return errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	return wrapWrite(s.db.WithContext(ctx).Create(m).Error)
}

func (s *MeetingTrainingService) UpdateOtherTraining(ctx context.Context, id int64, updates map[string]any) error {
	return updateRow(s.db.WithContext(ctx), &model.OtherTraining{}, id, updates)
}

func (s *MeetingTrainingService) DeleteOtherTraining(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.OtherTraining{}, id).Error
}
