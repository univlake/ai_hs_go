// Package blocexpert 实现集团专家库的 HTTP 入口：
//
//	GET    /bloc-experts            列表（带 user / company preload，按 nickname 模糊搜索）
//	GET    /bloc-experts/:id        Laravel show 实际取的是 User，frontend 用 user.company.title
//	POST   /bloc-experts            新增（company_id + expert_user_id；唯一约束）
//	DELETE /bloc-experts/:id        移除
//	GET    /bloc-experts-list       「添加专家」面板：列出 status=1 且尚未在该企业专家库中的用户
//
// 跟 Laravel BlocExpertsController 路由对照见 routes/api/statistics-reporting.php。
package blocexpert

import (
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
	svc *service.ExpertService
}

func NewHandler(svc *service.ExpertService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) List(c *gin.Context) {
	page, per := apiutil.ParsePagination(c, 15)
	cid, _ := strconv.ParseInt(c.Query("company_id"), 10, 64)
	if cid == 0 {
		response.BadRequest(c, "company_id 不能为空")
		return
	}
	rows, total, err := h.svc.ListBlocExpert(c.Request.Context(), service.BlocExpertListParams{
		Page: page, PerPage: per, CompanyID: cid, Nickname: c.Query("nickname"),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

// Get — Laravel: User::query()->with('company')->find($id)。
// 这里 :id 被当成 user.id（不是 bloc_expert_users.id），跟前端 BlocExpertRow.user 字段保持一致。
func (h *Handler) Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.svc.GetBlocExpertUserByID(c.Request.Context(), id)
	if err != nil {
		writeAppError(c, err, "专家不存在")
		return
	}
	response.Success(c, row)
}

type createReq struct {
	CompanyID    *int64 `json:"company_id"`
	ExpertUserID *int64 `json:"expert_user_id"`
}

func (h *Handler) Create(c *gin.Context) {
	var req createReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	if req.CompanyID == nil || *req.CompanyID == 0 {
		response.BadRequest(c, "company_id 不能为空")
		return
	}
	if req.ExpertUserID == nil || *req.ExpertUserID == 0 {
		response.BadRequest(c, "expert_user_id 不能为空")
		return
	}
	uid := middleware.CurrentUserID(c)
	m := &model.BlocExpertUser{
		CompanyID:    *req.CompanyID,
		ExpertUserID: *req.ExpertUserID,
	}
	m.UserID = uid
	m.LastUserID = uid
	if err := h.svc.CreateBlocExpert(c.Request.Context(), m); err != nil {
		writeAppError(c, err, "添加失败")
		return
	}
	response.Created(c, m)
}

func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.svc.DeleteBlocExpert(c.Request.Context(), id); err != nil {
		writeAppError(c, err, "移除失败")
		return
	}
	response.Success(c, gin.H{"message": "删除成功"})
}

// AvailableList — 「添加专家」面板：返回所有 status=1 且不在该企业专家库的 User。
func (h *Handler) AvailableList(c *gin.Context) {
	cid, _ := strconv.ParseInt(c.Query("company_id"), 10, 64)
	if cid == 0 {
		response.BadRequest(c, "company_id 不能为空")
		return
	}
	rows, err := h.svc.ListAvailableExperts(c.Request.Context(), cid, c.Query("nickname"))
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Success(c, rows)
}

func writeAppError(c *gin.Context, err error, fallback string) {
	if ae, ok := errs.As(err); ok {
		switch ae.Code {
		case errs.CodeNotFound:
			response.NotFound(c, ae.Message)
		case errs.CodeUnprocessable:
			response.Failed(c, http.StatusUnprocessableEntity, ae.Message)
		case errs.CodeConflict:
			response.Failed(c, http.StatusConflict, ae.Message)
		default:
			response.Internal(c, ae.Message)
		}
		return
	}
	response.Internal(c, fallback)
}
