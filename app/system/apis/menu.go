package apis

import (
	"sky/app/system/models"
	"sky/pkg/conn"
	"sky/pkg/tools/response"

	"github.com/gin-gonic/gin"
)

func MenuTree(c *gin.Context) {
	var (
		err        error
		menus      []*models.Menu
		result     []*models.MenuValue
		m          models.MenuTree
		buttonList []*models.Menu
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

	// 查询所有页面对应的按钮列表
	err = conn.Orm.Model(&models.Menu{}).
		Where("type = 3").
		Order("sort, id").
		Find(&buttonList).Error
	if err != nil {
		response.Error(c, err, response.GetMenuButtonError)
		return
	}

	buttonMap := make(map[int][]*models.Menu)
	for _, b := range buttonList {
		if _, ok := buttonMap[b.Parent]; ok {
			buttonMap[b.Parent] = append(buttonMap[b.Parent], b)
		} else {
			buttonMap[b.Parent] = []*models.Menu{b}
		}
	}

	response.OK(c, map[string]interface{}{
		"menu":   result,
		"button": buttonMap,
	}, "")
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

func BatchDeleteMenu(c *gin.Context) {
	var (
		err     error
		menuIds struct {
			MenuIds []int `json:"menu_ids"`
		}
		menuCount int64
	)

	err = c.ShouldBind(&menuIds)
	if err != nil {
		response.Error(c, err, response.InvalidParameterError)
		return
	}

	// 确认是否存在子节点
	err = conn.Orm.Model(&models.Menu{}).Where("parent in (?)", menuIds.MenuIds).Count(&menuCount).Error
	if err != nil {
		response.Error(c, err, response.GetMenuError)
		return
	}
	if menuCount > 0 {
		response.Error(c, err, response.SubmenuExistsError)
		return
	}

	err = conn.Orm.Delete(&models.Menu{}, menuIds.MenuIds).Error
	if err != nil {
		response.Error(c, err, response.DeleteMenuError)
		return
	}

	response.OK(c, "", "")
}

// MenuButton 查询菜单下的所有按钮
func MenuButton(c *gin.Context) {
	var (
		err        error
		menuId     string
		buttonList []*models.Menu
	)

	menuId = c.Param("id")

	err = conn.Orm.Model(&models.Menu{}).
		Where("parent = ?", menuId).
		Order("sort, id").
		Find(&buttonList).Error
	if err != nil {
		response.Error(c, err, response.GetMenuButtonError)
		return
	}

	response.OK(c, buttonList, "")
}
