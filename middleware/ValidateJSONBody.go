package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func ValidateJSONBody() gin.HandlerFunc {
	var user User

	return func(c *gin.Context) {
		err := c.ShouldBindJSON(&user)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "fuck you"})
		}

		c.Next()
	}
}
