package storage

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"

	"ai_hs_go/internal/config"
)

// Client wraps the Aliyun OSS Bucket and provides the small subset of
// operations the API actually uses (signed URL, multipart upload, web
// direct-upload policy).
type Client struct {
	bucket *oss.Bucket
	url    string

	// kept for browser-direct POST policy signing — the SDK only signs
	// PUT URLs but the frontend OssUploadField does form-POST against
	// the bucket host with a precomputed policy + signature.
	accessKeyID     string
	accessKeySecret string
	endpoint        string // for fallback host computation
}

func New(cfg config.OSSConfig) (*Client, error) {
	if cfg.AccessKeyID == "" || cfg.AccessKeySecret == "" || cfg.Bucket == "" {
		return nil, errors.New("oss config incomplete (access key/secret/bucket)")
	}
	cli, err := oss.New(cfg.Endpoint, cfg.AccessKeyID, cfg.AccessKeySecret)
	if err != nil {
		return nil, fmt.Errorf("oss client: %w", err)
	}
	bucket, err := cli.Bucket(cfg.Bucket)
	if err != nil {
		return nil, fmt.Errorf("oss bucket: %w", err)
	}
	return &Client{
		bucket:          bucket,
		url:             cfg.PublicURL,
		accessKeyID:     cfg.AccessKeyID,
		accessKeySecret: cfg.AccessKeySecret,
		endpoint:        cfg.Endpoint,
	}, nil
}

// PostPolicy describes a presigned form-POST upload that browsers can
// submit directly to the bucket host. Mirrors Laravel OssService::policy.
type PostPolicy struct {
	AccessKeyID string `json:"access_key_id"`
	Host        string `json:"host"`
	Policy      string `json:"policy"`
	Signature   string `json:"signature"`
	Expire      int64  `json:"expire"`
	Dir         string `json:"dir"`
}

// SignPostPolicy returns the parameters the frontend needs to do a
// browser-direct POST to the bucket. `dir` is the directory prefix the
// uploaded key must start with; the server enforces it via the policy
// `starts-with` condition.
func (c *Client) SignPostPolicy(dir string, ttl time.Duration) (*PostPolicy, error) {
	expire := time.Now().Add(ttl).Unix()
	expireISO := time.Unix(expire, 0).UTC().Format("2006-01-02T15:04:05.000Z")

	policy := map[string]any{
		"expiration": expireISO,
		"conditions": []any{
			map[string]string{"bucket": c.bucket.BucketName},
			[]any{"starts-with", "$key", dir},
			[]any{"content-length-range", 0, 100 * 1024 * 1024}, // up to 100 MB
		},
	}
	policyJSON, err := json.Marshal(policy)
	if err != nil {
		return nil, err
	}
	policyBase64 := base64.StdEncoding.EncodeToString(policyJSON)

	mac := hmac.New(sha1.New, []byte(c.accessKeySecret))
	mac.Write([]byte(policyBase64))
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	host := c.url
	if host == "" {
		// Fallback: <bucket>.<endpoint> with https://. Endpoint may already
		// have a scheme; strip it before joining.
		ep := c.endpoint
		ep = stripScheme(ep)
		host = "https://" + c.bucket.BucketName + "." + ep
	}

	return &PostPolicy{
		AccessKeyID: c.accessKeyID,
		Host:        host,
		Policy:      policyBase64,
		Signature:   signature,
		Expire:      expire,
		Dir:         dir,
	}, nil
}

func stripScheme(s string) string {
	for _, p := range []string{"https://", "http://"} {
		if len(s) > len(p) && s[:len(p)] == p {
			return s[len(p):]
		}
	}
	return s
}

// EnsureCORS configures bucket CORS so the browser can POST directly
// from `allowedOrigins` to the bucket host. Idempotent — overwrites the
// rules each time so config changes propagate without manual console
// work. Errors here are non-fatal; we log + continue so the server
// still starts when the access key lacks PutBucketCORS rights.
func (c *Client) EnsureCORS(allowedOrigins []string) error {
	rules := []oss.CORSRule{{
		AllowedOrigin: allowedOrigins,
		AllowedMethod: []string{"GET", "POST", "PUT", "HEAD"},
		AllowedHeader: []string{"*"},
		ExposeHeader:  []string{"ETag", "x-oss-request-id"},
		MaxAgeSeconds: 600,
	}}
	bucketName := c.bucket.BucketName
	return c.bucket.Client.SetBucketCORS(bucketName, rules)
}

// SignURL returns a presigned GET URL with a TTL (seconds). Mirrors
// app/Services/OssService.php signUrl(). Accepts either a bare key
// ("safety-architecture/foo.jpg") or a full bucket URL — historically
// some columns store the full public URL, and passing that through
// without normalization makes the SDK percent-encode the host into the
// path, producing a doubled URL like
// `https://bucket/https%3A%2F%2Fbucket%2F…` that 404s.
func (c *Client) SignURL(key string, ttl time.Duration, extraParams url.Values) (string, error) {
	options := []oss.Option{}
	for k, vs := range extraParams {
		for _, v := range vs {
			options = append(options, oss.AddParam(k, v))
		}
	}
	return c.bucket.SignURL(NormalizeKey(key, c.url), oss.HTTPGet, int64(ttl.Seconds()), options...)
}

// NormalizeKey turns a stored annex value into the bucket-relative key
// the SDK can sign. Accepts:
//   - bare key:   "safety-architecture/foo.jpg" → unchanged
//   - public URL: "https://bucket.oss-cn-…/safety-architecture/foo.jpg"
//     → "safety-architecture/foo.jpg" (matched against baseURL)
//   - any other absolute URL (custom CDN, expired signed URL, …) →
//     parse with net/url and take the path; otherwise the SDK appends
//     the input verbatim to the bucket host and produces a 404.
func NormalizeKey(raw, baseURL string) string {
	base := strings.TrimRight(baseURL, "/")
	if base != "" && strings.HasPrefix(raw, base+"/") {
		return strings.TrimPrefix(raw, base+"/")
	}
	if u, err := url.Parse(raw); err == nil && u.IsAbs() {
		return strings.TrimPrefix(u.Path, "/")
	}
	return raw
}

func (c *Client) InitMultipart(key string) (string, error) {
	imur, err := c.bucket.InitiateMultipartUpload(key)
	if err != nil {
		return "", err
	}
	return imur.UploadID, nil
}

func (c *Client) SignPartURL(key, uploadID string, partNumber int, ttl time.Duration) (string, error) {
	return c.bucket.SignURL(key, oss.HTTPPut, int64(ttl.Seconds()),
		oss.AddParam("uploadId", uploadID),
		oss.AddParam("partNumber", fmt.Sprintf("%d", partNumber)),
	)
}

type CompletePart struct {
	PartNumber int
	ETag       string
}

func (c *Client) CompleteMultipart(key, uploadID string, parts []CompletePart) error {
	imur := oss.InitiateMultipartUploadResult{Bucket: c.bucket.BucketName, Key: key, UploadID: uploadID}
	uploadParts := make([]oss.UploadPart, len(parts))
	for i, p := range parts {
		uploadParts[i] = oss.UploadPart{PartNumber: p.PartNumber, ETag: p.ETag}
	}
	_, err := c.bucket.CompleteMultipartUpload(imur, uploadParts)
	return err
}

func (c *Client) AbortMultipart(key, uploadID string) error {
	imur := oss.InitiateMultipartUploadResult{Bucket: c.bucket.BucketName, Key: key, UploadID: uploadID}
	return c.bucket.AbortMultipartUpload(imur)
}

func (c *Client) BaseURL() string {
	return c.url
}
