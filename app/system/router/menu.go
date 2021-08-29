package router

import (
	"sky/app/system/apis"

	"github.com/gin-gonic/gin"
)

/*
  @Author : lanyulei
*/

func MenuRouter(g *gin.RouterGroup) {
	router := g.Group("/menu")
	{
		router.GET("/tree", apis.MenuTree)
		router.POST("", apis.SaveMenu)
		router.DELETE("/:id", apis.DeleteMenu)
	}
}
