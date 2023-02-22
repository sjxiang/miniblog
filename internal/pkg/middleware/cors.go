package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS, HEAD")
		ctx.Header("Content-Type", "application/json")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, "+
			"Access-Control-Allow-Headers, Authorization, Cache-Control, Content-Language, Content-Type")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
			return 
		}

		ctx.Next()
	}
}


/*

	另一套实现，参考 github.com/gin-contrib/cors

	预检是与浏览器交互探讨，客户端能不能安全发送该请求

*/