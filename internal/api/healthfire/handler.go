// Package healthfire 实现职业健康（5 模块 + 1 小模块）+ 消防安全（3 模块）的 CRUD：
//
//	/occupational-hazards    — 职业危害因素清单（含 import/export/template）
//	/health-systems          — 职业健康体系（singleton）
//	/health-files            — 健康档案（含 import/export/template）
//	/health-reports          — 职业病报告
//	/health-checkups         — 职业健康体检
//	/hygiene-evaluations     — 职业卫生评价（关联 skill_units / skill_unit_kids）
//	/fire-safeties           — 消防安全（关联 skill_units / skill_unit_kids）
//	/fire-equipment          — 消防器材（含 copy 批量复制）
//	/fire-facilities         — 消防设施（含 copy 批量复制）
//
// 全部模块共用一份 HealthFireService。日期字段在请求体里走 *string，转 *model.Date 后写入。
package healthfire

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/excel"
	"ai_hs_go/internal/middleware"
	"ai_hs_go/internal/model"
	"ai_hs_go/internal/response"
	"ai_hs_go/internal/service"
	"ai_hs_go/internal/api/apiutil"
)

type Handler struct {
	svc *service.HealthFireService
}

func NewHandler(svc *service.HealthFireService) *Handler {
	return &Handler{svc: svc}
}

// ── 职业危害因素清单 ──────────────────────────────────────────────────

type occupationalHazardReq struct {
	Name        *string    `json:"name"`
	Type        *string    `json:"type"`
	Number      *flexInt32 `json:"number"`
	IsDeclare   *string    `json:"is_declare"`
	DeclareDate *string    `json:"declare_date"`
}

func (h *Handler) ListOccupationalHazard(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	page, per := apiutil.ParsePagination(c, 20)
	rows, total, err := h.svc.ListOccupationalHazard(c.Request.Context(), service.OccupationalHazardListParams{
		Page: page, PerPage: per, CompanyID: companyID,
		Name:      firstNonEmpty(c.Query("name"), c.Query("keyword")),
		IsDeclare: c.Query("is_declare"),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) GetOccupationalHazard(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.svc.GetOccupationalHazard(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

func (h *Handler) CreateOccupationalHazard(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	var req occupationalHazardReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	m := &model.OccupationalHazard{
		Name:        req.Name,
		Type:        req.Type,
		Number:      req.Number.Int32Ptr(),
		IsDeclare:   req.IsDeclare,
		DeclareDate: parseDate(req.DeclareDate),
		CompanyID:   companyID,
	}
	m.UserID = uid
	m.LastUserID = uid
	if err := h.svc.CreateOccupationalHazard(c.Request.Context(), m); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetOccupationalHazard(c.Request.Context(), m.ID)
	response.Created(c, row)
}

func (h *Handler) UpdateOccupationalHazard(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req occupationalHazardReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	updates := map[string]any{"last_user_id": middleware.CurrentUserID(c)}
	addNullableString(updates, "name", req.Name)
	addNullableString(updates, "type", req.Type)
	addNullableString(updates, "is_declare", req.IsDeclare)
	if req.Number != nil {
		updates["number"] = req.Number.Int32Ptr()
	}
	if req.DeclareDate != nil {
		updates["declare_date"] = parseDate(req.DeclareDate)
	}
	if err := h.svc.UpdateOccupationalHazard(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetOccupationalHazard(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteOccupationalHazard(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteOccupationalHazard(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

func (h *Handler) TemplateOccupationalHazard(c *gin.Context) {
	setXlsxAttachment(c, "职业病危害因素导入模板.xlsx", "occupational_hazards_template.xlsx")
	if err := excel.OccupationalHazardTemplate(c.Writer); err != nil {
		response.Internal(c, "生成模板失败")
		return
	}
}

func (h *Handler) ExportOccupationalHazard(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	rows, _, err := h.svc.ListOccupationalHazard(c.Request.Context(), service.OccupationalHazardListParams{
		Page: 1, PerPage: 10000, CompanyID: companyID,
		Name: c.Query("name"),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	setXlsxAttachment(c, "职业病危害因素列表.xlsx", "occupational_hazards.xlsx")
	if err := excel.OccupationalHazardExport(c.Writer, rows); err != nil {
		response.Internal(c, "导出失败")
		return
	}
}

func (h *Handler) ImportOccupationalHazard(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		response.BadRequest(c, "缺少 file 字段")
		return
	}
	defer file.Close()
	rows, err := excel.OccupationalHazardImport(file)
	if err != nil {
		response.BadRequest(c, "解析 Excel 失败: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	imported := 0
	for i := range rows {
		rows[i].CompanyID = companyID
		rows[i].UserID = uid
		rows[i].LastUserID = uid
		if err := h.svc.CreateOccupationalHazard(c.Request.Context(), &rows[i]); err == nil {
			imported++
		}
	}
	response.Success(c, gin.H{"imported": imported, "total": len(rows)})
}

// ── 职业健康体系（singleton per company） ─────────────────────────────

type healthSystemReq struct {
	CertificationTime *string `json:"certification_time"`
	AnnexURL          *string `json:"annex_url"`
	CompanyID         int64   `json:"company_id"`
}

func (h *Handler) GetHealthSystem(c *gin.Context) {
	cid := scopeOrCompanyID(c)
	if cid == 0 {
		response.BadRequest(c, "缺少 company_id")
		return
	}
	row, err := h.svc.FirstHealthSystem(c.Request.Context(), cid)
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Success(c, row)
}

func (h *Handler) CreateHealthSystem(c *gin.Context) {
	var req healthSystemReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.CompanyID == 0 {
		req.CompanyID = scopeOrCompanyID(c)
	}
	uid := middleware.CurrentUserID(c)
	m := &model.HealthSystem{
		CertificationTime: parseDate(req.CertificationTime),
		AnnexURL:          req.AnnexURL,
		CompanyID:         req.CompanyID,
	}
	m.UserID = uid
	m.LastUserID = uid
	if err := h.svc.CreateHealthSystem(c.Request.Context(), m); err != nil {
		writeAppError(c, err)
		return
	}
	response.Created(c, m)
}

func (h *Handler) UpdateHealthSystem(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req healthSystemReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	updates := map[string]any{"last_user_id": middleware.CurrentUserID(c)}
	addNullableString(updates, "annex_url", req.AnnexURL)
	if req.CertificationTime != nil {
		updates["certification_time"] = parseDate(req.CertificationTime)
	}
	if err := h.svc.UpdateHealthSystem(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, gin.H{"id": id})
}

// ── 健康档案 ──────────────────────────────────────────────────────────

type healthFileReq struct {
	Name             *string `json:"name"`
	Number           *string `json:"number"`
	Department       *string `json:"department"`
	Post             *string `json:"post"`
	CheckupDate      *string `json:"checkup_date"`
	CheckupExpiry    *string `json:"checkup_expiry"`
	ReviewSituation  *string `json:"review_situation"`
	HandlingOpinions *string `json:"handling_opinions"`
	ProcessResult    *string `json:"process_result"`
	AnnexURL         *string `json:"annex_url"`
}

func (h *Handler) ListHealthFile(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	page, per := apiutil.ParsePagination(c, 20)
	rows, total, err := h.svc.ListHealthFile(c.Request.Context(), service.HealthFileListParams{
		Page: page, PerPage: per, CompanyID: companyID,
		Name: firstNonEmpty(c.Query("name"), c.Query("keyword")),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) GetHealthFile(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.svc.GetHealthFile(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

func (h *Handler) CreateHealthFile(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	var req healthFileReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	m := &model.HealthFile{
		Name:             req.Name,
		Number:           req.Number,
		Department:       req.Department,
		Post:             req.Post,
		CheckupDate:      parseDate(req.CheckupDate),
		CheckupExpiry:    req.CheckupExpiry,
		ReviewSituation:  req.ReviewSituation,
		HandlingOpinions: req.HandlingOpinions,
		ProcessResult:    req.ProcessResult,
		AnnexURL:         req.AnnexURL,
		CompanyID:        companyID,
	}
	m.UserID = uid
	m.LastUserID = uid
	if err := h.svc.CreateHealthFile(c.Request.Context(), m); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetHealthFile(c.Request.Context(), m.ID)
	response.Created(c, row)
}

func (h *Handler) UpdateHealthFile(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req healthFileReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	updates := map[string]any{"last_user_id": middleware.CurrentUserID(c)}
	addNullableString(updates, "name", req.Name)
	addNullableString(updates, "number", req.Number)
	addNullableString(updates, "department", req.Department)
	addNullableString(updates, "post", req.Post)
	addNullableString(updates, "checkup_expiry", req.CheckupExpiry)
	addNullableString(updates, "review_situation", req.ReviewSituation)
	addNullableString(updates, "handling_opinions", req.HandlingOpinions)
	addNullableString(updates, "process_result", req.ProcessResult)
	addNullableString(updates, "annex_url", req.AnnexURL)
	if req.CheckupDate != nil {
		updates["checkup_date"] = parseDate(req.CheckupDate)
	}
	if err := h.svc.UpdateHealthFile(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetHealthFile(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteHealthFile(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteHealthFile(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

func (h *Handler) TemplateHealthFile(c *gin.Context) {
	setXlsxAttachment(c, "职业健康档案导入模板.xlsx", "health_files_template.xlsx")
	if err := excel.HealthFileTemplate(c.Writer); err != nil {
		response.Internal(c, "生成模板失败")
		return
	}
}

func (h *Handler) ExportHealthFile(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	rows, _, err := h.svc.ListHealthFile(c.Request.Context(), service.HealthFileListParams{
		Page: 1, PerPage: 10000, CompanyID: companyID,
		Name: c.Query("name"),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	setXlsxAttachment(c, "职业健康档案列表.xlsx", "health_files.xlsx")
	if err := excel.HealthFileExport(c.Writer, rows); err != nil {
		response.Internal(c, "导出失败")
		return
	}
}

func (h *Handler) ImportHealthFile(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		response.BadRequest(c, "缺少 file 字段")
		return
	}
	defer file.Close()
	rows, err := excel.HealthFileImport(file)
	if err != nil {
		response.BadRequest(c, "解析 Excel 失败: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	imported := 0
	for i := range rows {
		rows[i].CompanyID = companyID
		rows[i].UserID = uid
		rows[i].LastUserID = uid
		if err := h.svc.CreateHealthFile(c.Request.Context(), &rows[i]); err == nil {
			imported++
		}
	}
	response.Success(c, gin.H{"imported": imported, "total": len(rows)})
}

// ── 职业病报告 ────────────────────────────────────────────────────────

type healthReportReq struct {
	FindDate     *string `json:"find_date"`
	Name         *string `json:"name"`
	Post         *string `json:"post"`
	Type         *string `json:"type"`
	IsTrue       *int16  `json:"is_true"`
	Mechanism    *string `json:"mechanism"`
	IdentifyDate *string `json:"identify_date"`
	AnnexURL     *string `json:"annex_url"`
}

func (h *Handler) ListHealthReport(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	page, per := apiutil.ParsePagination(c, 20)
	rows, total, err := h.svc.ListHealthReport(c.Request.Context(), service.HealthReportListParams{
		Page: page, PerPage: per, CompanyID: companyID,
		Name: firstNonEmpty(c.Query("name"), c.Query("keyword")),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) GetHealthReport(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.svc.GetHealthReport(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

func (h *Handler) CreateHealthReport(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	var req healthReportReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	m := &model.HealthReport{
		FindDate:     parseDate(req.FindDate),
		Name:         req.Name,
		Post:         req.Post,
		Type:         req.Type,
		IsTrue:       req.IsTrue,
		Mechanism:    req.Mechanism,
		IdentifyDate: parseDate(req.IdentifyDate),
		AnnexURL:     req.AnnexURL,
		CompanyID:    companyID,
	}
	m.UserID = uid
	m.LastUserID = uid
	if err := h.svc.CreateHealthReport(c.Request.Context(), m); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetHealthReport(c.Request.Context(), m.ID)
	response.Created(c, row)
}

func (h *Handler) UpdateHealthReport(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req healthReportReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	updates := map[string]any{"last_user_id": middleware.CurrentUserID(c)}
	addNullableString(updates, "name", req.Name)
	addNullableString(updates, "post", req.Post)
	addNullableString(updates, "type", req.Type)
	addNullableString(updates, "mechanism", req.Mechanism)
	addNullableString(updates, "annex_url", req.AnnexURL)
	if req.IsTrue != nil {
		updates["is_true"] = req.IsTrue
	}
	if req.FindDate != nil {
		updates["find_date"] = parseDate(req.FindDate)
	}
	if req.IdentifyDate != nil {
		updates["identify_date"] = parseDate(req.IdentifyDate)
	}
	if err := h.svc.UpdateHealthReport(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetHealthReport(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteHealthReport(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteHealthReport(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

// ── 职业健康体检 ──────────────────────────────────────────────────────

type healthCheckupReq struct {
	Type          *string    `json:"type"`
	Department    *string    `json:"department"`
	Number        *flexInt32 `json:"number"`
	CheckupDate   *string    `json:"checkup_date"`
	CheckupExpiry *string    `json:"checkup_expiry"`
	Measure       *string    `json:"measure"`
	AnnexURL      *string    `json:"annex_url"`
}

func (h *Handler) ListHealthCheckup(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	page, per := apiutil.ParsePagination(c, 20)
	rows, total, err := h.svc.ListHealthCheckup(c.Request.Context(), service.HealthCheckupListParams{
		Page: page, PerPage: per, CompanyID: companyID,
		Type: firstNonEmpty(c.Query("type"), c.Query("keyword")),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) GetHealthCheckup(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.svc.GetHealthCheckup(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

func (h *Handler) CreateHealthCheckup(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	var req healthCheckupReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	m := &model.HealthCheckup{
		Type:          req.Type,
		Department:    req.Department,
		Number:        req.Number.Int32Ptr(),
		CheckupDate:   parseDate(req.CheckupDate),
		CheckupExpiry: req.CheckupExpiry,
		Measure:       req.Measure,
		AnnexURL:      req.AnnexURL,
		CompanyID:     companyID,
	}
	m.UserID = uid
	m.LastUserID = uid
	if err := h.svc.CreateHealthCheckup(c.Request.Context(), m); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetHealthCheckup(c.Request.Context(), m.ID)
	response.Created(c, row)
}

func (h *Handler) UpdateHealthCheckup(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req healthCheckupReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	updates := map[string]any{"last_user_id": middleware.CurrentUserID(c)}
	addNullableString(updates, "type", req.Type)
	addNullableString(updates, "department", req.Department)
	addNullableString(updates, "checkup_expiry", req.CheckupExpiry)
	addNullableString(updates, "measure", req.Measure)
	addNullableString(updates, "annex_url", req.AnnexURL)
	if req.Number != nil {
		updates["number"] = req.Number.Int32Ptr()
	}
	if req.CheckupDate != nil {
		updates["checkup_date"] = parseDate(req.CheckupDate)
	}
	if err := h.svc.UpdateHealthCheckup(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetHealthCheckup(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteHealthCheckup(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteHealthCheckup(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

// ── 职业卫生评价 ──────────────────────────────────────────────────────

type hygieneEvaluationReq struct {
	Name           *string    `json:"name"`
	Type           *string    `json:"type"`
	SkillUnitID    *flexInt64 `json:"skill_unit_id"`
	SkillUnitKidID *flexInt64 `json:"skill_unit_kid_id"`
	ValidPeriod    *string    `json:"valid_period"`
	AnnexURL       *string    `json:"annex_url"`
}

func (h *Handler) ListHygieneEvaluation(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	page, per := apiutil.ParsePagination(c, 20)
	rows, total, err := h.svc.ListHygieneEvaluation(c.Request.Context(), service.HygieneEvaluationListParams{
		Page: page, PerPage: per, CompanyID: companyID,
		Name: firstNonEmpty(c.Query("name"), c.Query("keyword")),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) GetHygieneEvaluation(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.svc.GetHygieneEvaluation(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

func (h *Handler) CreateHygieneEvaluation(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	var req hygieneEvaluationReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	m := &model.HygieneEvaluation{
		Name:           req.Name,
		Type:           req.Type,
		SkillUnitID:    req.SkillUnitID.Int64Ptr(),
		SkillUnitKidID: req.SkillUnitKidID.Int64Ptr(),
		ValidPeriod:    parseDate(req.ValidPeriod),
		AnnexURL:       req.AnnexURL,
		CompanyID:      companyID,
	}
	m.UserID = uid
	m.LastUserID = uid
	if err := h.svc.CreateHygieneEvaluation(c.Request.Context(), m); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetHygieneEvaluation(c.Request.Context(), m.ID)
	response.Created(c, row)
}

func (h *Handler) UpdateHygieneEvaluation(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req hygieneEvaluationReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	updates := map[string]any{"last_user_id": middleware.CurrentUserID(c)}
	addNullableString(updates, "name", req.Name)
	addNullableString(updates, "type", req.Type)
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
	if err := h.svc.UpdateHygieneEvaluation(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetHygieneEvaluation(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteHygieneEvaluation(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteHygieneEvaluation(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

// ── 消防安全 ──────────────────────────────────────────────────────────

type fireSafetyReq struct {
	Name           *string    `json:"name"`
	Type           *string    `json:"type"`
	SkillUnitID    *flexInt64 `json:"skill_unit_id"`
	SkillUnitKidID *flexInt64 `json:"skill_unit_kid_id"`
	ValidPeriod    *string    `json:"valid_period"`
	AnnexURL       *string    `json:"annex_url"`
}

func (h *Handler) ListFireSafety(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	page, per := apiutil.ParsePagination(c, 20)
	rows, total, err := h.svc.ListFireSafety(c.Request.Context(), service.FireSafetyListParams{
		Page: page, PerPage: per, CompanyID: companyID,
		Name: firstNonEmpty(c.Query("name"), c.Query("keyword")),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) GetFireSafety(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.svc.GetFireSafety(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

func (h *Handler) CreateFireSafety(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	var req fireSafetyReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	m := &model.FireSafety{
		Name:           req.Name,
		Type:           req.Type,
		SkillUnitID:    req.SkillUnitID.Int64Ptr(),
		SkillUnitKidID: req.SkillUnitKidID.Int64Ptr(),
		ValidPeriod:    req.ValidPeriod,
		AnnexURL:       req.AnnexURL,
		CompanyID:      companyID,
	}
	m.UserID = uid
	m.LastUserID = uid
	if err := h.svc.CreateFireSafety(c.Request.Context(), m); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetFireSafety(c.Request.Context(), m.ID)
	response.Created(c, row)
}

func (h *Handler) UpdateFireSafety(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req fireSafetyReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	updates := map[string]any{"last_user_id": middleware.CurrentUserID(c)}
	addNullableString(updates, "name", req.Name)
	addNullableString(updates, "type", req.Type)
	addNullableString(updates, "valid_period", req.ValidPeriod)
	addNullableString(updates, "annex_url", req.AnnexURL)
	if req.SkillUnitID != nil {
		updates["skill_unit_id"] = req.SkillUnitID.Int64Ptr()
	}
	if req.SkillUnitKidID != nil {
		updates["skill_unit_kid_id"] = req.SkillUnitKidID.Int64Ptr()
	}
	if err := h.svc.UpdateFireSafety(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetFireSafety(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteFireSafety(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteFireSafety(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

// ── 消防器材 ──────────────────────────────────────────────────────────

type fireEquipmentReq struct {
	EquipmentName    *string `json:"equipment_name"`
	UtilizedLocation *string `json:"utilized_location"`
	SpecialType      *string `json:"special_type"`
	UseTime          *string `json:"use_time"`
	AutomationType   *string `json:"automation_type"`
	MaintainInfo     *string `json:"maintain_info"`
	Comment          *string `json:"comment"`
}

func (h *Handler) ListFireEquipment(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	page, per := apiutil.ParsePagination(c, 20)
	rows, total, err := h.svc.ListFireEquipment(c.Request.Context(), service.FireEquipmentListParams{
		Page: page, PerPage: per, CompanyID: companyID,
		EquipmentName: firstNonEmpty(c.Query("equipment_name"), c.Query("keyword")),
		SpecialType:   c.Query("special_type"),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) GetFireEquipment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.svc.GetFireEquipment(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

func (h *Handler) CreateFireEquipment(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	var req fireEquipmentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	m := &model.FireEquipment{
		EquipmentName:    req.EquipmentName,
		UtilizedLocation: req.UtilizedLocation,
		SpecialType:      req.SpecialType,
		UseTime:          parseDate(req.UseTime),
		AutomationType:   req.AutomationType,
		MaintainInfo:     req.MaintainInfo,
		Comment:          req.Comment,
		CompanyID:        companyID,
	}
	m.UserID = uid
	m.LastUserID = uid
	if err := h.svc.CreateFireEquipment(c.Request.Context(), m); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetFireEquipment(c.Request.Context(), m.ID)
	response.Created(c, row)
}

func (h *Handler) UpdateFireEquipment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req fireEquipmentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	updates := map[string]any{"last_user_id": middleware.CurrentUserID(c)}
	addNullableString(updates, "equipment_name", req.EquipmentName)
	addNullableString(updates, "utilized_location", req.UtilizedLocation)
	addNullableString(updates, "special_type", req.SpecialType)
	addNullableString(updates, "automation_type", req.AutomationType)
	addNullableString(updates, "maintain_info", req.MaintainInfo)
	addNullableString(updates, "comment", req.Comment)
	if req.UseTime != nil {
		updates["use_time"] = parseDate(req.UseTime)
	}
	if err := h.svc.UpdateFireEquipment(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetFireEquipment(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteFireEquipment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteFireEquipment(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

type copyReq struct {
	IDs []int64 `json:"ids"`
}

func (h *Handler) CopyFireEquipment(c *gin.Context) {
	var req copyReq
	if err := c.ShouldBindJSON(&req); err != nil || len(req.IDs) == 0 {
		response.BadRequest(c, "缺少 ids")
		return
	}
	ids, err := h.svc.CopyFireEquipment(c.Request.Context(), req.IDs, middleware.CurrentUserID(c))
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, gin.H{"ids": ids, "count": len(ids)})
}

// ── 消防设施 ──────────────────────────────────────────────────────────

type fireFacilityReq = fireEquipmentReq

func (h *Handler) ListFireFacility(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	page, per := apiutil.ParsePagination(c, 20)
	rows, total, err := h.svc.ListFireFacility(c.Request.Context(), service.FireFacilityListParams{
		Page: page, PerPage: per, CompanyID: companyID,
		EquipmentName: firstNonEmpty(c.Query("equipment_name"), c.Query("keyword")),
		SpecialType:   c.Query("special_type"),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) GetFireFacility(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.svc.GetFireFacility(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

func (h *Handler) CreateFireFacility(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	var req fireFacilityReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	m := &model.FireFacility{
		EquipmentName:    req.EquipmentName,
		UtilizedLocation: req.UtilizedLocation,
		SpecialType:      req.SpecialType,
		UseTime:          parseDate(req.UseTime),
		AutomationType:   req.AutomationType,
		MaintainInfo:     req.MaintainInfo,
		Comment:          req.Comment,
		CompanyID:        companyID,
	}
	m.UserID = uid
	m.LastUserID = uid
	if err := h.svc.CreateFireFacility(c.Request.Context(), m); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetFireFacility(c.Request.Context(), m.ID)
	response.Created(c, row)
}

func (h *Handler) UpdateFireFacility(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req fireFacilityReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	updates := map[string]any{"last_user_id": middleware.CurrentUserID(c)}
	addNullableString(updates, "equipment_name", req.EquipmentName)
	addNullableString(updates, "utilized_location", req.UtilizedLocation)
	addNullableString(updates, "special_type", req.SpecialType)
	addNullableString(updates, "automation_type", req.AutomationType)
	addNullableString(updates, "maintain_info", req.MaintainInfo)
	addNullableString(updates, "comment", req.Comment)
	if req.UseTime != nil {
		updates["use_time"] = parseDate(req.UseTime)
	}
	if err := h.svc.UpdateFireFacility(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetFireFacility(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteFireFacility(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteFireFacility(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

func (h *Handler) CopyFireFacility(c *gin.Context) {
	var req copyReq
	if err := c.ShouldBindJSON(&req); err != nil || len(req.IDs) == 0 {
		response.BadRequest(c, "缺少 ids")
		return
	}
	ids, err := h.svc.CopyFireFacility(c.Request.Context(), req.IDs, middleware.CurrentUserID(c))
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, gin.H{"ids": ids, "count": len(ids)})
}

func setXlsxAttachment(c *gin.Context, chineseName string, asciiFallback string) {
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header(
		"Content-Disposition",
		fmt.Sprintf(
			"attachment; filename=\"%s\"; filename*=UTF-8''%s",
			asciiFallback,
			url.PathEscape(chineseName),
		),
	)
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
