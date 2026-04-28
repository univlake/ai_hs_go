// Package extramodules 实现剩余 5 个模块的 HTTP 入口：
//
//	/change-manages          变更管理 CRUD
//	/regulation-libraries    法律法规库 CRUD（preload regulation_type）
//	/training-files          培训资料·文件 CRUD（show 走 OSS 签名 + 下载流水自增）
//	/training-videos         培训资料·视频 CRUD（show 走 OSS 签名 + 观看流水自增）
//	/risks/portrait          单企业风险画像（inherent + float + risk_log）
//	/risks/log               单企业风险日志分页
//
// 跟 process 包同模式：所有字段用 *T，前端传 null 直接清空字段；日期字符串走 parseDate；
// company 上下文统一从 middleware.ScopeCompanyID 取（前端可显式传 query.company_id 兜底）。
package extramodules

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

type Handler struct {
	svc *service.ExtraModulesService
	oss *storage.Client // optional — nil 时 training-files/videos 的 show 接口直接返回 raw key
}

func NewHandler(svc *service.ExtraModulesService, oss *storage.Client) *Handler {
	return &Handler{svc: svc, oss: oss}
}

// ── 变更管理 ───────────────────────────────────────────────────────────

type changeManageReq struct {
	Title                   *string `json:"title"`
	DepartmentName          *string `json:"department_name"`
	ApplyDate               *string `json:"apply_date"`
	Reason                  *string `json:"reason"`
	ApplyLeaderOpinion      *string `json:"apply_leader_opinion"`
	ApplyLeaderOpinionDate  *string `json:"apply_leader_opinion_date"`
	RiskAnalyze             *string `json:"risk_analyze"`
	ManageLeaderOpinion     *string `json:"manage_leader_opinion"`
	ManageLeaderOpinionDate *string `json:"manage_leader_opinion_date"`
	Approve                 *string `json:"approve"`
	ApproveDate             *string `json:"approve_date"`
	AcceptanceInfo          *string `json:"acceptance_info"`
	AcceptanceInfoDate      *string `json:"acceptance_info_date"`
}

func (h *Handler) ListChangeManage(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "15"))
	rows, total, err := h.svc.ListChangeManage(c.Request.Context(), service.ChangeManageListParams{
		Page: page, PerPage: per, CompanyID: companyID,
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) GetChangeManage(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.svc.GetChangeManage(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

func (h *Handler) CreateChangeManage(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	var req changeManageReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	m := &model.ChangeManage{
		Title:                   req.Title,
		DepartmentName:          req.DepartmentName,
		ApplyDate:               parseDate(req.ApplyDate),
		Reason:                  req.Reason,
		ApplyLeaderOpinion:      req.ApplyLeaderOpinion,
		ApplyLeaderOpinionDate:  parseDate(req.ApplyLeaderOpinionDate),
		RiskAnalyze:             req.RiskAnalyze,
		ManageLeaderOpinion:     req.ManageLeaderOpinion,
		ManageLeaderOpinionDate: parseDate(req.ManageLeaderOpinionDate),
		Approve:                 req.Approve,
		ApproveDate:             parseDate(req.ApproveDate),
		AcceptanceInfo:          req.AcceptanceInfo,
		AcceptanceInfoDate:      parseDate(req.AcceptanceInfoDate),
		CompanyID:               companyID,
	}
	m.UserID = uid
	m.LastUserID = uid
	if err := h.svc.CreateChangeManage(c.Request.Context(), m); err != nil {
		writeAppError(c, err)
		return
	}
	response.Created(c, m)
}

func (h *Handler) UpdateChangeManage(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req changeManageReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	updates := map[string]any{"last_user_id": uid}
	addNullableString(updates, "title", req.Title)
	addNullableString(updates, "department_name", req.DepartmentName)
	addNullableString(updates, "reason", req.Reason)
	addNullableString(updates, "apply_leader_opinion", req.ApplyLeaderOpinion)
	addNullableString(updates, "risk_analyze", req.RiskAnalyze)
	addNullableString(updates, "manage_leader_opinion", req.ManageLeaderOpinion)
	addNullableString(updates, "approve", req.Approve)
	addNullableString(updates, "acceptance_info", req.AcceptanceInfo)
	if req.ApplyDate != nil {
		updates["apply_date"] = parseDate(req.ApplyDate)
	}
	if req.ApplyLeaderOpinionDate != nil {
		updates["apply_leader_opinion_date"] = parseDate(req.ApplyLeaderOpinionDate)
	}
	if req.ManageLeaderOpinionDate != nil {
		updates["manage_leader_opinion_date"] = parseDate(req.ManageLeaderOpinionDate)
	}
	if req.ApproveDate != nil {
		updates["approve_date"] = parseDate(req.ApproveDate)
	}
	if req.AcceptanceInfoDate != nil {
		updates["acceptance_info_date"] = parseDate(req.AcceptanceInfoDate)
	}
	if err := h.svc.UpdateChangeManage(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetChangeManage(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteChangeManage(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteChangeManage(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

// ── 法律法规库 ─────────────────────────────────────────────────────────

type regulationLibraryReq struct {
	Title                    *string    `json:"title"`
	RegulationLibraryTypeID  *flexInt64 `json:"regulation_library_type_id"`
	EstablishedDate          *string    `json:"established_date"`
	EffectiveDate            *string    `json:"effective_date"`
	PromulgatingDepartment   *string    `json:"promulgating_department"`
	StandardEncoding         *string    `json:"standard_encoding"`
	IdentifyDate             *string    `json:"identify_date"`
	Status                   *string    `json:"status"`
	AnnexURL                 *string    `json:"annex_url"`
	CompanyID                *flexInt64 `json:"company_id"`
}

func (h *Handler) ListRegulationLibrary(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "15"))
	params := service.RegulationLibraryListParams{
		Page: page, PerPage: per,
		Title: c.Query("title"),
	}
	if v := c.Query("regulation_library_type_id"); v != "" {
		if id, err := strconv.ParseInt(v, 10, 64); err == nil {
			params.RegulationLibraryTypeID = &id
		}
	}
	if v := c.Query("type_parent_id"); v != "" {
		if id, err := strconv.ParseInt(v, 10, 64); err == nil {
			params.TypeParentID = &id
		}
	}
	rows, total, err := h.svc.ListRegulationLibrary(c.Request.Context(), params)
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) GetRegulationLibrary(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.svc.GetRegulationLibrary(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

func (h *Handler) CreateRegulationLibrary(c *gin.Context) {
	var req regulationLibraryReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.Title == nil || *req.Title == "" {
		response.BadRequest(c, "标题不能为空")
		return
	}
	uid := middleware.CurrentUserID(c)
	m := &model.RegulationLibrary{
		Title:                   *req.Title,
		RegulationLibraryTypeID: req.RegulationLibraryTypeID.Int64Ptr(),
		EstablishedDate:         parseDate(req.EstablishedDate),
		EffectiveDate:           parseDate(req.EffectiveDate),
		PromulgatingDepartment:  req.PromulgatingDepartment,
		StandardEncoding:        req.StandardEncoding,
		IdentifyDate:            parseDate(req.IdentifyDate),
		Status:                  req.Status,
		AnnexURL:                req.AnnexURL,
	}
	if cid := req.CompanyID.Int64Ptr(); cid != nil {
		m.CompanyID = *cid
	}
	m.UserID = uid
	m.LastUserID = uid
	if err := h.svc.CreateRegulationLibrary(c.Request.Context(), m); err != nil {
		writeAppError(c, err)
		return
	}
	response.Created(c, m)
}

func (h *Handler) UpdateRegulationLibrary(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req regulationLibraryReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	updates := map[string]any{"last_user_id": uid}
	addNullableString(updates, "title", req.Title)
	addNullableString(updates, "promulgating_department", req.PromulgatingDepartment)
	addNullableString(updates, "standard_encoding", req.StandardEncoding)
	addNullableString(updates, "status", req.Status)
	addNullableString(updates, "annex_url", req.AnnexURL)
	if req.RegulationLibraryTypeID != nil {
		updates["regulation_library_type_id"] = req.RegulationLibraryTypeID.Int64Ptr()
	}
	if req.EstablishedDate != nil {
		updates["established_date"] = parseDate(req.EstablishedDate)
	}
	if req.EffectiveDate != nil {
		updates["effective_date"] = parseDate(req.EffectiveDate)
	}
	if req.IdentifyDate != nil {
		updates["identify_date"] = parseDate(req.IdentifyDate)
	}
	if req.CompanyID != nil {
		if cid := req.CompanyID.Int64Ptr(); cid != nil {
			updates["company_id"] = *cid
		}
	}
	if err := h.svc.UpdateRegulationLibrary(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetRegulationLibrary(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteRegulationLibrary(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteRegulationLibrary(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

// ── 培训资料·文件 ─────────────────────────────────────────────────────

type trainingFileReq struct {
	Title   *string `json:"title"`
	FileURL *string `json:"file_url"`
}

func (h *Handler) ListTrainingFile(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "15"))
	parent, _ := strconv.ParseInt(c.Query("parent_company_id"), 10, 64)
	rows, total, err := h.svc.ListTrainingFile(c.Request.Context(), service.TrainingFileListParams{
		Page: page, PerPage: per, CompanyID: companyID,
		ParentCompanyID: parent,
		Title:           c.Query("title"),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

// GetTrainingFile — Laravel show 在这条接口里同步做：自增 times_download +
// 写入 training_file_logs + 返回 OSS 签名 URL。前端拿到 URL 后 window.open 直接下载。
func (h *Handler) GetTrainingFile(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.svc.GetTrainingFile(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err)
		return
	}
	uid := middleware.CurrentUserID(c)
	cid, _ := middleware.ScopeCompanyID(c)
	// 计数 + 流水写失败不影响下载（沿 Laravel 隐式忽略错误）。
	_ = h.svc.IncrementTrainingFileDownload(c.Request.Context(), id, uid, cid)
	response.Success(c, h.signedOrRaw(row.FileURL))
}

func (h *Handler) CreateTrainingFile(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	var req trainingFileReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	m := &model.TrainingFile{
		Title:     req.Title,
		FileURL:   req.FileURL,
		CompanyID: companyID,
	}
	m.UserID = uid
	m.LastUserID = uid
	if err := h.svc.CreateTrainingFile(c.Request.Context(), m); err != nil {
		writeAppError(c, err)
		return
	}
	response.Created(c, m)
}

func (h *Handler) UpdateTrainingFile(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req trainingFileReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	updates := map[string]any{"last_user_id": uid}
	addNullableString(updates, "title", req.Title)
	addNullableString(updates, "file_url", req.FileURL)
	if err := h.svc.UpdateTrainingFile(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetTrainingFile(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteTrainingFile(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteTrainingFile(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

// ── 培训资料·视频 ─────────────────────────────────────────────────────

type trainingVideoReq struct {
	Title     *string `json:"title"`
	Introduce *string `json:"introduce"`
	ThumbURL  *string `json:"thumb_url"`
	VideoURL  *string `json:"video_url"`
	Duration  *string `json:"duration"`
	Lector    *string `json:"lector"`
}

func (h *Handler) ListTrainingVideo(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "15"))
	parent, _ := strconv.ParseInt(c.Query("parent_company_id"), 10, 64)
	rows, total, err := h.svc.ListTrainingVideo(c.Request.Context(), service.TrainingVideoListParams{
		Page: page, PerPage: per, CompanyID: companyID,
		ParentCompanyID: parent,
		Title:           c.Query("title"),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) GetTrainingVideo(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.svc.GetTrainingVideo(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err)
		return
	}
	uid := middleware.CurrentUserID(c)
	cid, _ := middleware.ScopeCompanyID(c)
	_ = h.svc.IncrementTrainingVideoWatch(c.Request.Context(), id, uid, cid)
	response.Success(c, h.signedOrRaw(row.VideoURL))
}

func (h *Handler) CreateTrainingVideo(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	var req trainingVideoReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	m := &model.TrainingVideo{
		Title:     req.Title,
		Introduce: req.Introduce,
		ThumbURL:  req.ThumbURL,
		VideoURL:  req.VideoURL,
		Duration:  req.Duration,
		Lector:    req.Lector,
		CompanyID: companyID,
	}
	m.UserID = uid
	m.LastUserID = uid
	if err := h.svc.CreateTrainingVideo(c.Request.Context(), m); err != nil {
		writeAppError(c, err)
		return
	}
	response.Created(c, m)
}

func (h *Handler) UpdateTrainingVideo(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req trainingVideoReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	updates := map[string]any{"last_user_id": uid}
	addNullableString(updates, "title", req.Title)
	addNullableString(updates, "introduce", req.Introduce)
	addNullableString(updates, "thumb_url", req.ThumbURL)
	addNullableString(updates, "video_url", req.VideoURL)
	addNullableString(updates, "duration", req.Duration)
	addNullableString(updates, "lector", req.Lector)
	if err := h.svc.UpdateTrainingVideo(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetTrainingVideo(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteTrainingVideo(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteTrainingVideo(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

// ── 风险画像 / 风险日志 ─────────────────────────────────────────────────

func (h *Handler) RisksPortrait(c *gin.Context) {
	cid := scopeOrCompanyID(c)
	if cid == 0 {
		response.BadRequest(c, "缺少 company_id")
		return
	}
	data, err := h.svc.CompanyPortrait(c.Request.Context(), cid)
	if err != nil {
		writeAppError(c, err)
		return
	}
	if data == nil {
		// 跟 Laravel 一致：固有或浮动任一缺失就返回空数组，前端 if (!log) 分支处理。
		response.Success(c, []any{})
		return
	}
	response.Success(c, data)
}

func (h *Handler) RisksLog(c *gin.Context) {
	cid := scopeOrCompanyID(c)
	if cid == 0 {
		response.BadRequest(c, "缺少 company_id")
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "15"))
	rows, total, err := h.svc.ListCompanyRiskLogs(c.Request.Context(), service.RiskLogListParams{
		Page: page, PerPage: per, CompanyID: cid,
	})
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Page(c, rows, page, per, total)
}

// ── helpers ─────────────────────────────────────────────────────────────

func (h *Handler) signedOrRaw(key *string) string {
	if key == nil || *key == "" {
		return ""
	}
	if h.oss == nil {
		return *key
	}
	signed, err := h.oss.SignURL(*key, time.Hour, nil)
	if err != nil {
		return *key
	}
	return signed
}

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

// flexInt64 — 容忍前端 <select> 把数字以字符串提交：""/null 当作未传值。
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
