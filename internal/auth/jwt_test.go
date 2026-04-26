package auth

import (
	"testing"
	"time"
)

func TestJWTService_IssueAndParse(t *testing.T) {
	svc := NewJWTService("test-secret-test-secret-test", time.Minute, "ai_hs_go")
	cid := int64(7)
	res, err := svc.Issue(IssueParams{
		UserID:          42,
		Username:        "alice",
		CompanyID:       &cid,
		IsPlatformAdmin: false,
	})
	if err != nil {
		t.Fatalf("issue: %v", err)
	}
	if res.Token == "" {
		t.Fatal("empty token")
	}
	claims, err := svc.Parse(res.Token)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if claims.UserID != 42 || claims.Username != "alice" || claims.CompanyID == nil || *claims.CompanyID != 7 {
		t.Fatalf("unexpected claims: %+v", claims)
	}

	if _, err := svc.Parse("not-a-token"); err == nil {
		t.Fatal("expected parse error")
	}

	expired := NewJWTService("s", -time.Minute, "x")
	res2, _ := expired.Issue(IssueParams{UserID: 1, Username: "a"})
	if _, err := expired.Parse(res2.Token); err == nil {
		t.Fatal("expected expired token to fail parse")
	}
}
