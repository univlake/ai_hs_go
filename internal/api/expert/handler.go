// Package expert 实现专家检查模块的 HTTP 入口：
//
//	/expert-inspections                          检查计划 CRUD（含 expert_user_id 三态过滤）
//	/expert-inspection-companies                 受检企业子表 CRUD
//	/expert-inspection-companies/:id/confirm     受检企业确认收到通知
//	/expert-inspection-companies/:id/inspection-notice  POST 生成 / GET 下载 检查通知单
//	/expert-inspection-users                     专家指派子表 CRUD
//	/expert-inspection-users/:id/confirm         专家确认接受指派
//
//	/expert-inspection-notifications/company     受检企业仪表板待办（已到达通知时间未确认）
//	/expert-inspection-notifications/expert      专家端仪表板待办（已指派未确认）
//	/my-check-inspections                        本企业 + 我作为专家 的检查记录合集
//	/my-expert-assignments                       本专家所有被指派的记录
//
// 跟 Laravel ExpertInspectionsController 路由对照见 routes/api/statistics-reporting.php。
package expert

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/middleware"
	"ai_hs_go/internal/model"
	"ai_hs_go/internal/response"
	"ai_hs_go/internal/service"
)

// osRemove 别名 — 单独抽出来便于测试 mock；仅 ExportHazardStatistics 使用 deferred cleanup。
func osRemove(p string) error { return os.Remove(p) }

type Handler struct {
	svc       *service.ExpertService
	rectifySvc *service.RectificationSheetService
}

func NewHandler(svc *service.ExpertService, rectifySvc *service.RectificationSheetService) *Handler {
	return &Handler{svc: svc, rectifySvc: rectifySvc}
}

// ── ExpertInspection（检查计划） ─────────────────────────────────────

type expertInspectionReq struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Comment     *string `json:"comment"`
	StartDate   *string `json:"start_date"`
	EndDate     *string `json:"end_date"`
	CompanyID   *int64  `json:"company_id"`
}

func (h *Handler) ListInspection(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "15"))
	cid, _ := strconv.ParseInt(c.Query("company_id"), 10, 64)
	uid, _ := strconv.ParseInt(c.Query("expert_user_id"), 10, 64)
	rows, total, err := h.svc.ListExpertInspection(c.Request.Context(), service.ExpertInspectionListParams{
		Page: page, PerPage: per, CompanyID: cid, ExpertUserID: uid,
		Title: c.Query("title"),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) GetInspection(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.svc.GetExpertInspection(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err, "专家检查不存在")
		return
	}
	response.Success(c, row)
}

func (h *Handler) CreateInspection(c *gin.Context) {
	var req expertInspectionReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.Title == nil || *req.Title == "" {
		response.BadRequest(c, "title 不能为空")
		return
	}
	if req.CompanyID == nil || *req.CompanyID == 0 {
		response.BadRequest(c, "company_id 不能为空")
		return
	}
	uid := middleware.CurrentUserID(c)
	m := &model.ExpertInspection{
		Title:       *req.Title,
		Description: req.Description,
		Comment:     req.Comment,
		StartDate:   parseDate(req.StartDate),
		EndDate:     parseDate(req.EndDate),
		CompanyID:   *req.CompanyID,
	}
	m.UserID = uid
	m.LastUserID = uid
	if err := h.svc.CreateExpertInspection(c.Request.Context(), m); err != nil {
		writeAppError(c, err, "创建失败")
		return
	}
	response.Created(c, m)
}

func (h *Handler) UpdateInspection(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req expertInspectionReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	updates := map[string]any{"last_user_id": uid}
	if req.Title != nil {
		updates["title"] = *req.Title
	}
	if req.Description != nil {
		updates["description"] = req.Description
	}
	if req.Comment != nil {
		updates["comment"] = req.Comment
	}
	if req.StartDate != nil {
		updates["start_date"] = parseDate(req.StartDate)
	}
	if req.EndDate != nil {
		updates["end_date"] = parseDate(req.EndDate)
	}
	if req.CompanyID != nil {
		updates["company_id"] = *req.CompanyID
	}
	if err := h.svc.UpdateExpertInspection(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err, "更新失败")
		return
	}
	row, _ := h.svc.GetExpertInspection(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteInspection(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteExpertInspection(c.Request.Context(), id); err != nil {
		writeAppError(c, err, "删除失败")
		return
	}
	response.Success(c, gin.H{"message": "删除成功"})
}

// ── ExpertCompany（受检企业） ───────────────────────────────────────

type expertCompanyReq struct {
	ExpertInspectionID *int64  `json:"expert_inspection_id"`
	CheckCompanyID     *int64  `json:"check_company_id"`
	CompanyID          *int64  `json:"company_id"`
	NotifyCompanyTime  *string `json:"notify_company_time"`
	CheckCompanyTime   *string `json:"check_company_time"`
	ConfirmTime        *string `json:"confirm_time"`
	Status             *int    `json:"status"`
}

func (h *Handler) ListCompany(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "15"))
	insID, _ := strconv.ParseInt(c.Query("expert_inspection_id"), 10, 64)
	rows, total, err := h.svc.ListExpertCompany(c.Request.Context(), service.ExpertCompanyListParams{
		Page: page, PerPage: per, ExpertInspectionID: insID,
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) CreateCompany(c *gin.Context) {
	var req expertCompanyReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.ExpertInspectionID == nil || *req.ExpertInspectionID == 0 {
		response.BadRequest(c, "expert_inspection_id 不能为空")
		return
	}
	if req.CheckCompanyID == nil || *req.CheckCompanyID == 0 {
		response.BadRequest(c, "check_company_id 不能为空")
		return
	}
	if req.CompanyID == nil || *req.CompanyID == 0 {
		response.BadRequest(c, "company_id 不能为空")
		return
	}
	uid := middleware.CurrentUserID(c)
	m := &model.ExpertCompany{
		ExpertInspectionID: *req.ExpertInspectionID,
		CheckCompanyID:     *req.CheckCompanyID,
		CompanyID:          *req.CompanyID,
		NotifyCompanyTime:  parseDateTime(req.NotifyCompanyTime),
		CheckCompanyTime:   parseDateTime(req.CheckCompanyTime),
		Status:             1,
	}
	if req.Status != nil {
		m.Status = *req.Status
	}
	m.UserID = uid
	m.LastUserID = uid
	if err := h.svc.CreateExpertCompany(c.Request.Context(), m); err != nil {
		writeAppError(c, err, "创建失败")
		return
	}
	response.Created(c, m)
}

func (h *Handler) UpdateCompany(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req expertCompanyReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	updates := map[string]any{"last_user_id": uid}
	if req.NotifyCompanyTime != nil {
		updates["notify_company_time"] = parseDateTime(req.NotifyCompanyTime)
	}
	if req.CheckCompanyTime != nil {
		updates["check_company_time"] = parseDateTime(req.CheckCompanyTime)
	}
	if req.ConfirmTime != nil {
		updates["confirm_time"] = parseDateTime(req.ConfirmTime)
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}
	if err := h.svc.UpdateExpertCompany(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err, "更新失败")
		return
	}
	row, _ := h.svc.GetExpertCompany(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteCompany(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteExpertCompany(c.Request.Context(), id); err != nil {
		writeAppError(c, err, "删除失败")
		return
	}
	response.Success(c, gin.H{"message": "删除成功"})
}

func (h *Handler) ConfirmCompany(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.svc.ConfirmExpertCompany(c.Request.Context(), id, middleware.CurrentUserID(c))
	if err != nil {
		writeAppError(c, err, "确认失败")
		return
	}
	response.Success(c, row)
}

// ── ExpertInspectionUser（专家指派） ────────────────────────────────

type expertInspectionUserReq struct {
	ExpertUserID    *int64  `json:"expert_user_id"`
	ExpertCompanyID *int64  `json:"expert_company_id"`
	CompanyID       *int64  `json:"company_id"`
	Status          *int    `json:"status"`
	Comment         *string `json:"comment"`
}

func (h *Handler) ListUser(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "15"))
	ecID, _ := strconv.ParseInt(c.Query("expert_company_id"), 10, 64)
	if ecID == 0 {
		response.BadRequest(c, "expert_company_id 不能为空")
		return
	}
	rows, total, err := h.svc.ListExpertInspectionUser(c.Request.Context(), service.ExpertInspectionUserListParams{
		Page: page, PerPage: per, ExpertCompanyID: ecID,
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) CreateUser(c *gin.Context) {
	var req expertInspectionUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.ExpertUserID == nil || *req.ExpertUserID == 0 {
		response.BadRequest(c, "expert_user_id 不能为空")
		return
	}
	if req.ExpertCompanyID == nil || *req.ExpertCompanyID == 0 {
		response.BadRequest(c, "expert_company_id 不能为空")
		return
	}
	if req.CompanyID == nil || *req.CompanyID == 0 {
		response.BadRequest(c, "company_id 不能为空")
		return
	}
	uid := middleware.CurrentUserID(c)
	m := &model.ExpertInspectionUser{
		ExpertUserID:    *req.ExpertUserID,
		ExpertCompanyID: *req.ExpertCompanyID,
		CompanyID:       *req.CompanyID,
		Comment:         req.Comment,
		Status:          1,
	}
	if req.Status != nil {
		m.Status = *req.Status
	}
	m.UserID = uid
	m.LastUserID = uid
	if err := h.svc.CreateExpertInspectionUser(c.Request.Context(), m); err != nil {
		writeAppError(c, err, "创建失败")
		return
	}
	response.Created(c, m)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req expertInspectionUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	updates := map[string]any{"last_user_id": uid}
	if req.Status != nil {
		updates["status"] = *req.Status
	}
	if req.Comment != nil {
		updates["comment"] = req.Comment
	}
	if err := h.svc.UpdateExpertInspectionUser(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err, "更新失败")
		return
	}
	response.Success(c, gin.H{"message": "ok"})
}

func (h *Handler) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteExpertInspectionUser(c.Request.Context(), id); err != nil {
		writeAppError(c, err, "删除失败")
		return
	}
	response.Success(c, gin.H{"message": "删除成功"})
}

func (h *Handler) ConfirmUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.svc.ConfirmExpertInspectionUser(c.Request.Context(), id, middleware.CurrentUserID(c))
	if err != nil {
		writeAppError(c, err, "确认失败")
		return
	}
	response.Success(c, row)
}

// ── 仪表板通知与「我的检查/指派」 ───────────────────────────────────

func (h *Handler) CompanyNotifications(c *gin.Context) {
	cid, _ := strconv.ParseInt(c.Query("company_id"), 10, 64)
	if cid == 0 {
		response.BadRequest(c, "company_id 不能为空")
		return
	}
	rows, err := h.svc.CompanyNotifications(c.Request.Context(), cid)
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Success(c, rows)
}

func (h *Handler) ExpertNotifications(c *gin.Context) {
	uid, _ := strconv.ParseInt(c.Query("expert_user_id"), 10, 64)
	if uid == 0 {
		uid = middleware.CurrentUserID(c)
	}
	rows, err := h.svc.ExpertNotifications(c.Request.Context(), uid)
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Success(c, rows)
}

func (h *Handler) MyCheckInspections(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "50"))
	cid, _ := strconv.ParseInt(c.Query("company_id"), 10, 64)
	if cid == 0 {
		response.BadRequest(c, "company_id 不能为空")
		return
	}
	uid, _ := strconv.ParseInt(c.Query("expert_user_id"), 10, 64)
	if uid == 0 {
		uid, _ = strconv.ParseInt(c.Query("user_id"), 10, 64)
	}
	rows, total, err := h.svc.MyCheckInspections(c.Request.Context(), service.MyCheckInspectionsParams{
		Page: page, PerPage: per, CompanyID: cid, ExpertUserID: uid,
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) MyExpertAssignments(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "50"))
	uid, _ := strconv.ParseInt(c.Query("expert_user_id"), 10, 64)
	if uid == 0 {
		uid = middleware.CurrentUserID(c)
	}
	rows, total, err := h.svc.MyExpertAssignments(c.Request.Context(), service.MyExpertAssignmentsParams{
		Page: page, PerPage: per, UserID: uid,
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

// ── 检查通知单 ──────────────────────────────────────────────────────

func (h *Handler) GenerateInspectionNotice(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	userCompanyID, _ := middleware.ScopeCompanyID(c)
	if userCompanyID == 0 {
		response.Forbidden(c, "缺少企业信息")
		return
	}
	row, err := h.svc.GenerateInspectionNotice(c.Request.Context(), id, userCompanyID)
	if err != nil {
		writeAppError(c, err, "生成失败")
		return
	}
	response.Success(c, row)
}

func (h *Handler) DownloadInspectionNotice(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	userID := middleware.CurrentUserID(c)
	userCompanyID, _ := middleware.ScopeCompanyID(c)
	abs, name, err := h.svc.InspectionNoticeFile(c.Request.Context(), id, userID, userCompanyID)
	if err != nil {
		writeAppError(c, err, "下载失败")
		return
	}
	c.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, name))
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
	c.File(abs)
}

// ── 整改通知单 (.xlsx) + 隐患分类统计 (.xlsx) ─────────────────────────

// 校验权限：计划方 / 已指派专家可生成；下载多了一个「受检企业本身」可下。
func (h *Handler) loadEcAndCheck(c *gin.Context, allowChecked bool) (*model.ExpertCompany, error) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return nil, errs.New(errs.CodeUnprocessable, "无效 id")
	}
	abs, _, ec, err := h.rectifySvc.RectificationSheetFile(c.Request.Context(), id)
	_ = abs
	// 这里我们只用 ec；err 不影响权限校验逻辑（即使尚未生成也要拿到 ec 决定权限）。
	if ec == nil {
		// 重新拉一次 ec（File 接口在 ec 不存在时返回 nil）
		ecRow, err2 := h.svc.GetExpertCompany(c.Request.Context(), id)
		if err2 != nil {
			return nil, err2
		}
		ec = ecRow
	}
	_ = err
	uid := middleware.CurrentUserID(c)
	userCompanyID, _ := middleware.ScopeCompanyID(c)
	planCompanyID := int64(0)
	if ec.ExpertInspection != nil {
		planCompanyID = ec.ExpertInspection.CompanyID
	}
	isOwner := planCompanyID != 0 && planCompanyID == userCompanyID
	isExpert := false
	for _, u := range ec.ExpertInspectionUsers {
		if u.ExpertUserID == uid {
			isExpert = true
			break
		}
	}
	isChecked := allowChecked && ec.CheckCompanyID == userCompanyID
	if !isOwner && !isExpert && !isChecked {
		return nil, errs.New(errs.CodeForbidden, "无权限")
	}
	return ec, nil
}

// GenerateRectificationSheet — POST /expert-inspection-companies/:id/rectification-sheet
func (h *Handler) GenerateRectificationSheet(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if _, err := h.loadEcAndCheck(c, false); err != nil {
		writeAppError(c, err, "权限校验失败")
		return
	}
	row, err := h.rectifySvc.GenerateRectificationSheet(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err, "生成整改通知单失败")
		return
	}
	response.Success(c, row)
}

// DownloadRectificationSheet — GET /expert-inspection-companies/:id/rectification-sheet
func (h *Handler) DownloadRectificationSheet(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if _, err := h.loadEcAndCheck(c, true); err != nil {
		writeAppError(c, err, "权限校验失败")
		return
	}
	abs, name, _, err := h.rectifySvc.RectificationSheetFile(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err, "下载失败")
		return
	}
	c.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, name))
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.File(abs)
}

// ExportHazardStatistics — GET /hazard-statistics
// 仅本企业（user.company_id == query.company_id）可导出。
func (h *Handler) ExportHazardStatistics(c *gin.Context) {
	cid, _ := strconv.ParseInt(c.Query("company_id"), 10, 64)
	if cid == 0 {
		response.BadRequest(c, "company_id 不能为空")
		return
	}
	userCompanyID, _ := middleware.ScopeCompanyID(c)
	if cid != userCompanyID {
		response.Forbidden(c, "仅可导出本企业的隐患分类统计")
		return
	}
	abs, err := h.rectifySvc.GenerateHazardStatistics(c.Request.Context(), service.HazardStatisticsParams{
		CompanyID: cid,
		StartTime: c.Query("start_time"),
		EndTime:   c.Query("end_time"),
		Type:      c.Query("type"),
	})
	if err != nil {
		writeAppError(c, err, "导出失败")
		return
	}
	defer func() { _ = osRemove(abs) }()
	c.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filepath.Base(abs)))
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.File(abs)
}

// ── helpers ─────────────────────────────────────────────────────────

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
		"2006-01-02T15:04",
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

func writeAppError(c *gin.Context, err error, fallback string) {
	if ae, ok := errs.As(err); ok {
		switch ae.Code {
		case errs.CodeNotFound:
			response.NotFound(c, ae.Message)
		case errs.CodeUnprocessable:
			response.Failed(c, http.StatusUnprocessableEntity, ae.Message)
		case errs.CodeForbidden:
			response.Forbidden(c, ae.Message)
		case errs.CodeConflict:
			response.Failed(c, http.StatusConflict, ae.Message)
		default:
			response.Internal(c, ae.Message)
		}
		return
	}
	response.Internal(c, fallback)
}
