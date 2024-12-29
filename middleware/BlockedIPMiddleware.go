package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func BlockedIPMiddleware() gin.HandlerFunc {

	bannedIPs := []string{"::1", "222.222.222.222"}

	return func(c *gin.Context) {
		clientIP := c.ClientIP()

		for _, bannedIP := range bannedIPs {
			if clientIP == bannedIP {
				log.Printf("[WARN] Blocked IP attempted access: %s", clientIP)

				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
					"message": "Access denied. Your IP is blocked.",
				})
				return
			}
		}
		
		c.Next()
	}
}
