package router

import (
	"sky/app/system/router"

	"github.com/gin-gonic/gin"
)

func registerSystemRouter(g *gin.RouterGroup) {
	group := g.Group("/system")

	router.UserRouter(group) // 用户管理
}
