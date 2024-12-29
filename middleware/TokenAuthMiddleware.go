package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		validToken := "1234"
		token := c.GetHeader("Authorization")

		if token == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"msg": "Token is nil",
			})
			return
		}

		if token != validToken {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"msg": "Token is not valid",
			})
			return
		}

		c.Next()

	}

}
