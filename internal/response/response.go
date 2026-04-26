package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Envelope struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
	Reason  string `json:"reason,omitempty"`
}

type Paginated struct {
	Data        any   `json:"data"`
	CurrentPage int   `json:"current_page"`
	LastPage    int   `json:"last_page"`
	PerPage     int   `json:"per_page"`
	Total       int64 `json:"total"`
}

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Envelope{Status: "success", Code: http.StatusOK, Data: data})
}

func Created(c *gin.Context, data any) {
	c.JSON(http.StatusCreated, Envelope{Status: "success", Code: http.StatusCreated, Data: data})
}

func NoContent(c *gin.Context) {
	c.JSON(http.StatusOK, Envelope{Status: "success", Code: http.StatusOK})
}

func Failed(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, Envelope{Status: "error", Code: code, Message: message})
}

func FailedWithReason(c *gin.Context, code int, message, reason string) {
	c.AbortWithStatusJSON(code, Envelope{Status: "error", Code: code, Message: message, Reason: reason})
}

func BadRequest(c *gin.Context, message string) {
	Failed(c, http.StatusBadRequest, message)
}

func Unauthorized(c *gin.Context, message string) {
	if message == "" {
		message = "未登录或登录已过期"
	}
	Failed(c, http.StatusUnauthorized, message)
}

func Forbidden(c *gin.Context, message string) {
	if message == "" {
		message = "无权访问"
	}
	Failed(c, http.StatusForbidden, message)
}

func NotFound(c *gin.Context, message string) {
	if message == "" {
		message = "资源不存在"
	}
	Failed(c, http.StatusNotFound, message)
}

func Internal(c *gin.Context, message string) {
	if message == "" {
		message = "服务器内部错误"
	}
	Failed(c, http.StatusInternalServerError, message)
}

func Page(c *gin.Context, items any, page, perPage int, total int64) {
	last := 1
	if perPage > 0 && total > 0 {
		last = int((total + int64(perPage) - 1) / int64(perPage))
	}
	Success(c, Paginated{
		Data:        items,
		CurrentPage: page,
		LastPage:    last,
		PerPage:     perPage,
		Total:       total,
	})
}
