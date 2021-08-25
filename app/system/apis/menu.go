package apis

import (
	"sky/app/system/models"
	"sky/pkg/conn"
	"sky/pkg/tools/response"

	"github.com/gin-gonic/gin"
)

func MenuTree(c *gin.Context) {
	var (
		err   error
		menus []models.Menu
	)

	err = conn.Orm.Model(&models.Menu{}).Where("type = 1 and parent = 0").Find(&menus).Error
	if err != nil {
		response.Error(c, err, response.GetMenuError)
		return
	}

	response.OK(c, menus, "")
}
