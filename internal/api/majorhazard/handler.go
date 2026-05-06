// Package majorhazard 实现 /major-hazards CRUD（重大危险源）。
//
// 职责切分参照 internal/api/safetyledger 的 hazardous_chemical.go：
//   - 主表 partial update 用 map[string]any，nil 字段不动
//   - 子表 constitutes 整体替换：前端每次提交都把全量数组发上来
package majorhazard

import (
	"bytes"
	"encoding/json"
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

// flexInt64 接受 JSON 里的数字或字符串形式的整数，方便前端 <select> 直接把
// option.value（永远是 string）丢进 payload，不需要在 React 侧手动 Number()。
// 空字符串 / null / 缺失字段一律视为未提供。
type flexInt64 struct {
	value *int64
}

func (f *flexInt64) UnmarshalJSON(data []byte) error {
	trimmed := bytes.TrimSpace(data)
	if len(trimmed) == 0 || bytes.Equal(trimmed, []byte("null")) || bytes.Equal(trimmed, []byte(`""`)) {
		return nil
	}
	if trimmed[0] == '"' {
		var s string
		if err := json.Unmarshal(trimmed, &s); err != nil {
			return err
		}
		if s == "" {
			return nil
		}
		v, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return err
		}
		f.value = &v
		return nil
	}
	var v int64
	if err := json.Unmarshal(trimmed, &v); err != nil {
		return err
	}
	f.value = &v
	return nil
}

func (f flexInt64) ptr() *int64 { return f.value }

type Handler struct {
	svc *service.MajorHazardService
}

func NewHandler(svc *service.MajorHazardService) *Handler {
	return &Handler{svc: svc}
}

type constituteReq struct {
	Name                *string   `json:"name"`
	HazardousChemicalID flexInt64 `json:"hazardous_chemical_id"`
	Number              *string   `json:"number"`
}

type majorHazardReq struct {
	Name            *string         `json:"name"`
	Level           *string         `json:"level"`
	Position        *string         `json:"position"`
	RValue          *string         `json:"r_value"`
	IsSIS           int16           `json:"is_sis"`
	SILLevel        *string         `json:"sil_level"`
	DetectParameter *string         `json:"detect_parameter"`
	Constitutes     []constituteReq `json:"constitutes"`
}

// List: GET /api/v1/major-hazards
func (h *Handler) List(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	page, per := apiutil.ParsePagination(c, 20)
	rows, total, err := h.svc.List(c.Request.Context(), service.MajorHazardListParams{
		Page: page, PerPage: per, CompanyID: companyID,
		// 前端搜索字段是 name（见 major-hazards/Index.tsx onSearchChange）；
		// 同时兼容通用的 keyword 参数，便于其它客户端调用。
		Name: firstNonEmpty(c.Query("name"), c.Query("keyword")),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

// Create: POST /api/v1/major-hazards
func (h *Handler) Create(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	var req majorHazardReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	hazard := &model.MajorHazard{
		Name:            req.Name,
		Level:           req.Level,
		Position:        req.Position,
		RValue:          req.RValue,
		IsSIS:           req.IsSIS,
		SILLevel:        req.SILLevel,
		DetectParameter: req.DetectParameter,
		CompanyID:       companyID,
	}
	hazard.UserID = uid
	hazard.LastUserID = uid
	constitutes := mapConstitutes(req.Constitutes, uid)
	if err := h.svc.Create(c.Request.Context(), service.CreateMajorHazardInput{
		Hazard:      hazard,
		Constitutes: constitutes,
	}); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.Get(c.Request.Context(), hazard.ID)
	response.Created(c, row)
}

// Update: PUT /api/v1/major-hazards/:id
func (h *Handler) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req majorHazardReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	updates := map[string]any{
		"last_user_id": uid,
		// is_sis 是 SMALLINT，零值即「未配置」，所以始终写入而不做 nil 判断
		"is_sis": req.IsSIS,
	}
	addNullable(updates, "name", req.Name)
	addNullable(updates, "level", req.Level)
	addNullable(updates, "position", req.Position)
	addNullable(updates, "r_value", req.RValue)
	addNullable(updates, "sil_level", req.SILLevel)
	addNullable(updates, "detect_parameter", req.DetectParameter)

	// constitutes 字段未提供时（nil）保留原数据；提供空数组时清空。
	var constitutes []model.MajorConstitute
	if req.Constitutes != nil {
		constitutes = mapConstitutes(req.Constitutes, uid)
	}
	if err := h.svc.Update(c.Request.Context(), id, updates, constitutes); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.Get(c.Request.Context(), id)
	response.Success(c, row)
}

// Delete: DELETE /api/v1/major-hazards/:id
func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

// ── helpers ─────────────────────────────────────────────────────────────

func mapConstitutes(items []constituteReq, uid int64) []model.MajorConstitute {
	out := make([]model.MajorConstitute, len(items))
	for i, it := range items {
		out[i] = model.MajorConstitute{
			Name:                it.Name,
			HazardousChemicalID: it.HazardousChemicalID.ptr(),
			Number:              it.Number,
		}
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
