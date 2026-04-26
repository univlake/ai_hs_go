package rbac

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

// NewEnforcer wires Casbin to the same GORM connection pool used by the rest
// of the app. Policy data lives in the casbin_rule table created by the 0002
// migration.
//
// IMPORTANT: gorm-adapter v3 calls `db.Scopes(...)` internally, which mutates
// the chained Statement of the passed-in `*gorm.DB`. To avoid leaking that
// scope onto subsequent app queries (causing "relation X does not exist" when
// every later query is silently rewritten to use the casbin_rule table), give
// the adapter its own Session-isolated handle.
func NewEnforcer(db *gorm.DB, modelPath string) (*casbin.Enforcer, error) {
	isolated := db.Session(&gorm.Session{NewDB: true})
	adapter, err := gormadapter.NewAdapterByDBUseTableName(isolated, "", "casbin_rule")
	if err != nil {
		return nil, fmt.Errorf("casbin adapter: %w", err)
	}
	enforcer, err := casbin.NewEnforcer(modelPath, adapter)
	if err != nil {
		return nil, fmt.Errorf("casbin enforcer: %w", err)
	}
	if err := enforcer.LoadPolicy(); err != nil {
		return nil, fmt.Errorf("load policy: %w", err)
	}
	return enforcer, nil
}

// SubjectForUser is the conventional subject string for a user (`u:<id>`).
// Casbin role inheritance uses `r:<role-name>` for groupings.
func SubjectForUser(userID int64) string {
	return fmt.Sprintf("u:%d", userID)
}

func SubjectForRole(role string) string {
	return "r:" + role
}
