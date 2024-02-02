package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// TODO
// Complete Authentication once the auth backend is ready.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check authentication logic here
		if true {
			fmt.Println("Auth Bypassed")
			c.Next()
		} else {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
		}
	}
}
