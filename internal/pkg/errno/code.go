package errno

var (
	// Ok 代表请求成功
	Ok = NewErrNo(200, "", "Success")

	// ErrPageNotFound 表示路由不匹配错误
	ErrPageNotFound = NewErrNo(404, "ResourceNotFound.PageNotFound", "Page not found.")

	// InternalServerError 表示所有未知的服务器端错误.
	InternalServerError = NewErrNo(500, "InternalError", "Internal server error.")

	// ErrBind 表示参数绑定错误.
	ErrBind = NewErrNo(400, "InvalidParam.BindError", "Error occurred while binding the request body to the struct.")

	// ErrInvalidParam 表示所有验证失败的错误.
	ErrInvalidParam = NewErrNo(400, "InvalidParam.ParamError", "Param verification failed.")

	// ErrSignToken 表示签发 JWT Token 时出错.
	ErrSignToken = NewErrNo(401, "AuthFailure.SignTokenError", "Error occurred while signing the JSON web token.")

	// ErrTokenInvalid 表示 JWT Token 格式错误.
	ErrTokenInvalid = NewErrNo(401, "AuthFailure.TokenInvalid", "Token was invalid.")

	// ErrUnauthorized 表示请求没有被授权.
	ErrUnauthorized = NewErrNo(401, "AuthFailure.Unauthorized", "Unauthorized.")
)

func NewErrNo(status int, bizCode string, msg string) *ErrNo {
	return &ErrNo{
		Status:  status,
		BizCode: bizCode,
		Message: msg,
	}
}
