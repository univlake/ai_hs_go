// Package meetingtraining 实现会议培训四个模块的 CRUD：
//
//	/internal-meetings   — 内部会议（含 kids 子表）
//	/internal-trainings  — 内部培训
//	/external-trainings  — 外部培训
//	/other-trainings     — 相关方培训
//
// 共用一份 MeetingTrainingService。日期字段（training_date / meeting_date）
// 在请求体里走 *string，转成 model.Date / model.DateTime 后写入。
package meetingtraining

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
	svc *service.MeetingTrainingService
}

func NewHandler(svc *service.MeetingTrainingService) *Handler {
	return &Handler{svc: svc}
}

// ── 内部会议（含 kids 子表） ───────────────────────────────────────────

type internalMeetingKidReq struct {
	JoinUserID *int64 `json:"join_user_id"`
}

type internalMeetingReq struct {
	Name             *string                 `json:"name"`
	Theme            *string                 `json:"theme"`
	Type             *string                 `json:"type"`
	MeetingDate      *string                 `json:"meeting_date"`
	Summary          *string                 `json:"summary"`
	AnnexURL         *string                 `json:"annex_url"`
	NoticeAnnexURL   *string                 `json:"notice_annex_url"`
	SummaryAnnexURL  *string                 `json:"summary_annex_url"`
	SignAnnexURL     *string                 `json:"sign_annex_url"`
	MaterialAnnexURL *string                 `json:"material_annex_url"`
	PhotoAnnexURL    *string                 `json:"photo_annex_url"`
	Kids             []internalMeetingKidReq `json:"kids"`
}

func (h *Handler) ListInternalMeeting(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "20"))
	rows, total, err := h.svc.ListInternalMeeting(c.Request.Context(), service.InternalMeetingListParams{
		Page: page, PerPage: per, CompanyID: companyID,
		Name: firstNonEmpty(c.Query("name"), c.Query("keyword")),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) GetInternalMeeting(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.svc.GetInternalMeeting(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

func (h *Handler) CreateInternalMeeting(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	var req internalMeetingReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	meeting := &model.InternalMeeting{
		Name:             req.Name,
		Theme:            req.Theme,
		Type:             req.Type,
		MeetingDate:      parseDateTime(req.MeetingDate),
		Summary:          req.Summary,
		AnnexURL:         req.AnnexURL,
		NoticeAnnexURL:   req.NoticeAnnexURL,
		SummaryAnnexURL:  req.SummaryAnnexURL,
		SignAnnexURL:     req.SignAnnexURL,
		MaterialAnnexURL: req.MaterialAnnexURL,
		PhotoAnnexURL:    req.PhotoAnnexURL,
		CompanyID:        companyID,
	}
	meeting.UserID = uid
	meeting.LastUserID = uid
	if err := h.svc.CreateInternalMeeting(c.Request.Context(), service.CreateInternalMeetingInput{
		Meeting: meeting,
		Kids:    mapInternalMeetingKids(req.Kids, uid),
	}); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetInternalMeeting(c.Request.Context(), meeting.ID)
	response.Created(c, row)
}

func (h *Handler) UpdateInternalMeeting(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req internalMeetingReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	updates := map[string]any{"last_user_id": uid}
	addNullableString(updates, "name", req.Name)
	addNullableString(updates, "theme", req.Theme)
	addNullableString(updates, "type", req.Type)
	addNullableString(updates, "summary", req.Summary)
	addNullableString(updates, "annex_url", req.AnnexURL)
	addNullableString(updates, "notice_annex_url", req.NoticeAnnexURL)
	addNullableString(updates, "summary_annex_url", req.SummaryAnnexURL)
	addNullableString(updates, "sign_annex_url", req.SignAnnexURL)
	addNullableString(updates, "material_annex_url", req.MaterialAnnexURL)
	addNullableString(updates, "photo_annex_url", req.PhotoAnnexURL)
	if req.MeetingDate != nil {
		updates["meeting_date"] = parseDateTime(req.MeetingDate)
	}

	var kids []model.InternalMeetingKid
	if req.Kids != nil {
		kids = mapInternalMeetingKids(req.Kids, uid)
	}
	if err := h.svc.UpdateInternalMeeting(c.Request.Context(), id, updates, kids); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetInternalMeeting(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteInternalMeeting(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteInternalMeeting(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

// ── 内部培训 ────────────────────────────────────────────────────────────

type internalTrainingReq struct {
	Name             *string `json:"name"`
	Comment          *string `json:"comment"`
	Object           *string `json:"object"`
	Number           *int32  `json:"number"`
	TrainingDate     *string `json:"training_date"`
	TrainingType     *string `json:"training_type"`
	Happening        *string `json:"happening"`
	AnnexURL         *string `json:"annex_url"`
	SummaryAnnexURL  *string `json:"summary_annex_url"`
	SignAnnexURL     *string `json:"sign_annex_url"`
	MaterialAnnexURL *string `json:"material_annex_url"`
	PhotoAnnexURL    *string `json:"photo_annex_url"`
}

func (h *Handler) ListInternalTraining(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "20"))
	rows, total, err := h.svc.ListInternalTraining(c.Request.Context(), service.TrainingListParams{
		Page: page, PerPage: per, CompanyID: companyID,
		Name: firstNonEmpty(c.Query("name"), c.Query("keyword")),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) GetInternalTraining(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.svc.GetInternalTraining(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

func (h *Handler) CreateInternalTraining(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	var req internalTrainingReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	m := &model.InternalTraining{
		Name:             req.Name,
		Comment:          req.Comment,
		Object:           req.Object,
		Number:           req.Number,
		TrainingDate:     parseDate(req.TrainingDate),
		TrainingType:     req.TrainingType,
		Happening:        req.Happening,
		AnnexURL:         req.AnnexURL,
		SummaryAnnexURL:  req.SummaryAnnexURL,
		SignAnnexURL:     req.SignAnnexURL,
		MaterialAnnexURL: req.MaterialAnnexURL,
		PhotoAnnexURL:    req.PhotoAnnexURL,
		CompanyID:        companyID,
	}
	m.UserID = uid
	m.LastUserID = uid
	if err := h.svc.CreateInternalTraining(c.Request.Context(), m); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetInternalTraining(c.Request.Context(), m.ID)
	response.Created(c, row)
}

func (h *Handler) UpdateInternalTraining(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req internalTrainingReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	updates := map[string]any{"last_user_id": middleware.CurrentUserID(c)}
	addNullableString(updates, "name", req.Name)
	addNullableString(updates, "comment", req.Comment)
	addNullableString(updates, "object", req.Object)
	addNullableString(updates, "training_type", req.TrainingType)
	addNullableString(updates, "happening", req.Happening)
	addNullableString(updates, "annex_url", req.AnnexURL)
	addNullableString(updates, "summary_annex_url", req.SummaryAnnexURL)
	addNullableString(updates, "sign_annex_url", req.SignAnnexURL)
	addNullableString(updates, "material_annex_url", req.MaterialAnnexURL)
	addNullableString(updates, "photo_annex_url", req.PhotoAnnexURL)
	if req.Number != nil {
		updates["number"] = req.Number
	}
	if req.TrainingDate != nil {
		updates["training_date"] = parseDate(req.TrainingDate)
	}
	if err := h.svc.UpdateInternalTraining(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetInternalTraining(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteInternalTraining(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteInternalTraining(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

// ── 外部培训 ────────────────────────────────────────────────────────────

type externalTrainingReq struct {
	Name             *string `json:"name"`
	Comment          *string `json:"comment"`
	Object           *string `json:"object"`
	Number           *int32  `json:"number"`
	TrainingDate     *string `json:"training_date"`
	TrainingUnit     *string `json:"training_unit"`
	AnnexURL         *string `json:"annex_url"`
	SummaryAnnexURL  *string `json:"summary_annex_url"`
	SignAnnexURL     *string `json:"sign_annex_url"`
	MaterialAnnexURL *string `json:"material_annex_url"`
	PhotoAnnexURL    *string `json:"photo_annex_url"`
}

func (h *Handler) ListExternalTraining(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "20"))
	rows, total, err := h.svc.ListExternalTraining(c.Request.Context(), service.TrainingListParams{
		Page: page, PerPage: per, CompanyID: companyID,
		Name: firstNonEmpty(c.Query("name"), c.Query("keyword")),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) GetExternalTraining(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.svc.GetExternalTraining(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

func (h *Handler) CreateExternalTraining(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	var req externalTrainingReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	m := &model.ExternalTraining{
		Name:             req.Name,
		Comment:          req.Comment,
		Object:           req.Object,
		Number:           req.Number,
		TrainingDate:     parseDate(req.TrainingDate),
		TrainingUnit:     req.TrainingUnit,
		AnnexURL:         req.AnnexURL,
		SummaryAnnexURL:  req.SummaryAnnexURL,
		SignAnnexURL:     req.SignAnnexURL,
		MaterialAnnexURL: req.MaterialAnnexURL,
		PhotoAnnexURL:    req.PhotoAnnexURL,
		CompanyID:        companyID,
	}
	m.UserID = uid
	m.LastUserID = uid
	if err := h.svc.CreateExternalTraining(c.Request.Context(), m); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetExternalTraining(c.Request.Context(), m.ID)
	response.Created(c, row)
}

func (h *Handler) UpdateExternalTraining(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req externalTrainingReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	updates := map[string]any{"last_user_id": middleware.CurrentUserID(c)}
	addNullableString(updates, "name", req.Name)
	addNullableString(updates, "comment", req.Comment)
	addNullableString(updates, "object", req.Object)
	addNullableString(updates, "training_unit", req.TrainingUnit)
	addNullableString(updates, "annex_url", req.AnnexURL)
	addNullableString(updates, "summary_annex_url", req.SummaryAnnexURL)
	addNullableString(updates, "sign_annex_url", req.SignAnnexURL)
	addNullableString(updates, "material_annex_url", req.MaterialAnnexURL)
	addNullableString(updates, "photo_annex_url", req.PhotoAnnexURL)
	if req.Number != nil {
		updates["number"] = req.Number
	}
	if req.TrainingDate != nil {
		updates["training_date"] = parseDate(req.TrainingDate)
	}
	if err := h.svc.UpdateExternalTraining(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetExternalTraining(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteExternalTraining(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteExternalTraining(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

// ── 相关方培训 ──────────────────────────────────────────────────────────

type otherTrainingReq struct {
	Name             *string `json:"name"`
	Comment          *string `json:"comment"`
	Object           *string `json:"object"`
	Number           *int32  `json:"number"`
	TrainingDate     *string `json:"training_date"`
	TaskContent      *string `json:"task_content"`
	AnnexURL         *string `json:"annex_url"`
	SummaryAnnexURL  *string `json:"summary_annex_url"`
	SignAnnexURL     *string `json:"sign_annex_url"`
	MaterialAnnexURL *string `json:"material_annex_url"`
	PhotoAnnexURL    *string `json:"photo_annex_url"`
}

func (h *Handler) ListOtherTraining(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "20"))
	rows, total, err := h.svc.ListOtherTraining(c.Request.Context(), service.TrainingListParams{
		Page: page, PerPage: per, CompanyID: companyID,
		Name: firstNonEmpty(c.Query("name"), c.Query("keyword")),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) GetOtherTraining(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.svc.GetOtherTraining(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

func (h *Handler) CreateOtherTraining(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	var req otherTrainingReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	m := &model.OtherTraining{
		Name:             req.Name,
		Comment:          req.Comment,
		Object:           req.Object,
		Number:           req.Number,
		TrainingDate:     parseDate(req.TrainingDate),
		TaskContent:      req.TaskContent,
		AnnexURL:         req.AnnexURL,
		SummaryAnnexURL:  req.SummaryAnnexURL,
		SignAnnexURL:     req.SignAnnexURL,
		MaterialAnnexURL: req.MaterialAnnexURL,
		PhotoAnnexURL:    req.PhotoAnnexURL,
		CompanyID:        companyID,
	}
	m.UserID = uid
	m.LastUserID = uid
	if err := h.svc.CreateOtherTraining(c.Request.Context(), m); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetOtherTraining(c.Request.Context(), m.ID)
	response.Created(c, row)
}

func (h *Handler) UpdateOtherTraining(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req otherTrainingReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	updates := map[string]any{"last_user_id": middleware.CurrentUserID(c)}
	addNullableString(updates, "name", req.Name)
	addNullableString(updates, "comment", req.Comment)
	addNullableString(updates, "object", req.Object)
	addNullableString(updates, "task_content", req.TaskContent)
	addNullableString(updates, "annex_url", req.AnnexURL)
	addNullableString(updates, "summary_annex_url", req.SummaryAnnexURL)
	addNullableString(updates, "sign_annex_url", req.SignAnnexURL)
	addNullableString(updates, "material_annex_url", req.MaterialAnnexURL)
	addNullableString(updates, "photo_annex_url", req.PhotoAnnexURL)
	if req.Number != nil {
		updates["number"] = req.Number
	}
	if req.TrainingDate != nil {
		updates["training_date"] = parseDate(req.TrainingDate)
	}
	if err := h.svc.UpdateOtherTraining(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetOtherTraining(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteOtherTraining(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteOtherTraining(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

// ── helpers ─────────────────────────────────────────────────────────────

func mapInternalMeetingKids(items []internalMeetingKidReq, uid int64) []model.InternalMeetingKid {
	out := make([]model.InternalMeetingKid, len(items))
	for i, it := range items {
		out[i] = model.InternalMeetingKid{JoinUserID: it.JoinUserID}
		out[i].UserID = uid
		out[i].LastUserID = uid
	}
	return out
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
