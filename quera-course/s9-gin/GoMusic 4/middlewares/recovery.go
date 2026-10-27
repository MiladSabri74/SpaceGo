package middlewares

import (
	"github.com/gin-gonic/gin"
)

func CustomRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Recover from panics and return 500 JSON: {"error": "internal server error", "message": "<panic>"}
		c.Next()
	}
}
