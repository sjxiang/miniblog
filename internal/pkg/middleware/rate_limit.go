package middleware

import "github.com/gin-gonic/gin"


func RateLimit() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}


// Redis 限流搞一个

// middlewares.RateLimit(2 * time.Second, 1)