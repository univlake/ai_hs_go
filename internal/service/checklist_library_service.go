package service

import (
	"context"
	"errors"
	"strings"

	"gorm.io/gorm"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/model"
)

// ChecklistLibraryService backs the 检查表库 CRUD. Records are tenant-
// scoped via company_id, with optional preview + actual annex URL pairs
// for the file uploaded to OSS.
type ChecklistLibraryService struct {
	db *gorm.DB
}

func NewChecklistLibraryService(db *gorm.DB) *ChecklistLibraryService {
	return &ChecklistLibraryService{db: db}
}

type ChecklistLibraryListParams struct {
	Page      int
	PerPage   int
	CompanyID int64
	Title     string
}

func (s *ChecklistLibraryService) List(ctx context.Context, p ChecklistLibraryListParams) ([]model.ChecklistLibrary, int64, error) {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PerPage <= 0 || p.PerPage > 200 {
		p.PerPage = 20
	}
	q := s.db.WithContext(ctx).Model(&model.ChecklistLibrary{})
	if p.CompanyID > 0 {
		q = q.Where("company_id = ?", p.CompanyID)
	}
	if t := strings.TrimSpace(p.Title); t != "" {
		q = q.Where("title ILIKE ?", "%"+t+"%")
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.ChecklistLibrary
	if err := q.Order("id DESC").
		Offset((p.Page - 1) * p.PerPage).
		Limit(p.PerPage).
		Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

type ChecklistLibraryInput struct {
	Title           string
	Status          *int16
	PreviewAnnexURL *string
	AnnexURL        *string
	CompanyID       int64
	UserID          int64
}

func (s *ChecklistLibraryService) Create(ctx context.Context, in ChecklistLibraryInput) (*model.ChecklistLibrary, error) {
	if strings.TrimSpace(in.Title) == "" {
		return nil, errs.New(errs.CodeUnprocessable, "标题不能为空")
	}
	if in.CompanyID == 0 {
		return nil, errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	row := &model.ChecklistLibrary{
		Title:           in.Title,
		Status:          1,
		PreviewAnnexURL: in.PreviewAnnexURL,
		AnnexURL:        in.AnnexURL,
		CompanyID:       in.CompanyID,
	}
	if in.Status != nil {
		row.Status = *in.Status
	}
	row.UserID = in.UserID
	row.LastUserID = in.UserID
	if err := s.db.WithContext(ctx).Create(row).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "创建失败", err)
	}
	return row, nil
}

func (s *ChecklistLibraryService) Update(ctx context.Context, id int64, in ChecklistLibraryInput) (*model.ChecklistLibrary, error) {
	var row model.ChecklistLibrary
	if err := s.db.WithContext(ctx).First(&row, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.New(errs.CodeNotFound, "记录不存在")
		}
		return nil, err
	}
	updates := map[string]any{"last_user_id": in.UserID}
	if t := strings.TrimSpace(in.Title); t != "" {
		updates["title"] = t
	}
	if in.Status != nil {
		updates["status"] = *in.Status
	}
	if in.PreviewAnnexURL != nil {
		updates["preview_annex_url"] = in.PreviewAnnexURL
	}
	if in.AnnexURL != nil {
		updates["annex_url"] = in.AnnexURL
	}
	if err := s.db.WithContext(ctx).Model(&row).Updates(updates).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "更新失败", err)
	}
	return &row, nil
}

func (s *ChecklistLibraryService) Delete(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.ChecklistLibrary{}, id).Error
}
