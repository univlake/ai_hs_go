package safetyledger

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"ai_hs_go/internal/middleware"
	"ai_hs_go/internal/model"
	"ai_hs_go/internal/response"
	"ai_hs_go/internal/service"
	"ai_hs_go/internal/api/apiutil"
)

func (h *Handler) ListHazardousChemicals(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	page, per := apiutil.ParsePagination(c, 20)
	rows, total, err := h.svc.ListHazardousChemicals(c.Request.Context(), service.ListParams{
		Page: page, PerPage: per, CompanyID: companyID, Keyword: c.Query("keyword"),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

type ingredientReq struct {
	Name    *string `json:"name"`
	Content *string `json:"content"`
	CAS     *string `json:"cas"`
}

type hazardousChemicalReq struct {
	Name           *string         `json:"name"`
	Alias          *string         `json:"alias"`
	IsSupervision  int16           `json:"is_supervision"`
	IsToxin        int16           `json:"is_toxin"`
	IsExplosion    int16           `json:"is_explosion"`
	IsEasyExplode  int16           `json:"is_easy_explode"`
	IsNarcotics    int16           `json:"is_narcotics"`
	IsMonitor      int16           `json:"is_monitor"`
	IsMixing       int16           `json:"is_mixing"`
	CAS            *string         `json:"cas"`
	Position       *string         `json:"position"`
	Temperature    *string         `json:"temperature"`
	Storage        *string         `json:"storage"`
	YearUse        *string         `json:"year_use"`
	YearUseUnit    *string         `json:"year_use_unit"`
	StorageUnit    *string         `json:"storage_unit"`
	Pressure       *string         `json:"pressure"`
	PressureUnit   *string         `json:"pressure_unit"`
	FireClass      *string         `json:"fire_class"`
	MatterState    *string         `json:"matter_state"`
	MSDS           *string         `json:"msds"`
	Ingredients    []ingredientReq `json:"ingredients"`
}

func (h *Handler) CreateHazardousChemical(c *gin.Context) {
	companyID := scopeCompanyOrFail(c)
	if companyID == 0 {
		return
	}
	var req hazardousChemicalReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	chem := &model.HazardousChemical{
		Name:          req.Name,
		Alias:         req.Alias,
		IsSupervision: req.IsSupervision,
		IsToxin:       req.IsToxin,
		IsExplosion:   req.IsExplosion,
		IsEasyExplode: req.IsEasyExplode,
		IsNarcotics:   req.IsNarcotics,
		IsMonitor:     req.IsMonitor,
		IsMixing:      req.IsMixing,
		CAS:           req.CAS,
		Position:      req.Position,
		Temperature:   req.Temperature,
		Storage:       req.Storage,
		YearUse:       req.YearUse,
		YearUseUnit:   req.YearUseUnit,
		StorageUnit:   req.StorageUnit,
		Pressure:      req.Pressure,
		PressureUnit:  req.PressureUnit,
		FireClass:     req.FireClass,
		MatterState:   req.MatterState,
		MSDS:          req.MSDS,
		CompanyID:     companyID,
	}
	chem.UserID = uid
	chem.LastUserID = uid
	ingredients := mapIngredients(req.Ingredients, uid)
	if err := h.svc.CreateHazardousChemical(c.Request.Context(), service.CreateHazardousChemicalInput{
		Chemical:    chem,
		Ingredients: ingredients,
	}); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetHazardousChemical(c.Request.Context(), chem.ID)
	response.Created(c, row)
}

func (h *Handler) UpdateHazardousChemical(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req hazardousChemicalReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	updates := map[string]any{
		"last_user_id":    uid,
		"is_supervision":  req.IsSupervision,
		"is_toxin":        req.IsToxin,
		"is_explosion":    req.IsExplosion,
		"is_easy_explode": req.IsEasyExplode,
		"is_narcotics":    req.IsNarcotics,
		"is_monitor":      req.IsMonitor,
		"is_mixing":       req.IsMixing,
	}
	addNullable(updates, "name", req.Name)
	addNullable(updates, "alias", req.Alias)
	addNullable(updates, "cas", req.CAS)
	addNullable(updates, "position", req.Position)
	addNullable(updates, "temperature", req.Temperature)
	addNullable(updates, "storage", req.Storage)
	addNullable(updates, "year_use", req.YearUse)
	addNullable(updates, "year_use_unit", req.YearUseUnit)
	addNullable(updates, "storage_unit", req.StorageUnit)
	addNullable(updates, "pressure", req.Pressure)
	addNullable(updates, "pressure_unit", req.PressureUnit)
	addNullable(updates, "fire_class", req.FireClass)
	addNullable(updates, "matter_state", req.MatterState)
	addNullable(updates, "msds", req.MSDS)

	var ingredients []model.ChemicalIngredient
	if req.Ingredients != nil {
		ingredients = mapIngredients(req.Ingredients, uid)
	}
	if err := h.svc.UpdateHazardousChemical(c.Request.Context(), id, updates, ingredients); err != nil {
		writeAppError(c, err)
		return
	}
	row, _ := h.svc.GetHazardousChemical(c.Request.Context(), id)
	response.Success(c, row)
}

func (h *Handler) DeleteHazardousChemical(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteHazardousChemical(c.Request.Context(), id); err != nil {
		response.Internal(c, "删除失败")
		return
	}
	response.NoContent(c)
}

func mapIngredients(items []ingredientReq, uid int64) []model.ChemicalIngredient {
	out := make([]model.ChemicalIngredient, len(items))
	for i, it := range items {
		out[i] = model.ChemicalIngredient{
			Name:    it.Name,
			Content: it.Content,
			CAS:     it.CAS,
		}
		out[i].UserID = uid
		out[i].LastUserID = uid
	}
	return out
}
