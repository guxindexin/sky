package router

import (
	"sky/app/public/router"

	"github.com/gin-gonic/gin"
)

func registerPublicRouter(g *gin.RouterGroup) {
	router.LoginRouter(g)
}
