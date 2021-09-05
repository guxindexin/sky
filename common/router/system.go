package router

import (
	"sky/app/system/router"
	"sky/common/middleware/auth"
	"sky/common/middleware/permission"

	"github.com/gin-gonic/gin"
)

func registerSystemRouter(g *gin.RouterGroup) {
	group := g.Group("/system", auth.JWTAuthMiddleware(), permission.CheckPermMiddleware())

	router.UserRouter(group)     // 用户管理
	router.MenuRouter(group)     // 菜单管理
	router.RoleRouter(group)     // 角色管理
	router.ApiRouter(group)      // 接口管理
	router.ApiGroupRouter(group) // 接口管理
}
