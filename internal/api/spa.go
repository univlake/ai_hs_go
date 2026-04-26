package api

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// mountSPA serves the Vite-built SPA from web/dist/. In dev the Vite server
// runs on :5173 and proxies /api/v1 to this binary, so this fallback only
// activates when the production bundle is present.
func (r *Router) mountSPA(engine *gin.Engine) {
	dist := r.deps.Config.App.Env
	_ = dist
	const distDir = "./web/dist"
	if _, err := os.Stat(distDir); err != nil {
		return // no built SPA — skip mounting silently
	}

	engine.Static("/assets", filepath.Join(distDir, "assets"))
	engine.StaticFile("/favicon.ico", filepath.Join(distDir, "favicon.ico"))

	// Anything not matched by /api/* or /healthz falls through to index.html so
	// React Router can take over (true SPA history routing).
	engine.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		if strings.HasPrefix(path, "/api/") || path == "/healthz" {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.File(filepath.Join(distDir, "index.html"))
	})
}
