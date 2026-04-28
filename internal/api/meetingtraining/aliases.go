package meetingtraining

import "ai_hs_go/internal/api/apiutil"

var (
	parseDate          = apiutil.ParseDate
	parseDateTime      = apiutil.ParseDateTime
	addNullableString  = apiutil.AddNullableString
	firstNonEmpty      = apiutil.FirstNonEmpty
	scopeCompanyOrFail = apiutil.ScopeCompanyOrFail
)
