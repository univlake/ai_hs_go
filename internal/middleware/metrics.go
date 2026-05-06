package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// HTTP request metrics. Cardinality is bounded by labelling on the route
// template (`c.FullPath()`) rather than the raw URL — so /companies/:id never
// explodes into a metric per id. Methods are uppercased ASCII; status is the
// numeric HTTP status code as a string.
//
// We deliberately omit the request-size / response-size histograms here: Gin
// doesn't expose them cheaply, and most ops dashboards only need RPS + p99 +
// status mix to start.
var (
	httpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Number of HTTP requests, partitioned by method, route, and status.",
		},
		[]string{"method", "route", "status"},
	)

	httpRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request latency in seconds.",
			Buckets: prometheus.DefBuckets, // .005…10s
		},
		[]string{"method", "route"},
	)
)

// Metrics increments per-request counters and observes latency. Place this
// before any handler that writes a response — the Gin chain runs c.Next() and
// returns here, where c.Writer.Status() is final.
func Metrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		// FullPath is empty for 404s; bucket those under "unmatched" so a flood
		// of probes doesn't create one timeseries per attempted path.
		route := c.FullPath()
		if route == "" {
			route = "unmatched"
		}
		status := strconv.Itoa(c.Writer.Status())
		httpRequestsTotal.WithLabelValues(c.Request.Method, route, status).Inc()
		httpRequestDuration.WithLabelValues(c.Request.Method, route).
			Observe(time.Since(start).Seconds())
	}
}
