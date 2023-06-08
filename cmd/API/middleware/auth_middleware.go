package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rafikmoreira/go-blog-api/cmd/API/util"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if util.VerifyJWT(c.GetHeader("Authorization")) {
			c.Next()
			return
		}

		c.AbortWithStatusJSON(401, gin.H{
			"message": "Unauthorized",
		})
	}
}
