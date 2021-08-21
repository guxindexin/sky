package router

import (
	"fmt"
	"sky/app/system/apis"

	"github.com/gin-gonic/gin"
)

/*
  @Author : lanyulei
*/

func UserRouter(g *gin.RouterGroup) {
	group := fmt.Sprintf("%s", "/user")
	router := g.Group(group)
	{
		router.GET("", apis.UserList)
	}
}
