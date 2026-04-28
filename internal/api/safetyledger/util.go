package safetyledger

import (
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"
)

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
