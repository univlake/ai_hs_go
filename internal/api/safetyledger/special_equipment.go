package safetyledger

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/excel"
	"ai_hs_go/internal/middleware"
	"ai_hs_go/internal/model"
	"ai_hs_go/internal/response"
	"ai_hs_go/internal/service"
	"ai_hs_go/internal/api/apiutil"
)

type Handler struct {
	svc *service.SafetyLedgerService
}

func NewHandler(svc *service.SafetyLedgerService) *Handler {
	return &Handler{svc: svc}
}

// ListSpecialEquipment: GET /api/v1/special-equipment
func (h *Handler) ListSpecialEquipment(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	page, per := apiutil.ParsePagination(c, 20)
	rows, total, err := h.svc.ListSpecialEquipment(c.Request.Context(), service.ListParams{
		Page: page, PerPage: per, CompanyID: companyID, Keyword: c.Query("keyword"),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

type specialEquipmentReq struct {
	EquipmentName    *string `json:"equipment_name"`
	SpecialTypeID    *int64  `json:"special_type_id"`
	UtilizedLocation *string `json:"utilized_location"`
	NormModel1       *string `json:"norm_model1"`
	NormModel2       *string `json:"norm_model2"`
	NormModel3       *string `json:"norm_model3"`
	NormModel4       *string `json:"norm_model4"`
	SerialNumber     *string `json:"serial_number"`
	UseTime          *string `json:"use_time"`
	ValidPeriod      *string `json:"valid_period"`
	InspectUnit      *string `json:"inspect_unit"`
	MaintainInfo     *string `json:"maintain_info"`
	Comment          *string `json:"comment"`
}

func (h *Handler) CreateSpecialEquipment(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	var req specialEquipmentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	row := &model.SpecialEquipment{
		EquipmentName:    req.EquipmentName,
		SpecialTypeID:    req.SpecialTypeID,
		UtilizedLocation: req.UtilizedLocation,
		NormModel1:       req.NormModel1,
		NormModel2:       req.NormModel2,
		NormModel3:       req.NormModel3,
		NormModel4:       req.NormModel4,
		SerialNumber:     req.SerialNumber,
		UseTime:          parseDate(req.UseTime),
		ValidPeriod:      parseDate(req.ValidPeriod),
		InspectUnit:      req.InspectUnit,
		MaintainInfo:     req.MaintainInfo,
		Comment:          req.Comment,
		CompanyID:        companyID,
	}
	row.UserID = uid
	row.LastUserID = uid
	if err := h.svc.CreateSpecialEquipment(c.Request.Context(), row); err != nil {
		writeAppError(c, err)
		return
	}
	response.Created(c, row)
}

func (h *Handler) UpdateSpecialEquipment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req specialEquipmentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	updates := buildSpecialUpdates(req, middleware.CurrentUserID(c))
	if err := h.svc.UpdateSpecialEquipment(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetSpecialEquipment(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteSpecialEquipment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteSpecialEquipment(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

// ExportSpecialEquipment: GET /api/v1/special-equipment/export
func (h *Handler) ExportSpecialEquipment(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	rows, _, err := h.svc.ListSpecialEquipment(c.Request.Context(), service.ListParams{
		Page: 1, PerPage: 10000, CompanyID: companyID,
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	setXlsxAttachment(c, "特种设备台账.xlsx", "special_equipment.xlsx")
	if err := excel.SpecialEquipmentExport(c.Writer, rows); err != nil {
		response.Internal(c, "导出失败")
		return
	}
}

// TemplateSpecialEquipment: GET /api/v1/special-equipment/template
func (h *Handler) TemplateSpecialEquipment(c *gin.Context) {
	setXlsxAttachment(c, "特种设备导入模板.xlsx", "special_equipment_template.xlsx")
	if err := excel.SpecialEquipmentTemplate(c.Writer); err != nil {
		response.Internal(c, "生成模板失败")
		return
	}
}

// CopySpecialEquipment: POST /api/v1/special-equipment/copy
func (h *Handler) CopySpecialEquipment(c *gin.Context) {
	var req copyReq
	if err := c.ShouldBindJSON(&req); err != nil || len(req.IDs) == 0 {
		response.BadRequest(c, "缺少 ids")
		return
	}
	ids, err := h.svc.CopySpecialEquipment(c.Request.Context(), req.IDs, middleware.CurrentUserID(c))
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, gin.H{"ids": ids, "count": len(ids)})
}

// copyReq — 与 fire-equipment / fire-facilities 复制接口共用的请求体形状。
type copyReq struct {
	IDs []int64 `json:"ids"`
}

// ImportSpecialEquipment: POST /api/v1/special-equipment/import (multipart form, field: file)
func (h *Handler) ImportSpecialEquipment(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		response.BadRequest(c, "缺少 file 字段")
		return
	}
	defer file.Close()
	rows, err := excel.SpecialEquipmentImport(file)
	if err != nil {
		response.BadRequest(c, "解析 Excel 失败: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	imported := 0
	for i := range rows {
		rows[i].CompanyID = companyID
		rows[i].UserID = uid
		rows[i].LastUserID = uid
		if err := h.svc.CreateSpecialEquipment(c.Request.Context(), &rows[i]); err == nil {
			imported++
		}
	}
	response.Success(c, gin.H{"imported": imported, "total": len(rows)})
}

func buildSpecialUpdates(req specialEquipmentReq, uid int64) map[string]any {
	u := map[string]any{"last_user_id": uid}
	if req.EquipmentName != nil {
		u["equipment_name"] = req.EquipmentName
	}
	if req.SpecialTypeID != nil {
		u["special_type_id"] = req.SpecialTypeID
	}
	if req.UtilizedLocation != nil {
		u["utilized_location"] = req.UtilizedLocation
	}
	if req.NormModel1 != nil {
		u["norm_model1"] = req.NormModel1
	}
	if req.NormModel2 != nil {
		u["norm_model2"] = req.NormModel2
	}
	if req.NormModel3 != nil {
		u["norm_model3"] = req.NormModel3
	}
	if req.NormModel4 != nil {
		u["norm_model4"] = req.NormModel4
	}
	if req.SerialNumber != nil {
		u["serial_number"] = req.SerialNumber
	}
	if t := parseDate(req.UseTime); t != nil {
		u["use_time"] = t
	}
	if t := parseDate(req.ValidPeriod); t != nil {
		u["valid_period"] = t
	}
	if req.InspectUnit != nil {
		u["inspect_unit"] = req.InspectUnit
	}
	if req.MaintainInfo != nil {
		u["maintain_info"] = req.MaintainInfo
	}
	if req.Comment != nil {
		u["comment"] = req.Comment
	}
	return u
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
