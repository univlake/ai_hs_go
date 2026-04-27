package company

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/middleware"
	"ai_hs_go/internal/response"
	"ai_hs_go/internal/service"
)

type Handler struct {
	svc *service.CompanyService
}

func NewHandler(svc *service.CompanyService) *Handler {
	return &Handler{svc: svc}
}

// PublicList: GET /api/v1/companies-public
func (h *Handler) PublicList(c *gin.Context) {
	rows, err := h.svc.PublicList(c.Request.Context())
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Success(c, rows)
}

// List: GET /api/v1/companies
//
// Query params (Laravel-compatible aliases for legacy frontend pages):
//   - parent_id (preferred) or pid → parent filter
//   - keyword (preferred) or title → fuzzy match against title / credit_code
//   - page, per_page → standard pagination
func (h *Handler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "20"))

	var parent *int64
	parentRaw := c.Query("parent_id")
	if parentRaw == "" {
		parentRaw = c.Query("pid")
	}
	if parentRaw != "" {
		if id, err := strconv.ParseInt(parentRaw, 10, 64); err == nil {
			parent = &id
		}
	}

	keyword := c.Query("keyword")
	if keyword == "" {
		keyword = c.Query("title")
	}

	var rootFilter *int64
	if !middleware.IsPlatformAdmin(c) {
		// non-platform admin always restricted to their own root
		if v, ok := c.Get(middleware.CtxKeyCompanyID); ok {
			if cid, ok := v.(int64); ok {
				rootFilter = &cid // simplification: scope middleware enforces real boundary
			}
		}
	}
	rows, total, err := h.svc.List(c.Request.Context(), service.CompanyListParams{
		Page: page, PerPage: per, ParentID: parent, RootIDFilter: rootFilter,
		Keyword: keyword,
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

// ReviewList: GET /api/v1/companies-review-list
//
// Returns the queue of pending companies (status = 0) the active user can
// approve. The "reviewer" is whichever company the request is currently
// scoped to (X-Scope-Company-Id) — for platform admins this is the
// impersonated tenant, so 集团一's queue stays out of 集团二's, and only
// when no scope is set does a platform admin see the unfiltered queue.
func (h *Handler) ReviewList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ := strconv.Atoi(c.DefaultQuery("per_page", "15"))
	params := service.ReviewListParams{
		Page:    page,
		PerPage: per,
		Title:   c.Query("title"),
	}
	if scopeID, ok := middleware.ScopeCompanyID(c); ok && scopeID > 0 {
		params.ReviewerCompanyID = scopeID
	} else if middleware.IsPlatformAdmin(c) {
		params.IsPlatform = true
	} else if v, ok := c.Get(middleware.CtxKeyCompanyID); ok {
		// Defensive fallback — scope middleware normally already resolved
		// this for tenant users.
		if cid, ok := v.(int64); ok {
			params.ReviewerCompanyID = cid
		}
	}
	rows, total, err := h.svc.ReviewList(c.Request.Context(), params)
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

// createReq accepts both `parent_id` (canonical) and `pid` (Laravel-era
// alias the legacy enterprise/subsidiaries page sends). When both are set
// `parent_id` wins.
type createReq struct {
	Title             string  `json:"title" binding:"required,max=255"`
	EnterpriseTypeID  *int64  `json:"enterprise_type_id"`
	OperationalTypeID int64   `json:"operational_type_id"`
	AreaID            *int64  `json:"area_id"`
	Address           *string `json:"address"`
	CreditCode        *string `json:"credit_code"`
	BusinessScope     *string `json:"business_scope"`
	ParentID          *int64  `json:"parent_id"`
	Pid               *int64  `json:"pid"`
	IsProduce         int16   `json:"is_produce"`
	IsBloc            int16   `json:"is_bloc"`
}

func (h *Handler) Create(c *gin.Context) {
	var req createReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	parent := req.ParentID
	if parent == nil {
		parent = req.Pid
	}
	if parent == nil && !middleware.IsPlatformAdmin(c) {
		response.Forbidden(c, "仅平台管理员可创建一级企业")
		return
	}
	in := service.CreateCompanyInput{
		Title:             req.Title,
		EnterpriseTypeID:  req.EnterpriseTypeID,
		OperationalTypeID: req.OperationalTypeID,
		AreaID:            req.AreaID,
		Address:           req.Address,
		CreditCode:        req.CreditCode,
		BusinessScope:     req.BusinessScope,
		ParentID:          parent,
		IsProduce:         req.IsProduce,
		IsBloc:            req.IsBloc,
		CreatedBy:         middleware.CurrentUserID(c),
	}
	if parent == nil {
		// Tenant root — CreateRoot still produces an admin user but the
		// caller (platform admin via /tenants flow) typically supplies its
		// own initial password and wants the admin object back, not the
		// short-lived plaintext. Return the company shape without
		// admin_credentials so /tenants and /companies stay symmetric.
		c2, _, err := h.svc.CreateRoot(c.Request.Context(), in, "")
		if err != nil {
			writeAppError(c, err)
			return
		}
		response.Created(c, c2)
		return
	}
	c2, creds, err := h.svc.Create(c.Request.Context(), in)
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Created(c, gin.H{
		"id":                  c2.ID,
		"title":               c2.Title,
		"parent_id":           c2.ParentID,
		"pid":                 c2.ParentID,
		"root_id":             c2.RootID,
		"status":              c2.Status,
		"is_bloc":             c2.IsBloc,
		"is_produce":          c2.IsProduce,
		"enterprise_type_id":  c2.EnterpriseTypeID,
		"operational_type_id": c2.OperationalTypeID,
		"credit_code":         c2.CreditCode,
		"address":             c2.Address,
		"created_at":          c2.CreatedAt,
		"admin_credentials":   creds,
	})
}

// updateReq accepts every editable field on the current-info page. Numeric
// columns (longitude / dimension / acreage) and id columns
// (enterprise_type_id / operational_type_id) come in as `flexNumber`
// because the frontend stores them as the raw <input> string and ships
// "120.5" or "" as JSON strings instead of numbers. flexNumber treats
// either form as valid and exposes a parsed float / int back to the
// service.
type updateReq struct {
	Title               *string     `json:"title"`
	EnterpriseTypeID    *flexNumber `json:"enterprise_type_id"`
	OperationalTypeID   *flexNumber `json:"operational_type_id"`
	AreaID              *flexNumber `json:"area_id"`
	Address             *string     `json:"address"`
	Established         *string     `json:"established"`
	CreditCode          *string     `json:"credit_code"`
	BusinessScope       *string     `json:"business_scope"`
	CompanyIntroduction *string     `json:"company_introduction"`
	Longitude           *flexNumber `json:"longitude"`
	Dimension           *flexNumber `json:"dimension"`
	GeographicRange     *string     `json:"geographic_range"`
	Acreage             *flexNumber `json:"acreage"`
	SuperviseName       *string     `json:"supervise_name"`
	SuperviseTel        *string     `json:"supervise_tel"`
	SuperviseFix        *string     `json:"supervise_fix"`
	SuperviseEmail      *string     `json:"supervise_email"`
	SupervisePosition   *string     `json:"supervise_position"`
	IsProduce           *int16      `json:"is_produce"`
	IsBloc              *int16      `json:"is_bloc"`
}

// flexNumber parses both JSON numbers and JSON strings as float64. Empty
// string / null become a zeroed value with `Empty=true` so the service
// can distinguish "user blanked the field" from "user did not touch it".
type flexNumber struct {
	Empty bool
	Float float64
}

func (f *flexNumber) UnmarshalJSON(data []byte) error {
	s := string(data)
	if s == "null" {
		f.Empty = true
		return nil
	}
	// JSON string form: "120.5" or "".
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		inner := s[1 : len(s)-1]
		if inner == "" {
			f.Empty = true
			return nil
		}
		v, err := strconv.ParseFloat(inner, 64)
		if err != nil {
			return fmt.Errorf("invalid number: %q", inner)
		}
		f.Float = v
		return nil
	}
	// Bare JSON number.
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	f.Float = v
	return nil
}

func (f *flexNumber) AsFloat() *float64 {
	if f == nil || f.Empty {
		return nil
	}
	v := f.Float
	return &v
}

func (f *flexNumber) AsInt64() *int64 {
	if f == nil || f.Empty {
		return nil
	}
	v := int64(f.Float)
	return &v
}

// Get: GET /api/v1/companies/:id — single company detail.
// Used by ScopeContext + Dashboard + the current-info edit page; preloads
// EnterpriseType / OperationalType so the type dropdowns render the correct
// label without a separate lookup.
func (h *Handler) Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.svc.Repo().GetByIDWithRelations(c.Request.Context(), id)
	if err != nil {
		response.NotFound(c, "企业不存在")
		return
	}
	response.Success(c, row)
}

func (h *Handler) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req updateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	row, err := h.svc.Update(c.Request.Context(), id, service.UpdateCompanyInput{
		Title:               req.Title,
		EnterpriseTypeID:    req.EnterpriseTypeID.AsInt64(),
		OperationalTypeID:   req.OperationalTypeID.AsInt64(),
		AreaID:              req.AreaID.AsInt64(),
		Address:             req.Address,
		Established:         req.Established,
		CreditCode:          req.CreditCode,
		BusinessScope:       req.BusinessScope,
		CompanyIntroduction: req.CompanyIntroduction,
		Longitude:           req.Longitude.AsFloat(),
		Dimension:           req.Dimension.AsFloat(),
		GeographicRange:     req.GeographicRange,
		Acreage:             req.Acreage.AsFloat(),
		SuperviseName:       req.SuperviseName,
		SuperviseTel:        req.SuperviseTel,
		SuperviseFix:        req.SuperviseFix,
		SuperviseEmail:      req.SuperviseEmail,
		SupervisePosition:   req.SupervisePosition,
		IsProduce:           req.IsProduce,
		IsBloc:              req.IsBloc,
		UpdatedBy:           middleware.CurrentUserID(c),
	})
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

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

// ToggleStatus: PATCH /companies/:id/toggle-status
func (h *Handler) ToggleStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	row, err := h.svc.ToggleStatus(c.Request.Context(), id, middleware.CurrentUserID(c))
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

// Review: POST /companies/:id/review
//
// Frontend approval page sends `{status: 1|2}` directly (1=enable / approve,
// 2=disabled / reject) — same shape as the Laravel project. We keep `approve`
// as a backwards-compatible boolean so old internal tools don't break.
type reviewReq struct {
	Status  *int16 `json:"status"`
	Approve *bool  `json:"approve"`
}

func (h *Handler) Review(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req reviewReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	var target int16
	switch {
	case req.Status != nil:
		target = *req.Status
	case req.Approve != nil:
		if *req.Approve {
			target = service.CompanyStatusEnable
		} else {
			target = service.CompanyStatusDisabled
		}
	default:
		response.BadRequest(c, "需提供 status 或 approve")
		return
	}
	row, err := h.svc.Review(c.Request.Context(), id, target, middleware.CurrentUserID(c))
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

// Tree: GET /companies-tree
//
// `root_id` (canonical) and `pid` (legacy alias) both pin the subtree to
// start from. With neither supplied, the active user's own company is used.
//
// Response shape mirrors Laravel: an **array of the root's direct children**
// (each with their own `companies`/`children` populated through the full
// depth) — NOT the root itself. The grouping page iterates this array, so
// returning a single object would break it ("nodes is not iterable").
func (h *Handler) Tree(c *gin.Context) {
	rootRaw := c.Query("root_id")
	if rootRaw == "" {
		rootRaw = c.Query("pid")
	}
	var rootID int64
	if rootRaw != "" {
		id, err := strconv.ParseInt(rootRaw, 10, 64)
		if err != nil {
			response.BadRequest(c, "无效 root_id")
			return
		}
		rootID = id
	} else {
		cid, ok := c.Get(middleware.CtxKeyCompanyID)
		if !ok {
			response.BadRequest(c, "缺少 root_id")
			return
		}
		rootID, _ = cid.(int64)
	}
	tree, err := h.svc.Tree(c.Request.Context(), rootID)
	if err != nil {
		writeAppError(c, err)
		return
	}
	if tree == nil {
		response.Success(c, []any{})
		return
	}
	// Children are populated alongside Companies in service.Tree; either
	// can serve as the array source. nil → empty array so the frontend's
	// `for ... of` doesn't blow up on no-children roots.
	out := tree.Children
	if out == nil {
		out = []*service.TreeNode{}
	}
	response.Success(c, out)
}

// Forest: GET /companies-forest (platform admin only)
func (h *Handler) Forest(c *gin.Context) {
	if !middleware.IsPlatformAdmin(c) {
		response.Forbidden(c, "仅平台管理员可访问")
		return
	}
	forest, err := h.svc.Forest(c.Request.Context())
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Success(c, forest)
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
