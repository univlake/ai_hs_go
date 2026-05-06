package basedata

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ai_hs_go/internal/errs"
	"ai_hs_go/internal/middleware"
	"ai_hs_go/internal/response"
	"ai_hs_go/internal/service"
	"ai_hs_go/internal/api/apiutil"
)

// flexInt64 unmarshals both JSON numbers (`5`) and JSON strings (`"5"`)
// into an int64. Empty string / null → 0. Used by request fields that the
// frontend Select sends as strings (option value keys are stringified
// before submit) but the backend models as int64.
type flexInt64 int64

func (f *flexInt64) UnmarshalJSON(data []byte) error {
	s := string(data)
	if s == "null" {
		*f = 0
		return nil
	}
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		inner := s[1 : len(s)-1]
		if inner == "" {
			*f = 0
			return nil
		}
		v, err := strconv.ParseInt(inner, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid integer: %q", inner)
		}
		*f = flexInt64(v)
		return nil
	}
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	*f = flexInt64(v)
	return nil
}

// Handler exposes the base-data dictionary endpoints — 9 type tables
// (enterprise / operational / record / process / special / space-work /
// safety-management / dangerous-craft / regulation-library) plus the
// checklist library. Every type table shares the same shape, so a single
// generic handler set is dispatched by URL slug; ChecklistLibrary has its
// own typed handler because the row carries extra columns.
type Handler struct {
	dictSvc      *service.DictService
	checklistSvc *service.ChecklistLibraryService
	regLibSvc    *service.RegulationLibraryTypeService
}

func NewHandler(
	dictSvc *service.DictService,
	checklistSvc *service.ChecklistLibraryService,
	regLibSvc *service.RegulationLibraryTypeService,
) *Handler {
	return &Handler{
		dictSvc:      dictSvc,
		checklistSvc: checklistSvc,
		regLibSvc:    regLibSvc,
	}
}

// dictHandler returns a closure that captures the dictionary slug for a
// specific endpoint group. Used by router.go to register
// `/enterprise-types`, `/operational-types`, …, all backed by the same
// service methods. Keeping the closure in a single helper means the
// permission middleware and slug whitelist live in exactly one place.
func (h *Handler) DictList(slug string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !h.dictSvc.IsKnownDict(slug) {
			response.NotFound(c, "未知字典")
			return
		}
		page, per := apiutil.ParsePagination(c, 20)
		rows, total, err := h.dictSvc.List(c.Request.Context(), slug, service.DictListParams{
			Page: page, PerPage: per, Title: c.Query("title"),
		})
		if err != nil {
			writeAppError(c, err)
			return
		}
		response.Page(c, rows, page, per, total)
	}
}

func (h *Handler) DictCreate(slug string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !h.dictSvc.IsKnownDict(slug) {
			response.NotFound(c, "未知字典")
			return
		}
		var req struct {
			Title string `json:"title" binding:"required,max=255"`
			Sort  int    `json:"sort"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			response.BadRequest(c, "参数无效: "+err.Error())
			return
		}
		row, err := h.dictSvc.Create(c.Request.Context(), slug, service.DictMutateInput{
			Title: req.Title, Sort: req.Sort,
		})
		if err != nil {
			writeAppError(c, err)
			return
		}
		response.Created(c, row)
	}
}

func (h *Handler) DictUpdate(slug string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !h.dictSvc.IsKnownDict(slug) {
			response.NotFound(c, "未知字典")
			return
		}
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			response.BadRequest(c, "无效 id")
			return
		}
		var req struct {
			Title string `json:"title" binding:"required,max=255"`
			Sort  int    `json:"sort"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			response.BadRequest(c, "参数无效: "+err.Error())
			return
		}
		row, err := h.dictSvc.Update(c.Request.Context(), slug, id, service.DictMutateInput{
			Title: req.Title, Sort: req.Sort,
		})
		if err != nil {
			writeAppError(c, err)
			return
		}
		response.Success(c, row)
	}
}

func (h *Handler) DictDelete(slug string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !h.dictSvc.IsKnownDict(slug) {
			response.NotFound(c, "未知字典")
			return
		}
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			response.BadRequest(c, "无效 id")
			return
		}
		if err := h.dictSvc.Delete(c.Request.Context(), slug, id); err != nil {
			writeAppError(c, err)
			return
		}
		response.NoContent(c)
	}
}

// ── Regulation library types (multi-level dictionary) ────────────

type regLibTypeReq struct {
	Title string    `json:"title" binding:"required,max=255"`
	Pid   flexInt64 `json:"pid"`
	Sort  int       `json:"sort"`
}

// RegLibTypeList: GET /api/v1/regulation-library-types
//
// Accepts ?pid=N to scope to one level (Laravel-style). Without pid we
// return everything paginated; the frontend filters / displays via
// `all_name` and the parent picker.
func (h *Handler) RegLibTypeList(c *gin.Context) {
	page, per := apiutil.ParsePagination(c, 200)
	params := service.RegLibTypeListParams{
		Page:    page,
		PerPage: per,
		Title:   c.Query("title"),
	}
	if v := c.Query("pid"); v != "" {
		if id, err := strconv.ParseInt(v, 10, 64); err == nil {
			params.Pid = &id
		}
	}
	rows, total, err := h.regLibSvc.List(c.Request.Context(), params)
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

func (h *Handler) RegLibTypeCreate(c *gin.Context) {
	var req regLibTypeReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	row, err := h.regLibSvc.Create(c.Request.Context(), service.RegLibTypeInput{
		Title:  req.Title,
		Pid:    int64(req.Pid),
		Sort:   req.Sort,
		UserID: middleware.CurrentUserID(c),
	})
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Created(c, row)
}

func (h *Handler) RegLibTypeUpdate(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req regLibTypeReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	row, err := h.regLibSvc.Update(c.Request.Context(), id, service.RegLibTypeInput{
		Title:  req.Title,
		Pid:    int64(req.Pid),
		Sort:   req.Sort,
		UserID: middleware.CurrentUserID(c),
	})
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

func (h *Handler) RegLibTypeDelete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.regLibSvc.Delete(c.Request.Context(), id); err != nil {
		writeAppError(c, err)
		return
	}
	response.NoContent(c)
}

// ── Checklist libraries ───────────────────────────────────────────

type checklistReq struct {
	Title           string  `json:"title"`
	Status          *int16  `json:"status"`
	PreviewAnnexURL *string `json:"preview_annex_url"`
	AnnexURL        *string `json:"annex_url"`
}

// ChecklistList: GET /api/v1/checklist-libraries
func (h *Handler) ChecklistList(c *gin.Context) {
	page, per := apiutil.ParsePagination(c, 20)
	cid, _ := middleware.ScopeCompanyID(c)
	rows, total, err := h.checklistSvc.List(c.Request.Context(), service.ChecklistLibraryListParams{
		Page:      page,
		PerPage:   per,
		CompanyID: cid,
		Title:     c.Query("title"),
	})
	if err != nil {
		response.Internal(c, "查询失败")
		return
	}
	response.Page(c, rows, page, per, total)
}

// ChecklistCreate: POST /api/v1/checklist-libraries
func (h *Handler) ChecklistCreate(c *gin.Context) {
	cid, ok := middleware.ScopeCompanyID(c)
	if !ok || cid == 0 {
		response.BadRequest(c, "缺少 X-Scope-Company-Id")
		return
	}
	var req checklistReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	row, err := h.checklistSvc.Create(c.Request.Context(), service.ChecklistLibraryInput{
		Title:           req.Title,
		Status:          req.Status,
		PreviewAnnexURL: req.PreviewAnnexURL,
		AnnexURL:        req.AnnexURL,
		CompanyID:       cid,
		UserID:          middleware.CurrentUserID(c),
	})
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Created(c, row)
}

// ChecklistUpdate: PUT /api/v1/checklist-libraries/:id
func (h *Handler) ChecklistUpdate(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	var req checklistReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	row, err := h.checklistSvc.Update(c.Request.Context(), id, service.ChecklistLibraryInput{
		Title:           req.Title,
		Status:          req.Status,
		PreviewAnnexURL: req.PreviewAnnexURL,
		AnnexURL:        req.AnnexURL,
		UserID:          middleware.CurrentUserID(c),
	})
	if err != nil {
		writeAppError(c, err)
		return
	}
	response.Success(c, row)
}

// ChecklistDelete: DELETE /api/v1/checklist-libraries/:id
func (h *Handler) ChecklistDelete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, "无效 id")
		return
	}
	if err := h.checklistSvc.Delete(c.Request.Context(), id); err != nil {
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
