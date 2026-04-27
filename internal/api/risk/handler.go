// Package risk 实现风险管理 5 个模块的 CRUD：
//
//	/risk-distribution-maps    — 风险源管理（per-company singleton）
//	/monthly-investigations    — 隐患管理（含 reviews/views 子表整体替换）
//	/job-management            — 作业管理
//	/three-togethers           — 三同时管理
//	/assess-detects            — 安全评价与检测
//
// 全部模块共用一份 RiskService。日期字段在请求体里走 *string，转 *model.Date 后写入。
package risk

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/middleware"
	"ai_hs_go/internal/model"
	"ai_hs_go/internal/response"
	"ai_hs_go/internal/service"
	"ai_hs_go/internal/storage"
)

// 子表替换语义：在 monthlyInvestigationReq 里 Reviews/Views 都是 *[]investigationAnnexReq —
// nil 指针表示请求体未携带该 key（不动子表），非 nil（含空切片）表示按本次值整体覆盖。
// Laravel 原写法 `$reviews = $data['monthlyInvestigationReviews'] ?? null` 同语义。

type Handler struct {
	svc *service.RiskService
	// OSS 是可选 dep — 仅 risk-distribution-maps GET 路径用，nil 时返回原始 key。
	oss *storage.Client
}

func NewHandler(svc *service.RiskService, oss *storage.Client) *Handler {
	return &Handler{svc: svc, oss: oss}
}

// ── 风险源管理（singleton per company） ─────────────────────────────────

type riskDistributionMapReq struct {
	AnnexURL  *string `json:"annex_url"`
	CompanyID int64   `json:"company_id"`
}

func (h *Handler) GetRiskDistributionMap(c *gin.Context) {
	cid := scopeOrCompanyID(c)
	if cid == 0 {
		response.BadRequest(c, "缺少 company_id")
		return
	}
	row, err := h.svc.FirstRiskDistributionMap(c.Request.Context(), cid)
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	if row != nil && h.oss != nil && row.AnnexURL != nil && *row.AnnexURL != "" {
		if signed, err := h.oss.SignURL(*row.AnnexURL, time.Hour, nil); err == nil {
			row.SignedAnnexURL = &signed
		}
	}
	response.Success(c, row)
}

func (h *Handler) CreateRiskDistributionMap(c *gin.Context) {
	var req riskDistributionMapReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.CompanyID == 0 {
		req.CompanyID = scopeOrCompanyID(c)
	}
	uid := middleware.CurrentUserID(c)
	m := &model.RiskDistributionMap{
		AnnexURL:  req.AnnexURL,
		CompanyID: req.CompanyID,
	}
	m.UserID = uid
	m.LastUserID = uid
	if err := h.svc.CreateRiskDistributionMap(c.Request.Context(), m); err != nil {
		writeAppError(c, err)
		return
	}
	response.Created(c, m)
}

func (h *Handler) UpdateRiskDistributionMap(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req riskDistributionMapReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	updates := map[string]any{"last_user_id": middleware.CurrentUserID(c)}
	addNullableString(updates, "annex_url", req.AnnexURL)
	if err := h.svc.UpdateRiskDistributionMap(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, gin.H{"id": id})
}

// ── 隐患管理 ───────────────────────────────────────────────────────────

type investigationAnnexReq struct {
	AnnexURL *string `json:"annex_url"`
}

type monthlyInvestigationReq struct {
	Type              *string `json:"type"`
	IsFireFighting    *int16  `json:"is_fire_fighting"`
	IsEmergency       *int16  `json:"is_emergency"`
	Content           *string `json:"content"`
	InvTime           *string `json:"inv_time"`
	InvName           *string `json:"inv_name"`
	Level             *string `json:"level"`
	Proposal          *string `json:"proposal"`
	Legalbasis        *string `json:"legalbasis"`
	CorrectDate       *string `json:"correct_date"`
	ReviewDiscription *string `json:"review_discription"`
	ReviewDate        *string `json:"review_date"`
	ReviewContent     *string `json:"review_content"`
	ReviewUser        *string `json:"review_user"`
	CorrectResult     *string `json:"correct_result"`
	IsGovern          *int16  `json:"is_govern"`
	GovernTitle       *string `json:"govern_title"`
	GovernSupplies    *string `json:"govern_supplies"`
	GovernUser        *string `json:"govern_user"`
	GovernDate        *string `json:"govern_date"`
	GovernText        *string `json:"govern_text"`
	GovernMoney       *string `json:"govern_money"`
	// 子表整体替换：nil 指针 = 请求体未带该 key（不动子表），非 nil（含空切片）= 按本次值整体覆盖。
	Reviews *[]investigationAnnexReq `json:"monthlyInvestigationReviews"`
	Views   *[]investigationAnnexReq `json:"monthlyInvestigationViews"`
}

func (h *Handler) ListMonthlyInvestigation(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "20"))
	rows, total, err := h.svc.ListMonthlyInvestigation(c.Request.Context(), service.MonthlyInvestigationListParams{
		Page: page, PerPage: per, CompanyID: companyID,
		Type: c.Query("type"),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) GetMonthlyInvestigation(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.svc.GetMonthlyInvestigation(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

func (h *Handler) CreateMonthlyInvestigation(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	var req monthlyInvestigationReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	inv := &model.MonthlyInvestigation{
		Type:              req.Type,
		IsFireFighting:    req.IsFireFighting,
		IsEmergency:       req.IsEmergency,
		Content:           req.Content,
		InvTime:           parseDate(req.InvTime),
		InvName:           req.InvName,
		Level:             req.Level,
		Proposal:          req.Proposal,
		Legalbasis:        req.Legalbasis,
		CorrectDate:       parseDate(req.CorrectDate),
		ReviewDiscription: req.ReviewDiscription,
		ReviewDate:        parseDate(req.ReviewDate),
		ReviewContent:     req.ReviewContent,
		ReviewUser:        req.ReviewUser,
		CorrectResult:     req.CorrectResult,
		IsGovern:          req.IsGovern,
		GovernTitle:       req.GovernTitle,
		GovernSupplies:    req.GovernSupplies,
		GovernUser:        req.GovernUser,
		GovernDate:        req.GovernDate,
		GovernText:        req.GovernText,
		GovernMoney:       req.GovernMoney,
		CompanyID:         companyID,
	}
	inv.UserID = uid
	inv.LastUserID = uid
	if err := h.svc.CreateMonthlyInvestigation(c.Request.Context(), service.MonthlyInvestigationInput{
		Investigation: inv,
		Reviews:       mapInvestigationAnnexReviews(derefAnnex(req.Reviews)),
		Views:         mapInvestigationAnnexViews(derefAnnex(req.Views)),
	}); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetMonthlyInvestigation(c.Request.Context(), inv.ID)
	response.Created(c, row)
}

func (h *Handler) UpdateMonthlyInvestigation(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req monthlyInvestigationReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	updates := map[string]any{"last_user_id": uid}
	addNullableString(updates, "type", req.Type)
	addNullableString(updates, "content", req.Content)
	addNullableString(updates, "inv_name", req.InvName)
	addNullableString(updates, "level", req.Level)
	addNullableString(updates, "proposal", req.Proposal)
	addNullableString(updates, "legalbasis", req.Legalbasis)
	addNullableString(updates, "review_discription", req.ReviewDiscription)
	addNullableString(updates, "review_content", req.ReviewContent)
	addNullableString(updates, "review_user", req.ReviewUser)
	addNullableString(updates, "correct_result", req.CorrectResult)
	addNullableString(updates, "govern_title", req.GovernTitle)
	addNullableString(updates, "govern_supplies", req.GovernSupplies)
	addNullableString(updates, "govern_user", req.GovernUser)
	addNullableString(updates, "govern_date", req.GovernDate)
	addNullableString(updates, "govern_text", req.GovernText)
	addNullableString(updates, "govern_money", req.GovernMoney)
	if req.IsFireFighting != nil {
		updates["is_fire_fighting"] = req.IsFireFighting
	}
	if req.IsEmergency != nil {
		updates["is_emergency"] = req.IsEmergency
	}
	if req.IsGovern != nil {
		updates["is_govern"] = req.IsGovern
	}
	if req.InvTime != nil {
		updates["inv_time"] = parseDate(req.InvTime)
	}
	if req.CorrectDate != nil {
		updates["correct_date"] = parseDate(req.CorrectDate)
	}
	if req.ReviewDate != nil {
		updates["review_date"] = parseDate(req.ReviewDate)
	}
	// 子表替换：req.Reviews / req.Views 是否非 nil 决定是否替换 — *[] 解码后
	// nil 指针表示 JSON 体里没出现该 key（不动子表），空切片表示清空。
	if err := h.svc.UpdateMonthlyInvestigation(
		c.Request.Context(), id, updates,
		mapInvestigationAnnexReviews(derefAnnex(req.Reviews)),
		mapInvestigationAnnexViews(derefAnnex(req.Views)),
		req.Reviews != nil, req.Views != nil, uid,
	); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetMonthlyInvestigation(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteMonthlyInvestigation(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteMonthlyInvestigation(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

func mapInvestigationAnnexReviews(items []investigationAnnexReq) []model.MonthlyInvestigationReview {
	out := make([]model.MonthlyInvestigationReview, len(items))
	for i, it := range items {
		out[i] = model.MonthlyInvestigationReview{AnnexURL: it.AnnexURL}
	}
	return out
}

func mapInvestigationAnnexViews(items []investigationAnnexReq) []model.MonthlyInvestigationView {
	out := make([]model.MonthlyInvestigationView, len(items))
	for i, it := range items {
		out[i] = model.MonthlyInvestigationView{AnnexURL: it.AnnexURL}
	}
	return out
}

func derefAnnex(p *[]investigationAnnexReq) []investigationAnnexReq {
	if p == nil {
		return nil
	}
	return *p
}

// ── 作业管理 ───────────────────────────────────────────────────────────

type jobManagementReq struct {
	Department       *string    `json:"department"`
	Content          *string    `json:"content"`
	Type             *string    `json:"type"`
	SpecialType      *string    `json:"special_type"`
	Position         *string    `json:"position"`
	JobTime          *string    `json:"job_time"`
	Unit             *string    `json:"unit"`
	IsOutsourcing    *int16     `json:"is_outsourcing"`
	EvaluationUnitID *flexInt64 `json:"evaluation_unit_id"`
	JobPeopleNumber  *flexInt32 `json:"job_people_number"`
	JobPeopleName    *string    `json:"job_people_name"`
	RiskAnalysis     *string    `json:"risk_analysis"`
	RiskProtection   *string    `json:"risk_protection"`
}

func (h *Handler) ListJobManagement(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "20"))
	rows, total, err := h.svc.ListJobManagement(c.Request.Context(), service.JobManagementListParams{
		Page: page, PerPage: per, CompanyID: companyID,
		Type:          c.Query("type"),
		IsOutsourcing: c.Query("is_outsourcing"),
		JobTime:       c.Query("job_time"),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) GetJobManagement(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.svc.GetJobManagement(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

func (h *Handler) CreateJobManagement(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	var req jobManagementReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	m := &model.JobManagement{
		Department:       req.Department,
		Content:          req.Content,
		Type:             req.Type,
		SpecialType:      req.SpecialType,
		Position:         req.Position,
		JobTime:          parseDate(req.JobTime),
		Unit:             req.Unit,
		IsOutsourcing:    req.IsOutsourcing,
		EvaluationUnitID: req.EvaluationUnitID.Int64Ptr(),
		JobPeopleNumber:  req.JobPeopleNumber.Int32Ptr(),
		JobPeopleName:    req.JobPeopleName,
		RiskAnalysis:     req.RiskAnalysis,
		RiskProtection:   req.RiskProtection,
		CompanyID:        companyID,
	}
	m.UserID = uid
	m.LastUserID = uid
	if err := h.svc.CreateJobManagement(c.Request.Context(), m); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetJobManagement(c.Request.Context(), m.ID)
	response.Created(c, row)
}

func (h *Handler) UpdateJobManagement(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req jobManagementReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	updates := map[string]any{"last_user_id": middleware.CurrentUserID(c)}
	addNullableString(updates, "department", req.Department)
	addNullableString(updates, "content", req.Content)
	addNullableString(updates, "type", req.Type)
	addNullableString(updates, "special_type", req.SpecialType)
	addNullableString(updates, "position", req.Position)
	addNullableString(updates, "unit", req.Unit)
	addNullableString(updates, "job_people_name", req.JobPeopleName)
	addNullableString(updates, "risk_analysis", req.RiskAnalysis)
	addNullableString(updates, "risk_protection", req.RiskProtection)
	if req.IsOutsourcing != nil {
		updates["is_outsourcing"] = req.IsOutsourcing
	}
	if req.EvaluationUnitID != nil {
		updates["evaluation_unit_id"] = req.EvaluationUnitID.Int64Ptr()
	}
	if req.JobPeopleNumber != nil {
		updates["job_people_number"] = req.JobPeopleNumber.Int32Ptr()
	}
	if req.JobTime != nil {
		updates["job_time"] = parseDate(req.JobTime)
	}
	if err := h.svc.UpdateJobManagement(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetJobManagement(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteJobManagement(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteJobManagement(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

// ── 三同时管理 ─────────────────────────────────────────────────────────

type threeTogetherReq struct {
	Name                *string    `json:"name"`
	ThreeTogetherType   *string    `json:"three_together_type"`
	EstablishTime       *string    `json:"establish_time"`
	EstablishSection    *string    `json:"establish_section"`
	Status              *string    `json:"status"`
	DesignUnitID        *flexInt64 `json:"design_unit_id"`
	DesignAnnexURL      *string    `json:"design_annex_url"`
	EvaluationUnitID    *flexInt64 `json:"evaluation_unit_id"`
	WriteTime           *string    `json:"write_time"`
	ConstructionComment *string    `json:"construction_comment"`
	SafetyRequirements  *string    `json:"safety_requirements"`
	CompleteComment     *string    `json:"complete_comment"`
	Acceptance          *string    `json:"acceptance"`
}

func (h *Handler) ListThreeTogether(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "20"))
	rows, total, err := h.svc.ListThreeTogether(c.Request.Context(), service.ThreeTogetherListParams{
		Page: page, PerPage: per, CompanyID: companyID,
		Name: firstNonEmpty(c.Query("name"), c.Query("keyword")),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) GetThreeTogether(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.svc.GetThreeTogether(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

func (h *Handler) CreateThreeTogether(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	var req threeTogetherReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	m := &model.ThreeTogether{
		Name:                req.Name,
		ThreeTogetherType:   req.ThreeTogetherType,
		EstablishTime:       parseDate(req.EstablishTime),
		EstablishSection:    req.EstablishSection,
		Status:              req.Status,
		DesignUnitID:        req.DesignUnitID.Int64Ptr(),
		DesignAnnexURL:      req.DesignAnnexURL,
		EvaluationUnitID:    req.EvaluationUnitID.Int64Ptr(),
		WriteTime:           parseDate(req.WriteTime),
		ConstructionComment: req.ConstructionComment,
		SafetyRequirements:  req.SafetyRequirements,
		CompleteComment:     req.CompleteComment,
		Acceptance:          req.Acceptance,
		CompanyID:           companyID,
	}
	m.UserID = uid
	m.LastUserID = uid
	if err := h.svc.CreateThreeTogether(c.Request.Context(), m); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetThreeTogether(c.Request.Context(), m.ID)
	response.Created(c, row)
}

func (h *Handler) UpdateThreeTogether(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req threeTogetherReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	updates := map[string]any{"last_user_id": middleware.CurrentUserID(c)}
	addNullableString(updates, "name", req.Name)
	addNullableString(updates, "three_together_type", req.ThreeTogetherType)
	addNullableString(updates, "establish_section", req.EstablishSection)
	addNullableString(updates, "status", req.Status)
	addNullableString(updates, "design_annex_url", req.DesignAnnexURL)
	addNullableString(updates, "construction_comment", req.ConstructionComment)
	addNullableString(updates, "safety_requirements", req.SafetyRequirements)
	addNullableString(updates, "complete_comment", req.CompleteComment)
	addNullableString(updates, "acceptance", req.Acceptance)
	if req.DesignUnitID != nil {
		updates["design_unit_id"] = req.DesignUnitID.Int64Ptr()
	}
	if req.EvaluationUnitID != nil {
		updates["evaluation_unit_id"] = req.EvaluationUnitID.Int64Ptr()
	}
	if req.EstablishTime != nil {
		updates["establish_time"] = parseDate(req.EstablishTime)
	}
	if req.WriteTime != nil {
		updates["write_time"] = parseDate(req.WriteTime)
	}
	if err := h.svc.UpdateThreeTogether(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetThreeTogether(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteThreeTogether(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteThreeTogether(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

// ── 安全评价与检测 ─────────────────────────────────────────────────────

type assessDetectReq struct {
	Name             *string    `json:"name"`
	AssessDetectType *string    `json:"assess_detect_type"`
	SkillUnitID      *flexInt64 `json:"skill_unit_id"`
	SkillUnitKidID   *flexInt64 `json:"skill_unit_kid_id"`
	ValidPeriod      *string    `json:"valid_period"`
	AnnexURL         *string    `json:"annex_url"`
}

func (h *Handler) ListAssessDetect(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "20"))
	rows, total, err := h.svc.ListAssessDetect(c.Request.Context(), service.AssessDetectListParams{
		Page: page, PerPage: per, CompanyID: companyID,
		Name: firstNonEmpty(c.Query("name"), c.Query("keyword")),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) GetAssessDetect(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.svc.GetAssessDetect(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

func (h *Handler) CreateAssessDetect(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	var req assessDetectReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	m := &model.AssessDetect{
		Name:             req.Name,
		AssessDetectType: req.AssessDetectType,
		SkillUnitID:      req.SkillUnitID.Int64Ptr(),
		SkillUnitKidID:   req.SkillUnitKidID.Int64Ptr(),
		ValidPeriod:      parseDate(req.ValidPeriod),
		AnnexURL:         req.AnnexURL,
		CompanyID:        companyID,
	}
	m.UserID = uid
	m.LastUserID = uid
	if err := h.svc.CreateAssessDetect(c.Request.Context(), m); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetAssessDetect(c.Request.Context(), m.ID)
	response.Created(c, row)
}

func (h *Handler) UpdateAssessDetect(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req assessDetectReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	updates := map[string]any{"last_user_id": middleware.CurrentUserID(c)}
	addNullableString(updates, "name", req.Name)
	addNullableString(updates, "assess_detect_type", req.AssessDetectType)
	addNullableString(updates, "annex_url", req.AnnexURL)
	if req.SkillUnitID != nil {
		updates["skill_unit_id"] = req.SkillUnitID.Int64Ptr()
	}
	if req.SkillUnitKidID != nil {
		updates["skill_unit_kid_id"] = req.SkillUnitKidID.Int64Ptr()
	}
	if req.ValidPeriod != nil {
		updates["valid_period"] = parseDate(req.ValidPeriod)
	}
	if err := h.svc.UpdateAssessDetect(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetAssessDetect(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteAssessDetect(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteAssessDetect(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

// ── helpers ──────────────────────────────────────────────────────────

func parseDate(p *string) *model.Date {
	if p == nil || *p == "" {
		return nil
	}
	for _, f := range []string{"2006-01-02", "2006/01/02", "2006/1/2", time.RFC3339} {
		if t, err := time.Parse(f, *p); err == nil {
			d := model.Date(t)
			return &d
		}
	}
	return nil
}

func addNullableString(m map[string]any, key string, v *string) {
	if v != nil {
		m[key] = v
	}
}

func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}

func scopeCompanyOrFail(c *gin.Context) int64 {
	if cid, ok := middleware.ScopeCompanyID(c); ok {
		return cid
	}
	response.BadRequest(c, "缺少 company_id")
	return 0
}

func scopeOrCompanyID(c *gin.Context) int64 {
	if v := c.Query("company_id"); v != "" {
		if id, err := strconv.ParseInt(v, 10, 64); err == nil {
			return id
		}
	}
	if cid, ok := middleware.ScopeCompanyID(c); ok {
		return cid
	}
	return 0
}

// flexInt32 / flexInt64 — 容忍前端 <select>/<input type=number> 把数字以字符串
// 形式提交的场景。空字符串 / null 当作"未传值"，对应 *int32 / *int64 留 nil。
type flexInt32 int32

func (f *flexInt32) UnmarshalJSON(data []byte) error {
	s := string(data)
	if s == "null" || s == `""` {
		return nil
	}
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		s = s[1 : len(s)-1]
	}
	if s == "" {
		return nil
	}
	v, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return fmt.Errorf("invalid int32: %q", s)
	}
	*f = flexInt32(v)
	return nil
}

func (f *flexInt32) Int32Ptr() *int32 {
	if f == nil {
		return nil
	}
	v := int32(*f)
	return &v
}

type flexInt64 int64

func (f *flexInt64) UnmarshalJSON(data []byte) error {
	s := string(data)
	if s == "null" || s == `""` {
		return nil
	}
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		s = s[1 : len(s)-1]
	}
	if s == "" {
		return nil
	}
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid int64: %q", s)
	}
	*f = flexInt64(v)
	return nil
}

func (f *flexInt64) Int64Ptr() *int64 {
	if f == nil {
		return nil
	}
	v := int64(*f)
	return &v
}

func writeAppError(c *gin.Context, err error) {
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
	response.Internal(c, err.Error())
}
