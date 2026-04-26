package permission

import (
	"github.com/gin-gonic/gin"

	"ai_hs_go/internal/response"
	"ai_hs_go/internal/service"
)

type Handler struct {
	rbac *service.RBACService
}

func NewHandler(rbac *service.RBACService) *Handler {
	return &Handler{rbac: rbac}
}

// List: GET /api/v1/permissions
// Returns flat permission list and a grouped view for the frontend dropdowns.
func (h *Handler) List(c *gin.Context) {
	perms, err := h.rbac.ListPermissions(c.Request.Context())
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	grouped := map[string][]string{}
	for _, p := range perms {
		mod := "*"
		for i := len(p.Name) - 1; i >= 0; i-- {
			if p.Name[i] == '.' {
				mod = p.Name[:i]
				break
			}
		}
		grouped[mod] = append(grouped[mod], p.Name)
	}
	response.Success(c, gin.H{
		"permissions": perms,
		"grouped":     grouped,
	})
}
