package router

import (
	"sky/app/system/apis"

	"github.com/gin-gonic/gin"
)

func RoleRouter(g *gin.RouterGroup) {
	router := g.Group("/role")
	{
		router.GET("", apis.RoleList)
		router.POST("", apis.SaveRole)
		router.DELETE("/:id", apis.DeleteRole)
		router.POST("/permission/:id", apis.UpdateRolePermission)
		router.GET("/permission/:id", apis.GetRolePermission)
	}
}
