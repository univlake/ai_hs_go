package upload

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"ai_hs_go/internal/middleware"
	"ai_hs_go/internal/response"
	"ai_hs_go/internal/storage"
)

type Handler struct {
	oss *storage.Client
}

// NewHandler tolerates a nil OSS client (dev environments without keys); the
// public-key/sign endpoints will return 503 when not configured.
func NewHandler(oss *storage.Client) *Handler {
	return &Handler{oss: oss}
}

type initReq struct {
	Filename string `json:"filename" binding:"required"`
}

type initResp struct {
	Key      string `json:"key"`
	UploadID string `json:"upload_id"`
}

// Init: POST /api/v1/upload-oss/init
func (h *Handler) Init(c *gin.Context) {
	if h.oss == nil {
		response.Failed(c, 503, "OSS 未配置")
		return
	}
	var req initReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	uid := middleware.CurrentUserID(c)
	scopeID, _ := middleware.ScopeCompanyID(c)
	key := buildKey(scopeID, uid, req.Filename)

	uploadID, err := h.oss.InitMultipart(key)
	if err != nil {
		response.Internal(c, "初始化上传失败: "+err.Error())
		return
	}
	response.Success(c, initResp{Key: key, UploadID: uploadID})
}

// SignPart: POST /api/v1/upload-oss/sign-part
type signPartReq struct {
	Key        string `json:"key" binding:"required"`
	UploadID   string `json:"upload_id" binding:"required"`
	PartNumber int    `json:"part_number" binding:"required,min=1"`
}

func (h *Handler) SignPart(c *gin.Context) {
	if h.oss == nil {
		response.Failed(c, 503, "OSS 未配置")
		return
	}
	var req signPartReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	url, err := h.oss.SignPartURL(req.Key, req.UploadID, req.PartNumber, 2*time.Hour)
	if err != nil {
		response.Internal(c, "签名失败")
		return
	}
	response.Success(c, gin.H{"url": url})
}

// Complete: POST /api/v1/upload-oss/complete
type completeReq struct {
	Key      string `json:"key" binding:"required"`
	UploadID string `json:"upload_id" binding:"required"`
	Parts    []struct {
		PartNumber int    `json:"PartNumber" binding:"required"`
		ETag       string `json:"ETag" binding:"required"`
	} `json:"parts" binding:"required"`
}

func (h *Handler) Complete(c *gin.Context) {
	if h.oss == nil {
		response.Failed(c, 503, "OSS 未配置")
		return
	}
	var req completeReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	parts := make([]storage.CompletePart, len(req.Parts))
	for i, p := range req.Parts {
		parts[i] = storage.CompletePart{PartNumber: p.PartNumber, ETag: p.ETag}
	}
	if err := h.oss.CompleteMultipart(req.Key, req.UploadID, parts); err != nil {
		response.Internal(c, "完成上传失败: "+err.Error())
		return
	}
	signed, err := h.oss.SignURL(req.Key, time.Hour, nil)
	if err != nil {
		signed = ""
	}
	response.Success(c, gin.H{"key": req.Key, "url": signed})
}

// Abort: POST /api/v1/upload-oss/abort
func (h *Handler) Abort(c *gin.Context) {
	if h.oss == nil {
		response.Failed(c, 503, "OSS 未配置")
		return
	}
	var req struct {
		Key      string `json:"key" binding:"required"`
		UploadID string `json:"upload_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "参数无效: "+err.Error())
		return
	}
	_ = h.oss.AbortMultipart(req.Key, req.UploadID)
	response.Success(c, gin.H{"message": "已终止"})
}

// SignedURL: GET /api/v1/upload-oss/signed?key=xxx&ttl=3600
func (h *Handler) SignedURL(c *gin.Context) {
	if h.oss == nil {
		response.Failed(c, 503, "OSS 未配置")
		return
	}
	key := c.Query("key")
	if key == "" {
		response.BadRequest(c, "缺少 key")
		return
	}
	ttl := time.Hour
	if v := c.Query("ttl"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			ttl = time.Duration(n) * time.Second
		}
	}
	url, err := h.oss.SignURL(key, ttl, nil)
	if err != nil {
		response.Internal(c, "签名失败")
		return
	}
	response.Success(c, gin.H{"url": url})
}

func buildKey(companyID, userID int64, filename string) string {
	return fmt.Sprintf("uploads/c%d/u%d/%s_%s", companyID, userID, time.Now().Format("20060102"), uuid.NewString()+"_"+filename)
}
