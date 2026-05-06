// Package accidentemergency 实现事故应急管理 5 个模块：
//
//	/accident-bulletins         — 生产事故快报（含 kids 整体替换）
//	/bloc-accident-bulletins    — 集团事故快报（跨 company subtree 聚合 + annual-stats）
//	/emergency-drills           — 应急演练
//	/emergency-supplies         — 应急物资
//	/emergency-teams            — 应急队伍（含 kids 整体替换）
//
// 子表替换语义跟 risk 模块一致：*[]kidReq 用 nil 指针表示请求体未携带 key（不动子表），
// 非 nil（含空切片）表示按本次值整体覆盖。Laravel 原写法 `$kids = $request->input('kids')` 同语义。
package accidentemergency

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/middleware"
	"ai_hs_go/internal/model"
	"ai_hs_go/internal/response"
	"ai_hs_go/internal/service"
	"ai_hs_go/internal/api/apiutil"
)

type Handler struct {
	svc *service.AccidentEmergencyService
}

func NewHandler(svc *service.AccidentEmergencyService) *Handler {
	return &Handler{svc: svc}
}

// ── 事故快报（生产视图） ──────────────────────────────────────────────

type accidentKidReq struct {
	Name           *string    `json:"name"`
	DegreeInjury   *string    `json:"degree_injury"`
	Department     *string    `json:"department"`
	Age            *flexInt32 `json:"age"`
	Sex            *string    `json:"sex"`
	WorkType       *string    `json:"work_type"`
	WorkAge        *flexInt32 `json:"work_age"`
	EducationLevel *string    `json:"education_level"`
	DaysOff        *flexInt32 `json:"days_off"`
}

type accidentBulletinReq struct {
	Type               *string    `json:"type"`
	DirectCause        *string    `json:"direct_cause"`
	IndirectCause      *string    `json:"indirect_cause"`
	AccidentLevel      *string    `json:"accident_level"`
	Department         *string    `json:"department"`
	HappenTime         *string    `json:"happen_time"`
	HappenArea         *string    `json:"happen_area"`
	Content            *string    `json:"content"`
	DegreeInjury       *string    `json:"degree_injury"`
	DirectLossProperty *string    `json:"direct_loss_property"`
	GaugeLossProperty  *string    `json:"gauge_loss_property"`
	StaffDie           *flexInt32 `json:"staff_die"`
	StaffHurt          *flexInt32 `json:"staff_hurt"`
	NotStaffDie        *flexInt32 `json:"not_staff_die"`
	NotStaffHurt       *flexInt32 `json:"not_staff_hurt"`
	AccidentReportURL  *string    `json:"accident_report_url"`
	AccidentPhotoURL   *string    `json:"accident_photo_url"`
	// nil = 不动子表；非 nil（含空切片）= 整体替换
	Kids *[]accidentKidReq `json:"kids"`
}

func (h *Handler) ListAccidentBulletin(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	page, per := apiutil.ParsePagination(c, 15)
	rows, total, err := h.svc.ListAccidentBulletin(c.Request.Context(), service.AccidentBulletinListParams{
		Page: page, PerPage: per, CompanyID: companyID,
		Type:          c.Query("type"),
		AccidentLevel: c.Query("accident_level"),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) GetAccidentBulletin(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.svc.GetAccidentBulletin(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

func (h *Handler) CreateAccidentBulletin(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	var req accidentBulletinReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	bulletin := &model.AccidentBulletin{
		Type:               req.Type,
		DirectCause:        req.DirectCause,
		IndirectCause:      req.IndirectCause,
		AccidentLevel:      req.AccidentLevel,
		Department:         req.Department,
		HappenTime:         parseDate(req.HappenTime),
		HappenArea:         req.HappenArea,
		Content:            req.Content,
		DegreeInjury:       req.DegreeInjury,
		DirectLossProperty: req.DirectLossProperty,
		GaugeLossProperty:  req.GaugeLossProperty,
		StaffDie:           req.StaffDie.Int32Ptr(),
		StaffHurt:          req.StaffHurt.Int32Ptr(),
		NotStaffDie:        req.NotStaffDie.Int32Ptr(),
		NotStaffHurt:       req.NotStaffHurt.Int32Ptr(),
		AccidentReportURL:  req.AccidentReportURL,
		AccidentPhotoURL:   req.AccidentPhotoURL,
		CompanyID:          companyID,
	}
	bulletin.UserID = uid
	bulletin.LastUserID = uid
	if err := h.svc.CreateAccidentBulletin(c.Request.Context(), service.AccidentBulletinInput{
		Bulletin: bulletin,
		Kids:     mapAccidentKids(derefAccidentKids(req.Kids)),
	}); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetAccidentBulletin(c.Request.Context(), bulletin.ID)
	response.Created(c, row)
}

func (h *Handler) UpdateAccidentBulletin(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req accidentBulletinReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	updates := map[string]any{"last_user_id": uid}
	addNullableString(updates, "type", req.Type)
	addNullableString(updates, "direct_cause", req.DirectCause)
	addNullableString(updates, "indirect_cause", req.IndirectCause)
	addNullableString(updates, "accident_level", req.AccidentLevel)
	addNullableString(updates, "department", req.Department)
	addNullableString(updates, "happen_area", req.HappenArea)
	addNullableString(updates, "content", req.Content)
	addNullableString(updates, "degree_injury", req.DegreeInjury)
	addNullableString(updates, "direct_loss_property", req.DirectLossProperty)
	addNullableString(updates, "gauge_loss_property", req.GaugeLossProperty)
	addNullableString(updates, "accident_report_url", req.AccidentReportURL)
	addNullableString(updates, "accident_photo_url", req.AccidentPhotoURL)
	if req.HappenTime != nil {
		updates["happen_time"] = parseDate(req.HappenTime)
	}
	if req.StaffDie != nil {
		updates["staff_die"] = req.StaffDie.Int32Ptr()
	}
	if req.StaffHurt != nil {
		updates["staff_hurt"] = req.StaffHurt.Int32Ptr()
	}
	if req.NotStaffDie != nil {
		updates["not_staff_die"] = req.NotStaffDie.Int32Ptr()
	}
	if req.NotStaffHurt != nil {
		updates["not_staff_hurt"] = req.NotStaffHurt.Int32Ptr()
	}
	if err := h.svc.UpdateAccidentBulletin(
		c.Request.Context(), id, updates,
		mapAccidentKids(derefAccidentKids(req.Kids)),
		req.Kids != nil, uid,
	); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetAccidentBulletin(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteAccidentBulletin(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteAccidentBulletin(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

func mapAccidentKids(items []accidentKidReq) []model.AccidentBulletinKid {
	out := make([]model.AccidentBulletinKid, len(items))
	for i, it := range items {
		out[i] = model.AccidentBulletinKid{
			Name:           it.Name,
			DegreeInjury:   it.DegreeInjury,
			Department:     it.Department,
			Age:            it.Age.Int32Ptr(),
			Sex:            it.Sex,
			WorkType:       it.WorkType,
			WorkAge:        it.WorkAge.Int32Ptr(),
			EducationLevel: it.EducationLevel,
			DaysOff:        it.DaysOff.Int32Ptr(),
		}
	}
	return out
}

func derefAccidentKids(p *[]accidentKidReq) []accidentKidReq {
	if p == nil {
		return nil
	}
	return *p
}

// ── 集团事故快报 ─────────────────────────────────────────────────────

func (h *Handler) ListBlocAccidentBulletin(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	page, per := apiutil.ParsePagination(c, 15)
	rows, total, err := h.svc.ListBlocAccidentBulletin(c.Request.Context(), service.BlocAccidentBulletinListParams{
		Page: page, PerPage: per, CompanyID: companyID,
		Type:          c.Query("type"),
		AccidentLevel: c.Query("accident_level"),
		HappenYear:    c.Query("happen_year"),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) BlocAccidentAnnualStats(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	rows, err := h.svc.BlocAccidentAnnualStats(c.Request.Context(), companyID)
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Success(c, rows)
}

// ── 应急演练 ─────────────────────────────────────────────────────────

type emergencyDrillReq struct {
	Name                   *string `json:"name"`
	DrillTime              *string `json:"drill_time"`
	DrillLocal             *string `json:"drill_local"`
	OrganizationDepartment *string `json:"organization_department"`
	JoinDepartment         *string `json:"join_department"`
	DrillNumber            *string `json:"drill_number"`
	DrillType              *string `json:"drill_type"`
	DrillForm              *string `json:"drill_form"`
	DrillKind              *string `json:"drill_kind"`
	Event                  *string `json:"event"`
	AnnexURL               *string `json:"annex_url"`
}

func (h *Handler) ListEmergencyDrill(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	page, per := apiutil.ParsePagination(c, 15)
	rows, total, err := h.svc.ListEmergencyDrill(c.Request.Context(), service.EmergencyDrillListParams{
		Page: page, PerPage: per, CompanyID: companyID,
		Name: firstNonEmpty(c.Query("name"), c.Query("keyword")),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) GetEmergencyDrill(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.svc.GetEmergencyDrill(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

func (h *Handler) CreateEmergencyDrill(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	var req emergencyDrillReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	m := &model.EmergencyDrill{
		Name:                   req.Name,
		DrillTime:              parseDate(req.DrillTime),
		DrillLocal:             req.DrillLocal,
		OrganizationDepartment: req.OrganizationDepartment,
		JoinDepartment:         req.JoinDepartment,
		DrillNumber:            req.DrillNumber,
		DrillType:              req.DrillType,
		DrillForm:              req.DrillForm,
		DrillKind:              req.DrillKind,
		Event:                  req.Event,
		AnnexURL:               req.AnnexURL,
		CompanyID:              companyID,
	}
	m.UserID = uid
	m.LastUserID = uid
	if err := h.svc.CreateEmergencyDrill(c.Request.Context(), m); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetEmergencyDrill(c.Request.Context(), m.ID)
	response.Created(c, row)
}

func (h *Handler) UpdateEmergencyDrill(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req emergencyDrillReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	updates := map[string]any{"last_user_id": middleware.CurrentUserID(c)}
	addNullableString(updates, "name", req.Name)
	addNullableString(updates, "drill_local", req.DrillLocal)
	addNullableString(updates, "organization_department", req.OrganizationDepartment)
	addNullableString(updates, "join_department", req.JoinDepartment)
	addNullableString(updates, "drill_number", req.DrillNumber)
	addNullableString(updates, "drill_type", req.DrillType)
	addNullableString(updates, "drill_form", req.DrillForm)
	addNullableString(updates, "drill_kind", req.DrillKind)
	addNullableString(updates, "event", req.Event)
	addNullableString(updates, "annex_url", req.AnnexURL)
	if req.DrillTime != nil {
		updates["drill_time"] = parseDate(req.DrillTime)
	}
	if err := h.svc.UpdateEmergencyDrill(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetEmergencyDrill(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteEmergencyDrill(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteEmergencyDrill(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

// ── 应急物资 ─────────────────────────────────────────────────────────

type emergencySuppliesReq struct {
	Name     *string `json:"name"`
	Number   *string `json:"number"`
	SaveUser *string `json:"save_user"`
	Tel      *string `json:"tel"`
	Comment  *string `json:"comment"`
}

func (h *Handler) ListEmergencySupplies(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	page, per := apiutil.ParsePagination(c, 15)
	rows, total, err := h.svc.ListEmergencySupplies(c.Request.Context(), service.EmergencySuppliesListParams{
		Page: page, PerPage: per, CompanyID: companyID,
		Name: firstNonEmpty(c.Query("name"), c.Query("keyword")),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) GetEmergencySupplies(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.svc.GetEmergencySupplies(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

func (h *Handler) CreateEmergencySupplies(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	var req emergencySuppliesReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	m := &model.EmergencySupplies{
		Name:      req.Name,
		Number:    req.Number,
		SaveUser:  req.SaveUser,
		Tel:       req.Tel,
		Comment:   req.Comment,
		CompanyID: companyID,
	}
	m.UserID = uid
	m.LastUserID = uid
	if err := h.svc.CreateEmergencySupplies(c.Request.Context(), m); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetEmergencySupplies(c.Request.Context(), m.ID)
	response.Created(c, row)
}

func (h *Handler) UpdateEmergencySupplies(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req emergencySuppliesReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	updates := map[string]any{"last_user_id": middleware.CurrentUserID(c)}
	addNullableString(updates, "name", req.Name)
	addNullableString(updates, "number", req.Number)
	addNullableString(updates, "save_user", req.SaveUser)
	addNullableString(updates, "tel", req.Tel)
	addNullableString(updates, "comment", req.Comment)
	if err := h.svc.UpdateEmergencySupplies(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetEmergencySupplies(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteEmergencySupplies(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteEmergencySupplies(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

// ── 应急队伍 ─────────────────────────────────────────────────────────

type emergencyTeamKidReq struct {
	Name       *string `json:"name"`
	Department *string `json:"department"`
	Content    *string `json:"content"`
	Tel        *string `json:"tel"`
}

type emergencyTeamReq struct {
	Name       *string                `json:"name"`
	TeamNumber *string                `json:"team_number"`
	SaveUser   *string                `json:"save_user"`
	Tel        *string                `json:"tel"`
	Kids       *[]emergencyTeamKidReq `json:"kids"`
}

func (h *Handler) ListEmergencyTeam(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	page, per := apiutil.ParsePagination(c, 15)
	rows, total, err := h.svc.ListEmergencyTeam(c.Request.Context(), service.EmergencyTeamListParams{
		Page: page, PerPage: per, CompanyID: companyID,
		Name: firstNonEmpty(c.Query("name"), c.Query("keyword")),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) GetEmergencyTeam(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.svc.GetEmergencyTeam(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

func (h *Handler) CreateEmergencyTeam(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	var req emergencyTeamReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	team := &model.EmergencyTeam{
		Name:       req.Name,
		TeamNumber: req.TeamNumber,
		SaveUser:   req.SaveUser,
		Tel:        req.Tel,
		CompanyID:  companyID,
	}
	team.UserID = uid
	team.LastUserID = uid
	if err := h.svc.CreateEmergencyTeam(c.Request.Context(), service.EmergencyTeamInput{
		Team: team,
		Kids: mapTeamKids(derefTeamKids(req.Kids)),
	}); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetEmergencyTeam(c.Request.Context(), team.ID)
	response.Created(c, row)
}

func (h *Handler) UpdateEmergencyTeam(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req emergencyTeamReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	updates := map[string]any{"last_user_id": uid}
	addNullableString(updates, "name", req.Name)
	addNullableString(updates, "team_number", req.TeamNumber)
	addNullableString(updates, "save_user", req.SaveUser)
	addNullableString(updates, "tel", req.Tel)
	if err := h.svc.UpdateEmergencyTeam(
		c.Request.Context(), id, updates,
		mapTeamKids(derefTeamKids(req.Kids)),
		req.Kids != nil, uid,
	); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetEmergencyTeam(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteEmergencyTeam(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteEmergencyTeam(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

func mapTeamKids(items []emergencyTeamKidReq) []model.EmergencyTeamKid {
	out := make([]model.EmergencyTeamKid, len(items))
	for i, it := range items {
		out[i] = model.EmergencyTeamKid{
			Name:       it.Name,
			Department: it.Department,
			Content:    it.Content,
			Tel:        it.Tel,
		}
	}
	return out
}

func derefTeamKids(p *[]emergencyTeamKidReq) []emergencyTeamKidReq {
	if p == nil {
		return nil
	}
	return *p
}

// flexInt32 — 容忍前端 <input type=number> 把数字以字符串形式提交。
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
