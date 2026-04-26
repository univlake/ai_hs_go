package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID                int64  `json:"sub_id"`
	Username              string `json:"username"`
	CompanyID             *int64 `json:"company_id,omitempty"`
	IsPlatformAdmin       bool   `json:"is_platform_admin"`
	MustChangeCredentials bool   `json:"mcc,omitempty"`
	jwt.RegisteredClaims
}

type JWTService struct {
	secret []byte
	ttl    time.Duration
	issuer string
}

func NewJWTService(secret string, ttl time.Duration, issuer string) *JWTService {
	return &JWTService{secret: []byte(secret), ttl: ttl, issuer: issuer}
}

type IssueParams struct {
	UserID                int64
	Username              string
	CompanyID             *int64
	IsPlatformAdmin       bool
	MustChangeCredentials bool
}

type IssueResult struct {
	Token     string
	ExpiresAt time.Time
}

func (s *JWTService) Issue(p IssueParams) (IssueResult, error) {
	now := time.Now()
	exp := now.Add(s.ttl)
	claims := Claims{
		UserID:                p.UserID,
		Username:              p.Username,
		CompanyID:             p.CompanyID,
		IsPlatformAdmin:       p.IsPlatformAdmin,
		MustChangeCredentials: p.MustChangeCredentials,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.issuer,
			Subject:   fmt.Sprintf("%d", p.UserID),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(s.secret)
	if err != nil {
		return IssueResult{}, fmt.Errorf("sign jwt: %w", err)
	}
	return IssueResult{Token: signed, ExpiresAt: exp}, nil
}

func (s *JWTService) Parse(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return s.secret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
