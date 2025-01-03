package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	router := gin.New()

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("| httpMethod : %v | 경로 : %v | 핸들러이름 : %v | 핸들러 갯수 : %v |", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	//핸들러 이름이 main.main.func2 이지랄 나는 이유
	//함수가 속한 패키지(main).
	//	정의된 함수(main()).
	//	익명 함수의 자동 이름(func2).

	//로그를 파일로 저장 How to write log file
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	/* 커스텀 미들웨어
	// BlockedIPMiddleware
	router.Use(middleware.BlockedIPMiddleware())

	//TokenAuthMiddleware
	router.Use(middleware.TokenAuthMiddleware())

	//ValidateJSONBody
	router.Use(middleware.ValidateJSONBody())

	//RateLimiterMiddleware
	router.Use(middleware.RateLimiterMiddleware())

	middleware := router.Group("/middleware")
	{
		// BlockedIPMiddleware
		middleware.GET("/BlockedIPMiddleware", func(c *gin.Context) {
			log.Printf("|%s", c.ClientIP())
			c.JSON(http.StatusOK, gin.H{
				"message": "You are allowed to access this endpoint!",
			})
		})

		//TokenAuthMiddleware
		middleware.GET("/TokenAuthMiddleware", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "You have accessed secure data!",
			})
		})

		//ValidateJSONBody
		middleware.POST("/ValidateJSONBody", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "No middleware, body validation skipped!",
			})
		})

		//RateLimiterMiddleware
		middleware.GET("/RateLimiterMiddleware", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Good Request",
			})
		})
	}
	*/

	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/index.tmpl" , ~~)

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main WebSite",
		})
	})

	router.Run(":8080") // 서버 실행
}
