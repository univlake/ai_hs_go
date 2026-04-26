package middleware

// Context keys used across middleware to share request-scoped data.
const (
	CtxKeyUserID           = "auth.user_id"
	CtxKeyUsername         = "auth.username"
	CtxKeyCompanyID        = "auth.company_id"
	CtxKeyRootCompanyID    = "auth.root_company_id"
	CtxKeyIsPlatformAdmin  = "auth.is_platform_admin"
	CtxKeyScopeCompanyID   = "scope.company_id"
	CtxKeyMustChangeCreds  = "auth.must_change_credentials"
)
