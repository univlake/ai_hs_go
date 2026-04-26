package service

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/model"
)

// RegulationLibraryTypeService handles the multi-level 法规库类型 dictionary.
// Tree shape lives in the `pid` column (0 = top-level); the `all_name`
// column caches the slash-joined full path, recomputed by Create/Update so
// the list view can display "电气方面 / 防火防爆安全技术" without doing
// recursive lookups in the frontend.
type RegulationLibraryTypeService struct {
	db *gorm.DB
}

func NewRegulationLibraryTypeService(db *gorm.DB) *RegulationLibraryTypeService {
	return &RegulationLibraryTypeService{db: db}
}

type RegLibTypeListParams struct {
	Page    int
	PerPage int
	Title   string
	Pid     *int64 // nil = no filter; 0 = top-level only
}

func (s *RegulationLibraryTypeService) List(ctx context.Context, p RegLibTypeListParams) ([]model.RegulationLibraryType, int64, error) {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PerPage <= 0 || p.PerPage > 500 {
		p.PerPage = 200
	}
	q := s.db.WithContext(ctx).Model(&model.RegulationLibraryType{})
	if p.Pid != nil {
		q = q.Where("pid = ?", *p.Pid)
	}
	if t := strings.TrimSpace(p.Title); t != "" {
		q = q.Where("title ILIKE ? OR all_name ILIKE ?", "%"+t+"%", "%"+t+"%")
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var rows []model.RegulationLibraryType
	if err := q.Order("sort ASC, id ASC").
		Offset((p.Page - 1) * p.PerPage).
		Limit(p.PerPage).
		Find(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

type RegLibTypeInput struct {
	Title  string
	Pid    int64
	Sort   int
	UserID int64
}

func (s *RegulationLibraryTypeService) Create(ctx context.Context, in RegLibTypeInput) (*model.RegulationLibraryType, error) {
	if strings.TrimSpace(in.Title) == "" {
		return nil, errs.New(errs.CodeUnprocessable, "标题不能为空")
	}
	allName, err := s.computeAllName(ctx, in.Pid, in.Title)
	if err != nil {
		return nil, err
	}
	row := &model.RegulationLibraryType{
		Title:      in.Title,
		Pid:        in.Pid,
		AllName:    &allName,
		Sort:       in.Sort,
		UserID:     in.UserID,
		LastUserID: in.UserID,
	}
	if err := s.db.WithContext(ctx).Create(row).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "创建失败", err)
	}
	return row, nil
}

func (s *RegulationLibraryTypeService) Update(ctx context.Context, id int64, in RegLibTypeInput) (*model.RegulationLibraryType, error) {
	var row model.RegulationLibraryType
	if err := s.db.WithContext(ctx).First(&row, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.New(errs.CodeNotFound, "记录不存在")
		}
		return nil, err
	}
	if strings.TrimSpace(in.Title) == "" {
		return nil, errs.New(errs.CodeUnprocessable, "标题不能为空")
	}
	if in.Pid == id {
		return nil, errs.New(errs.CodeUnprocessable, "父级不能是自己")
	}
	if cycle, err := s.wouldCycle(ctx, id, in.Pid); err != nil {
		return nil, err
	} else if cycle {
		return nil, errs.New(errs.CodeUnprocessable, "父级不能是自己的下级")
	}
	allName, err := s.computeAllName(ctx, in.Pid, in.Title)
	if err != nil {
		return nil, err
	}
	updates := map[string]any{
		"title":        in.Title,
		"pid":          in.Pid,
		"all_name":     allName,
		"sort":         in.Sort,
		"last_user_id": in.UserID,
	}
	if err := s.db.WithContext(ctx).Model(&row).Updates(updates).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "更新失败", err)
	}
	// Re-read to surface fresh row to caller.
	if err := s.db.WithContext(ctx).First(&row, id).Error; err != nil {
		return nil, err
	}
	// Cascade all_name refresh into descendants — title or pid may have
	// just shifted, so any node whose all_name starts with the old prefix
	// needs to be rebuilt. Done in a small recursive sweep; the dictionary
	// is small (dozens of rows) so the simple approach is fine.
	if err := s.cascadeAllName(ctx, id); err != nil {
		return nil, err
	}
	return &row, nil
}

func (s *RegulationLibraryTypeService) Delete(ctx context.Context, id int64) error {
	// Block delete when this node still has children — otherwise the
	// orphaned descendants would silently lose their parent and `all_name`
	// would be wrong.
	var children int64
	if err := s.db.WithContext(ctx).Model(&model.RegulationLibraryType{}).
		Where("pid = ?", id).Count(&children).Error; err != nil {
		return err
	}
	if children > 0 {
		return errs.New(errs.CodeUnprocessable, "请先删除该类型下的子项")
	}
	return s.db.WithContext(ctx).Delete(&model.RegulationLibraryType{}, id).Error
}

// computeAllName wraps the title in full-width parentheses appended to
// the parent's all_name — matches the seed catalog's
// `安全生产部门规章（劳动防护用品）` shape so seeded rows and
// user-created descendants stay visually consistent. Top-level (pid=0)
// just returns the title.
func (s *RegulationLibraryTypeService) computeAllName(ctx context.Context, pid int64, title string) (string, error) {
	if pid == 0 {
		return title, nil
	}
	var parent model.RegulationLibraryType
	if err := s.db.WithContext(ctx).First(&parent, pid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errs.New(errs.CodeUnprocessable, fmt.Sprintf("父级 %d 不存在", pid))
		}
		return "", err
	}
	prefix := parent.Title
	if parent.AllName != nil && *parent.AllName != "" {
		prefix = *parent.AllName
	}
	return prefix + "（" + title + "）", nil
}

// wouldCycle returns true when setting candidateParent as the parent of
// candidateChild would form a cycle (i.e. candidateParent is itself a
// descendant of candidateChild).
func (s *RegulationLibraryTypeService) wouldCycle(ctx context.Context, candidateChild, candidateParent int64) (bool, error) {
	if candidateParent == 0 {
		return false, nil
	}
	cur := candidateParent
	for i := 0; i < 32 && cur != 0; i++ { // bounded for safety
		if cur == candidateChild {
			return true, nil
		}
		var node model.RegulationLibraryType
		if err := s.db.WithContext(ctx).First(&node, cur).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return false, nil
			}
			return false, err
		}
		cur = node.Pid
	}
	return false, nil
}

// cascadeAllName rebuilds all_name for every descendant of `rootID`.
func (s *RegulationLibraryTypeService) cascadeAllName(ctx context.Context, rootID int64) error {
	var children []model.RegulationLibraryType
	if err := s.db.WithContext(ctx).Where("pid = ?", rootID).Find(&children).Error; err != nil {
		return err
	}
	for _, child := range children {
		newName, err := s.computeAllName(ctx, child.Pid, child.Title)
		if err != nil {
			return err
		}
		if err := s.db.WithContext(ctx).Model(&model.RegulationLibraryType{}).
			Where("id = ?", child.ID).
			Update("all_name", newName).Error; err != nil {
			return err
		}
		if err := s.cascadeAllName(ctx, child.ID); err != nil {
			return err
		}
	}
	return nil
}
