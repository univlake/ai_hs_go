package healthfire

import "ai_hs_go/internal/api/apiutil"

var (
	parseDate          = apiutil.ParseDate
	addNullableString  = apiutil.AddNullableString
	firstNonEmpty      = apiutil.FirstNonEmpty
	scopeCompanyOrFail = apiutil.ScopeCompanyOrFail
	scopeOrCompanyID   = apiutil.ScopeOrCompanyID
)
