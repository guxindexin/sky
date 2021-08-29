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

	err = conn.Orm.Model(&models.Menu{}).
		Where("type in (1, 2)").
		Order("sort, id").
		Find(&menus).Error
	if err != nil {
		response.Error(c, err, response.GetMenuError)
		return
	}

	m = models.MenuTree{Menus: menus}
	result = m.GetMenuTree()

	response.OK(c, result, "")
}

func CreateMenu(c *gin.Context) {
	response.OK(c, "", "")
}

func SaveMenu(c *gin.Context) {
	var (
		err  error
		menu models.Menu
	)

	err = c.ShouldBind(&menu)
	if err != nil {
		response.Error(c, err, response.InvalidParameterError)
		return
	}

	db := conn.Orm.Model(&models.Menu{})

	if menu.Id != 0 {
		db = db.Where("id = ?", menu.Id)
	}

	err = db.Save(&menu).Error
	if err != nil {
		response.Error(c, err, response.SaveMenuError)
		return
	}

	response.OK(c, menu, "")
}

func DeleteMenu(c *gin.Context) {
	var (
		err       error
		menuId    string
		menuCount int64
	)

	menuId = c.Param("id")

	// 确认是否有菜单节点，若有，则不允许删除
	err = conn.Orm.Model(&models.Menu{}).Where("parent = ?", menuId).Count(&menuCount).Error
	if err != nil {
		response.Error(c, err, response.GetMenuError)
		return
	}

	if menuCount > 0 {
		response.Error(c, err, response.SubmenuExistsError)
		return
	}

	err = conn.Orm.Delete(&models.Menu{}, menuId).Error
	if err != nil {
		response.Error(c, err, response.DeleteMenuError)
		return
	}

	response.OK(c, "", "")
}
