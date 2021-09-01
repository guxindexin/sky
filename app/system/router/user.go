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
		router.GET("/info", apis.UserInfo)
		router.POST("", apis.CreateUser)
		router.PUT("/:id", apis.UpdateUser)
		router.DELETE("/:id", apis.DeleteUser)
	}
}
