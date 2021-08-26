package apis

import (
	"sky/app/system/models"
	"sky/pkg/conn"
	"sky/pkg/tools/response"

	"github.com/gin-gonic/gin"
)

func MenuTree(c *gin.Context) {
	var (
		err    error
		menus  []*models.Menu
		result []*models.MenuValue
		m      models.MenuTree
	)

	err = conn.Orm.Model(&models.Menu{}).Where("type = 1").Find(&menus).Error
	if err != nil {
		response.Error(c, err, response.GetMenuError)
		return
	}

	m = models.MenuTree{Menus: menus}
	result = m.GetMenuTree()

	response.OK(c, result, "")
}
