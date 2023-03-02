package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS, HEAD")
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")  
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Content-Type", "application/json")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return 
		}

		c.Next()
	}
}


/*

	Web 开发趋势：

		前后端分离，双方部署的 ip 地址不同，通过 nginx 转发请求
		
			vue 192.168.0.2
			gin 192.168.0.3


		但也有安全问题，CSRF 跨站请求伪造（用户在前端误操作，不可以瑟瑟）
		
		
		应对，浏览器同源策略

			浏览器会加塞个 options 请求，header 里有个 Origin 字段，询问下策略；
			那就回复它，有什么要求，让浏览器自己掂量。

*/