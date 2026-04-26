// Package thirdparty 实现第三方管理三个模块的 CRUD：
//
//	/design-units      — 设计单位
//	/evaluation-units  — 施工/评价单位
//	/skill-units       — 技术服务单位（含 kids 子表）
//
// 三模块共用一份 ThirdPartyService，写入语义参考 internal/api/process。
package thirdparty

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
	svc *service.ThirdPartyService
}

func NewHandler(svc *service.ThirdPartyService) *Handler {
	return &Handler{svc: svc}
}

// ── 设计单位 ────────────────────────────────────────────────────────────

type designUnitReq struct {
	Name              *string `json:"name"`
	QualificationName *string `json:"qualification_name"`
	Contact           *string `json:"contact"`
	ContactTel        *string `json:"contact_tel"`
	AnnexURL          *string `json:"annex_url"`
}

func (h *Handler) ListDesignUnit(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "20"))
	rows, total, err := h.svc.ListDesignUnit(c.Request.Context(), service.DesignUnitListParams{
		Page: page, PerPage: per, CompanyID: companyID,
		Name: firstNonEmpty(c.Query("name"), c.Query("keyword")),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) CreateDesignUnit(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	var req designUnitReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	m := &model.DesignUnit{
		Name:              req.Name,
		QualificationName: req.QualificationName,
		Contact:           req.Contact,
		ContactTel:        req.ContactTel,
		AnnexURL:          req.AnnexURL,
		CompanyID:         companyID,
	}
	m.UserID = uid
	m.LastUserID = uid
	if err := h.svc.CreateDesignUnit(c.Request.Context(), m); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetDesignUnit(c.Request.Context(), m.ID)
	response.Created(c, row)
}

func (h *Handler) UpdateDesignUnit(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req designUnitReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	updates := map[string]any{"last_user_id": middleware.CurrentUserID(c)}
	addNullable(updates, "name", req.Name)
	addNullable(updates, "qualification_name", req.QualificationName)
	addNullable(updates, "contact", req.Contact)
	addNullable(updates, "contact_tel", req.ContactTel)
	addNullable(updates, "annex_url", req.AnnexURL)
	if err := h.svc.UpdateDesignUnit(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetDesignUnit(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteDesignUnit(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteDesignUnit(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

// ── 评价单位 ────────────────────────────────────────────────────────────

type evaluationUnitReq struct {
	Name                  *string `json:"name"`
	IsQualification       int16   `json:"is_qualification"`
	QualificationName     *string `json:"qualification_name"`
	QualificationAnnexURL *string `json:"qualification_annex_url"`
	EvaluationArea        *string `json:"evaluation_area"`
	Contact               *string `json:"contact"`
	ContactTel            *string `json:"contact_tel"`
	IsAgreement           int16   `json:"is_agreement"`
	AgreementAnnexURL     *string `json:"agreement_annex_url"`
	SupervisionURL        *string `json:"supervision_url"`
	AnnexURL              *string `json:"annex_url"`
}

func (h *Handler) ListEvaluationUnit(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "20"))
	rows, total, err := h.svc.ListEvaluationUnit(c.Request.Context(), service.EvaluationUnitListParams{
		Page: page, PerPage: per, CompanyID: companyID,
		Name: firstNonEmpty(c.Query("name"), c.Query("keyword")),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) CreateEvaluationUnit(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	var req evaluationUnitReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	m := &model.EvaluationUnit{
		Name:                  req.Name,
		IsQualification:       defaultIfZero(req.IsQualification, 2),
		QualificationName:     req.QualificationName,
		QualificationAnnexURL: req.QualificationAnnexURL,
		EvaluationArea:        req.EvaluationArea,
		Contact:               req.Contact,
		ContactTel:            req.ContactTel,
		IsAgreement:           defaultIfZero(req.IsAgreement, 2),
		AgreementAnnexURL:     req.AgreementAnnexURL,
		SupervisionURL:        req.SupervisionURL,
		AnnexURL:              req.AnnexURL,
		CompanyID:             companyID,
	}
	m.UserID = uid
	m.LastUserID = uid
	if err := h.svc.CreateEvaluationUnit(c.Request.Context(), m); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetEvaluationUnit(c.Request.Context(), m.ID)
	response.Created(c, row)
}

func (h *Handler) UpdateEvaluationUnit(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req evaluationUnitReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	updates := map[string]any{
		"last_user_id":     middleware.CurrentUserID(c),
		"is_qualification": defaultIfZero(req.IsQualification, 2),
		"is_agreement":     defaultIfZero(req.IsAgreement, 2),
	}
	addNullable(updates, "name", req.Name)
	addNullable(updates, "qualification_name", req.QualificationName)
	addNullable(updates, "qualification_annex_url", req.QualificationAnnexURL)
	addNullable(updates, "evaluation_area", req.EvaluationArea)
	addNullable(updates, "contact", req.Contact)
	addNullable(updates, "contact_tel", req.ContactTel)
	addNullable(updates, "agreement_annex_url", req.AgreementAnnexURL)
	addNullable(updates, "supervision_url", req.SupervisionURL)
	addNullable(updates, "annex_url", req.AnnexURL)
	if err := h.svc.UpdateEvaluationUnit(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetEvaluationUnit(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteEvaluationUnit(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteEvaluationUnit(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

// ── 技术服务单位（含 kids 子表） ───────────────────────────────────────

type skillUnitKidReq struct {
	QualificationName     *string `json:"qualification_name"`
	QualificationAnnexURL *string `json:"qualification_annex_url"`
}

type skillUnitReq struct {
	Name            *string           `json:"name"`
	ServingType     *string           `json:"serving_type"`
	IsQualification int16             `json:"is_qualification"`
	EvaluationArea  *string           `json:"evaluation_area"`
	Contact         *string           `json:"contact"`
	ContactTel      *string           `json:"contact_tel"`
	Kids            []skillUnitKidReq `json:"kids"`
}

func (h *Handler) ListSkillUnit(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "20"))
	rows, total, err := h.svc.ListSkillUnit(c.Request.Context(), service.SkillUnitListParams{
		Page: page, PerPage: per, CompanyID: companyID,
		Name: firstNonEmpty(c.Query("name"), c.Query("keyword")),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) CreateSkillUnit(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	var req skillUnitReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	unit := &model.SkillUnit{
		Name:            req.Name,
		ServingType:     req.ServingType,
		IsQualification: defaultIfZero(req.IsQualification, 2),
		EvaluationArea:  req.EvaluationArea,
		Contact:         req.Contact,
		ContactTel:      req.ContactTel,
		CompanyID:       companyID,
	}
	unit.UserID = uid
	unit.LastUserID = uid
	if err := h.svc.CreateSkillUnit(c.Request.Context(), service.CreateSkillUnitInput{
		Unit: unit,
		Kids: mapSkillUnitKids(req.Kids, uid),
	}); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetSkillUnit(c.Request.Context(), unit.ID)
	response.Created(c, row)
}

func (h *Handler) UpdateSkillUnit(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req skillUnitReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	updates := map[string]any{
		"last_user_id":     uid,
		"is_qualification": defaultIfZero(req.IsQualification, 2),
	}
	addNullable(updates, "name", req.Name)
	addNullable(updates, "serving_type", req.ServingType)
	addNullable(updates, "evaluation_area", req.EvaluationArea)
	addNullable(updates, "contact", req.Contact)
	addNullable(updates, "contact_tel", req.ContactTel)

	// kids != nil 时整体替换；nil 时保留原数据。
	var kids []model.SkillUnitKid
	if req.Kids != nil {
		kids = mapSkillUnitKids(req.Kids, uid)
	}
	if err := h.svc.UpdateSkillUnit(c.Request.Context(), id, updates, kids); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetSkillUnit(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteSkillUnit(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteSkillUnit(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

// ── helpers ─────────────────────────────────────────────────────────────

func mapSkillUnitKids(items []skillUnitKidReq, uid int64) []model.SkillUnitKid {
	out := make([]model.SkillUnitKid, len(items))
	for i, it := range items {
		out[i] = model.SkillUnitKid{
			QualificationName:     it.QualificationName,
			QualificationAnnexURL: it.QualificationAnnexURL,
		}
		out[i].UserID = uid
		out[i].LastUserID = uid
	}
	return out
}

func defaultIfZero(v, def int16) int16 {
	if v == 0 {
		return def
	}
	return v
}

func addNullable(m map[string]any, key string, v *string) {
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
	cid, ok := middleware.ScopeCompanyID(c)
	if !ok || cid == 0 {
		response.BadRequest(c, "缺少 X-Scope-Company-Id 或 company_id")
		return 0
	}
	return cid
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
