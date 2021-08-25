package router

import (
	"sky/app/system/router"
	"sky/common/middleware/auth"

	"github.com/gin-gonic/gin"
)

func registerSystemRouter(g *gin.RouterGroup) {
	//group := g.Group("/system", auth.JWTAuthMiddleware(), permission.CheckPermMiddleware())
	group := g.Group("/system", auth.JWTAuthMiddleware())

	router.UserRouter(group) // 用户管理
	router.MenuRouter(group) // 菜单管理
}
