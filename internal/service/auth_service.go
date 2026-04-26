package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"ai_hs_go/internal/auth"
	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/model"
)

const (
	UserStatusEnable    int16 = 1
	UserStatusDisable   int16 = 2
	CompanyStatusEnable int16 = 1
)

type AuthService struct {
	db  *gorm.DB
	jwt *auth.JWTService
	rsa *auth.RSAService
}

func NewAuthService(db *gorm.DB, jwtSvc *auth.JWTService, rsaSvc *auth.RSAService) *AuthService {
	return &AuthService{db: db, jwt: jwtSvc, rsa: rsaSvc}
}

type LoginInput struct {
	Username           string
	EncryptedPassword  string
}

type LoginOutput struct {
	AccessToken string    `json:"access_token"`
	TokenType   string    `json:"token_type"`
	ExpiresAt   time.Time `json:"expires_at"`
	User        *model.User
}

func (s *AuthService) Login(ctx context.Context, in LoginInput) (*LoginOutput, error) {
	plain, err := s.rsa.Decrypt(in.EncryptedPassword)
	if err != nil {
		return nil, errs.New(errs.CodeUnprocessable, "密码格式无效")
	}

	var user model.User
	if err := s.db.WithContext(ctx).Where("username = ?", in.Username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.New(errs.CodeUnauthorized, "用户名或密码错误")
		}
		return nil, errs.Wrap(errs.CodeInternal, "查询用户失败", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(plain)); err != nil {
		return nil, errs.New(errs.CodeUnauthorized, "用户名或密码错误")
	}

	if user.Status != UserStatusEnable {
		return nil, errs.New(errs.CodeForbidden, "该账号已被禁用")
	}

	if !user.IsPlatformAdmin && user.CompanyID != nil {
		var company model.Company
		if err := s.db.WithContext(ctx).First(&company, *user.CompanyID).Error; err == nil {
			if company.Status != CompanyStatusEnable {
				return nil, errs.WithReason(errs.CodeForbidden, "该企业已被停用，请联系管理员", string(errs.CodeCompanyDisabled))
			}
		}
	}

	res, err := s.jwt.Issue(auth.IssueParams{
		UserID:                user.ID,
		Username:              user.Username,
		CompanyID:             user.CompanyID,
		IsPlatformAdmin:       user.IsPlatformAdmin,
		MustChangeCredentials: user.MustChangeCredentials,
	})
	if err != nil {
		return nil, errs.Wrap(errs.CodeInternal, "签发 token 失败", err)
	}

	// best-effort login_times bump
	go func(id int64) {
		_ = s.db.Model(&model.User{}).Where("id = ?", id).
			Updates(map[string]any{"login_times": time.Now().Format(time.RFC3339)}).Error
	}(user.ID)

	return &LoginOutput{
		AccessToken: res.Token,
		TokenType:   "Bearer",
		ExpiresAt:   res.ExpiresAt,
		User:        &user,
	}, nil
}

type ProfileUpdateInput struct {
	Nickname  string
	Phone     *string
	Email     *string
	HeaderImg *string
	Username  *string
}

func (s *AuthService) UpdateProfile(ctx context.Context, userID int64, in ProfileUpdateInput) error {
	var user model.User
	if err := s.db.WithContext(ctx).First(&user, userID).Error; err != nil {
		return errs.New(errs.CodeNotFound, "用户不存在")
	}

	updates := map[string]any{
		"nickname":   in.Nickname,
		"phone":      in.Phone,
		"email":      in.Email,
		"header_img": in.HeaderImg,
	}

	if in.Username != nil && *in.Username != user.Username {
		check := CheckUsername(ctx, s.db, *in.Username, user.ID)
		if !check.Available {
			return errs.WithReason(errs.CodeUnprocessable, check.Reason, string(errs.CodeUsernameInvalid))
		}
		updates["username"] = strings.TrimSpace(*in.Username)
	}

	if err := s.db.WithContext(ctx).Model(&user).Updates(updates).Error; err != nil {
		return errs.Wrap(errs.CodeInternal, "更新失败", err)
	}
	return nil
}

type FirstLoginInput struct {
	NewUsername string
	NewPassword string
}

func (s *AuthService) FirstLoginSetup(ctx context.Context, userID int64, in FirstLoginInput) error {
	var user model.User
	if err := s.db.WithContext(ctx).First(&user, userID).Error; err != nil {
		return errs.New(errs.CodeNotFound, "用户不存在")
	}
	if !user.MustChangeCredentials {
		return errs.New(errs.CodeForbidden, "无需首次登录设置")
	}
	if in.NewUsername == user.Username {
		return errs.New(errs.CodeUnprocessable, "新用户名不能与当前用户名相同")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.NewPassword)); err == nil {
		return errs.New(errs.CodeUnprocessable, "新密码不能与当前密码相同")
	}

	check := CheckUsername(ctx, s.db, in.NewUsername, user.ID)
	if !check.Available {
		return errs.WithReason(errs.CodeUnprocessable, check.Reason, string(errs.CodeUsernameInvalid))
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(in.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return errs.Wrap(errs.CodeInternal, "密码加密失败", err)
	}

	updates := map[string]any{
		"username":                 in.NewUsername,
		"password":                 string(hashed),
		"must_change_credentials":  false,
	}
	if err := s.db.WithContext(ctx).Model(&user).Updates(updates).Error; err != nil {
		return errs.Wrap(errs.CodeInternal, "更新失败", err)
	}
	return nil
}

type ResetPasswordInput struct {
	OldPassword string
	NewPassword string
}

func (s *AuthService) ResetPassword(ctx context.Context, userID int64, in ResetPasswordInput) error {
	var user model.User
	if err := s.db.WithContext(ctx).First(&user, userID).Error; err != nil {
		return errs.New(errs.CodeNotFound, "用户不存在")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.OldPassword)); err != nil {
		return errs.New(errs.CodeUnprocessable, "原密码错误")
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(in.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return errs.Wrap(errs.CodeInternal, "密码加密失败", err)
	}
	return s.db.WithContext(ctx).Model(&user).Update("password", string(hashed)).Error
}

func (s *AuthService) GetUser(ctx context.Context, userID int64) (*model.User, error) {
	var user model.User
	if err := s.db.WithContext(ctx).First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetCompanyByID returns the company a user belongs to. Used by /user-info
// so the frontend can bootstrap ScopeContext without a second round-trip.
// Returns (nil, nil) when not found rather than wrapping ErrRecordNotFound —
// the caller treats absence as "not bound".
func (s *AuthService) GetCompanyByID(ctx context.Context, id int64) (*model.Company, error) {
	var co model.Company
	if err := s.db.WithContext(ctx).First(&co, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &co, nil
}

func (s *AuthService) RSAPublicKey() string {
	return s.rsa.PublicKeyPEM()
}

func (s *AuthService) CheckUsername(ctx context.Context, username string, excludeUserID int64) UsernameCheck {
	return CheckUsername(ctx, s.db, username, excludeUserID)
}
