package safetyarchive

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"ai_hs_go/internal/middleware"
	"ai_hs_go/internal/response"
	"ai_hs_go/internal/service"
	"ai_hs_go/internal/api/apiutil"
)

// ── Safety architecture ─────────────────────────────────────────

// ArchitectureGet: GET /api/v1/safety-architectures?company_id=X
//
// Single-record-per-company. When no row exists yet we return null so the
// frontend can decide between POST and PUT.
func (h *Handler) ArchitectureGet(c *gin.Context) {
	cid := scopeOrCompanyID(c)
	if cid == 0 {
		response.BadRequest(c, "缺少 company_id")
		return
	}
	row, err := h.archSvc.First(c.Request.Context(), cid)
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	if row == nil {
		response.Success(c, nil)
		return
	}
	// Pre-sign the annex URL when OSS is configured so the frontend
	// renders a previewable link without a separate /upload-oss/signed
	// round-trip.
	if h.oss != nil && row.AnnexURL != nil && *row.AnnexURL != "" {
		if signed, err := h.oss.SignURL(*row.AnnexURL, time.Hour, nil); err == nil {
			row.SignedAnnexURL = &signed
		}
	}
	response.Success(c, row)
}

type architectureReq struct {
	ThumbnailURL *string `json:"thumbnail_url"`
	AnnexURL     *string `json:"annex_url"`
	CompanyID    int64   `json:"company_id"`
}

// ArchitectureCreate: POST /api/v1/safety-architectures
func (h *Handler) ArchitectureCreate(c *gin.Context) {
	var req architectureReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.CompanyID == 0 {
		req.CompanyID = scopeOrCompanyID(c)
	}
	row, err := h.archSvc.Create(c.Request.Context(), service.SafetyArchitectureInput{
		ThumbnailURL: req.ThumbnailURL,
		AnnexURL:     req.AnnexURL,
		CompanyID:    req.CompanyID,
		UserID:       middleware.CurrentUserID(c),
	})
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Created(c, row)
}

// ArchitectureUpdate: PUT /api/v1/safety-architectures/:id
func (h *Handler) ArchitectureUpdate(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req architectureReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	row, err := h.archSvc.Update(c.Request.Context(), id, service.SafetyArchitectureInput{
		ThumbnailURL: req.ThumbnailURL,
		AnnexURL:     req.AnnexURL,
		UserID:       middleware.CurrentUserID(c),
	})
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

// ── Filing permissions ──────────────────────────────────────────

type filingReq struct {
	RecordTypeID      *int64  `json:"record_type_id"`
	IssuingDepartment *string `json:"issuing_department"`
	FilingRange       *string `json:"filing_range"`
	FilingTime        *string `json:"filing_time"`
	ValidPeriod       *string `json:"valid_period"`
	Comment           *string `json:"comment"`
	AnnexURL          *string `json:"annex_url"`
	CompanyID         int64   `json:"company_id"`
}

func (h *Handler) FilingList(c *gin.Context) {
	page, per := apiutil.ParsePagination(c, 20)
	rows, total, err := h.filingSvc.List(c.Request.Context(), service.FilingPermissionListParams{
		Page: page, PerPage: per, CompanyID: scopeOrCompanyID(c),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) FilingCreate(c *gin.Context) {
	var req filingReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.CompanyID == 0 {
		req.CompanyID = scopeOrCompanyID(c)
	}
	row, err := h.filingSvc.Create(c.Request.Context(), service.FilingPermissionInput{
		RecordTypeID:      req.RecordTypeID,
		IssuingDepartment: req.IssuingDepartment,
		FilingRange:       req.FilingRange,
		FilingTime:        req.FilingTime,
		ValidPeriod:       req.ValidPeriod,
		Comment:           req.Comment,
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

func (h *Handler) FilingUpdate(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req filingReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	row, err := h.filingSvc.Update(c.Request.Context(), id, service.FilingPermissionInput{
		RecordTypeID:      req.RecordTypeID,
		IssuingDepartment: req.IssuingDepartment,
		FilingRange:       req.FilingRange,
		FilingTime:        req.FilingTime,
		ValidPeriod:       req.ValidPeriod,
		Comment:           req.Comment,
		AnnexURL:          req.AnnexURL,
		UserID:            middleware.CurrentUserID(c),
	})
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

func (h *Handler) FilingDelete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.filingSvc.Delete(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

// ── Buildings ───────────────────────────────────────────────────

type buildingReq struct {
	Title     *string `json:"title"`
	Structure *string `json:"structure"`
	Area      *string `json:"area"`
	FireClass *string `json:"fire_class"`
	CompanyID int64   `json:"company_id"`
}

func (h *Handler) BuildingList(c *gin.Context) {
	page, per := apiutil.ParsePagination(c, 20)
	rows, total, err := h.buildingSvc.List(c.Request.Context(), service.BuildingListParams{
		Page: page, PerPage: per,
		CompanyID: scopeOrCompanyID(c),
		Keyword:   c.Query("keyword"),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) BuildingCreate(c *gin.Context) {
	var req buildingReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.CompanyID == 0 {
		req.CompanyID = scopeOrCompanyID(c)
	}
	row, err := h.buildingSvc.Create(c.Request.Context(), service.BuildingInput{
		Title:     req.Title,
		Structure: req.Structure,
		Area:      req.Area,
		FireClass: req.FireClass,
		CompanyID: req.CompanyID,
		UserID:    middleware.CurrentUserID(c),
	})
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Created(c, row)
}

func (h *Handler) BuildingUpdate(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req buildingReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	row, err := h.buildingSvc.Update(c.Request.Context(), id, service.BuildingInput{
		Title:     req.Title,
		Structure: req.Structure,
		Area:      req.Area,
		FireClass: req.FireClass,
		UserID:    middleware.CurrentUserID(c),
	})
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

func (h *Handler) BuildingDelete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.buildingSvc.Delete(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

// ── Security investments ───────────────────────────────────────

type investmentReq struct {
	Year           *string  `json:"year"`
	Type           *string  `json:"type"`
	Title          *string  `json:"title"`
	Content        *string  `json:"content"`
	Budget         *float64 `json:"budget"`
	LastYearIncome *float64 `json:"last_year_income"`
	CompanyID      int64    `json:"company_id"`
}

func (h *Handler) InvestmentList(c *gin.Context) {
	page, per := apiutil.ParsePagination(c, 20)
	rows, total, err := h.investmentSvc.List(c.Request.Context(), service.SecurityInvestmentListParams{
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

func (h *Handler) InvestmentCreate(c *gin.Context) {
	var req investmentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.CompanyID == 0 {
		req.CompanyID = scopeOrCompanyID(c)
	}
	row, err := h.investmentSvc.Create(c.Request.Context(), service.SecurityInvestmentInput{
		Year:           req.Year,
		Type:           req.Type,
		Title:          req.Title,
		Content:        req.Content,
		Budget:         req.Budget,
		LastYearIncome: req.LastYearIncome,
		CompanyID:      req.CompanyID,
		UserID:         middleware.CurrentUserID(c),
	})
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Created(c, row)
}

func (h *Handler) InvestmentUpdate(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req investmentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	row, err := h.investmentSvc.Update(c.Request.Context(), id, service.SecurityInvestmentInput{
		Year:           req.Year,
		Type:           req.Type,
		Title:          req.Title,
		Content:        req.Content,
		Budget:         req.Budget,
		LastYearIncome: req.LastYearIncome,
		UserID:         middleware.CurrentUserID(c),
	})
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

func (h *Handler) InvestmentDelete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.investmentSvc.Delete(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

// ── Personnels ──────────────────────────────────────────────────

type personnelReq struct {
	CompanyID              int64 `json:"company_id"`
	TotalPeople            int   `json:"total_people"`
	Staff                  int   `json:"staff"`
	Service                int   `json:"service"`
	Manager                int   `json:"manager"`
	HighSchoolDiploma      int   `json:"high_school_diploma"`
	SecurityOfficer        int   `json:"security_officer"`
	SafetyEngineer         int   `json:"safety_engineer"`
	SafetyEngineeringMajor int   `json:"safety_engineering_major"`
	FireEngineer           int   `json:"fire_engineer"`
	Other                  int   `json:"other"`
}

func (h *Handler) PersonnelList(c *gin.Context) {
	page, per := apiutil.ParsePagination(c, 20)
	rows, total, err := h.personnelSvc.List(c.Request.Context(), service.PersonnelListParams{
		Page: page, PerPage: per, CompanyID: scopeOrCompanyID(c),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) PersonnelGet(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.personnelSvc.Get(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

func (h *Handler) PersonnelCreate(c *gin.Context) {
	var req personnelReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.CompanyID == 0 {
		req.CompanyID = scopeOrCompanyID(c)
	}
	row, err := h.personnelSvc.Create(c.Request.Context(), service.PersonnelInput{
		CompanyID:              req.CompanyID,
		TotalPeople:            req.TotalPeople,
		Staff:                  req.Staff,
		Service:                req.Service,
		Manager:                req.Manager,
		HighSchoolDiploma:      req.HighSchoolDiploma,
		SecurityOfficer:        req.SecurityOfficer,
		SafetyEngineer:         req.SafetyEngineer,
		SafetyEngineeringMajor: req.SafetyEngineeringMajor,
		FireEngineer:           req.FireEngineer,
		Other:                  req.Other,
		UserID:                 middleware.CurrentUserID(c),
	})
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Created(c, row)
}

func (h *Handler) PersonnelUpdate(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req personnelReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	row, err := h.personnelSvc.Update(c.Request.Context(), id, service.PersonnelInput{
		CompanyID:              req.CompanyID,
		TotalPeople:            req.TotalPeople,
		Staff:                  req.Staff,
		Service:                req.Service,
		Manager:                req.Manager,
		HighSchoolDiploma:      req.HighSchoolDiploma,
		SecurityOfficer:        req.SecurityOfficer,
		SafetyEngineer:         req.SafetyEngineer,
		SafetyEngineeringMajor: req.SafetyEngineeringMajor,
		FireEngineer:           req.FireEngineer,
		Other:                  req.Other,
		UserID:                 middleware.CurrentUserID(c),
	})
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}
