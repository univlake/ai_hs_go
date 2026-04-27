// Package todo 实现待办事项的占位入口：
//
//	GET /todos              占位返回 []
//	GET /todos-list-table   占位返回 []
//
// Laravel 的 TodosController 同样是占位实现（仅校验 company_id / type 参数后
// 返回空数组）。todos 业务尚未排期，先保持 API 形态可用，后续模块上线时再
// 把空数组替换成真实聚合查询即可。前端仪表板对空数组渲染为「暂无待办」。
package todo

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"ai_hs_go/internal/response"
)

type Handler struct{}

func NewHandler() *Handler { return &Handler{} }

// Index — Laravel 校验 (company_id required + type required)；占位实现保持同样的入参语义，
// 否则前端在传缺失参数时会拿到 200 + [] 而不是 400，可能掩盖调用错误。
func (h *Handler) Index(c *gin.Context) {
	cid, _ := strconv.ParseInt(c.Query("company_id"), 10, 64)
	if cid == 0 {
		response.BadRequest(c, "company_id 不能为空")
		return
	}
	if c.Query("type") == "" {
		response.BadRequest(c, "type 不能为空")
		return
	}
	response.Success(c, []any{})
}

// ListTable — Laravel 同样的占位，仅校验 company_id。
func (h *Handler) ListTable(c *gin.Context) {
	cid, _ := strconv.ParseInt(c.Query("company_id"), 10, 64)
	if cid == 0 {
		response.BadRequest(c, "company_id 不能为空")
		return
	}
	response.Success(c, []any{})
}
