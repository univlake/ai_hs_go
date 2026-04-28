package risk

import "ai_hs_go/internal/api/apiutil"

// Local aliases for the helpers that used to be defined privately in this
// package. Centralised in internal/api/apiutil so per-module copies can no
// longer drift apart.
var (
	parseDate          = apiutil.ParseDate
	addNullableString  = apiutil.AddNullableString
	firstNonEmpty      = apiutil.FirstNonEmpty
	scopeCompanyOrFail = apiutil.ScopeCompanyOrFail
	scopeOrCompanyID   = apiutil.ScopeOrCompanyID
)
