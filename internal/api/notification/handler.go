// Package notification 实现通知公告 / 会议通知 / 信息稿管理三个模块的 HTTP 入口：
//
//	/office-notifies              通知公告 - 任务型 CRUD（生产视图 + 集团视图共用）
//	/office-notify-scopes/:id/feedback     接收企业提交反馈
//	/office-notify-notifications/pending   仪表板待办：本企业未反馈
//
//	/new-notifies                 通知公告 - 通知型 CRUD
//	/new-notify-scopes/:id/read           接收企业标记已读
//	/new-notify-notifications/pending     仪表板待办：本企业未读
//
//	/meetings                     会议通知 CRUD（含 kids = 接收企业整体替换）
//	/meetings/:id                 GET 含 meeting_notices + meeting_sign_ups
//	/meeting-notices/:id/notify           触发会议通知（写 notify_time）
//	/meeting-notices/:id/confirm          接收企业确认通知
//	/meeting-notifications/pending        仪表板待办：本企业未确认会议
//
//	/meeting-sign-ups             会议报名 CRUD（接收企业的报名记录）
//
//	/information-drafts                          信息稿 CRUD（按 scope=published/received 过滤）
//	/information-drafts/dashboard/readings       仪表板「最新动态」
//	/information-drafts/dashboard/pending        仪表板「待办」（下级上报未处理）
//
// 三个模块共用 NotificationService。kids 整体替换的语义：传 nil 不动子表，
// 传非 nil（含空切片）= 整体替换。
package notification

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
)

type Handler struct {
	svc *service.NotificationService
}

func NewHandler(svc *service.NotificationService) *Handler {
	return &Handler{svc: svc}
}

// ── 公共 kid 请求结构 ────────────────────────────────────────────────

type scopeKidReq struct {
	CompanyID        flexInt64 `json:"company_id"`
	CompanyName      *string   `json:"company_name"`
	CompanyGroupID   *flexInt64 `json:"company_group_id"`
	CompanyGroupName *string   `json:"company_group_name"`
}

func mapScopeKids(items []scopeKidReq) []service.ScopeKid {
	out := make([]service.ScopeKid, 0, len(items))
	for _, it := range items {
		k := service.ScopeKid{
			CompanyID:        int64(it.CompanyID),
			CompanyName:      it.CompanyName,
			CompanyGroupName: it.CompanyGroupName,
		}
		if it.CompanyGroupID != nil {
			gid := int64(*it.CompanyGroupID)
			k.CompanyGroupID = &gid
		}
		out = append(out, k)
	}
	return out
}

func derefScopeKids(p *[]scopeKidReq) []scopeKidReq {
	if p == nil {
		return nil
	}
	return *p
}

// ── 任务型通知（office_notifies） ────────────────────────────────────

type officeNotifyReq struct {
	Name         *string        `json:"name"`
	FeedbackTime *string        `json:"feedback_time"`
	AlertOthers  *string        `json:"alert_others"`
	Comment      *string        `json:"comment"`
	AnnexURL     *string        `json:"annex_url"`
	ExcelURL     *string        `json:"excel_url"`
	Kids         *[]scopeKidReq `json:"kids"`
}

func (h *Handler) ListOfficeNotify(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "15"))
	cid := companyFromQueryOrScope(c)
	notifyCID, _ := strconv.ParseInt(c.Query("notify_company_id"), 10, 64)
	rows, total, err := h.svc.ListOfficeNotify(c.Request.Context(), service.OfficeNotifyListParams{
		Page: page, PerPage: per, CompanyID: cid, NotifyCompanyID: notifyCID,
		Name: c.Query("name"),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) GetOfficeNotify(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.svc.GetOfficeNotify(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

func (h *Handler) CreateOfficeNotify(c *gin.Context) {
	companyID := companyFromBodyOrScope(c)
	if companyID == 0 {
		response.BadRequest(c, "缺少 company_id")
		return
	}
	var req officeNotifyReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	notify := &model.OfficeNotify{
		Name:         req.Name,
		FeedbackTime: parseDate(req.FeedbackTime),
		AlertOthers:  req.AlertOthers,
		Comment:      req.Comment,
		AnnexURL:     req.AnnexURL,
		ExcelURL:     req.ExcelURL,
		CompanyID:    companyID,
	}
	notify.UserID = uid
	notify.LastUserID = uid
	if err := h.svc.CreateOfficeNotify(c.Request.Context(), service.OfficeNotifyInput{
		Notify: notify,
		Kids:   mapScopeKids(derefScopeKids(req.Kids)),
	}); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetOfficeNotify(c.Request.Context(), notify.ID)
	response.Created(c, row)
}

func (h *Handler) UpdateOfficeNotify(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req officeNotifyReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	updates := map[string]any{"last_user_id": uid}
	addNullableString(updates, "name", req.Name)
	addNullableString(updates, "alert_others", req.AlertOthers)
	addNullableString(updates, "comment", req.Comment)
	addNullableString(updates, "annex_url", req.AnnexURL)
	addNullableString(updates, "excel_url", req.ExcelURL)
	if req.FeedbackTime != nil {
		updates["feedback_time"] = parseDate(req.FeedbackTime)
	}
	if err := h.svc.UpdateOfficeNotify(c.Request.Context(), id, updates,
		mapScopeKids(derefScopeKids(req.Kids)), req.Kids != nil, uid); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetOfficeNotify(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteOfficeNotify(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteOfficeNotify(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

func (h *Handler) PendingOfficeNotifyScopes(c *gin.Context) {
	cid := companyFromQueryOrScope(c)
	if cid == 0 {
		response.BadRequest(c, "缺少 company_id")
		return
	}
	rows, err := h.svc.PendingOfficeNotifyScopes(c.Request.Context(), cid)
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Success(c, rows)
}

type officeFeedbackReq struct {
	FeedbackAnnexURL string `json:"feedback_annex_url" binding:"required"`
}

func (h *Handler) SubmitOfficeNotifyFeedback(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req officeFeedbackReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	scope, err := h.svc.SubmitOfficeNotifyFeedback(c.Request.Context(), id, req.FeedbackAnnexURL, middleware.CurrentUserID(c))
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, scope)
}

// ── 通知型通知（new_notifies） ───────────────────────────────────────

type newNotifyReq struct {
	Name        *string        `json:"name"`
	AlertOthers *string        `json:"alert_others"`
	Comment     *string        `json:"comment"`
	AnnexURL    *string        `json:"annex_url"`
	Kids        *[]scopeKidReq `json:"kids"`
}

func (h *Handler) ListNewNotify(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "15"))
	cid := companyFromQueryOrScope(c)
	notifyCID, _ := strconv.ParseInt(c.Query("notify_company_id"), 10, 64)
	rows, total, err := h.svc.ListNewNotify(c.Request.Context(), service.NewNotifyListParams{
		Page: page, PerPage: per, CompanyID: cid, NotifyCompanyID: notifyCID,
		Name: c.Query("name"),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) GetNewNotify(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.svc.GetNewNotify(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

func (h *Handler) CreateNewNotify(c *gin.Context) {
	companyID := companyFromBodyOrScope(c)
	if companyID == 0 {
		response.BadRequest(c, "缺少 company_id")
		return
	}
	var req newNotifyReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	notify := &model.NewNotify{
		Name:        req.Name,
		AlertOthers: req.AlertOthers,
		Comment:     req.Comment,
		AnnexURL:    req.AnnexURL,
		CompanyID:   companyID,
	}
	notify.UserID = uid
	notify.LastUserID = uid
	if err := h.svc.CreateNewNotify(c.Request.Context(), service.NewNotifyInput{
		Notify: notify,
		Kids:   mapScopeKids(derefScopeKids(req.Kids)),
	}); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetNewNotify(c.Request.Context(), notify.ID)
	response.Created(c, row)
}

func (h *Handler) UpdateNewNotify(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req newNotifyReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	updates := map[string]any{"last_user_id": uid}
	addNullableString(updates, "name", req.Name)
	addNullableString(updates, "alert_others", req.AlertOthers)
	addNullableString(updates, "comment", req.Comment)
	addNullableString(updates, "annex_url", req.AnnexURL)
	if err := h.svc.UpdateNewNotify(c.Request.Context(), id, updates,
		mapScopeKids(derefScopeKids(req.Kids)), req.Kids != nil, uid); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetNewNotify(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteNewNotify(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteNewNotify(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

func (h *Handler) PendingNewNotifyScopes(c *gin.Context) {
	cid := companyFromQueryOrScope(c)
	if cid == 0 {
		response.BadRequest(c, "缺少 company_id")
		return
	}
	rows, err := h.svc.PendingNewNotifyScopes(c.Request.Context(), cid)
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Success(c, rows)
}

func (h *Handler) MarkNewNotifyRead(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	scope, err := h.svc.MarkNewNotifyRead(c.Request.Context(), id, middleware.CurrentUserID(c))
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, scope)
}

// ── 会议通知（meetings） ──────────────────────────────────────────────

type meetingReq struct {
	Name             *string        `json:"name"`
	MeetingDate      *string        `json:"meeting_date"`
	SignUpDate       *string        `json:"sign_up_date"`
	SignEndDate      *string        `json:"sign_end_date"`
	AlertOthers      *string        `json:"alert_others"`
	Theme            *string        `json:"theme"`
	Place            *string        `json:"place"`
	Summary          *string        `json:"summary"`
	AnnexURL         *string        `json:"annex_url"`
	SummaryAnnexURL  *string        `json:"summary_annex_url"`
	SignAnnexURL     *string        `json:"sign_annex_url"`
	MaterialAnnexURL *string        `json:"material_annex_url"`
	PhotoAnnexURL    *string        `json:"photo_annex_url"`
	Kids             *[]scopeKidReq `json:"kids"`
}

func (h *Handler) ListMeeting(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "15"))
	cid := companyFromQueryOrScope(c)
	notifyCID, _ := strconv.ParseInt(c.Query("notify_company_id"), 10, 64)
	rows, total, err := h.svc.ListMeeting(c.Request.Context(), service.MeetingListParams{
		Page: page, PerPage: per, CompanyID: cid, NotifyCompanyID: notifyCID,
		Name:      c.Query("name"),
		StartDate: c.Query("start_date"),
		EndDate:   c.Query("end_date"),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) GetMeeting(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.svc.GetMeeting(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

func (h *Handler) CreateMeeting(c *gin.Context) {
	companyID := companyFromBodyOrScope(c)
	if companyID == 0 {
		response.BadRequest(c, "缺少 company_id")
		return
	}
	var req meetingReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	meeting := &model.Meeting{
		Name:             req.Name,
		MeetingDate:      parseDateTime(req.MeetingDate),
		SignUpDate:       parseDateTime(req.SignUpDate),
		SignEndDate:      parseDateTime(req.SignEndDate),
		AlertOthers:      req.AlertOthers,
		Theme:            req.Theme,
		Place:            req.Place,
		Summary:          req.Summary,
		AnnexURL:         req.AnnexURL,
		SummaryAnnexURL:  req.SummaryAnnexURL,
		SignAnnexURL:     req.SignAnnexURL,
		MaterialAnnexURL: req.MaterialAnnexURL,
		PhotoAnnexURL:    req.PhotoAnnexURL,
		CompanyID:        companyID,
	}
	meeting.UserID = uid
	meeting.LastUserID = uid
	if err := h.svc.CreateMeeting(c.Request.Context(), service.MeetingInput{
		Meeting: meeting,
		Kids:    mapScopeKids(derefScopeKids(req.Kids)),
	}); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetMeeting(c.Request.Context(), meeting.ID)
	response.Created(c, row)
}

func (h *Handler) UpdateMeeting(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req meetingReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	updates := map[string]any{"last_user_id": uid}
	addNullableString(updates, "name", req.Name)
	addNullableString(updates, "alert_others", req.AlertOthers)
	addNullableString(updates, "theme", req.Theme)
	addNullableString(updates, "place", req.Place)
	addNullableString(updates, "summary", req.Summary)
	addNullableString(updates, "annex_url", req.AnnexURL)
	addNullableString(updates, "summary_annex_url", req.SummaryAnnexURL)
	addNullableString(updates, "sign_annex_url", req.SignAnnexURL)
	addNullableString(updates, "material_annex_url", req.MaterialAnnexURL)
	addNullableString(updates, "photo_annex_url", req.PhotoAnnexURL)
	if req.MeetingDate != nil {
		updates["meeting_date"] = parseDateTime(req.MeetingDate)
	}
	if req.SignUpDate != nil {
		updates["sign_up_date"] = parseDateTime(req.SignUpDate)
	}
	if req.SignEndDate != nil {
		updates["sign_end_date"] = parseDateTime(req.SignEndDate)
	}
	if err := h.svc.UpdateMeeting(c.Request.Context(), id, updates,
		mapScopeKids(derefScopeKids(req.Kids)), req.Kids != nil, uid); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetMeeting(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteMeeting(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteMeeting(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

func (h *Handler) NotifyMeetingNotice(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	notice, err := h.svc.NotifyMeetingNotice(c.Request.Context(), id, middleware.CurrentUserID(c))
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, notice)
}

func (h *Handler) ConfirmMeetingNotice(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	notice, err := h.svc.ConfirmMeetingNotice(c.Request.Context(), id, middleware.CurrentUserID(c))
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, notice)
}

func (h *Handler) PendingMeetingNotices(c *gin.Context) {
	cid := companyFromQueryOrScope(c)
	if cid == 0 {
		response.BadRequest(c, "缺少 company_id")
		return
	}
	rows, err := h.svc.PendingMeetingNotices(c.Request.Context(), cid)
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Success(c, rows)
}

// ── 会议报名（meeting_sign_ups） ─────────────────────────────────────

type meetingSignUpReq struct {
	CompanyID   *flexInt64 `json:"company_id"`
	MeetingID   *flexInt64 `json:"meeting_id"`
	CompanyName *string    `json:"company_name"`
	Name        *string    `json:"name"`
	Sex         *string    `json:"sex"`
	JobTitle    *string    `json:"job_title"`
	Tel         *string    `json:"tel"`
	CardNumber  *string    `json:"card_number"`
	IsDriving   *string    `json:"is_driving"`
}

func (h *Handler) ListMeetingSignUp(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "15"))
	mid, _ := strconv.ParseInt(c.Query("meeting_id"), 10, 64)
	cid, _ := strconv.ParseInt(c.Query("company_id"), 10, 64)
	rows, total, err := h.svc.ListMeetingSignUp(c.Request.Context(), service.MeetingSignUpListParams{
		Page: page, PerPage: per, MeetingID: mid, CompanyID: cid,
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) GetMeetingSignUp(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.svc.GetMeetingSignUp(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

func (h *Handler) CreateMeetingSignUp(c *gin.Context) {
	var req meetingSignUpReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.CompanyID == nil || req.MeetingID == nil {
		response.BadRequest(c, "company_id / meeting_id 不能为空")
		return
	}
	uid := middleware.CurrentUserID(c)
	su := &model.MeetingSignUp{
		CompanyName: req.CompanyName,
		Name:        req.Name,
		Sex:         req.Sex,
		JobTitle:    req.JobTitle,
		Tel:         req.Tel,
		CardNumber:  req.CardNumber,
		IsDriving:   req.IsDriving,
		MeetingID:   int64(*req.MeetingID),
		CompanyID:   int64(*req.CompanyID),
	}
	su.UserID = uid
	su.LastUserID = uid
	if err := h.svc.CreateMeetingSignUp(c.Request.Context(), su); err != nil {
		writeAppError(c, err)
		return
	}
	response.Created(c, su)
}

func (h *Handler) UpdateMeetingSignUp(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req meetingSignUpReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	updates := map[string]any{"last_user_id": middleware.CurrentUserID(c)}
	addNullableString(updates, "company_name", req.CompanyName)
	addNullableString(updates, "name", req.Name)
	addNullableString(updates, "sex", req.Sex)
	addNullableString(updates, "job_title", req.JobTitle)
	addNullableString(updates, "tel", req.Tel)
	addNullableString(updates, "card_number", req.CardNumber)
	addNullableString(updates, "is_driving", req.IsDriving)
	if err := h.svc.UpdateMeetingSignUp(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetMeetingSignUp(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteMeetingSignUp(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteMeetingSignUp(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

// ── 信息稿（information_drafts） ─────────────────────────────────────

type informationDraftReq struct {
	Name            *string    `json:"name"`
	CompanyName     *string    `json:"company_name"`
	DraftDepartment *string    `json:"draft_department"`
	DraftOthers     *string    `json:"draft_others"`
	Comment         *string    `json:"comment"`
	AnnexURL        *string    `json:"annex_url"`
	NextReviewer    *flexInt64 `json:"next_reviewer"`
	IsTop           *flexInt32 `json:"is_top"`
	Status          *flexInt32 `json:"status"`
}

func (h *Handler) ListInformationDraft(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "15"))
	cid := companyFromQueryOrScope(c)
	rows, total, err := h.svc.ListInformationDraft(c.Request.Context(), service.InformationDraftListParams{
		Page: page, PerPage: per, CompanyID: cid,
		Scope:  c.Query("scope"),
		Status: c.Query("status"),
		Name:   c.Query("name"),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) GetInformationDraft(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.svc.GetInformationDraft(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

func (h *Handler) CreateInformationDraft(c *gin.Context) {
	companyID := companyFromBodyOrScope(c)
	if companyID == 0 {
		response.BadRequest(c, "缺少 company_id")
		return
	}
	var req informationDraftReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.Name == nil || *req.Name == "" {
		response.BadRequest(c, "name 不能为空")
		return
	}
	if req.CompanyName == nil || *req.CompanyName == "" {
		response.BadRequest(c, "company_name 不能为空")
		return
	}
	uid := middleware.CurrentUserID(c)
	draft := &model.InformationDraft{
		Name:            *req.Name,
		CompanyName:     *req.CompanyName,
		DraftDepartment: req.DraftDepartment,
		DraftOthers:     req.DraftOthers,
		Comment:         req.Comment,
		AnnexURL:        req.AnnexURL,
		CompanyID:       companyID,
	}
	if req.NextReviewer != nil {
		nr := int64(*req.NextReviewer)
		draft.NextReviewer = &nr
	}
	if req.IsTop != nil {
		draft.IsTop = int(*req.IsTop)
	}
	if req.Status != nil {
		draft.Status = int(*req.Status)
	}
	draft.UserID = uid
	draft.LastUserID = uid
	if err := h.svc.CreateInformationDraft(c.Request.Context(), draft); err != nil {
		writeAppError(c, err)
		return
	}
	response.Created(c, draft)
}

func (h *Handler) UpdateInformationDraft(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req informationDraftReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	updates := map[string]any{"last_user_id": middleware.CurrentUserID(c)}
	addNullableString(updates, "name", req.Name)
	addNullableString(updates, "company_name", req.CompanyName)
	addNullableString(updates, "draft_department", req.DraftDepartment)
	addNullableString(updates, "draft_others", req.DraftOthers)
	addNullableString(updates, "comment", req.Comment)
	addNullableString(updates, "annex_url", req.AnnexURL)
	if req.NextReviewer != nil {
		nr := int64(*req.NextReviewer)
		updates["next_reviewer"] = nr
	}
	if req.IsTop != nil {
		updates["is_top"] = int(*req.IsTop)
	}
	if req.Status != nil {
		updates["status"] = int(*req.Status)
	}
	if err := h.svc.UpdateInformationDraft(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetInformationDraft(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteInformationDraft(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteInformationDraft(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

func (h *Handler) DashboardInformationDraftReadings(c *gin.Context) {
	cid := companyFromQueryOrScope(c)
	if cid == 0 {
		response.BadRequest(c, "缺少 company_id")
		return
	}
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "6"))
	rows, err := h.svc.DashboardInformationDraftReadings(c.Request.Context(), cid, limit)
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Success(c, rows)
}

func (h *Handler) DashboardInformationDraftPending(c *gin.Context) {
	cid := companyFromQueryOrScope(c)
	if cid == 0 {
		response.BadRequest(c, "缺少 company_id")
		return
	}
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	rows, err := h.svc.DashboardInformationDraftPending(c.Request.Context(), cid, limit)
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Success(c, rows)
}

// ── 共用 helpers ─────────────────────────────────────────────────────

// companyFromQueryOrScope — 三个模块的 list / dashboard endpoint 都允许前端通过
// ?company_id= 显式指定（前端在切换企业前后保持一致），未传时回落到中间件解析的 scope。
func companyFromQueryOrScope(c *gin.Context) int64 {
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

// companyFromBodyOrScope — 创建/更新接受 body 里的 company_id（前端表单显式传），
// 否则回落到 scope 中间件。三个模块的 form 都包含 company_id 字段。
func companyFromBodyOrScope(c *gin.Context) int64 {
	// Avoid consuming the request body — bind into a probe struct via
	// ShouldBindBodyWith would be cleaner but introduces a dep; instead
	// trust scope + let the per-handler json.Unmarshal pick up other fields.
	if cid, ok := middleware.ScopeCompanyID(c); ok && cid > 0 {
		return cid
	}
	return 0
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

func parseDateTime(p *string) *model.DateTime {
	if p == nil || *p == "" {
		return nil
	}
	for _, layout := range []string{
		"2006-01-02T15:04",       // <input type="datetime-local">
		"2006-01-02T15:04:05",
		"2006-01-02 15:04:05",
		"2006-01-02 15:04",
		"2006-01-02",
		"2006-01-02T15:04:05Z07:00",
	} {
		if t, err := time.Parse(layout, *p); err == nil {
			dt := model.DateTime(t)
			return &dt
		}
	}
	return nil
}

func addNullableString(m map[string]any, key string, v *string) {
	if v != nil {
		m[key] = v
	}
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

// flexInt32 / flexInt64 — 容忍前端 <input type=number> 把数字以字符串提交。
type flexInt32 int32
type flexInt64 int64

func (f *flexInt32) UnmarshalJSON(data []byte) error {
	v, err := unmarshalFlexInt(data, 32)
	if err != nil {
		return err
	}
	*f = flexInt32(v)
	return nil
}

func (f *flexInt64) UnmarshalJSON(data []byte) error {
	v, err := unmarshalFlexInt(data, 64)
	if err != nil {
		return err
	}
	*f = flexInt64(v)
	return nil
}

func unmarshalFlexInt(data []byte, bits int) (int64, error) {
	s := string(data)
	if s == "null" || s == `""` {
		return 0, nil
	}
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		s = s[1 : len(s)-1]
	}
	if s == "" {
		return 0, nil
	}
	v, err := strconv.ParseInt(s, 10, bits)
	if err != nil {
		return 0, fmt.Errorf("invalid int: %q", s)
	}
	return v, nil
}
