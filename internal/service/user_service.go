package service

import (
	"context"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/model"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

type UserListParams struct {
	Page      int
	PerPage   int
	CompanyID *int64 // when set, filter by this company (after scope check)
	Keyword   string // matches username/nickname/phone
	Status    *int16 // optional status filter
	Role      string // optional role-name filter
}

func (s *UserService) List(ctx context.Context, p UserListParams) ([]model.User, int64, error) {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PerPage <= 0 || p.PerPage > 100 {
		p.PerPage = 20
	}

	q := s.db.WithContext(ctx).Model(&model.User{})
	if p.CompanyID != nil {
		q = q.Where("company_id = ?", *p.CompanyID)
	}
	if p.Status != nil {
		q = q.Where("status = ?", *p.Status)
	}
	if kw := strings.TrimSpace(p.Keyword); kw != "" {
		like := "%" + kw + "%"
		q = q.Where("username ILIKE ? OR nickname ILIKE ? OR phone ILIKE ?", like, like, like)
	}
	if role := strings.TrimSpace(p.Role); role != "" {
		// Filter by role name through model_has_roles → roles
		q = q.Where(`id IN (
            SELECT mhr.model_id FROM model_has_roles mhr
            JOIN roles r ON r.id = mhr.role_id
            WHERE mhr.model_type = ? AND r.name = ? AND r.guard_name = 'api'
        )`, UserModelType, role)
	}

	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var users []model.User
	if err := q.Order("id DESC").Offset((p.Page - 1) * p.PerPage).Limit(p.PerPage).Find(&users).Error; err != nil {
		return nil, 0, err
	}
	return users, total, nil
}

type CreateUserInput struct {
	Username       string
	Password       string // plaintext
	Nickname       string
	CompanyID      *int64
	Phone          *string
	Email          *string
	Status         int16
	IsExpert       int16
	HeaderImg      *string
	DepartmentName *string
	CreatedBy      int64
}

func (s *UserService) Create(ctx context.Context, in CreateUserInput) (*model.User, error) {
	check := CheckUsername(ctx, s.db, in.Username, 0)
	if !check.Available {
		return nil, errs.WithReason(errs.CodeUnprocessable, check.Reason, string(errs.CodeUsernameInvalid))
	}
	if len(in.Password) < 6 {
		return nil, errs.New(errs.CodeUnprocessable, "密码至少 6 位")
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "密码加密失败", err)
	}
	user := &model.User{
		Username:       in.Username,
		Password:       string(hashed),
		Nickname:       in.Nickname,
		CompanyID:      in.CompanyID,
		Phone:          in.Phone,
		Email:          in.Email,
		Status:         coalesceInt16(in.Status, UserStatusEnable),
		IsExpert:       in.IsExpert,
		HeaderImg:      in.HeaderImg,
		DepartmentName: in.DepartmentName,
	}
	user.UserID = in.CreatedBy
	user.LastUserID = in.CreatedBy
	if err := s.db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "创建用户失败", err)
	}
	return user, nil
}

type UpdateUserInput struct {
	Nickname       *string
	CompanyID      *int64
	Phone          *string
	Email          *string
	Status         *int16
	IsExpert       *int16
	HeaderImg      *string
	DepartmentName *string
	UpdatedBy      int64
}

func (s *UserService) Update(ctx context.Context, id int64, in UpdateUserInput) (*model.User, error) {
	var user model.User
	if err := s.db.WithContext(ctx).First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.New(errs.CodeNotFound, "用户不存在")
		}
		return nil, err
	}
	updates := map[string]any{"last_user_id": in.UpdatedBy}
	if in.Nickname != nil {
		updates["nickname"] = *in.Nickname
	}
	if in.CompanyID != nil {
		updates["company_id"] = *in.CompanyID
	}
	if in.Phone != nil {
		updates["phone"] = in.Phone
	}
	if in.Email != nil {
		updates["email"] = in.Email
	}
	if in.Status != nil {
		updates["status"] = *in.Status
	}
	if in.IsExpert != nil {
		updates["is_expert"] = *in.IsExpert
	}
	if in.HeaderImg != nil {
		updates["header_img"] = in.HeaderImg
	}
	if in.DepartmentName != nil {
		updates["department_name"] = in.DepartmentName
	}
	if err := s.db.WithContext(ctx).Model(&user).Updates(updates).Error; err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "更新失败", err)
	}
	return &user, nil
}

func (s *UserService) Delete(ctx context.Context, id int64) error {
	if err := s.db.WithContext(ctx).Delete(&model.User{}, id).Error; err != nil {
		return errs.Wrap(errs.CodeInternal, "删除失败", err)
	}
	return nil
}

func (s *UserService) Get(ctx context.Context, id int64) (*model.User, error) {
	var u model.User
	if err := s.db.WithContext(ctx).First(&u, id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

// AdminResetPassword lets a privileged user replace a child account's password.
// The user submitting the request is recorded as last_user_id for audit.
func (s *UserService) AdminResetPassword(ctx context.Context, id int64, newPassword string, by int64) error {
	if len(newPassword) < 6 {
		return errs.New(errs.CodeUnprocessable, "新密码至少 6 位")
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return errs.Wrap(errs.CodeInternal, "密码加密失败", err)
	}
	res := s.db.WithContext(ctx).Model(&model.User{}).
		Where("id = ?", id).
		Updates(map[string]any{
			"password":     string(hashed),
			"last_user_id": by,
		})
	if res.Error != nil {
		return errs.Wrap(errs.CodeInternal, "重置密码失败", res.Error)
	}
	if res.RowsAffected == 0 {
		return errs.New(errs.CodeNotFound, "用户不存在")
	}
	return nil
}

func coalesceInt16(v, fallback int16) int16 {
	if v == 0 {
		return fallback
	}
	return v
}
