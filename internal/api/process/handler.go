// Package process 实现工艺信息三个模块的 CRUD：
//
//	/major-hazardous-processes — 重点监管危险工艺
//	/hazardous-processes       — 较大危险工艺
//	/space-works               — 有限空间辨识
//
// 三个模块互无关联，但都共用同一份 ProcessService（避免 router 里挂三个零行为
// 的构造函数）。请求结构体里所有字段都用指针，前端 `form.x || null` 提交语义
// 直接对应：JSON null → nil pointer → 写入 SQL NULL，可以清空字段。
package process

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
	svc *service.ProcessService
}

func NewHandler(svc *service.ProcessService) *Handler {
	return &Handler{svc: svc}
}

// ── 重点监管危险工艺 ────────────────────────────────────────────────────

type majorHazardousProcessReq struct {
	Name                 *string `json:"name"`
	DangerousCraftTypeID *int64  `json:"dangerous_craft_type_id"`
	Position             *string `json:"position"`
	HazardousChemicalID  *int64  `json:"hazardous_chemical_id"`
	DangerousEquipmentID *int64  `json:"dangerous_equipment_id"`
	OperatingPeopleNum   *string `json:"operating_people_num"`
	Temperature          *string `json:"temperature"`
	Pressure             *string `json:"pressure"`
	PressureUnit         *string `json:"pressure_unit"`
	IsSIS                int16   `json:"is_sis"`
	SILLevel             *string `json:"sil_level"`
	DetectParameter      *string `json:"detect_parameter"`
}

func (h *Handler) ListMajorHazardous(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "20"))
	rows, total, err := h.svc.ListMajorHazardousProcess(c.Request.Context(), service.MajorHazardousProcessListParams{
		Page: page, PerPage: per, CompanyID: companyID,
		Name: firstNonEmpty(c.Query("name"), c.Query("keyword")),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) CreateMajorHazardous(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	var req majorHazardousProcessReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	m := &model.MajorHazardousProcess{
		Name:                 req.Name,
		DangerousCraftTypeID: req.DangerousCraftTypeID,
		Position:             req.Position,
		HazardousChemicalID:  req.HazardousChemicalID,
		DangerousEquipmentID: req.DangerousEquipmentID,
		OperatingPeopleNum:   req.OperatingPeopleNum,
		Temperature:          req.Temperature,
		Pressure:             req.Pressure,
		PressureUnit:         req.PressureUnit,
		IsSIS:                req.IsSIS,
		SILLevel:             req.SILLevel,
		DetectParameter:      req.DetectParameter,
		CompanyID:            companyID,
	}
	m.UserID = uid
	m.LastUserID = uid
	if err := h.svc.CreateMajorHazardousProcess(c.Request.Context(), m); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetMajorHazardousProcess(c.Request.Context(), m.ID)
	response.Created(c, row)
}

func (h *Handler) UpdateMajorHazardous(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req majorHazardousProcessReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	updates := map[string]any{
		"last_user_id":            uid,
		"name":                    req.Name,
		"dangerous_craft_type_id": req.DangerousCraftTypeID,
		"position":                req.Position,
		"hazardous_chemical_id":   req.HazardousChemicalID,
		"dangerous_equipment_id":  req.DangerousEquipmentID,
		"operating_people_num":    req.OperatingPeopleNum,
		"temperature":             req.Temperature,
		"pressure":                req.Pressure,
		"pressure_unit":           req.PressureUnit,
		"is_sis":                  req.IsSIS,
		"sil_level":               req.SILLevel,
		"detect_parameter":        req.DetectParameter,
	}
	if err := h.svc.UpdateMajorHazardousProcess(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetMajorHazardousProcess(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteMajorHazardous(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteMajorHazardousProcess(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

// ── 较大危险工艺 ───────────────────────────────────────────────────────

type hazardousProcessReq struct {
	ProcessTypeID        *int64  `json:"process_type_id"`
	Position             *string `json:"position"`
	HazardousChemicalID  *int64  `json:"hazardous_chemical_id"`
	DangerousEquipmentID *int64  `json:"dangerous_equipment_id"`
	OtherEquipmentName   *string `json:"other_equipment_name"`
	OperatingPeopleNum   *string `json:"operating_people_num"`
	IsOnlineTest         int16   `json:"is_online_test"`
	IsAutomatic          int16   `json:"is_automatic"`
	IsControl            int16   `json:"is_control"`
	IsDust               int16   `json:"is_dust"`
}

func (h *Handler) ListHazardous(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "20"))
	rows, total, err := h.svc.ListHazardousProcess(c.Request.Context(), service.HazardousProcessListParams{
		Page: page, PerPage: per, CompanyID: companyID,
		Position: firstNonEmpty(c.Query("position"), c.Query("keyword")),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) CreateHazardous(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	var req hazardousProcessReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	m := &model.HazardousProcess{
		ProcessTypeID:        req.ProcessTypeID,
		Position:             req.Position,
		HazardousChemicalID:  req.HazardousChemicalID,
		DangerousEquipmentID: req.DangerousEquipmentID,
		OtherEquipmentName:   req.OtherEquipmentName,
		OperatingPeopleNum:   req.OperatingPeopleNum,
		IsOnlineTest:         req.IsOnlineTest,
		IsAutomatic:          req.IsAutomatic,
		IsControl:            req.IsControl,
		IsDust:               req.IsDust,
		CompanyID:            companyID,
	}
	m.UserID = uid
	m.LastUserID = uid
	if err := h.svc.CreateHazardousProcess(c.Request.Context(), m); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetHazardousProcess(c.Request.Context(), m.ID)
	response.Created(c, row)
}

func (h *Handler) UpdateHazardous(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req hazardousProcessReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	updates := map[string]any{
		"last_user_id":           uid,
		"process_type_id":        req.ProcessTypeID,
		"position":               req.Position,
		"hazardous_chemical_id":  req.HazardousChemicalID,
		"dangerous_equipment_id": req.DangerousEquipmentID,
		"other_equipment_name":   req.OtherEquipmentName,
		"operating_people_num":   req.OperatingPeopleNum,
		"is_online_test":         req.IsOnlineTest,
		"is_automatic":           req.IsAutomatic,
		"is_control":             req.IsControl,
		"is_dust":                req.IsDust,
	}
	if err := h.svc.UpdateHazardousProcess(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetHazardousProcess(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteHazardous(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteHazardousProcess(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

// ── 有限空间辨识 ───────────────────────────────────────────────────────

type spaceWorkReq struct {
	Name            *string `json:"name"`
	Position        *string `json:"position"`
	SpaceWorkTypeID *int64  `json:"space_work_type_id"`
	Size            *string `json:"size"`
	PeopleNum       *int32  `json:"people_num"`
	HazardID        *string `json:"hazard_id"`
	WordContent     *string `json:"word_content"`
}

func (h *Handler) ListSpaceWork(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "20"))
	rows, total, err := h.svc.ListSpaceWork(c.Request.Context(), service.SpaceWorkListParams{
		Page: page, PerPage: per, CompanyID: companyID,
		Name: firstNonEmpty(c.Query("name"), c.Query("keyword")),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) CreateSpaceWork(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	var req spaceWorkReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	m := &model.SpaceWork{
		Name:            req.Name,
		Position:        req.Position,
		SpaceWorkTypeID: req.SpaceWorkTypeID,
		Size:            req.Size,
		PeopleNum:       req.PeopleNum,
		HazardID:        req.HazardID,
		WordContent:     req.WordContent,
		CompanyID:       companyID,
	}
	m.UserID = uid
	m.LastUserID = uid
	if err := h.svc.CreateSpaceWork(c.Request.Context(), m); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetSpaceWork(c.Request.Context(), m.ID)
	response.Created(c, row)
}

func (h *Handler) UpdateSpaceWork(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req spaceWorkReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	updates := map[string]any{
		"last_user_id":       uid,
		"name":               req.Name,
		"position":           req.Position,
		"space_work_type_id": req.SpaceWorkTypeID,
		"size":               req.Size,
		"people_num":         req.PeopleNum,
		"hazard_id":          req.HazardID,
		"word_content":       req.WordContent,
	}
	if err := h.svc.UpdateSpaceWork(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetSpaceWork(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteSpaceWork(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteSpaceWork(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
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
