package router

import (
	"sky/app/public/apis"

	"github.com/gin-gonic/gin"
)

/*
  @Author : lanyulei
*/

func LoginRouter(g *gin.RouterGroup) {
	g.POST("/login", apis.Login)
}
