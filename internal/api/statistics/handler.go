// Package statistics 实现数据统计模块（集团视图 + 生产视图共享）的 HTTP 入口：
//
//	GET    /month-statistics                 月报列表（按 company_id / pid 过滤；pid 走集团视图）
//	GET    /month-statistics/stuff           当月预填充数据（聚合本企业 + 直接下级）
//	GET    /month-statistics/statistics      月报跨直接下级汇总
//	GET    /month-statistics/:id             月报详情含子表
//	POST   /month-statistics                 新增 + 子表整体替换
//	PUT    /month-statistics/:id             部分更新 + 子表整体替换
//	DELETE /month-statistics/:id             软删除
//
//	GET    /year-statistics                  年报列表
//	GET    /year-statistics/stuff            当年预填充数据
//	GET    /year-statistics/statistics       年报跨直接下级汇总
//	GET    /year-statistics/:id              年报详情含 kids
//	POST   /year-statistics                  新增 + kids 整体替换
//	PUT    /year-statistics/:id              部分更新 + kids 整体替换
//	DELETE /year-statistics/:id              软删除
//
//	GET    /notice-stats/summary             下级企业通知公告反馈/阅读情况汇总
//	GET    /notice-stats/pending             指定下级未反馈/未读明细
//
// 集团视图 vs 生产视图共用同一组路由 / 权限（month-statistics.* 等），仅由前端
// 切换 pid 还是 company_id 入参，service 层据此选择过滤集合。
package statistics

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/middleware"
	"ai_hs_go/internal/model"
	"ai_hs_go/internal/response"
	"ai_hs_go/internal/service"
	"ai_hs_go/internal/api/apiutil"
)

type Handler struct {
	svc *service.StatisticsService
}

func NewHandler(svc *service.StatisticsService) *Handler {
	return &Handler{svc: svc}
}

// ════════════════════════════════════════════════════════════════════
//   月度数据上报
// ════════════════════════════════════════════════════════════════════

func (h *Handler) ListMonthStatistic(c *gin.Context) {
	page, per := apiutil.ParsePagination(c, 15)
	cid, _ := strconv.ParseInt(c.Query("company_id"), 10, 64)
	pid, _ := strconv.ParseInt(c.Query("pid"), 10, 64)
	includePid := boolQuery(c, "include_pid")
	rows, total, err := h.svc.ListMonthStatistic(c.Request.Context(), service.MonthStatisticListParams{
		Page: page, PerPage: per, CompanyID: cid, Pid: pid, IncludePid: includePid,
	})
	if err != nil {
		writeAppError(c, err, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) GetMonthStatistic(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	row, err := h.svc.GetMonthStatistic(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err, "月报不存在")
		return
	}
	response.Success(c, row)
}

func (h *Handler) MonthStatisticStuff(c *gin.Context) {
	cid, _ := strconv.ParseInt(c.Query("company_id"), 10, 64)
	if cid <= 0 {
		response.Failed(c, http.StatusUnprocessableEntity, "company_id 必填")
		return
	}
	resp, err := h.svc.MonthStuff(c.Request.Context(), cid)
	if err != nil {
		writeAppError(c, err, "预填数据查询失败")
		return
	}
	response.Success(c, resp)
}

func (h *Handler) MonthStatisticStatistics(c *gin.Context) {
	cid, _ := strconv.ParseInt(c.Query("company_id"), 10, 64)
	dateTime := c.Query("date_time")
	resp, err := h.svc.MonthStatisticStatistics(c.Request.Context(), cid, dateTime)
	if err != nil {
		writeAppError(c, err, "汇总查询失败")
		return
	}
	response.Success(c, resp)
}

type monthControlReq struct {
	GovernTitle    *string `json:"govern_title"`
	GovernSupplies *string `json:"govern_supplies"`
	GovernUser     *string `json:"govern_user"`
	GovernDate     *string `json:"govern_date"`
	GovernText     *string `json:"govern_text"`
	GovernMoney    *string `json:"govern_money"`
}

type monthAccidentReq struct {
	HappenTime         *string `json:"happen_time"`
	Department         *string `json:"department"`
	Name               *string `json:"name"`
	Sex                *string `json:"sex"`
	Age                *int    `json:"age"`
	WorkType           *string `json:"work_type"`
	WorkAge            *int    `json:"work_age"`
	EducationLevel     *string `json:"education_level"`
	DegreeInjury       *string `json:"degree_injury"`
	Type               *string `json:"type"`
	DirectCause        *string `json:"direct_cause"`
	Content            *string `json:"content"`
	DaysOff            *int    `json:"days_off"`
	DirectLossProperty *string `json:"direct_loss_property"`
}

type monthStatisticReq struct {
	AverageEmployee                 *int                 `json:"average_employee"`
	SafetyMeetingCount              *int                 `json:"safety_meeting_count"`
	FireMeetingCount                *int                 `json:"fire_meeting_count"`
	SafetyCheckCount                *int                 `json:"safety_check_count"`
	FireCheckCount                  *int                 `json:"fire_check_count"`
	SafetyFunding                   *int                 `json:"safety_funding"`
	FireFunding                     *int                 `json:"fire_funding"`
	HydratingSeedlings              *int                 `json:"hydrating_seedlings"`
	RetrofitElectrical              *int                 `json:"retrofit_electrical"`
	FireExtinguisher                *int                 `json:"fire_extinguisher"`
	SmartFire                       *int                 `json:"smart_fire"`
	SafetyTrainingCount             *int                 `json:"safety_training_count"`
	SafetyTrainingNumber            *int                 `json:"safety_training_number"`
	NewPeopleTrainingNumber         *int                 `json:"new_people_training_number"`
	FireTrainingCount               *int                 `json:"fire_training_count"`
	FireTrainingNumber              *int                 `json:"fire_training_number"`
	SafetyEmergencyCount            *int                 `json:"safety_emergency_count"`
	SafetyEmergencyNumber           *int                 `json:"safety_emergency_number"`
	FireEmergencyCount              *int                 `json:"fire_emergency_count"`
	FireEmergencyNumber             *int                 `json:"fire_emergency_number"`
	SomeHiddenDanger                *int                 `json:"some_hidden_danger"`
	SomeRectifyHiddenDanger         *int                 `json:"some_rectify_hidden_danger"`
	NoRectifyInfo                   *string              `json:"no_rectify_info"`
	BigHiddenDanger                 *int                 `json:"big_hidden_danger"`
	BigRectifyHiddenDanger          *int                 `json:"big_rectify_hidden_danger"`
	GovernHiddenDangerCount         *int                 `json:"govern_hidden_danger_count"`
	FireFightingDanger              *int                 `json:"fire_fighting_danger"`
	FireFightingHiddenDanger        *int                 `json:"fire_fighting_hidden_danger"`
	FireFightingHiddenDangerRectify *int                 `json:"fire_fighting_hidden_danger_rectify"`
	AccidentBulletinCount           *int                 `json:"accident_bulletin_count"`
	CompanyID                       *int64               `json:"company_id"`
	GovernHiddenDanger              *[]monthControlReq   `json:"govern_hidden_danger"`
	AccidentBulletin                *[]monthAccidentReq  `json:"accident_bulletin"`
}

func (h *Handler) CreateMonthStatistic(c *gin.Context) {
	var req monthStatisticReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.CompanyID == nil || *req.CompanyID == 0 {
		response.Failed(c, http.StatusUnprocessableEntity, "company_id 必填")
		return
	}
	rec := &model.MonthStatistic{CompanyID: *req.CompanyID}
	applyMonthStatisticReq(rec, req)
	uid := middleware.CurrentUserID(c)
	row, err := h.svc.CreateMonthStatistic(c.Request.Context(), service.MonthStatisticInput{
		Record:    rec,
		Controls:  mapMonthControls(req.GovernHiddenDanger),
		Accidents: mapMonthAccidents(req.AccidentBulletin),
		UserID:    uid,
	})
	if err != nil {
		writeAppError(c, err, "创建失败")
		return
	}
	response.Created(c, row)
}

func (h *Handler) UpdateMonthStatistic(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	var req monthStatisticReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	rec := &model.MonthStatistic{}
	if req.CompanyID != nil {
		rec.CompanyID = *req.CompanyID
	}
	applyMonthStatisticReq(rec, req)
	uid := middleware.CurrentUserID(c)
	row, err := h.svc.UpdateMonthStatistic(c.Request.Context(), id, service.MonthStatisticInput{
		Record:    rec,
		Controls:  mapMonthControls(req.GovernHiddenDanger),
		Accidents: mapMonthAccidents(req.AccidentBulletin),
		UserID:    uid,
	})
	if err != nil {
		writeAppError(c, err, "更新失败")
		return
	}
	response.Success(c, row)
}

func (h *Handler) DeleteMonthStatistic(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	if err := h.svc.DeleteMonthStatistic(c.Request.Context(), id); err != nil {
		writeAppError(c, err, "删除失败")
		return
	}
	response.Success(c, gin.H{"message": "删除成功"})
}

// ════════════════════════════════════════════════════════════════════
//   年度数据上报
// ════════════════════════════════════════════════════════════════════

func (h *Handler) ListYearStatistic(c *gin.Context) {
	page, per := apiutil.ParsePagination(c, 15)
	cid, _ := strconv.ParseInt(c.Query("company_id"), 10, 64)
	pid, _ := strconv.ParseInt(c.Query("pid"), 10, 64)
	includePid := boolQuery(c, "include_pid")
	rows, total, err := h.svc.ListYearStatistic(c.Request.Context(), service.YearStatisticListParams{
		Page: page, PerPage: per, CompanyID: cid, Pid: pid, IncludePid: includePid,
	})
	if err != nil {
		writeAppError(c, err, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) GetYearStatistic(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	row, err := h.svc.GetYearStatistic(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err, "年度统计不存在")
		return
	}
	response.Success(c, row)
}

func (h *Handler) YearStatisticStuff(c *gin.Context) {
	cid, _ := strconv.ParseInt(c.Query("company_id"), 10, 64)
	if cid <= 0 {
		response.Failed(c, http.StatusUnprocessableEntity, "company_id 必填")
		return
	}
	resp, err := h.svc.YearStuff(c.Request.Context(), cid)
	if err != nil {
		writeAppError(c, err, "预填数据查询失败")
		return
	}
	response.Success(c, resp)
}

func (h *Handler) YearStatisticStatistics(c *gin.Context) {
	cid, _ := strconv.ParseInt(c.Query("company_id"), 10, 64)
	dateTime := c.Query("date_time")
	resp, err := h.svc.YearStatisticStatistics(c.Request.Context(), cid, dateTime)
	if err != nil {
		writeAppError(c, err, "汇总查询失败")
		return
	}
	response.Success(c, resp)
}

type yearKidReq struct {
	Name                *string `json:"name"`
	Type                *int    `json:"type"`
	ShouldObtainPeople  *int    `json:"should_obtain_people"`
	AlreadyObtainPeople *int    `json:"already_obtain_people"`
	NoObtainPeople      *string `json:"no_obtain_people"`
	TrainPeople         *int    `json:"train_people"`
}

type yearStatisticReq struct {
	CompanyName     *string       `json:"company_name"`
	BaseArea        *string       `json:"base_area"`
	EnterpriseType  *string       `json:"enterprise_type"`
	SafetyFunding   *string       `json:"safety_funding"`
	FireFighting    *string       `json:"fire_fighting"`
	ComplianceTime  *string       `json:"compliance_time"`
	RevaluationTime *string       `json:"revaluation_time"`
	CompanyID       *int64        `json:"company_id"`
	Kids            *[]yearKidReq `json:"kids"`
}

func (h *Handler) CreateYearStatistic(c *gin.Context) {
	var req yearStatisticReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.CompanyID == nil || *req.CompanyID == 0 {
		response.Failed(c, http.StatusUnprocessableEntity, "company_id 必填")
		return
	}
	rec := &model.YearStatistic{CompanyID: *req.CompanyID}
	applyYearStatisticReq(rec, req)
	uid := middleware.CurrentUserID(c)
	row, err := h.svc.CreateYearStatistic(c.Request.Context(), service.YearStatisticInput{
		Record: rec,
		Kids:   mapYearKids(req.Kids),
		UserID: uid,
	})
	if err != nil {
		writeAppError(c, err, "创建失败")
		return
	}
	response.Created(c, row)
}

func (h *Handler) UpdateYearStatistic(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	var req yearStatisticReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	rec := &model.YearStatistic{}
	if req.CompanyID != nil {
		rec.CompanyID = *req.CompanyID
	}
	applyYearStatisticReq(rec, req)
	uid := middleware.CurrentUserID(c)
	row, err := h.svc.UpdateYearStatistic(c.Request.Context(), id, service.YearStatisticInput{
		Record: rec,
		Kids:   mapYearKids(req.Kids),
		UserID: uid,
	})
	if err != nil {
		writeAppError(c, err, "更新失败")
		return
	}
	response.Success(c, row)
}

func (h *Handler) DeleteYearStatistic(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	if err := h.svc.DeleteYearStatistic(c.Request.Context(), id); err != nil {
		writeAppError(c, err, "删除失败")
		return
	}
	response.Success(c, gin.H{"message": "删除成功"})
}

// ════════════════════════════════════════════════════════════════════
//   通知公告反馈统计
// ════════════════════════════════════════════════════════════════════

func (h *Handler) NoticeStatsSummary(c *gin.Context) {
	cid, _ := strconv.ParseInt(c.Query("company_id"), 10, 64)
	rows, err := h.svc.NoticeStatsSummary(c.Request.Context(), cid)
	if err != nil {
		writeAppError(c, err, "查询失败")
		return
	}
	response.Success(c, rows)
}

func (h *Handler) NoticeStatsPending(c *gin.Context) {
	sid, _ := strconv.ParseInt(c.Query("subordinate_id"), 10, 64)
	resp, err := h.svc.NoticeStatsPending(c.Request.Context(), sid)
	if err != nil {
		writeAppError(c, err, "查询失败")
		return
	}
	response.Success(c, resp)
}

// ════════════════════════════════════════════════════════════════════
//   helpers
// ════════════════════════════════════════════════════════════════════

func parseID(c *gin.Context) (int64, bool) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		response.BadRequest(c, "无效 id")
		return 0, false
	}
	return id, true
}

func boolQuery(c *gin.Context, key string) bool {
	v := c.Query(key)
	switch v {
	case "1", "true", "True", "TRUE", "yes":
		return true
	}
	return false
}

func writeAppError(c *gin.Context, err error, fallback string) {
	if ae, ok := errs.As(err); ok {
		switch ae.Code {
		case errs.CodeNotFound:
			response.NotFound(c, ae.Message)
		case errs.CodeUnprocessable:
			response.Failed(c, http.StatusUnprocessableEntity, ae.Message)
		case errs.CodeForbidden:
			response.Forbidden(c, ae.Message)
		default:
			response.Internal(c, ae.Message)
		}
		return
	}
	response.Internal(c, fallback)
}

func applyMonthStatisticReq(rec *model.MonthStatistic, req monthStatisticReq) {
	if req.AverageEmployee != nil {
		rec.AverageEmployee = *req.AverageEmployee
	}
	if req.SafetyMeetingCount != nil {
		rec.SafetyMeetingCount = *req.SafetyMeetingCount
	}
	if req.FireMeetingCount != nil {
		rec.FireMeetingCount = *req.FireMeetingCount
	}
	if req.SafetyCheckCount != nil {
		rec.SafetyCheckCount = *req.SafetyCheckCount
	}
	if req.FireCheckCount != nil {
		rec.FireCheckCount = *req.FireCheckCount
	}
	if req.SafetyFunding != nil {
		rec.SafetyFunding = *req.SafetyFunding
	}
	if req.FireFunding != nil {
		rec.FireFunding = *req.FireFunding
	}
	if req.HydratingSeedlings != nil {
		rec.HydratingSeedlings = *req.HydratingSeedlings
	}
	if req.RetrofitElectrical != nil {
		rec.RetrofitElectrical = *req.RetrofitElectrical
	}
	if req.FireExtinguisher != nil {
		rec.FireExtinguisher = *req.FireExtinguisher
	}
	if req.SmartFire != nil {
		rec.SmartFire = *req.SmartFire
	}
	if req.SafetyTrainingCount != nil {
		rec.SafetyTrainingCount = *req.SafetyTrainingCount
	}
	if req.SafetyTrainingNumber != nil {
		rec.SafetyTrainingNumber = *req.SafetyTrainingNumber
	}
	if req.NewPeopleTrainingNumber != nil {
		rec.NewPeopleTrainingNumber = *req.NewPeopleTrainingNumber
	}
	if req.FireTrainingCount != nil {
		rec.FireTrainingCount = *req.FireTrainingCount
	}
	if req.FireTrainingNumber != nil {
		rec.FireTrainingNumber = *req.FireTrainingNumber
	}
	if req.SafetyEmergencyCount != nil {
		rec.SafetyEmergencyCount = *req.SafetyEmergencyCount
	}
	if req.SafetyEmergencyNumber != nil {
		rec.SafetyEmergencyNumber = *req.SafetyEmergencyNumber
	}
	if req.FireEmergencyCount != nil {
		rec.FireEmergencyCount = *req.FireEmergencyCount
	}
	if req.FireEmergencyNumber != nil {
		rec.FireEmergencyNumber = *req.FireEmergencyNumber
	}
	if req.SomeHiddenDanger != nil {
		rec.SomeHiddenDanger = *req.SomeHiddenDanger
	}
	if req.SomeRectifyHiddenDanger != nil {
		rec.SomeRectifyHiddenDanger = *req.SomeRectifyHiddenDanger
	}
	if req.NoRectifyInfo != nil {
		rec.NoRectifyInfo = req.NoRectifyInfo
	}
	if req.BigHiddenDanger != nil {
		rec.BigHiddenDanger = *req.BigHiddenDanger
	}
	if req.BigRectifyHiddenDanger != nil {
		rec.BigRectifyHiddenDanger = *req.BigRectifyHiddenDanger
	}
	if req.GovernHiddenDangerCount != nil {
		rec.GovernHiddenDangerCount = *req.GovernHiddenDangerCount
	}
	if req.FireFightingDanger != nil {
		rec.FireFightingDanger = *req.FireFightingDanger
	}
	if req.FireFightingHiddenDanger != nil {
		rec.FireFightingHiddenDanger = *req.FireFightingHiddenDanger
	}
	if req.FireFightingHiddenDangerRectify != nil {
		rec.FireFightingHiddenDangerRectify = *req.FireFightingHiddenDangerRectify
	}
	if req.AccidentBulletinCount != nil {
		rec.AccidentBulletinCount = *req.AccidentBulletinCount
	}
}

func applyYearStatisticReq(rec *model.YearStatistic, req yearStatisticReq) {
	if req.CompanyName != nil {
		rec.CompanyName = req.CompanyName
	}
	if req.BaseArea != nil {
		rec.BaseArea = req.BaseArea
	}
	if req.EnterpriseType != nil {
		rec.EnterpriseType = req.EnterpriseType
	}
	if req.SafetyFunding != nil {
		rec.SafetyFunding = req.SafetyFunding
	}
	if req.FireFighting != nil {
		rec.FireFighting = req.FireFighting
	}
	if req.ComplianceTime != nil {
		rec.ComplianceTime = parseDateTime(req.ComplianceTime)
	}
	if req.RevaluationTime != nil {
		rec.RevaluationTime = parseDateTime(req.RevaluationTime)
	}
}

func mapMonthControls(p *[]monthControlReq) *[]model.MonthStatisticControl {
	if p == nil {
		return nil
	}
	out := make([]model.MonthStatisticControl, 0, len(*p))
	for _, v := range *p {
		out = append(out, model.MonthStatisticControl{
			GovernTitle:    v.GovernTitle,
			GovernSupplies: v.GovernSupplies,
			GovernUser:     v.GovernUser,
			GovernDate:     v.GovernDate,
			GovernText:     v.GovernText,
			GovernMoney:    v.GovernMoney,
		})
	}
	return &out
}

func mapMonthAccidents(p *[]monthAccidentReq) *[]model.MonthStatisticAccident {
	if p == nil {
		return nil
	}
	out := make([]model.MonthStatisticAccident, 0, len(*p))
	for _, v := range *p {
		out = append(out, model.MonthStatisticAccident{
			HappenTime:         parseDate(v.HappenTime),
			Department:         v.Department,
			Name:               v.Name,
			Sex:                v.Sex,
			Age:                v.Age,
			WorkType:           v.WorkType,
			WorkAge:            v.WorkAge,
			EducationLevel:     v.EducationLevel,
			DegreeInjury:       v.DegreeInjury,
			Type:               v.Type,
			DirectCause:        v.DirectCause,
			Content:            v.Content,
			DaysOff:            v.DaysOff,
			DirectLossProperty: v.DirectLossProperty,
		})
	}
	return &out
}

func mapYearKids(p *[]yearKidReq) *[]model.YearStatisticKid {
	if p == nil {
		return nil
	}
	out := make([]model.YearStatisticKid, 0, len(*p))
	for _, v := range *p {
		k := model.YearStatisticKid{
			Name:                v.Name,
			Type:                1,
			ShouldObtainPeople:  v.ShouldObtainPeople,
			AlreadyObtainPeople: v.AlreadyObtainPeople,
			NoObtainPeople:      v.NoObtainPeople,
			TrainPeople:         v.TrainPeople,
		}
		if v.Type != nil {
			k.Type = *v.Type
		}
		out = append(out, k)
	}
	return &out
}
