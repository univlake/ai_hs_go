package safetyledger

import (
	"fmt"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

func parseDate(p *string) *time.Time {
	if p == nil || *p == "" {
		return nil
	}
	formats := []string{"2006-01-02", "2006/01/02", "2006/1/2", time.RFC3339}
	for _, f := range formats {
		if t, err := time.Parse(f, *p); err == nil {
			return &t
		}
	}
	return nil
}

// setXlsxAttachment writes Content-Type + RFC 5987-encoded Content-Disposition
// so curl / direct browser hits get a Chinese filename. The frontend's
// `<a download="...">` overrides the filename anyway, but the headers must
// stay consistent for non-browser clients and dev-tool downloads.
func setXlsxAttachment(c *gin.Context, chineseName string, asciiFallback string) {
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header(
		"Content-Disposition",
		fmt.Sprintf(
			"attachment; filename=\"%s\"; filename*=UTF-8''%s",
			asciiFallback,
			url.PathEscape(chineseName),
		),
	)
}
