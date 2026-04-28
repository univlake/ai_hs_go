// Package riskmonitor 风险监测栏目（集团视图）下 5 个模块共用的 HTTP 入口：
//
//	GET    /bloc-risks-inherent              集团固有风险综合指数
//	GET    /bloc-risks-float                 集团浮动风险综合指数
//	GET    /bloc-risks-inherent-statistics   集团固有风险等级分布（饼图）
//	GET    /bloc-risks-float-statistics      集团浮动风险等级分布（饼图）
//	GET    /bloc-risks-log                   最近 7 日双类型风险曲线
//	GET    /bloc-hidden-danger-sort          下属企业隐患整改完成率排序
//
//	GET    /focus-companies                  重点关注企业列表（按 company_id 过滤）
//	POST   /focus-companies                  新增关注关系
//	DELETE /focus-companies/:id              取消关注
//
//	GET    /inspection-alerts                在线巡检告警计数（过期证书 / 设备 / 未整改隐患）
//	GET    /special-equipment-statistics     按下属企业 × 设备类型汇总特种设备数量
//
// 五个模块在前端 web/src/pages/risk-monitoring/ 下分别为 enterprise-map / hazard-map /
// online-inspection / risk-monitor / special-equipment。前两者纯属现有 /companies、
// /major-hazards 接口的可视化，后三者依赖本包。
package riskmonitor

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/middleware"
	"ai_hs_go/internal/model"
	"ai_hs_go/internal/response"
	"ai_hs_go/internal/service"
)

type Handler struct {
	svc *service.RiskMonitorService
}

func NewHandler(svc *service.RiskMonitorService) *Handler {
	return &Handler{svc: svc}
}

// ─── helpers ───────────────────────────────────────────────────────────

// rootCompanyOrFail 取 query.company_id（集团视图必传）。无则 400。
// 集团视图 5 模块都以集团根企业为聚合源头，跟生产视图通用 ScopeCompanyID 不同，
// 这里直接读 query 而不强制 X-Scope-Company-Id（前端 useScope 把 root 作为参数传入）。
func rootCompanyOrFail(c *gin.Context) int64 {
	v := c.Query("company_id")
	if v == "" {
		response.BadRequest(c, "缺少 company_id")
		return 0
	}
	id, err := strconv.ParseInt(v, 10, 64)
	if err != nil || id <= 0 {
		response.BadRequest(c, "company_id 非法")
		return 0
	}
	return id
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

// ─── 风险监控（综合指数 + 饼图） ─────────────────────────────────────

// InherentRiskNum: GET /bloc-risks-inherent
func (h *Handler) InherentRiskNum(c *gin.Context) {
	rid := rootCompanyOrFail(c)
	if rid == 0 {
		return
	}
	r, err := h.svc.InherentRiskNum(c.Request.Context(), rid)
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Success(c, gin.H{
		"inherent_risk":       r.RiskNum,
		"inherent_risk_grade": r.RiskGrade,
	})
}

// FloatRiskNum: GET /bloc-risks-float
func (h *Handler) FloatRiskNum(c *gin.Context) {
	rid := rootCompanyOrFail(c)
	if rid == 0 {
		return
	}
	r, err := h.svc.FloatRiskNum(c.Request.Context(), rid)
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Success(c, gin.H{
		"float_risk":       r.RiskNum,
		"float_risk_grade": r.RiskGrade,
	})
}

// enterpriseTypeID 解析可选的 enterprise_type_id 过滤参数，缺失或非法时返回 0。
func enterpriseTypeID(c *gin.Context) int64 {
	v := c.Query("enterprise_type_id")
	if v == "" {
		return 0
	}
	id, err := strconv.ParseInt(v, 10, 64)
	if err != nil || id <= 0 {
		return 0
	}
	return id
}

// InherentRiskStatistics: GET /bloc-risks-inherent-statistics
// 返回 [{inherent_risk_grade, num}, ...]
func (h *Handler) InherentRiskStatistics(c *gin.Context) {
	rid := rootCompanyOrFail(c)
	if rid == 0 {
		return
	}
	rows, err := h.svc.InherentRiskStatistics(c.Request.Context(), rid, enterpriseTypeID(c))
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	out := make([]gin.H, len(rows))
	for i, r := range rows {
		out[i] = gin.H{"inherent_risk_grade": r.Grade, "num": r.Num}
	}
	response.Success(c, out)
}

// FloatRiskStatistics: GET /bloc-risks-float-statistics
func (h *Handler) FloatRiskStatistics(c *gin.Context) {
	rid := rootCompanyOrFail(c)
	if rid == 0 {
		return
	}
	rows, err := h.svc.FloatRiskStatistics(c.Request.Context(), rid, enterpriseTypeID(c))
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	out := make([]gin.H, len(rows))
	for i, r := range rows {
		out[i] = gin.H{"float_risk_grade": r.Grade, "num": r.Num}
	}
	response.Success(c, out)
}

// RiskLog: GET /bloc-risks-log
func (h *Handler) RiskLog(c *gin.Context) {
	rid := rootCompanyOrFail(c)
	if rid == 0 {
		return
	}
	rows, err := h.svc.RiskLog7Days(c.Request.Context(), rid)
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Success(c, rows)
}

// HiddenDangerSort: GET /bloc-hidden-danger-sort
func (h *Handler) HiddenDangerSort(c *gin.Context) {
	rid := rootCompanyOrFail(c)
	if rid == 0 {
		return
	}
	rows, err := h.svc.HiddenDangerSort(c.Request.Context(), rid)
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Success(c, rows)
}

// ─── 重点关注企业 ──────────────────────────────────────────────────────

// ListFocusCompanies: GET /focus-companies
func (h *Handler) ListFocusCompanies(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "15"))
	companyID, _ := strconv.ParseInt(c.Query("company_id"), 10, 64)
	focusID, _ := strconv.ParseInt(c.Query("focus_id"), 10, 64)
	rows, total, err := h.svc.ListFocusCompanies(c.Request.Context(), service.FocusCompanyListParams{
		Page:      page,
		PerPage:   per,
		CompanyID: companyID,
		FocusID:   focusID,
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

type focusCompanyReq struct {
	FocusID   *int64 `json:"focus_id"`
	CompanyID int64  `json:"company_id"`
}

// CreateFocusCompany: POST /focus-companies
func (h *Handler) CreateFocusCompany(c *gin.Context) {
	var req focusCompanyReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.CompanyID == 0 || req.FocusID == nil || *req.FocusID == 0 {
		response.Failed(c, http.StatusUnprocessableEntity, "company_id 与 focus_id 不能为空")
		return
	}
	uid := middleware.CurrentUserID(c)
	row := &model.FocusCompany{
		FocusID:   req.FocusID,
		CompanyID: req.CompanyID,
	}
	row.UserID = uid
	row.LastUserID = uid
	if err := h.svc.CreateFocusCompany(c.Request.Context(), row); err != nil {
		writeAppError(c, err)
		return
	}
	response.Created(c, row)
}

// DeleteFocusCompany: DELETE /focus-companies/:id
func (h *Handler) DeleteFocusCompany(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteFocusCompany(c.Request.Context(), id); err != nil {
		writeAppError(c, err)
		return
	}
	response.NoContent(c)
}

// ─── 在线巡检告警 ──────────────────────────────────────────────────────

// InspectionAlerts: GET /inspection-alerts
func (h *Handler) InspectionAlerts(c *gin.Context) {
	rid := rootCompanyOrFail(c)
	if rid == 0 {
		return
	}
	out, err := h.svc.InspectionAlerts(c.Request.Context(), rid)
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Success(c, out)
}

// ─── 特种设备数据统计 ──────────────────────────────────────────────────

// SpecialEquipmentStatistics: GET /special-equipment-statistics
func (h *Handler) SpecialEquipmentStatistics(c *gin.Context) {
	rid := rootCompanyOrFail(c)
	if rid == 0 {
		return
	}
	out, err := h.svc.SpecialEquipmentStatistics(c.Request.Context(), rid)
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Success(c, out)
}
