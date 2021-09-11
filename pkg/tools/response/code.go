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
	GetMenuError           = Response{Code: 40010, Message: "查询菜单数据失败"}
	SaveMenuError          = Response{Code: 40011, Message: "保存菜单数据失败"}
	SubmenuExistsError     = Response{Code: 40012, Message: "当前菜单存在子节点，无法直接删除"}
	DeleteMenuError        = Response{Code: 40013, Message: "删除菜单失败"}
	GetMenuButtonError     = Response{Code: 40014, Message: "查询菜单按钮数据失败"}
	UserListError          = Response{Code: 40015, Message: "查询用户列表失败"}
	DeleteUserError        = Response{Code: 40016, Message: "删除用户失败"}
	UpdateUserError        = Response{Code: 40017, Message: "更新用户失败"}
	UserUnavailableError   = Response{Code: 40018, Message: "用户不可用"}
	RoleListError          = Response{Code: 40019, Message: "查询角色列表失败"}
	SaveRoleError          = Response{Code: 40020, Message: "保存角色失败"}
	RoleUsedError          = Response{Code: 40022, Message: "角色被其他用户关联，无法删除"}
	DeleteRoleError        = Response{Code: 40023, Message: "角色删除失败"}
	ApiListError           = Response{Code: 40024, Message: "查询API接口列表失败"}
	SaveApiError           = Response{Code: 40025, Message: "保存API接口失败"}
	DeleteApiError         = Response{Code: 40026, Message: "API接口删除失败"}
	GetApiMenuError        = Response{Code: 40027, Message: "获取API接口对应的菜单或按钮失败"}
	ApiUsedError           = Response{Code: 40028, Message: "API接口被其他角色关联，无法删除"}
	SaveApiGroupError      = Response{Code: 40029, Message: "保存API分组失败"}
	ApiGroupListError      = Response{Code: 40030, Message: "获取API分组列表失败"}
	ApiGroupExistError     = Response{Code: 40031, Message: "API分组已存在"}
	RoleExistError         = Response{Code: 40032, Message: "角色已存在"}
	GetApiError            = Response{Code: 40033, Message: "获取API失败"}
	ApiGroupUsedError      = Response{Code: 40034, Message: "API分组被使用，无法删除"}
	DeleteApiGroupError    = Response{Code: 40035, Message: "删除API分组失败"}
	GetRoleMenuError       = Response{Code: 40036, Message: "查询角色对应的菜单列表失败"}
	CreateRoleMenuError    = Response{Code: 40037, Message: "角色关联菜单权限失败"}
	DeleteRoleMenuError    = Response{Code: 40037, Message: "删除角色关联菜单权限失败"}
	GetRolePermissionError = Response{Code: 40038, Message: "查询角色对应的菜单权限失败"}
	GetRoleButtonError     = Response{Code: 40038, Message: "查询角色对应的菜单按钮权限失败"}
	GetMenuParentError     = Response{Code: 40039, Message: "查询所有父级别菜单失败"}
	GetMenuApiError        = Response{Code: 40040, Message: "查询菜单绑定的API失败"}
	MenuBindApiError       = Response{Code: 40041, Message: "菜单绑定API失败"}
	MenuUnBindApiError     = Response{Code: 40042, Message: "菜单结束绑定API失败"}
	CreateUserRoleError    = Response{Code: 40043, Message: "创建用户与角色的关联到Casbin中失败"}
	DeleteUserRoleError    = Response{Code: 40044, Message: "删除用户与角色的关联到Casbin中失败"}
	GetRoleError           = Response{Code: 40045, Message: "获取角色信息失败"}
	UserRoleError          = Response{Code: 40046, Message: "当前用户关联了其他角色，无法直接删除"}
	RoleBindApiError       = Response{Code: 40046, Message: "角色绑定接口权限失败"}
	RoleUnBindApiError     = Response{Code: 40046, Message: "角色解除接口权限失败"}
)
