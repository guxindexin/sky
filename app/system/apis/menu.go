package apis

import (
	"sky/app/system/models"
	"sky/pkg/conn"
	"sky/pkg/tools/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// MenuTree 菜单树
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

// SaveMenu 保存菜单
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

// DeleteMenu 删除菜单
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

// BatchDeleteMenu 批量删除菜单
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

// MenuBindApi 菜单绑定API
func MenuBindApi(c *gin.Context) {
	var (
		err    error
		params struct {
			Type int   `json:"type"` // 1 绑定， 0 解绑
			Apis []int `json:"apis"`
		}
		existApis []int
		apiMaps   map[int]struct{}
		menuApi   []models.MenuApi
		menuId    string
	)

	menuId = c.Param("id")
	menuIdInt, _ := strconv.Atoi(menuId)

	err = c.ShouldBind(&params)
	if err != nil {
		response.Error(c, err, response.InvalidParameterError)
		return
	}

	if params.Type == 1 {
		err = conn.Orm.Model(&models.MenuApi{}).Where("menu = ? and api in (?)", menuIdInt, params.Apis).Pluck("api", &existApis).Error
		if err != nil {
			response.Error(c, err, response.GetMenuApiError)
			return
		}
		apiMaps = make(map[int]struct{})
		if len(existApis) != 0 {
			for _, i := range existApis {
				apiMaps[i] = struct{}{}
			}
		}

		menuApi = make([]models.MenuApi, 0)
		for _, i := range params.Apis {
			if _, ok := apiMaps[i]; !ok {
				menuApi = append(menuApi, models.MenuApi{
					Menu: menuIdInt,
					Api:  i,
				})
			}
		}

		if len(menuApi) > 0 {
			err = conn.Orm.Create(&menuApi).Error
			if err != nil {
				response.Error(c, err, response.MenuBindApiError)
				return
			}
		}
	} else if params.Type == 0 {
		err = conn.Orm.Where("menu = ? and api in (?)", menuIdInt, params.Apis).Delete(&models.MenuApi{}).Error
		if err != nil {
			response.Error(c, err, response.MenuUnBindApiError)
			return
		}
	}

	response.OK(c, "", "")
}

// MenuApis 查询菜单绑定的API
func MenuApis(c *gin.Context) {
	var (
		err     error
		menuId  string
		apiList []int
	)

	menuId = c.Param("id")

	err = conn.Orm.Model(&models.MenuApi{}).Select("distinct api").Where("menu = ?", menuId).Pluck("api", &apiList).Error
	if err != nil {
		response.Error(c, err, response.GetMenuApiError)
		return
	}

	response.OK(c, apiList, "")
}

// MenuApiList 查询菜单绑定的API列表
func MenuApiList(c *gin.Context) {
	var (
		err     error
		menuId  string
		apiIds  []int
		apiList []models.Api
	)

	menuId = c.Param("id")

	err = conn.Orm.Model(&models.MenuApi{}).
		Select("distinct api").
		Where("menu = ?", menuId).
		Pluck("api", &apiIds).Error
	if err != nil {
		response.Error(c, err, response.GetMenuApiError)
		return
	}

	err = conn.Orm.Model(&models.Api{}).Where("id in (?)", apiIds).Find(&apiList).Error
	if err != nil {
		response.Error(c, err, response.GetMenuApiError)
		return
	}

	response.OK(c, apiList, "")
}
