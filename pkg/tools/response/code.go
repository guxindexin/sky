package response

var (
	Success = Response{Code: 20000, Message: "success"}

	AuthorizationNullError   = Response{Code: 30001, Message: "请求头中 Authorization 为空"}
	AuthorizationFormatError = Response{Code: 30002, Message: "请求头中 Authorization 格式有误"}
	InvalidTokenError        = Response{Code: 30003, Message: "Token 无效"}
	NoPermissionError        = Response{Code: 30004, Message: "暂无请求权限"}

	InvalidParameterError  = Response{Code: 40001, Message: "无效参数"}
	QueryUserError         = Response{Code: 40002, Message: "查询用户失败"}
	UserNotExistError      = Response{Code: 40003, Message: "用户名不存在"}
	UserExistError         = Response{Code: 40004, Message: "用户名已存在"}
	IncorrectPasswordError = Response{Code: 40005, Message: "密码不正确"}
	GenerateTokenError     = Response{Code: 40006, Message: "生成 Token 失败"}
	CreateUserError        = Response{Code: 40007, Message: "创建用户失败"}
	EncryptPasswordError   = Response{Code: 40008, Message: "加密密码失败"}
	GetUserInfoError       = Response{Code: 40009, Message: "获取用户详情失败"}

	GetMenuError       = Response{Code: 40010, Message: "查询菜单数据失败"}
	SaveMenuError      = Response{Code: 40011, Message: "保存菜单数据失败"}
	SubmenuExistsError = Response{Code: 40012, Message: "当前菜单存在子节点，无法直接删除"}
	DeleteMenuError    = Response{Code: 40013, Message: "删除菜单失败"}
	GetMenuButtonError = Response{Code: 40014, Message: "查询菜单按钮数据失败"}

	UserListError   = Response{Code: 40015, Message: "查询用户列表失败"}
	DeleteUserError = Response{Code: 40016, Message: "删除用户失败"}
	UpdateUserError = Response{Code: 40017, Message: "更新用户失败"}
)
