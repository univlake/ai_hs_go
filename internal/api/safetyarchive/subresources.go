package safetyarchive

import (
	"fmt"
	"io"
	"net/url"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"ai_hs_go/internal/excel"
	"ai_hs_go/internal/middleware"
	"ai_hs_go/internal/response"
	"ai_hs_go/internal/service"
)

// writeXlsxTemplate is the shared response helper for the 3 staff
// template downloads. Sets Content-Type + RFC 5987 Content-Disposition
// (Chinese filename + ASCII fallback) and streams the build function's
// xlsx output straight to the response body.
func writeXlsxTemplate(c *gin.Context, chineseName, asciiFallback string, build func(io.Writer) error) {
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header(
		"Content-Disposition",
		fmt.Sprintf(
			"attachment; filename=\"%s\"; filename*=UTF-8''%s",
			asciiFallback,
			url.PathEscape(chineseName),
		),
	)
	if err := build(c.Writer); err != nil {
		response.Internal(c, "生成模板失败")
		return
	}
}

// ── Certificate statuses ────────────────────────────────────────

type certStatusReq struct {
	Name                *string `json:"name"`
	Type                int16   `json:"type"`
	ShouldObtainPeople  *int    `json:"should_obtain_people"`
	AlreadyObtainPeople *int    `json:"already_obtain_people"`
	NoObtainPeople      *string `json:"no_obtain_people"`
	TrainPeople         *int    `json:"train_people"`
	CompanyID           int64   `json:"company_id"`
}

func (h *Handler) CertStatusList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "15"))
	params := service.CertificateStatusListParams{
		Page:      page,
		PerPage:   per,
		CompanyID: scopeOrCompanyID(c),
	}
	if v := c.Query("type"); v != "" {
		if t, err := strconv.Atoi(v); err == nil {
			t16 := int16(t)
			params.Type = &t16
		}
	}
	rows, total, err := h.certStatusSvc.List(c.Request.Context(), params)
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) CertStatusCreate(c *gin.Context) {
	var req certStatusReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.CompanyID == 0 {
		req.CompanyID = scopeOrCompanyID(c)
	}
	row, err := h.certStatusSvc.Create(c.Request.Context(), service.CertificateStatusInput{
		Name: req.Name, Type: req.Type,
		ShouldObtainPeople:  req.ShouldObtainPeople,
		AlreadyObtainPeople: req.AlreadyObtainPeople,
		NoObtainPeople:      req.NoObtainPeople,
		TrainPeople:         req.TrainPeople,
		CompanyID:           req.CompanyID,
		UserID:              middleware.CurrentUserID(c),
	})
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Created(c, row)
}

func (h *Handler) CertStatusUpdate(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req certStatusReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	row, err := h.certStatusSvc.Update(c.Request.Context(), id, service.CertificateStatusInput{
		Name: req.Name, Type: req.Type,
		ShouldObtainPeople:  req.ShouldObtainPeople,
		AlreadyObtainPeople: req.AlreadyObtainPeople,
		NoObtainPeople:      req.NoObtainPeople,
		TrainPeople:         req.TrainPeople,
		UserID:              middleware.CurrentUserID(c),
	})
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

// ── Production standardization (singleton) ─────────────────────

type productionStdReq struct {
	Grade       *string `json:"grade"`
	GetTime     *string `json:"get_time"`
	ValidPeriod *string `json:"valid_period"`
	AnnexURL    *string `json:"annex_url"`
	CompanyID   int64   `json:"company_id"`
}

func (h *Handler) ProductionStdGet(c *gin.Context) {
	cid := scopeOrCompanyID(c)
	if cid == 0 {
		response.BadRequest(c, "缺少 company_id")
		return
	}
	row, err := h.productionStdSvc.First(c.Request.Context(), cid)
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Success(c, row)
}

func (h *Handler) ProductionStdCreate(c *gin.Context) {
	var req productionStdReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.CompanyID == 0 {
		req.CompanyID = scopeOrCompanyID(c)
	}
	row, err := h.productionStdSvc.Create(c.Request.Context(), service.ProductionStandardizationInput{
		Grade:       req.Grade,
		GetTime:     req.GetTime,
		ValidPeriod: req.ValidPeriod,
		AnnexURL:    req.AnnexURL,
		CompanyID:   req.CompanyID,
		UserID:      middleware.CurrentUserID(c),
	})
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Created(c, row)
}

func (h *Handler) ProductionStdUpdate(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req productionStdReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	row, err := h.productionStdSvc.Update(c.Request.Context(), id, service.ProductionStandardizationInput{
		Grade:       req.Grade,
		GetTime:     req.GetTime,
		ValidPeriod: req.ValidPeriod,
		AnnexURL:    req.AnnexURL,
		UserID:      middleware.CurrentUserID(c),
	})
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

// ── Occupational health safety (singleton) ─────────────────────

type ohsReq struct {
	CertificationBody *string `json:"certification_body"`
	StandardEdition   *string `json:"standard_edition"`
	GetTime           *string `json:"get_time"`
	ValidPeriod       *string `json:"valid_period"`
	AnnexURL          *string `json:"annex_url"`
	CompanyID         int64   `json:"company_id"`
}

func (h *Handler) OHSGet(c *gin.Context) {
	cid := scopeOrCompanyID(c)
	if cid == 0 {
		response.BadRequest(c, "缺少 company_id")
		return
	}
	row, err := h.ohsSvc.First(c.Request.Context(), cid)
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Success(c, row)
}

func (h *Handler) OHSCreate(c *gin.Context) {
	var req ohsReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.CompanyID == 0 {
		req.CompanyID = scopeOrCompanyID(c)
	}
	row, err := h.ohsSvc.Create(c.Request.Context(), service.OccupationalHealthSafetyInput{
		CertificationBody: req.CertificationBody,
		StandardEdition:   req.StandardEdition,
		GetTime:           req.GetTime,
		ValidPeriod:       req.ValidPeriod,
		AnnexURL:          req.AnnexURL,
		CompanyID:         req.CompanyID,
		UserID:            middleware.CurrentUserID(c),
	})
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Created(c, row)
}

func (h *Handler) OHSUpdate(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req ohsReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	row, err := h.ohsSvc.Update(c.Request.Context(), id, service.OccupationalHealthSafetyInput{
		CertificationBody: req.CertificationBody,
		StandardEdition:   req.StandardEdition,
		GetTime:           req.GetTime,
		ValidPeriod:       req.ValidPeriod,
		AnnexURL:          req.AnnexURL,
		UserID:            middleware.CurrentUserID(c),
	})
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

// ── Funding executions ─────────────────────────────────────────

type fundingExecReq struct {
	Year                 *string  `json:"year"`
	SecurityInvestmentID *int64   `json:"security_investment_id"`
	SimpleContent        *string  `json:"simple_content"`
	Content              *string  `json:"content"`
	Budget               *float64 `json:"budget"`
	CompanyID            int64    `json:"company_id"`
}

func (h *Handler) FundingExecList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "15"))
	rows, total, err := h.fundingExecSvc.List(c.Request.Context(), service.FundingExecutionListParams{
		Page: page, PerPage: per,
		CompanyID: scopeOrCompanyID(c),
		Year:      c.Query("year"),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) FundingExecCreate(c *gin.Context) {
	var req fundingExecReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.CompanyID == 0 {
		req.CompanyID = scopeOrCompanyID(c)
	}
	row, err := h.fundingExecSvc.Create(c.Request.Context(), service.FundingExecutionInput{
		Year:                 req.Year,
		SecurityInvestmentID: req.SecurityInvestmentID,
		SimpleContent:        req.SimpleContent,
		Content:              req.Content,
		Budget:               req.Budget,
		CompanyID:            req.CompanyID,
		UserID:               middleware.CurrentUserID(c),
	})
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Created(c, row)
}

func (h *Handler) FundingExecUpdate(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req fundingExecReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	row, err := h.fundingExecSvc.Update(c.Request.Context(), id, service.FundingExecutionInput{
		Year:                 req.Year,
		SecurityInvestmentID: req.SecurityInvestmentID,
		SimpleContent:        req.SimpleContent,
		Content:              req.Content,
		Budget:               req.Budget,
		UserID:               middleware.CurrentUserID(c),
	})
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

func (h *Handler) FundingExecDelete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.fundingExecSvc.Delete(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

// ── Safety management staff ────────────────────────────────────

type safetyMgmtStaffReq struct {
	Name *string `json:"name"`
	// flexInt64: front-end Select sends "1" (string), bare int64 wouldn't
	// unmarshal that. flexInt64 accepts both shapes + null + empty string.
	SafetyManagementTypeID flexInt64 `json:"safety_management_type_id"`
	Tel                    *string   `json:"tel"`
	CertTitle              *string   `json:"cert_title"`
	CertNumber             *string   `json:"cert_number"`
	ValidPeriod            *string   `json:"valid_period"`
	CompanyID              int64     `json:"company_id"`
}

func (h *Handler) SafetyMgmtStaffList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "20"))
	params := service.SafetyManagementStaffListParams{
		Page: page, PerPage: per, CompanyID: scopeOrCompanyID(c),
	}
	if v := c.Query("safety_management_type_id"); v != "" {
		if id, err := strconv.ParseInt(v, 10, 64); err == nil {
			params.TypeID = &id
		}
	}
	rows, total, err := h.safetyMgmtStaffSvc.List(c.Request.Context(), params)
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) SafetyMgmtStaffCreate(c *gin.Context) {
	var req safetyMgmtStaffReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.CompanyID == 0 {
		req.CompanyID = scopeOrCompanyID(c)
	}
	row, err := h.safetyMgmtStaffSvc.Create(c.Request.Context(), service.SafetyManagementStaffInput{
		Name: req.Name, SafetyManagementTypeID: int64(req.SafetyManagementTypeID),
		Tel:         req.Tel,
		CertTitle:   req.CertTitle,
		CertNumber:  req.CertNumber,
		ValidPeriod: req.ValidPeriod,
		CompanyID:   req.CompanyID,
		UserID:      middleware.CurrentUserID(c),
	})
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Created(c, row)
}

func (h *Handler) SafetyMgmtStaffUpdate(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req safetyMgmtStaffReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	row, err := h.safetyMgmtStaffSvc.Update(c.Request.Context(), id, service.SafetyManagementStaffInput{
		Name: req.Name, SafetyManagementTypeID: int64(req.SafetyManagementTypeID),
		Tel:         req.Tel,
		CertTitle:   req.CertTitle,
		CertNumber:  req.CertNumber,
		ValidPeriod: req.ValidPeriod,
		UserID:      middleware.CurrentUserID(c),
	})
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

func (h *Handler) SafetyMgmtStaffDelete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.safetyMgmtStaffSvc.Delete(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

// ── Special equipment staff ────────────────────────────────────

type specialEquipStaffReq struct {
	Name        *string `json:"name"`
	CertTitle   *string `json:"cert_title"`
	CertNumber  *string `json:"cert_number"`
	ValidPeriod *string `json:"valid_period"`
	CompanyID   int64   `json:"company_id"`
}

func (h *Handler) SpecialEquipStaffList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "20"))
	rows, total, err := h.specialEquipStaffSvc.List(c.Request.Context(), service.SpecialEquipmentStaffListParams{
		Page: page, PerPage: per, CompanyID: scopeOrCompanyID(c),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) SpecialEquipStaffCreate(c *gin.Context) {
	var req specialEquipStaffReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.CompanyID == 0 {
		req.CompanyID = scopeOrCompanyID(c)
	}
	row, err := h.specialEquipStaffSvc.Create(c.Request.Context(), service.SpecialEquipmentStaffInput{
		Name: req.Name, CertTitle: req.CertTitle, CertNumber: req.CertNumber,
		ValidPeriod: req.ValidPeriod, CompanyID: req.CompanyID,
		UserID: middleware.CurrentUserID(c),
	})
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Created(c, row)
}

func (h *Handler) SpecialEquipStaffUpdate(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req specialEquipStaffReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	row, err := h.specialEquipStaffSvc.Update(c.Request.Context(), id, service.SpecialEquipmentStaffInput{
		Name: req.Name, CertTitle: req.CertTitle, CertNumber: req.CertNumber,
		ValidPeriod: req.ValidPeriod,
		UserID:      middleware.CurrentUserID(c),
	})
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

func (h *Handler) SpecialEquipStaffDelete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.specialEquipStaffSvc.Delete(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

// ── Elimination control staff ──────────────────────────────────

type elimCtrlStaffReq struct {
	Name        *string `json:"name"`
	Type        int16   `json:"type"`
	CertTitle   *string `json:"cert_title"`
	CertNumber  *string `json:"cert_number"`
	ValidPeriod *string `json:"valid_period"`
	CompanyID   int64   `json:"company_id"`
}

func (h *Handler) ElimCtrlStaffList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "20"))
	params := service.EliminationControlStaffListParams{
		Page: page, PerPage: per, CompanyID: scopeOrCompanyID(c),
	}
	if v := c.Query("type"); v != "" {
		if t, err := strconv.Atoi(v); err == nil {
			t16 := int16(t)
			params.Type = &t16
		}
	}
	rows, total, err := h.elimCtrlStaffSvc.List(c.Request.Context(), params)
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) ElimCtrlStaffCreate(c *gin.Context) {
	var req elimCtrlStaffReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.CompanyID == 0 {
		req.CompanyID = scopeOrCompanyID(c)
	}
	row, err := h.elimCtrlStaffSvc.Create(c.Request.Context(), service.EliminationControlStaffInput{
		Name: req.Name, Type: req.Type,
		CertTitle: req.CertTitle, CertNumber: req.CertNumber,
		ValidPeriod: req.ValidPeriod, CompanyID: req.CompanyID,
		UserID: middleware.CurrentUserID(c),
	})
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Created(c, row)
}

func (h *Handler) ElimCtrlStaffUpdate(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req elimCtrlStaffReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	row, err := h.elimCtrlStaffSvc.Update(c.Request.Context(), id, service.EliminationControlStaffInput{
		Name: req.Name, Type: req.Type,
		CertTitle: req.CertTitle, CertNumber: req.CertNumber,
		ValidPeriod: req.ValidPeriod,
		UserID:      middleware.CurrentUserID(c),
	})
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

func (h *Handler) ElimCtrlStaffDelete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.elimCtrlStaffSvc.Delete(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

// ── Excel templates for the 3 staff sub-tabs ───────────────────

func (h *Handler) SafetyMgmtStaffTemplate(c *gin.Context) {
	writeXlsxTemplate(c, "安全管理人员导入模板.xlsx", "safety_management_staff_template.xlsx",
		excel.SafetyManagementStaffTemplate)
}

func (h *Handler) SpecialEquipStaffTemplate(c *gin.Context) {
	writeXlsxTemplate(c, "特种作业人员导入模板.xlsx", "special_equipment_staff_template.xlsx",
		excel.SpecialEquipmentStaffTemplate)
}

func (h *Handler) ElimCtrlStaffTemplate(c *gin.Context) {
	writeXlsxTemplate(c, "消防控制室人员导入模板.xlsx", "elimination_control_staff_template.xlsx",
		excel.EliminationControlStaffTemplate)
}

// ── Excel importers ─────────────────────────────────────────────

// importStaffFile is the multipart-form -> *os.File adapter. Reads the
// `file` field, returns the io.Reader + the operator's company_id (the
// frontend posts company_id as a form field along with the file).
func openStaffUpload(c *gin.Context) (io.ReadCloser, int64, error) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		return nil, 0, fmt.Errorf("缺少 file 字段")
	}
	cid := scopeOrCompanyID(c)
	if v := c.PostForm("company_id"); v != "" {
		if id, err := strconv.ParseInt(v, 10, 64); err == nil && id > 0 {
			cid = id
		}
	}
	if cid == 0 {
		_ = file.Close()
		return nil, 0, fmt.Errorf("缺少 company_id")
	}
	return file, cid, nil
}

func (h *Handler) SafetyMgmtStaffImport(c *gin.Context) {
	file, cid, err := openStaffUpload(c)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	defer file.Close()
	rows, err := excel.SafetyManagementStaffImport(file)
	if err != nil {
		response.BadRequest(c, "解析 Excel 失败: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	imported := 0
	for _, r := range rows {
		in := service.SafetyManagementStaffInput{
			Name:                   r.Name,
			SafetyManagementTypeID: r.SafetyManagementTypeID,
			Tel:                    r.Tel,
			CertTitle:              r.CertTitle,
			CertNumber:             r.CertNumber,
			CompanyID:              cid,
			UserID:                 uid,
		}
		if r.ValidPeriod != nil {
			s := time.Time(*r.ValidPeriod).Format("2006-01-02")
			in.ValidPeriod = &s
		}
		if _, err := h.safetyMgmtStaffSvc.Create(c.Request.Context(), in); err == nil {
			imported++
		}
	}
	response.Success(c, gin.H{"imported": imported, "total": len(rows)})
}

func (h *Handler) SpecialEquipStaffImport(c *gin.Context) {
	file, cid, err := openStaffUpload(c)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	defer file.Close()
	rows, err := excel.SpecialEquipmentStaffImport(file)
	if err != nil {
		response.BadRequest(c, "解析 Excel 失败: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	imported := 0
	for _, r := range rows {
		in := service.SpecialEquipmentStaffInput{
			Name:       r.Name,
			CertTitle:  r.CertTitle,
			CertNumber: r.CertNumber,
			CompanyID:  cid,
			UserID:     uid,
		}
		if r.ValidPeriod != nil {
			s := time.Time(*r.ValidPeriod).Format("2006-01-02")
			in.ValidPeriod = &s
		}
		if _, err := h.specialEquipStaffSvc.Create(c.Request.Context(), in); err == nil {
			imported++
		}
	}
	response.Success(c, gin.H{"imported": imported, "total": len(rows)})
}

func (h *Handler) ElimCtrlStaffImport(c *gin.Context) {
	file, cid, err := openStaffUpload(c)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	defer file.Close()
	rows, err := excel.EliminationControlStaffImport(file)
	if err != nil {
		response.BadRequest(c, "解析 Excel 失败: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	imported := 0
	for _, r := range rows {
		in := service.EliminationControlStaffInput{
			Name:       r.Name,
			CertTitle:  r.CertTitle,
			CertNumber: r.CertNumber,
			CompanyID:  cid,
			UserID:     uid,
		}
		if r.ValidPeriod != nil {
			s := time.Time(*r.ValidPeriod).Format("2006-01-02")
			in.ValidPeriod = &s
		}
		if _, err := h.elimCtrlStaffSvc.Create(c.Request.Context(), in); err == nil {
			imported++
		}
	}
	response.Success(c, gin.H{"imported": imported, "total": len(rows)})
}
