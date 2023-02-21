package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sjxiang/miniblog/internal/pkg/errno"
)

// 序列化器 serializer

// 定义了发生错误时的返回信息
type ErrResponse struct {
	// 业务错误码
	BizCode string `json:"biz_code"`
	// 直接对外展示的错误信息
	Message string `json:"message"`
}

func WithResponse(ctx *gin.Context, err error, data interface{}) {
	if err != nil {
		status, bizCode, message := errno.Decode(err)
		ctx.JSON(status, ErrResponse{
			BizCode: bizCode,
			Message: message,
		})

		return
	}

	ctx.JSON(http.StatusOK, data)
}
