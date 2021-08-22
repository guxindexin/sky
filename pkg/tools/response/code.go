package response

var (
	Success = &Response{Code: 20000, Message: "success"}

	AuthorizationNullError   = &Response{Code: 30001, Message: "请求头中 Authorization 为空"}
	AuthorizationFormatError = &Response{Code: 30002, Message: "请求头中 Authorization 格式有误"}
	InvalidTokenError        = &Response{Code: 30003, Message: "Token 无效"}

	InvalidParameterError  = &Response{Code: 40001, Message: "无效参数"}
	QueryUserError         = &Response{Code: 40002, Message: "查询用户失败"}
	UserNotExistError      = &Response{Code: 40003, Message: "用户名不存在"}
	IncorrectPasswordError = &Response{Code: 40004, Message: "密码不正确"}
	GenerateTokenError     = &Response{Code: 40005, Message: "生成 Token 失败"}
)
