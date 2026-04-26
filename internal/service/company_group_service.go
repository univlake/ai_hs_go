package service

import (
	"context"
	"errors"
	"strings"

	"gorm.io/gorm"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/model"
)

// CompanyGroupService backs the 集团视图/企业分组 page. Each group is
// scoped to a single owner company (typically a tenant root) and references
// arbitrary descendants via a join table.
type CompanyGroupService struct {
	db *gorm.DB
}

func NewCompanyGroupService(db *gorm.DB) *CompanyGroupService {
	return &CompanyGroupService{db: db}
}

// List returns every group owned by `companyID` together with their member
// rows (preloaded for the right-pane render). Newest first. The
// `kid_companies` field on each row is populated post-fetch so the frontend
// gets a flat list of company objects instead of join rows.
func (s *CompanyGroupService) List(ctx context.Context, companyID int64) ([]model.CompanyGroup, error) {
	if companyID == 0 {
		return nil, errs.New(errs.CodeUnprocessable, "缺少 company_id")
	}
	var rows []model.CompanyGroup
	err := s.db.WithContext(ctx).
		Where("company_id = ?", companyID).
		Preload("Kids").
		Preload("Kids.Company").
		Order("id DESC").
		Find(&rows).Error
	if err != nil {
		return nil, err
	}
	for i := range rows {
		rows[i].KidCompanies = make([]model.Company, 0, len(rows[i].Kids))
		for _, k := range rows[i].Kids {
			if k.Company != nil {
				rows[i].KidCompanies = append(rows[i].KidCompanies, *k.Company)
			}
		}
	}
	return rows, nil
}

type CreateCompanyGroupInput struct {
	Name      string
	CompanyID int64
	UserID    int64
}

func (s *CompanyGroupService) Create(ctx context.Context, in CreateCompanyGroupInput) (*model.CompanyGroup, error) {
	if strings.TrimSpace(in.Name) == "" {
		return nil, errs.New(errs.CodeUnprocessable, "分组名不能为空")
	}
	if in.CompanyID == 0 {
		return nil, errs.New(errs.CodeUnprocessable, "缺少所属企业")
	}
	g := &model.CompanyGroup{
		Name:      in.Name,
		CompanyID: in.CompanyID,
	}
	g.UserID = in.UserID
	g.LastUserID = in.UserID
	if err := s.db.WithContext(ctx).Create(g).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "创建分组失败", err)
	}
	return g, nil
}

func (s *CompanyGroupService) Update(ctx context.Context, id int64, name string, by int64) (*model.CompanyGroup, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errs.New(errs.CodeUnprocessable, "分组名不能为空")
	}
	var g model.CompanyGroup
	if err := s.db.WithContext(ctx).First(&g, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.New(errs.CodeNotFound, "分组不存在")
		}
		return nil, err
	}
	if err := s.db.WithContext(ctx).Model(&g).
		Updates(map[string]any{"name": name, "last_user_id": by}).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "更新分组失败", err)
	}
	return &g, nil
}

func (s *CompanyGroupService) Delete(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Delete(&model.CompanyGroup{}, id).Error
}

// SyncKids replaces the group's member set with the given company ids.
// Atomic: deletes existing rows + reinserts the new set inside a tx so the
// page never observes a half-applied state.
func (s *CompanyGroupService) SyncKids(ctx context.Context, groupID int64, kidCompanyIDs []int64) error {
	var g model.CompanyGroup
	if err := s.db.WithContext(ctx).First(&g, groupID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.New(errs.CodeNotFound, "分组不存在")
		}
		return err
	}
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.
			Where("company_group_id = ?", groupID).
			Delete(&model.CompanyGroupKid{}).Error; err != nil {
			return err
		}
		if len(kidCompanyIDs) == 0 {
			return nil
		}
		// Dedupe in case the frontend posts duplicates.
		seen := make(map[int64]bool, len(kidCompanyIDs))
		rows := make([]model.CompanyGroupKid, 0, len(kidCompanyIDs))
		for _, cid := range kidCompanyIDs {
			if cid == 0 || seen[cid] {
				continue
			}
			seen[cid] = true
			rows = append(rows, model.CompanyGroupKid{
				CompanyGroupID: groupID,
				CompanyID:      cid,
			})
		}
		if len(rows) == 0 {
			return nil
		}
		return tx.Create(&rows).Error
	})
}
