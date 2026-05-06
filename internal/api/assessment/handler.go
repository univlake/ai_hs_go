// Package assessment 实现 Laravel inspections.php 路由的 HTTP 入口：
//
//	/daily-check-templates          模板树 CRUD
//	/daily-checks/...                填报流（startForm / saveDraft / complete / score / submission / destroyForm）
//	/accident-check-templates       同上
//	/accident-checks/...             直接评分（startForm / saveScore / submissions）
//	/add-points-templates           同上
//	/add-points/...                  直接评分（同 accident-checks 形状，但 startForm 仅 level=1）
//	/system-checks/...               活动参与及材料上报（CRUD + compute + submissions）
//	/norm-checks/...                 安全生产标准化（CRUD + saveScore + submissions + templateLink）
//	/assessment-templates/years      三表年度并集
//	/assessment-templates/copy       一键年度拷贝
//	/assessment-summary              6 维度汇总
//	/assessment-summary/export       Word 导出
//
// 跟 hiddendanger / accidentemergency handler 行为对齐：
//   - JSON 解析失败一律 BadRequest；errs.AppError → response 对应状态码。
//   - List 分页统一走 response.Page。
//   - URL 携带的 :id / company_id / year 缺省时按 0 处理（service 层会忽略）。
package assessment

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/middleware"
	"ai_hs_go/internal/model"
	"ai_hs_go/internal/response"
	"ai_hs_go/internal/service"
	"ai_hs_go/internal/api/apiutil"
)

// Handler 持有 service + 模板/落盘根目录。
type Handler struct {
	svc         *service.AssessmentService
	templateDir string
	storageDir  string
}

func NewHandler(svc *service.AssessmentService, templateDir, storageDir string) *Handler {
	return &Handler{svc: svc, templateDir: templateDir, storageDir: storageDir}
}

// ── 通用 helpers ───────────────────────────────────────────────────

func intQuery(c *gin.Context, key string) int {
	v, _ := strconv.Atoi(c.Query(key))
	return v
}

func int64Query(c *gin.Context, key string) int64 {
	v, _ := strconv.ParseInt(c.Query(key), 10, 64)
	return v
}

func parseID(c *gin.Context) (int64, bool) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		response.BadRequest(c, "无效 id")
		return 0, false
	}
	return id, true
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

// 三张模板表的写入 payload 共形：title + level/sort/pid/year/score_range。
type templateReq struct {
	Title      *string `json:"title"`
	Level      *int    `json:"level"`
	Sort       *int    `json:"sort"`
	Pid        *int64  `json:"pid"`
	Year       *int    `json:"year"`
	ScoreRange *string `json:"score_range"`
}

func (r templateReq) require(c *gin.Context) bool {
	if r.Title == nil || strings.TrimSpace(*r.Title) == "" {
		response.BadRequest(c, "title 不能为空")
		return false
	}
	if r.Year == nil || *r.Year == 0 {
		response.BadRequest(c, "year 不能为空")
		return false
	}
	return true
}

func (r templateReq) updates() map[string]any {
	m := make(map[string]any)
	if r.Title != nil {
		m["title"] = *r.Title
	}
	if r.Level != nil {
		m["level"] = *r.Level
	}
	if r.Sort != nil {
		m["sort"] = *r.Sort
	}
	if r.Pid != nil {
		m["pid"] = *r.Pid
	}
	if r.Year != nil {
		m["year"] = *r.Year
	}
	if r.ScoreRange != nil {
		m["score_range"] = *r.ScoreRange
	}
	return m
}

// ── 日常检查模板 ───────────────────────────────────────────────────

func (h *Handler) ListDailyCheckTemplates(c *gin.Context) {
	year := intQuery(c, "year")
	rows, err := h.svc.ListDailyCheckTemplates(c.Request.Context(), year)
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Success(c, rows)
}

func (h *Handler) GetDailyCheckTemplate(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	row, err := h.svc.GetDailyCheckTemplate(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err, "日常检查模板不存在")
		return
	}
	response.Success(c, row)
}

func (h *Handler) CreateDailyCheckTemplate(c *gin.Context) {
	var req templateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if !req.require(c) {
		return
	}
	rec := &model.DailyCheckTemplate{Title: *req.Title, Year: *req.Year, Level: 1, ScoreRange: "1"}
	if req.Level != nil {
		rec.Level = *req.Level
	}
	if req.Sort != nil {
		rec.Sort = *req.Sort
	}
	if req.Pid != nil {
		rec.Pid = *req.Pid
	}
	if req.ScoreRange != nil {
		rec.ScoreRange = *req.ScoreRange
	}
	if err := h.svc.CreateDailyCheckTemplate(c.Request.Context(), rec); err != nil {
		writeAppError(c, err, "创建失败")
		return
	}
	response.Created(c, rec)
}

func (h *Handler) UpdateDailyCheckTemplate(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	var req templateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if err := h.svc.UpdateDailyCheckTemplate(c.Request.Context(), id, req.updates()); err != nil {
		writeAppError(c, err, "更新失败")
		return
	}
	row, err := h.svc.GetDailyCheckTemplate(c.Request.Context(), id)
	if err != nil {
		response.Success(c, gin.H{"id": id})
		return
	}
	response.Success(c, row)
}

func (h *Handler) DeleteDailyCheckTemplate(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	if err := h.svc.DeleteDailyCheckTemplate(c.Request.Context(), id); err != nil {
		writeAppError(c, err, "删除失败")
		return
	}
	response.Success(c, gin.H{"message": "删除成功"})
}

// ── 日常检查（filing 流） ─────────────────────────────────────────

func (h *Handler) ListDailySubmissions(c *gin.Context) {
	pid := int64Query(c, "pid")
	companyID := int64Query(c, "company_id")
	if pid == 0 && companyID == 0 {
		response.Failed(c, http.StatusUnprocessableEntity, "company_id 或 pid 至少提供一个")
		return
	}
	year := intQuery(c, "year")
	var status *int
	if v := c.Query("status"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			status = &n
		}
	}
	rows, err := h.svc.ListDailySubmissions(c.Request.Context(), pid, companyID, status, year)
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Success(c, rows)
}

func (h *Handler) StartDailyForm(c *gin.Context) {
	companyID := int64Query(c, "company_id")
	year := intQuery(c, "year")
	if companyID == 0 || year == 0 {
		response.Failed(c, http.StatusUnprocessableEntity, "company_id / year 不能为空")
		return
	}
	isNew := c.Query("is_new") == "1" || c.Query("is_new") == "true"
	rows, err := h.svc.StartDailyForm(c.Request.Context(), companyID, year, isNew)
	if err != nil {
		writeAppError(c, err, "查询失败")
		return
	}
	response.Success(c, rows)
}

type dailyDraftReq struct {
	CompanyID int64 `json:"company_id"`
	Year      int   `json:"year"`
	Items     []struct {
		DailyCheckTemplateID int64   `json:"daily_check_template_id"`
		Path                 *string `json:"path"`
	} `json:"items"`
}

func (h *Handler) SaveDailyDraft(c *gin.Context) {
	var req dailyDraftReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.CompanyID == 0 || req.Year == 0 {
		response.Failed(c, http.StatusUnprocessableEntity, "company_id / year 不能为空")
		return
	}
	uid := middleware.CurrentUserID(c)
	items := make([]service.DailyDraftItem, 0, len(req.Items))
	for _, it := range req.Items {
		items = append(items, service.DailyDraftItem{DailyCheckTemplateID: it.DailyCheckTemplateID, Path: it.Path})
	}
	if err := h.svc.SaveDailyDraft(c.Request.Context(), req.CompanyID, req.Year, uid, items); err != nil {
		writeAppError(c, err, "暂存失败")
		return
	}
	response.Success(c, gin.H{"message": "暂存成功"})
}

type yearAndCompanyReq struct {
	CompanyID int64 `json:"company_id"`
	Year      int   `json:"year"`
}

func (h *Handler) CompleteDaily(c *gin.Context) {
	var req yearAndCompanyReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.CompanyID == 0 || req.Year == 0 {
		response.Failed(c, http.StatusUnprocessableEntity, "company_id / year 不能为空")
		return
	}
	uid := middleware.CurrentUserID(c)
	if err := h.svc.CompleteDaily(c.Request.Context(), req.CompanyID, req.Year, uid); err != nil {
		writeAppError(c, err, "提交失败")
		return
	}
	response.Success(c, gin.H{"message": "填报已提交，等待上级评分"})
}

type dailyScoreReq struct {
	CompanyID int64 `json:"company_id"`
	Year      int   `json:"year"`
	Items     []struct {
		DailyCheckTemplateID int64    `json:"daily_check_template_id"`
		Number               *float64 `json:"number"`
		Comment              *string  `json:"comment"`
	} `json:"items"`
}

func (h *Handler) ScoreDaily(c *gin.Context) {
	var req dailyScoreReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.CompanyID == 0 || req.Year == 0 {
		response.Failed(c, http.StatusUnprocessableEntity, "company_id / year 不能为空")
		return
	}
	uid := middleware.CurrentUserID(c)
	items := make([]service.DailyScoreItem, 0, len(req.Items))
	for _, it := range req.Items {
		items = append(items, service.DailyScoreItem{DailyCheckTemplateID: it.DailyCheckTemplateID, Number: it.Number, Comment: it.Comment})
	}
	if err := h.svc.ScoreDaily(c.Request.Context(), req.CompanyID, req.Year, uid, items); err != nil {
		writeAppError(c, err, "评分失败")
		return
	}
	response.Success(c, gin.H{"message": "评分完成"})
}

func (h *Handler) ShowDailySubmission(c *gin.Context) {
	companyID := int64Query(c, "company_id")
	year := intQuery(c, "year")
	if companyID == 0 || year == 0 {
		response.Failed(c, http.StatusUnprocessableEntity, "company_id / year 不能为空")
		return
	}
	row, err := h.svc.ShowDailySubmission(c.Request.Context(), companyID, year)
	if err != nil {
		writeAppError(c, err, "查询失败")
		return
	}
	response.Success(c, row)
}

func (h *Handler) DestroyDailyForm(c *gin.Context) {
	var req yearAndCompanyReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.CompanyID == 0 || req.Year == 0 {
		response.Failed(c, http.StatusUnprocessableEntity, "company_id / year 不能为空")
		return
	}
	if err := h.svc.DestroyDailyForm(c.Request.Context(), req.CompanyID, req.Year); err != nil {
		writeAppError(c, err, "删除失败")
		return
	}
	response.Success(c, gin.H{"message": "删除成功"})
}

func (h *Handler) ListDailyChecks(c *gin.Context) {
	page, per := apiutil.ParsePagination(c, 15)
	rows, total, err := h.svc.ListDailyChecks(c.Request.Context(), page, per, int64Query(c, "company_id"), intQuery(c, "year"))
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) GetDailyCheck(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	row, err := h.svc.GetDailyCheck(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err, "日常检查记录不存在")
		return
	}
	response.Success(c, row)
}

// ── 事故检查模板 ───────────────────────────────────────────────────

func (h *Handler) ListAccidentCheckTemplates(c *gin.Context) {
	year := intQuery(c, "year")
	rows, err := h.svc.ListAccidentCheckTemplates(c.Request.Context(), year)
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Success(c, rows)
}

func (h *Handler) GetAccidentCheckTemplate(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	row, err := h.svc.GetAccidentCheckTemplate(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err, "事故检查模板不存在")
		return
	}
	response.Success(c, row)
}

func (h *Handler) CreateAccidentCheckTemplate(c *gin.Context) {
	var req templateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if !req.require(c) {
		return
	}
	rec := &model.AccidentCheckTemplate{Title: *req.Title, Year: *req.Year, Level: 1, ScoreRange: "1"}
	if req.Level != nil {
		rec.Level = *req.Level
	}
	if req.Sort != nil {
		rec.Sort = *req.Sort
	}
	if req.Pid != nil {
		rec.Pid = *req.Pid
	}
	if req.ScoreRange != nil {
		rec.ScoreRange = *req.ScoreRange
	}
	if err := h.svc.CreateAccidentCheckTemplate(c.Request.Context(), rec); err != nil {
		writeAppError(c, err, "创建失败")
		return
	}
	response.Created(c, rec)
}

func (h *Handler) UpdateAccidentCheckTemplate(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	var req templateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if err := h.svc.UpdateAccidentCheckTemplate(c.Request.Context(), id, req.updates()); err != nil {
		writeAppError(c, err, "更新失败")
		return
	}
	row, err := h.svc.GetAccidentCheckTemplate(c.Request.Context(), id)
	if err != nil {
		response.Success(c, gin.H{"id": id})
		return
	}
	response.Success(c, row)
}

func (h *Handler) DeleteAccidentCheckTemplate(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	if err := h.svc.DeleteAccidentCheckTemplate(c.Request.Context(), id); err != nil {
		writeAppError(c, err, "删除失败")
		return
	}
	response.Success(c, gin.H{"message": "删除成功"})
}

// ── 事故检查（一次性评分） ────────────────────────────────────────

func (h *Handler) ListAccidentSubmissions(c *gin.Context) {
	pid := int64Query(c, "pid")
	if pid == 0 {
		response.Failed(c, http.StatusUnprocessableEntity, "pid 不能为空")
		return
	}
	rows, err := h.svc.ListAccidentSubmissions(c.Request.Context(), pid, intQuery(c, "year"))
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Success(c, rows)
}

func (h *Handler) StartAccidentForm(c *gin.Context) {
	companyID := int64Query(c, "company_id")
	year := intQuery(c, "year")
	if companyID == 0 || year == 0 {
		response.Failed(c, http.StatusUnprocessableEntity, "company_id / year 不能为空")
		return
	}
	rows, err := h.svc.StartAccidentForm(c.Request.Context(), companyID, year)
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Success(c, rows)
}

type accidentScoreReq struct {
	CompanyID int64 `json:"company_id"`
	Year      int   `json:"year"`
	Items     []struct {
		AccidentCheckTemplateID int64    `json:"accident_check_template_id"`
		Number                  *float64 `json:"number"`
		Comment                 *string  `json:"comment"`
	} `json:"items"`
}

func (h *Handler) SaveAccidentScore(c *gin.Context) {
	var req accidentScoreReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.CompanyID == 0 || req.Year == 0 {
		response.Failed(c, http.StatusUnprocessableEntity, "company_id / year 不能为空")
		return
	}
	uid := middleware.CurrentUserID(c)
	items := make([]service.AccidentScoreItem, 0, len(req.Items))
	for _, it := range req.Items {
		items = append(items, service.AccidentScoreItem{AccidentCheckTemplateID: it.AccidentCheckTemplateID, Number: it.Number, Comment: it.Comment})
	}
	if err := h.svc.SaveAccidentScore(c.Request.Context(), req.CompanyID, req.Year, uid, items); err != nil {
		writeAppError(c, err, "评分失败")
		return
	}
	response.Success(c, gin.H{"message": "评分已保存"})
}

func (h *Handler) ListAccidentChecks(c *gin.Context) {
	page, per := apiutil.ParsePagination(c, 15)
	rows, total, err := h.svc.ListAccidentChecks(c.Request.Context(), page, per, int64Query(c, "company_id"), intQuery(c, "year"))
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) GetAccidentCheck(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	row, err := h.svc.GetAccidentCheck(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err, "事故检查记录不存在")
		return
	}
	response.Success(c, row)
}

type accidentCheckReq struct {
	AccidentCheckTemplateID *int64   `json:"accident_check_template_id"`
	Path                    *string  `json:"path"`
	Number                  *float64 `json:"number"`
	Comment                 *string  `json:"comment"`
	CompanyID               *int64   `json:"company_id"`
	Year                    *int     `json:"year"`
}

func (h *Handler) CreateAccidentCheck(c *gin.Context) {
	var req accidentCheckReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.AccidentCheckTemplateID == nil || req.CompanyID == nil || req.Year == nil {
		response.Failed(c, http.StatusUnprocessableEntity, "accident_check_template_id / company_id / year 不能为空")
		return
	}
	uid := middleware.CurrentUserID(c)
	rec := &model.AccidentCheck{
		AccidentCheckTemplateID: *req.AccidentCheckTemplateID,
		CompanyID:               *req.CompanyID,
		Year:                    *req.Year,
		Path:                    req.Path,
		Comment:                 req.Comment,
	}
	if req.Number != nil {
		rec.Number = *req.Number
	}
	rec.UserID = uid
	rec.LastUserID = uid
	if err := h.svc.CreateAccidentCheck(c.Request.Context(), rec); err != nil {
		writeAppError(c, err, "创建失败")
		return
	}
	response.Created(c, rec)
}

func (h *Handler) UpdateAccidentCheck(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	var req accidentCheckReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	updates := map[string]any{"last_user_id": uid}
	if req.AccidentCheckTemplateID != nil {
		updates["accident_check_template_id"] = *req.AccidentCheckTemplateID
	}
	if req.Path != nil {
		updates["path"] = req.Path
	}
	if req.Number != nil {
		updates["number"] = *req.Number
	}
	if req.Comment != nil {
		updates["comment"] = req.Comment
	}
	if req.CompanyID != nil {
		updates["company_id"] = *req.CompanyID
	}
	if req.Year != nil {
		updates["year"] = *req.Year
	}
	if err := h.svc.UpdateAccidentCheck(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err, "更新失败")
		return
	}
	row, _ := h.svc.GetAccidentCheck(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteAccidentCheck(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	if err := h.svc.DeleteAccidentCheck(c.Request.Context(), id); err != nil {
		writeAppError(c, err, "删除失败")
		return
	}
	response.Success(c, gin.H{"message": "删除成功"})
}

// ── 加分模板 ───────────────────────────────────────────────────────

func (h *Handler) ListAddPointsTemplates(c *gin.Context) {
	year := intQuery(c, "year")
	rows, err := h.svc.ListAddPointsTemplates(c.Request.Context(), year)
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Success(c, rows)
}

func (h *Handler) GetAddPointsTemplate(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	row, err := h.svc.GetAddPointsTemplate(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err, "加分项模板不存在")
		return
	}
	response.Success(c, row)
}

func (h *Handler) CreateAddPointsTemplate(c *gin.Context) {
	var req templateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if !req.require(c) {
		return
	}
	rec := &model.AddPointsTemplate{Title: *req.Title, Year: *req.Year, Level: 1, ScoreRange: "1"}
	if req.Level != nil {
		rec.Level = *req.Level
	}
	if req.Sort != nil {
		rec.Sort = *req.Sort
	}
	if req.Pid != nil {
		rec.Pid = *req.Pid
	}
	if req.ScoreRange != nil {
		rec.ScoreRange = *req.ScoreRange
	}
	if err := h.svc.CreateAddPointsTemplate(c.Request.Context(), rec); err != nil {
		writeAppError(c, err, "创建失败")
		return
	}
	response.Created(c, rec)
}

func (h *Handler) UpdateAddPointsTemplate(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	var req templateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if err := h.svc.UpdateAddPointsTemplate(c.Request.Context(), id, req.updates()); err != nil {
		writeAppError(c, err, "更新失败")
		return
	}
	row, err := h.svc.GetAddPointsTemplate(c.Request.Context(), id)
	if err != nil {
		response.Success(c, gin.H{"id": id})
		return
	}
	response.Success(c, row)
}

func (h *Handler) DeleteAddPointsTemplate(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	if err := h.svc.DeleteAddPointsTemplate(c.Request.Context(), id); err != nil {
		writeAppError(c, err, "删除失败")
		return
	}
	response.Success(c, gin.H{"message": "删除成功"})
}

// ── 加分项 ────────────────────────────────────────────────────────

func (h *Handler) ListAddPointsSubmissions(c *gin.Context) {
	pid := int64Query(c, "pid")
	if pid == 0 {
		response.Failed(c, http.StatusUnprocessableEntity, "pid 不能为空")
		return
	}
	rows, err := h.svc.ListAddPointsSubmissions(c.Request.Context(), pid, intQuery(c, "year"))
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Success(c, rows)
}

func (h *Handler) StartAddPointsForm(c *gin.Context) {
	companyID := int64Query(c, "company_id")
	year := intQuery(c, "year")
	if companyID == 0 || year == 0 {
		response.Failed(c, http.StatusUnprocessableEntity, "company_id / year 不能为空")
		return
	}
	rows, err := h.svc.StartAddPointsForm(c.Request.Context(), companyID, year)
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Success(c, rows)
}

type addPointsScoreReq struct {
	CompanyID int64 `json:"company_id"`
	Year      int   `json:"year"`
	Items     []struct {
		AddPointsTemplateID int64    `json:"add_points_template_id"`
		Number              *float64 `json:"number"`
		Comment             *string  `json:"comment"`
	} `json:"items"`
}

func (h *Handler) SaveAddPointsScore(c *gin.Context) {
	var req addPointsScoreReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.CompanyID == 0 || req.Year == 0 {
		response.Failed(c, http.StatusUnprocessableEntity, "company_id / year 不能为空")
		return
	}
	uid := middleware.CurrentUserID(c)
	items := make([]service.AddPointsScoreItem, 0, len(req.Items))
	for _, it := range req.Items {
		items = append(items, service.AddPointsScoreItem{AddPointsTemplateID: it.AddPointsTemplateID, Number: it.Number, Comment: it.Comment})
	}
	if err := h.svc.SaveAddPointsScore(c.Request.Context(), req.CompanyID, req.Year, uid, items); err != nil {
		writeAppError(c, err, "评分失败")
		return
	}
	response.Success(c, gin.H{"message": "评分已保存"})
}

func (h *Handler) ListAddPoints(c *gin.Context) {
	page, per := apiutil.ParsePagination(c, 15)
	rows, total, err := h.svc.ListAddPoints(c.Request.Context(), page, per, int64Query(c, "company_id"), intQuery(c, "year"))
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) GetAddPoint(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	row, err := h.svc.GetAddPoint(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err, "加分项记录不存在")
		return
	}
	response.Success(c, row)
}

type addPointReq struct {
	AddPointsTemplateID *int64   `json:"add_points_template_id"`
	Path                *string  `json:"path"`
	Number              *float64 `json:"number"`
	Comment             *string  `json:"comment"`
	CompanyID           *int64   `json:"company_id"`
	Year                *int     `json:"year"`
}

func (h *Handler) CreateAddPoint(c *gin.Context) {
	var req addPointReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.AddPointsTemplateID == nil || req.CompanyID == nil || req.Year == nil {
		response.Failed(c, http.StatusUnprocessableEntity, "add_points_template_id / company_id / year 不能为空")
		return
	}
	uid := middleware.CurrentUserID(c)
	rec := &model.AddPoint{
		AddPointsTemplateID: *req.AddPointsTemplateID,
		CompanyID:           *req.CompanyID,
		Year:                *req.Year,
		Path:                req.Path,
		Comment:             req.Comment,
	}
	if req.Number != nil {
		rec.Number = *req.Number
	}
	rec.UserID = uid
	rec.LastUserID = uid
	if err := h.svc.CreateAddPoint(c.Request.Context(), rec); err != nil {
		writeAppError(c, err, "创建失败")
		return
	}
	response.Created(c, rec)
}

func (h *Handler) UpdateAddPoint(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	var req addPointReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	updates := map[string]any{"last_user_id": uid}
	if req.AddPointsTemplateID != nil {
		updates["add_points_template_id"] = *req.AddPointsTemplateID
	}
	if req.Path != nil {
		updates["path"] = req.Path
	}
	if req.Number != nil {
		updates["number"] = *req.Number
	}
	if req.Comment != nil {
		updates["comment"] = req.Comment
	}
	if req.CompanyID != nil {
		updates["company_id"] = *req.CompanyID
	}
	if req.Year != nil {
		updates["year"] = *req.Year
	}
	if err := h.svc.UpdateAddPoint(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err, "更新失败")
		return
	}
	row, _ := h.svc.GetAddPoint(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteAddPoint(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	if err := h.svc.DeleteAddPoint(c.Request.Context(), id); err != nil {
		writeAppError(c, err, "删除失败")
		return
	}
	response.Success(c, gin.H{"message": "删除成功"})
}

// DeleteAddPointsSubmission removes the entire (company, year) submission group —
// every add_points row for that subordinate company in that year. PointsTab
// renders one row per (company, year), so this matches the user-visible delete.
func (h *Handler) DeleteAddPointsSubmission(c *gin.Context) {
	companyID := int64Query(c, "company_id")
	year := intQuery(c, "year")
	if companyID <= 0 || year <= 0 {
		response.BadRequest(c, "company_id 和 year 必填")
		return
	}
	if err := h.svc.DeleteAddPointsSubmission(c.Request.Context(), companyID, year); err != nil {
		writeAppError(c, err, "删除失败")
		return
	}
	response.Success(c, gin.H{"message": "删除成功"})
}

// ── 系统检查 ──────────────────────────────────────────────────────

func (h *Handler) ListSystemChecks(c *gin.Context) {
	page, per := apiutil.ParsePagination(c, 15)
	rows, total, err := h.svc.ListSystemChecks(c.Request.Context(), page, per, int64Query(c, "company_id"), intQuery(c, "year"))
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) GetSystemCheck(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	row, err := h.svc.GetSystemCheck(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err, "系统考核记录不存在")
		return
	}
	response.Success(c, row)
}

type systemCheckReq struct {
	NumberV1  *float32 `json:"number_v1"`
	NumberV2  *float32 `json:"number_v2"`
	NumberV3  *float32 `json:"number_v3"`
	NumberV4  *float32 `json:"number_v4"`
	NumberV5  *float32 `json:"number_v5"`
	Number    *int     `json:"number"`
	Year      *int     `json:"year"`
	CompanyID *int64   `json:"company_id"`
}

func (h *Handler) CreateSystemCheck(c *gin.Context) {
	var req systemCheckReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.CompanyID == nil || req.Year == nil {
		response.Failed(c, http.StatusUnprocessableEntity, "company_id / year 不能为空")
		return
	}
	uid := middleware.CurrentUserID(c)
	rec := &model.SystemCheck{
		CompanyID: *req.CompanyID,
		Year:      *req.Year,
	}
	if req.NumberV1 != nil {
		rec.NumberV1 = *req.NumberV1
	}
	if req.NumberV2 != nil {
		rec.NumberV2 = *req.NumberV2
	}
	if req.NumberV3 != nil {
		rec.NumberV3 = *req.NumberV3
	}
	if req.NumberV4 != nil {
		rec.NumberV4 = *req.NumberV4
	}
	if req.NumberV5 != nil {
		rec.NumberV5 = *req.NumberV5
	}
	if req.Number != nil {
		rec.Number = *req.Number
	}
	rec.UserID = uid
	rec.LastUserID = uid
	if err := h.svc.CreateSystemCheck(c.Request.Context(), rec); err != nil {
		writeAppError(c, err, "创建失败")
		return
	}
	response.Created(c, rec)
}

func (h *Handler) UpdateSystemCheck(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	var req systemCheckReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	updates := map[string]any{"last_user_id": uid}
	if req.NumberV1 != nil {
		updates["number_v1"] = *req.NumberV1
	}
	if req.NumberV2 != nil {
		updates["number_v2"] = *req.NumberV2
	}
	if req.NumberV3 != nil {
		updates["number_v3"] = *req.NumberV3
	}
	if req.NumberV4 != nil {
		updates["number_v4"] = *req.NumberV4
	}
	if req.NumberV5 != nil {
		updates["number_v5"] = *req.NumberV5
	}
	if req.Number != nil {
		updates["number"] = *req.Number
	}
	if req.Year != nil {
		updates["year"] = *req.Year
	}
	if req.CompanyID != nil {
		updates["company_id"] = *req.CompanyID
	}
	if err := h.svc.UpdateSystemCheck(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err, "更新失败")
		return
	}
	row, _ := h.svc.GetSystemCheck(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteSystemCheck(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	if err := h.svc.DeleteSystemCheck(c.Request.Context(), id); err != nil {
		writeAppError(c, err, "删除失败")
		return
	}
	response.Success(c, gin.H{"message": "删除成功"})
}

func (h *Handler) ListSystemSubmissions(c *gin.Context) {
	pid := int64Query(c, "pid")
	if pid == 0 {
		response.Failed(c, http.StatusUnprocessableEntity, "pid 不能为空")
		return
	}
	rows, err := h.svc.ListSystemSubmissions(c.Request.Context(), pid, intQuery(c, "year"))
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Success(c, rows)
}

func (h *Handler) ComputeSystemCheck(c *gin.Context) {
	var req yearAndCompanyReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.CompanyID == 0 || req.Year == 0 {
		response.Failed(c, http.StatusUnprocessableEntity, "company_id / year 不能为空")
		return
	}
	uid := middleware.CurrentUserID(c)
	row, err := h.svc.ComputeSystemCheck(c.Request.Context(), req.CompanyID, req.Year, uid)
	if err != nil {
		writeAppError(c, err, "计算失败")
		return
	}
	response.Success(c, row)
}

// ── 标准化检查 ────────────────────────────────────────────────────

func (h *Handler) ListNormChecks(c *gin.Context) {
	page, per := apiutil.ParsePagination(c, 15)
	rows, total, err := h.svc.ListNormChecks(c.Request.Context(), page, per, int64Query(c, "company_id"), intQuery(c, "year"), c.Query("type"))
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) GetNormCheck(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	row, err := h.svc.GetNormCheck(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err, "标准化检查记录不存在")
		return
	}
	response.Success(c, row)
}

type normCheckReq struct {
	Path      *string `json:"path"`
	Number    *int    `json:"number"`
	Year      *int    `json:"year"`
	Type      *string `json:"type"`
	CompanyID *int64  `json:"company_id"`
}

func (h *Handler) CreateNormCheck(c *gin.Context) {
	var req normCheckReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.CompanyID == nil || req.Year == nil || req.Type == nil || strings.TrimSpace(*req.Type) == "" {
		response.Failed(c, http.StatusUnprocessableEntity, "company_id / year / type 不能为空")
		return
	}
	uid := middleware.CurrentUserID(c)
	rec := &model.NormCheck{
		CompanyID: *req.CompanyID,
		Year:      *req.Year,
		Type:      *req.Type,
		Path:      req.Path,
	}
	if req.Number != nil {
		rec.Number = *req.Number
	}
	rec.UserID = uid
	rec.LastUserID = uid
	if err := h.svc.CreateNormCheck(c.Request.Context(), rec); err != nil {
		writeAppError(c, err, "创建失败")
		return
	}
	response.Created(c, rec)
}

func (h *Handler) UpdateNormCheck(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	var req normCheckReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	updates := map[string]any{"last_user_id": uid}
	if req.Path != nil {
		updates["path"] = req.Path
	}
	if req.Number != nil {
		updates["number"] = *req.Number
	}
	if req.Year != nil {
		updates["year"] = *req.Year
	}
	if req.Type != nil {
		updates["type"] = *req.Type
	}
	if req.CompanyID != nil {
		updates["company_id"] = *req.CompanyID
	}
	if err := h.svc.UpdateNormCheck(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err, "更新失败")
		return
	}
	row, _ := h.svc.GetNormCheck(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteNormCheck(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	if err := h.svc.DeleteNormCheck(c.Request.Context(), id); err != nil {
		writeAppError(c, err, "删除失败")
		return
	}
	response.Success(c, gin.H{"message": "删除成功"})
}

func (h *Handler) ListNormSubmissions(c *gin.Context) {
	pid := int64Query(c, "pid")
	if pid == 0 {
		response.Failed(c, http.StatusUnprocessableEntity, "pid 不能为空")
		return
	}
	rows, err := h.svc.ListNormSubmissions(c.Request.Context(), pid, intQuery(c, "year"))
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Success(c, rows)
}

type normSaveScoreReq struct {
	CompanyID int64   `json:"company_id"`
	Year      int     `json:"year"`
	Type      string  `json:"type"`
	Number    int     `json:"number"`
	Path      *string `json:"path"`
}

func (h *Handler) SaveNormScore(c *gin.Context) {
	var req normSaveScoreReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.CompanyID == 0 || req.Year == 0 || strings.TrimSpace(req.Type) == "" {
		response.Failed(c, http.StatusUnprocessableEntity, "company_id / year / type 不能为空")
		return
	}
	if req.Number < 0 || req.Number > 100 {
		response.Failed(c, http.StatusUnprocessableEntity, "number 必须在 0-100 之间")
		return
	}
	uid := middleware.CurrentUserID(c)
	row, err := h.svc.SaveNormScore(c.Request.Context(), req.CompanyID, req.Year, uid, req.Type, req.Number, req.Path)
	if err != nil {
		writeAppError(c, err, "保存失败")
		return
	}
	response.Success(c, row)
}

func (h *Handler) NormTemplateLink(c *gin.Context) {
	kind := c.Query("type")
	if strings.TrimSpace(kind) == "" {
		response.Failed(c, http.StatusUnprocessableEntity, "type 不能为空")
		return
	}
	url, t := service.NormTemplateLink(kind)
	response.Success(c, gin.H{"url": url, "type": t})
}

// ── 考核模板（年度并集 + 一键拷贝） ───────────────────────────────

func (h *Handler) AssessmentYears(c *gin.Context) {
	rows, err := h.svc.AssessmentYears(c.Request.Context())
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Success(c, rows)
}

type copyReq struct {
	FromYear  int      `json:"from_year"`
	ToYear    int      `json:"to_year"`
	Kinds     []string `json:"kinds"`
	Overwrite bool     `json:"overwrite"`
}

func (h *Handler) CopyAssessmentTemplates(c *gin.Context) {
	var req copyReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.FromYear < 2000 || req.FromYear > 2100 || req.ToYear < 2000 || req.ToYear > 2100 {
		response.Failed(c, http.StatusUnprocessableEntity, "from_year / to_year 必须在 2000-2100 之间")
		return
	}
	if req.FromYear == req.ToYear {
		response.Failed(c, http.StatusUnprocessableEntity, "from_year 与 to_year 不能相同")
		return
	}
	res, err := h.svc.CopyAssessmentTemplates(c.Request.Context(), req.FromYear, req.ToYear, req.Kinds, req.Overwrite)
	if err != nil {
		writeAppError(c, err, "拷贝失败")
		return
	}
	response.Success(c, res)
}

// ── 考核汇总 ──────────────────────────────────────────────────────

func (h *Handler) Summary(c *gin.Context) {
	companyID := int64Query(c, "company_id")
	year := intQuery(c, "year")
	kind := c.Query("type")
	if companyID == 0 || year == 0 || (kind != "bloc" && kind != "company") {
		response.Failed(c, http.StatusUnprocessableEntity, "company_id / year / type(bloc|company) 不能为空")
		return
	}
	if kind == "bloc" {
		rows, err := h.svc.SummaryBloc(c.Request.Context(), companyID, year)
		if err != nil {
			writeAppError(c, err, "查询失败")
			return
		}
		response.Success(c, rows)
		return
	}
	row, err := h.svc.SummaryCompany(c.Request.Context(), companyID, year)
	if err != nil {
		writeAppError(c, err, "查询失败")
		return
	}
	response.Success(c, row)
}

func (h *Handler) ExportSummary(c *gin.Context) {
	companyID := int64Query(c, "company_id")
	year := intQuery(c, "year")
	kind := c.Query("type")
	if companyID == 0 || year == 0 || (kind != "bloc" && kind != "company") {
		response.Failed(c, http.StatusUnprocessableEntity, "company_id / year / type(bloc|company) 不能为空")
		return
	}
	absPath, filename, err := h.svc.ExportSummary(c.Request.Context(), companyID, year, kind, h.templateDir, h.storageDir)
	if err != nil {
		writeAppError(c, err, "导出失败")
		return
	}
	c.FileAttachment(absPath, filename)
}
