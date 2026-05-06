package safetyledger

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"ai_hs_go/internal/excel"
	"ai_hs_go/internal/middleware"
	"ai_hs_go/internal/model"
	"ai_hs_go/internal/response"
	"ai_hs_go/internal/service"
	"ai_hs_go/internal/api/apiutil"
)

// ListDangerousEquipment: GET /api/v1/dangerous-equipment
func (h *Handler) ListDangerousEquipment(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	page, per := apiutil.ParsePagination(c, 20)
	rows, total, err := h.svc.ListDangerousEquipment(c.Request.Context(), service.ListParams{
		Page: page, PerPage: per, CompanyID: companyID, Keyword: c.Query("keyword"),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

type dangerousEquipmentReq struct {
	EquipmentName    *string `json:"equipment_name"`
	UtilizedLocation *string `json:"utilized_location"`
	NormModel        *string `json:"norm_model"`
	SerialNumber     *string `json:"serial_number"`
	UseTime          *string `json:"use_time"`
	AutomationType   *string `json:"automation_type"`
	MaintainInfo     *string `json:"maintain_info"`
	Comment          *string `json:"comment"`
}

func (h *Handler) CreateDangerousEquipment(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	var req dangerousEquipmentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	row := &model.DangerousEquipment{
		EquipmentName:    req.EquipmentName,
		UtilizedLocation: req.UtilizedLocation,
		NormModel:        req.NormModel,
		SerialNumber:     req.SerialNumber,
		UseTime:          parseDate(req.UseTime),
		AutomationType:   req.AutomationType,
		MaintainInfo:     req.MaintainInfo,
		Comment:          req.Comment,
		CompanyID:        companyID,
	}
	row.UserID = uid
	row.LastUserID = uid
	if err := h.svc.CreateDangerousEquipment(c.Request.Context(), row); err != nil {
		writeAppError(c, err)
		return
	}
	response.Created(c, row)
}

func (h *Handler) UpdateDangerousEquipment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req dangerousEquipmentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	updates := map[string]any{"last_user_id": middleware.CurrentUserID(c)}
	if req.EquipmentName != nil {
		updates["equipment_name"] = req.EquipmentName
	}
	if req.UtilizedLocation != nil {
		updates["utilized_location"] = req.UtilizedLocation
	}
	if req.NormModel != nil {
		updates["norm_model"] = req.NormModel
	}
	if req.SerialNumber != nil {
		updates["serial_number"] = req.SerialNumber
	}
	if t := parseDate(req.UseTime); t != nil {
		updates["use_time"] = t
	}
	if req.AutomationType != nil {
		updates["automation_type"] = req.AutomationType
	}
	if req.MaintainInfo != nil {
		updates["maintain_info"] = req.MaintainInfo
	}
	if req.Comment != nil {
		updates["comment"] = req.Comment
	}
	if err := h.svc.UpdateDangerousEquipment(c.Request.Context(), id, updates); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetDangerousEquipment(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteDangerousEquipment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteDangerousEquipment(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

// TemplateDangerousEquipment: GET /api/v1/dangerous-equipment/template
func (h *Handler) TemplateDangerousEquipment(c *gin.Context) {
	setXlsxAttachment(c, "重大危险设备导入模板.xlsx", "dangerous_equipment_template.xlsx")
	if err := excel.DangerousEquipmentTemplate(c.Writer); err != nil {
		response.Internal(c, "生成模板失败")
		return
	}
}

// ExportDangerousEquipment: GET /api/v1/dangerous-equipment/export
func (h *Handler) ExportDangerousEquipment(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	rows, _, err := h.svc.ListDangerousEquipment(c.Request.Context(), service.ListParams{
		Page: 1, PerPage: 10000, CompanyID: companyID,
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	setXlsxAttachment(c, "重大危险设备台账.xlsx", "dangerous_equipment.xlsx")
	if err := excel.DangerousEquipmentExport(c.Writer, rows); err != nil {
		response.Internal(c, "导出失败")
		return
	}
}

// CopyDangerousEquipment: POST /api/v1/dangerous-equipment/copy
func (h *Handler) CopyDangerousEquipment(c *gin.Context) {
	var req copyReq
	if err := c.ShouldBindJSON(&req); err != nil || len(req.IDs) == 0 {
		response.BadRequest(c, "缺少 ids")
		return
	}
	ids, err := h.svc.CopyDangerousEquipment(c.Request.Context(), req.IDs, middleware.CurrentUserID(c))
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, gin.H{"ids": ids, "count": len(ids)})
}

// ImportDangerousEquipment: POST /api/v1/dangerous-equipment/import
func (h *Handler) ImportDangerousEquipment(c *gin.Context) {
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
	rows, err := excel.DangerousEquipmentImport(file)
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
		if err := h.svc.CreateDangerousEquipment(c.Request.Context(), &rows[i]); err == nil {
			imported++
		}
	}
	response.Success(c, gin.H{"imported": imported, "total": len(rows)})
}
