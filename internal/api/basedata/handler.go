package basedata

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"ai_hs_go/internal/model"
	"ai_hs_go/internal/response"
)

// Handler exposes shared dictionary endpoints (special types, dangerous
// craft types, etc.). These are tenant-agnostic — anyone authenticated can
// read them so dropdowns work across the app.
type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{db: db}
}

// SpecialTypes: GET /api/v1/special-types
func (h *Handler) SpecialTypes(c *gin.Context) {
	var rows []model.SpecialType
	if err := h.db.WithContext(c.Request.Context()).
		Order("id").
		Find(&rows).Error; err != nil {
		response.Internal(c, "查询特种类型失败")
		return
	}
	response.Success(c, rows)
}

// EnterpriseTypes: GET /api/v1/enterprise-types — backs the type dropdown
// on the enterprise create / edit forms.
func (h *Handler) EnterpriseTypes(c *gin.Context) {
	var rows []model.EnterpriseType
	if err := h.db.WithContext(c.Request.Context()).
		Order("id").
		Find(&rows).Error; err != nil {
		response.Internal(c, "查询企业类型失败")
		return
	}
	response.Success(c, rows)
}

// OperationalTypes: GET /api/v1/operational-types — backs the operational
// type dropdown.
func (h *Handler) OperationalTypes(c *gin.Context) {
	var rows []model.OperationalType
	if err := h.db.WithContext(c.Request.Context()).
		Order("id").
		Find(&rows).Error; err != nil {
		response.Internal(c, "查询经营类型失败")
		return
	}
	response.Success(c, rows)
}
