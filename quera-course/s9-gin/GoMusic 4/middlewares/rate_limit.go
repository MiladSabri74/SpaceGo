package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var lastTimeAccess map[string]time.Time
var countAcess map[string]int

func RateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Implement rate limiting (max 5 req/sec per IP) with sync.Mutex, return 429 on limit
		cip := c.ClientIP()
		if t, exists := lastTimeAccess[cip]; exists {
			if time.Since(t) <= 1*time.Second {
				countAcess[cip]++
			} else {
				lastTimeAccess[cip] = time.Now()
				countAcess[cip] = 1
			}
		} else {
			lastTimeAccess[cip] = time.Now()
			countAcess[cip] = 1
		}

		if countAcess[cip] > 5 {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "too many requests"})

			c.Abort()
			return
		}
		c.Next()
	}
}
