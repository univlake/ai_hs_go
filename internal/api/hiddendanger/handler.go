// Package hiddendanger 实现隐患信息（hidden_danger_infos）的 HTTP 入口：
//
//	GET    /hidden-danger-infos                  按 expert_inspection_id / company_id 等过滤
//	GET    /hidden-danger-infos/:id              详情含三类附件
//	POST   /hidden-danger-infos                  新增 + 一并提交三类附件
//	PUT    /hidden-danger-infos/:id              partial update + 三类附件整体替换
//	DELETE /hidden-danger-infos/:id              软删除（附件硬删除随主表）
//	POST   /hidden-danger-infos/:id/rectify      整改提交：写 zg_user_id / correct_status=2
//	POST   /hidden-danger-infos/:id/confirm-rectification  整改确认：写 review_user_id / correct_status=3
//	POST   /hidden-danger-infos/:id/archive      归档：写 archive_user_id / invest_status=2
//
// 跟 Laravel HiddenDangerInfoController 行为对齐。
package hiddendanger

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/middleware"
	"ai_hs_go/internal/model"
	"ai_hs_go/internal/response"
	"ai_hs_go/internal/service"
)

type Handler struct {
	svc *service.HiddenDangerService
}

func NewHandler(svc *service.HiddenDangerService) *Handler {
	return &Handler{svc: svc}
}

// ── List / Get ───────────────────────────────────────────────────────

func (h *Handler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "15"))
	insID, _ := strconv.ParseInt(c.Query("expert_inspection_id"), 10, 64)
	ecID, _ := strconv.ParseInt(c.Query("expert_company_id"), 10, 64)
	chkID, _ := strconv.ParseInt(c.Query("check_company_id"), 10, 64)
	cid, _ := strconv.ParseInt(c.Query("company_id"), 10, 64)
	uid, _ := strconv.ParseInt(c.Query("expert_user_id"), 10, 64)
	var cs *int
	if v := c.Query("correct_status"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			cs = &n
		}
	}
	rows, total, err := h.svc.List(c.Request.Context(), service.HiddenDangerListParams{
		Page: page, PerPage: per,
		ExpertInspectionID: insID, ExpertCompanyID: ecID, CheckCompanyID: chkID,
		CompanyID: cid, ExpertUserID: uid,
		Area: c.Query("area"), HiddenLevel: c.Query("hidden_level"), HiddenType: c.Query("hidden_type"),
		CorrectStatus: cs,
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.svc.Get(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err, "隐患信息不存在")
		return
	}
	response.Success(c, row)
}

// ── Create / Update / Delete ────────────────────────────────────────

type annexReq struct {
	Geolocation *string `json:"geolocation"`
	AnnexURL    *string `json:"annex_url"`
}

type dangerReq struct {
	ExpertInspectionID *int64      `json:"expert_inspection_id"`
	ExpertCompanyID    *int64      `json:"expert_company_id"`
	CheckCompanyID     *int64      `json:"check_company_id"`
	Area               *string     `json:"area"`
	Session            *string     `json:"session"`
	Description        *string     `json:"description"`
	HiddenType         *string     `json:"hidden_type"`
	HiddenLevel        *string     `json:"hidden_level"`
	Legalises          *string     `json:"legalises"`
	Proposal           *string     `json:"proposal"`
	CorrectDeadline    *string     `json:"correct_deadline"`
	InvestStatus       *int        `json:"invest_status"`
	CorrectStatus      *int        `json:"correct_status"`
	CorrectDate        *string     `json:"correct_date"`
	ReviewDescription  *string     `json:"review_description"`
	ReviewUserID       *int64      `json:"review_user_id"`
	ArchiveUserID      *int64      `json:"archive_user_id"`
	Comment            *string     `json:"comment"`
	CompanyID          *int64      `json:"company_id"`
	ExamineAnnex       *[]annexReq `json:"hiddenExamineAnnex"`
	ReviewAnnex        *[]annexReq `json:"hiddenReviewAnnex"`
	RectifyAnnex       *[]annexReq `json:"hiddenRectifyAnnex"`
}

func (h *Handler) Create(c *gin.Context) {
	var req dangerReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.Area == nil || *req.Area == "" {
		response.BadRequest(c, "area 不能为空")
		return
	}
	if req.Description == nil || *req.Description == "" {
		response.BadRequest(c, "description 不能为空")
		return
	}
	if req.HiddenType == nil || *req.HiddenType == "" {
		response.BadRequest(c, "hidden_type 不能为空")
		return
	}
	if req.HiddenLevel == nil || *req.HiddenLevel == "" {
		response.BadRequest(c, "hidden_level 不能为空")
		return
	}
	if req.CompanyID == nil || *req.CompanyID == 0 {
		response.BadRequest(c, "company_id 不能为空")
		return
	}
	uid := middleware.CurrentUserID(c)
	rec := &model.HiddenDangerInfo{
		ExpertInspectionID: req.ExpertInspectionID,
		ExpertCompanyID:    req.ExpertCompanyID,
		CheckCompanyID:     req.CheckCompanyID,
		Area:               *req.Area,
		Session:            "第一轮",
		Description:        *req.Description,
		HiddenType:         *req.HiddenType,
		HiddenLevel:        *req.HiddenLevel,
		Legalises:          req.Legalises,
		Proposal:           req.Proposal,
		CorrectDeadline:    parseDate(req.CorrectDeadline),
		Comment:            req.Comment,
		CompanyID:          *req.CompanyID,
	}
	if req.Session != nil && *req.Session != "" {
		rec.Session = *req.Session
	}
	if req.InvestStatus != nil {
		rec.InvestStatus = *req.InvestStatus
	}
	if req.CorrectStatus != nil {
		rec.CorrectStatus = *req.CorrectStatus
	}
	rec.UserID = uid
	rec.LastUserID = uid
	row, err := h.svc.Create(c.Request.Context(), service.HiddenDangerInput{
		Record:         rec,
		ExamineAnnexes: mapAnnexes(req.ExamineAnnex),
		ReviewAnnexes:  mapAnnexes(req.ReviewAnnex),
		RectifyAnnexes: mapAnnexes(req.RectifyAnnex),
	})
	if err != nil {
		writeAppError(c, err, "创建失败")
		return
	}
	response.Created(c, row)
}

func (h *Handler) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req dangerReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	updates := map[string]any{"last_user_id": uid}
	addNullableString(updates, "area", req.Area)
	addNullableString(updates, "session", req.Session)
	addNullableString(updates, "description", req.Description)
	addNullableString(updates, "hidden_type", req.HiddenType)
	addNullableString(updates, "hidden_level", req.HiddenLevel)
	addNullableString(updates, "legalises", req.Legalises)
	addNullableString(updates, "proposal", req.Proposal)
	addNullableString(updates, "review_description", req.ReviewDescription)
	addNullableString(updates, "comment", req.Comment)
	if req.CorrectDeadline != nil {
		updates["correct_deadline"] = parseDate(req.CorrectDeadline)
	}
	if req.CorrectDate != nil {
		updates["correct_date"] = parseDate(req.CorrectDate)
	}
	if req.InvestStatus != nil {
		updates["invest_status"] = *req.InvestStatus
	}
	if req.CorrectStatus != nil {
		updates["correct_status"] = *req.CorrectStatus
	}
	if req.ReviewUserID != nil {
		updates["review_user_id"] = *req.ReviewUserID
		updates["review_time"] = time.Now()
	}
	if req.ArchiveUserID != nil {
		updates["archive_user_id"] = *req.ArchiveUserID
	}
	if req.ReviewDescription != nil && *req.ReviewDescription != "" {
		updates["zg_user_id"] = uid
	}
	row, err := h.svc.Update(c.Request.Context(), id, updates,
		mapAnnexes(req.ExamineAnnex), mapAnnexes(req.ReviewAnnex), mapAnnexes(req.RectifyAnnex), uid)
	if err != nil {
		writeAppError(c, err, "更新失败")
		return
	}
	response.Success(c, row)
}

func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		writeAppError(c, err, "删除失败")
		return
	}
	response.Success(c, gin.H{"message": "删除成功"})
}

// ── Workflow endpoints ──────────────────────────────────────────────

type rectifyReq struct {
	ReviewDescription string      `json:"review_description"`
	CorrectDate       *string     `json:"correct_date"`
	RectifyAnnex      []annexReq  `json:"hiddenRectifyAnnex"`
}

func (h *Handler) Rectify(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req rectifyReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	annexes := make([]service.AnnexInput, 0, len(req.RectifyAnnex))
	for _, a := range req.RectifyAnnex {
		annexes = append(annexes, service.AnnexInput{
			Geolocation: deref(a.Geolocation),
			AnnexURL:    deref(a.AnnexURL),
		})
	}
	row, err := h.svc.Rectify(c.Request.Context(), id, uid, req.ReviewDescription, parseDate(req.CorrectDate), annexes)
	if err != nil {
		writeAppError(c, err, "提交整改失败")
		return
	}
	response.Success(c, row)
}

type confirmReq struct {
	ReviewAnnex []annexReq `json:"hiddenReviewAnnex"`
}

func (h *Handler) ConfirmRectification(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req confirmReq
	if err := c.ShouldBindJSON(&req); err != nil {
		// allow empty body
		req = confirmReq{}
	}
	uid := middleware.CurrentUserID(c)
	annexes := make([]service.AnnexInput, 0, len(req.ReviewAnnex))
	for _, a := range req.ReviewAnnex {
		annexes = append(annexes, service.AnnexInput{
			Geolocation: deref(a.Geolocation),
			AnnexURL:    deref(a.AnnexURL),
		})
	}
	row, err := h.svc.ConfirmRectification(c.Request.Context(), id, uid, annexes)
	if err != nil {
		writeAppError(c, err, "整改确认失败")
		return
	}
	response.Success(c, row)
}

func (h *Handler) Archive(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	uid := middleware.CurrentUserID(c)
	row, err := h.svc.Archive(c.Request.Context(), id, uid)
	if err != nil {
		writeAppError(c, err, "归档失败")
		return
	}
	response.Success(c, row)
}

// ── helpers ─────────────────────────────────────────────────────────

func mapAnnexes(p *[]annexReq) *[]service.AnnexInput {
	if p == nil {
		return nil
	}
	out := make([]service.AnnexInput, 0, len(*p))
	for _, a := range *p {
		out = append(out, service.AnnexInput{
			Geolocation: deref(a.Geolocation),
			AnnexURL:    deref(a.AnnexURL),
		})
	}
	return &out
}

func deref(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}

func parseDate(p *string) *model.Date {
	if p == nil || *p == "" {
		return nil
	}
	for _, layout := range []string{"2006-01-02", "2006/01/02", "2006/1/2", "2006-01-02T15:04:05Z07:00"} {
		if t, err := time.Parse(layout, *p); err == nil {
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
