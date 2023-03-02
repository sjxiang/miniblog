package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/sjxiang/miniblog/internal/pkg/consts"
)

// RequestID 在每一个 HTTP 请求的 context、response 中注入 `X-Request-ID` 键值对
func RequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 检查请求头中是否有 `X-Request-ID`，如果有则复用，没有则新建
		requestID := ctx.Request.Header.Get(consts.XRequestIDKey)

		if requestID == "" {
			requestID = uuid.New().String()
		}

		// 将 RequestID 保存在 gin.Context 中，方便后面程序使用
		ctx.Set(consts.XRequestIDKey, requestID)

		// 将 RequestID 保存在 HTTP 返回头中，Header 的键为 `X-Request-ID` （方便联调）
		ctx.Writer.Header().Set(consts.XRequestIDKey, requestID)

		ctx.Next()
	}
}
