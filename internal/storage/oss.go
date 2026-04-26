package storage

import (
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"

	"ai_hs_go/internal/config"
)

// Client wraps the Aliyun OSS Bucket and provides the small subset of
// operations the API actually uses (signed URL, multipart upload).
type Client struct {
	bucket *oss.Bucket
	url    string
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
	return &Client{bucket: bucket, url: cfg.PublicURL}, nil
}

// SignURL returns a presigned GET URL with a TTL (seconds). Mirrors
// app/Services/OssService.php signUrl().
func (c *Client) SignURL(key string, ttl time.Duration, extraParams url.Values) (string, error) {
	options := []oss.Option{}
	for k, vs := range extraParams {
		for _, v := range vs {
			options = append(options, oss.AddParam(k, v))
		}
	}
	return c.bucket.SignURL(key, oss.HTTPGet, int64(ttl.Seconds()), options...)
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
