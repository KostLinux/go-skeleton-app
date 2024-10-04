package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/khaaleoo/gin-rate-limiter/core"
)

func rateLimiterMiddleware() gin.HandlerFunc {
	// Create an IP rate limiter middleware
	rateLimiterMiddleware := core.RequireRateLimiter(core.RateLimiter{
		RateLimiterType: core.IPRateLimiter,
		Key:             "this_is_a_test_key",
		Option: core.RateLimiterOption{
			Limit: 100,
			Burst: 5,
			Len:   10 * time.Second,
		},
	})

	return rateLimiterMiddleware
}
