package main

import (
	"Study_Go/middleware"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.New()

	// 차단된 IP 미들웨어 추가
	router.Use(middleware.BlockedIPMiddleware())

	// 테스트 엔드포인트
	router.GET("/check", func(c *gin.Context) {
		log.Printf("|%s", c.ClientIP())
		c.JSON(http.StatusOK, gin.H{
			"message": "You are allowed to access this endpoint!",
		})
	})

	router.Run(":8080") // 서버 실행
}
