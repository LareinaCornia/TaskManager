package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

var lastRequest = make(map[string]time.Time)

func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		now := time.Now()
		if t, ok := lastRequest[ip]; ok && now.Sub(t) < time.Second {
			c.AbortWithStatusJSON(429, gin.H{"error": "Too Many Requests"})
			return
		}
		lastRequest[ip] = now
		c.Next()
	}
}
