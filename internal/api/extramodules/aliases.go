package extramodules

import "ai_hs_go/internal/api/apiutil"

var (
	parseDate          = apiutil.ParseDate
	addNullableString  = apiutil.AddNullableString
	scopeCompanyOrFail = apiutil.ScopeCompanyOrFail
	scopeOrCompanyID   = apiutil.ScopeOrCompanyID
)
