package middleware

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	"ai_hs_go/internal/response"
)

// fixedWindowLimiter is an in-memory fixed-window per-key rate limiter. Suitable
// for single-process deployments and dev. For multi-replica production, swap to
// a Redis-backed implementation.
type fixedWindowLimiter struct {
	mu      sync.Mutex
	buckets map[string]*window
	limit   int
	window  time.Duration
}

type window struct {
	count int
	reset time.Time
}

func newFixedWindowLimiter(limit int, w time.Duration) *fixedWindowLimiter {
	return &fixedWindowLimiter{
		buckets: make(map[string]*window),
		limit:   limit,
		window:  w,
	}
}

func (l *fixedWindowLimiter) allow(key string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	now := time.Now()
	b, ok := l.buckets[key]
	if !ok || now.After(b.reset) {
		l.buckets[key] = &window{count: 1, reset: now.Add(l.window)}
		return true
	}
	if b.count >= l.limit {
		return false
	}
	b.count++
	return true
}

// rateLimitDisabled is set true in dev so a single browser tab behind the
// vite proxy (where every request looks like 127.0.0.1) doesn't get throttled.
// Toggle via DisableInDev / EnableLimiter from main.go.
var rateLimitDisabled = false

// SetRateLimitEnabled lets the bootstrap code disable limiting for `local`
// environments without requiring callers to thread a flag through every
// middleware constructor.
func SetRateLimitEnabled(enabled bool) {
	rateLimitDisabled = !enabled
}

// LoginRateLimit limits per IP for login/register routes.
func LoginRateLimit(perMinute int) gin.HandlerFunc {
	limiter := newFixedWindowLimiter(perMinute, time.Minute)
	return func(c *gin.Context) {
		if rateLimitDisabled {
			c.Next()
			return
		}
		if !limiter.allow(c.ClientIP()) {
			response.Failed(c, 429, "请求过于频繁，请稍后再试")
			return
		}
		c.Next()
	}
}

// APIRateLimit applies a higher quota for authenticated users (keyed by user_id
// in context) and a stricter quota for guests (keyed by IP).
func APIRateLimit(authPerMinute, guestPerMinute int) gin.HandlerFunc {
	authLimiter := newFixedWindowLimiter(authPerMinute, time.Minute)
	guestLimiter := newFixedWindowLimiter(guestPerMinute, time.Minute)
	return func(c *gin.Context) {
		if rateLimitDisabled {
			c.Next()
			return
		}
		if uid, ok := c.Get(CtxKeyUserID); ok {
			if !authLimiter.allow(formatUserKey(uid)) {
				response.Failed(c, 429, "请求过于频繁，请稍后再试")
				return
			}
		} else {
			if !guestLimiter.allow(c.ClientIP()) {
				response.Failed(c, 429, "请求过于频繁，请稍后再试")
				return
			}
		}
		c.Next()
	}
}

func formatUserKey(v any) string {
	switch x := v.(type) {
	case string:
		return "u:" + x
	case int64:
		return "u:" + intToStr(x)
	case uint:
		return "u:" + intToStr(int64(x))
	case uint64:
		return "u:" + intToStr(int64(x))
	default:
		return "u:?"
	}
}

func intToStr(n int64) string {
	if n == 0 {
		return "0"
	}
	neg := n < 0
	if neg {
		n = -n
	}
	var buf [20]byte
	i := len(buf)
	for n > 0 {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
	}
	if neg {
		i--
		buf[i] = '-'
	}
	return string(buf[i:])
}
