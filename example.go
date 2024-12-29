package main

import (
	"Study_Go/middleware"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.New()

	// BlockedIPMiddleware
	router.Use(middleware.BlockedIPMiddleware())
	//TokenAuthMiddleware
	router.Use(middleware.TokenAuthMiddleware())

	// BlockedIPMiddleware
	router.GET("/BlockedIPMiddleware", func(c *gin.Context) {
		log.Printf("|%s", c.ClientIP())
		c.JSON(http.StatusOK, gin.H{
			"message": "You are allowed to access this endpoint!",
		})
	})

	//TokenAuthMiddleware
	router.GET("/TokenAuthMiddleware", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "You have accessed secure data!",
		})
	})

	router.Run(":8080") // 서버 실행
}
