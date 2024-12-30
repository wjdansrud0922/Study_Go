package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func RateLimiterMiddleware() gin.HandlerFunc {
	var requestCounts = make(map[string]int)
	var requestTime = make(map[string]time.Time)

	maxRequest := 5

	return func(c *gin.Context) {
		ClientIp := c.ClientIP()
		v, ok := requestCounts[ClientIp]
		if ok == false { //맵 안에 ClientIp 가 존재하지 않으면
			requestCounts[ClientIp] = 1
			requestTime[ClientIp] = time.Now()
		} else {
			if v < maxRequest { //ClientIp 에서 들어온 요청이 5개 이하일때만 카운트 +
				requestCounts[ClientIp] += 1
			} else {
				if time.Since(requestTime[ClientIp]) > time.Minute { //이미 저장해둔 time 이랑 지금 time 이랑 비교해서 첫 요청이 1분 지났으면 맵카운트를 1로 초기화시킴 안지났으면 오륲 반환
					requestCounts[ClientIp] = 1
				} else {
					c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"msg": "Too many requests"})

				}
			}
		}
		fmt.Println(v, "|", ok, "|", time.Since(requestTime[ClientIp]), "|", time.Minute)
		c.Next()
	}
}
