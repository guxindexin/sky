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
		router.DELETE("/batch", apis.BatchDeleteMenu)
		router.GET("/button/:id", apis.MenuButton)
		router.POST("/bind/api", apis.MenuBindApi)
		router.DELETE("/unbind/api", apis.MenuUnBindApi)
		router.GET("/api/:id", apis.MenuApis)
		router.GET("/api-list/:id", apis.MenuApiList)
	}
}
