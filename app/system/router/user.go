package router

import (
	"sky/app/system/apis"

	"github.com/gin-gonic/gin"
)

/*
  @Author : lanyulei
*/

func UserRouter(g *gin.RouterGroup) {
	router := g.Group("/user")
	{
		router.GET("", apis.UserList)
		router.POST("", apis.CreateUser)
	}
}
