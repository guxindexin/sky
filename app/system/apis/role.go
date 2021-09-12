package apis

import (
	"fmt"
	"sky/app/system/models"
	"sky/common/middleware/permission"
	"sky/pkg/conn"
	"sky/pkg/pagination"
	"sky/pkg/tools/response"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// RoleList 角色列表
func RoleList(c *gin.Context) {
	var (
		err      error
		roleList []*models.Role
		result   interface{}
	)

	SearchParams := map[string]map[string]interface{}{
		"like": pagination.RequestParams(c),
	}

	result, err = pagination.Paging(&pagination.Param{
		C:  c,
		DB: conn.Orm.Model(&models.Role{}),
	}, &roleList, SearchParams)
	if err != nil {
		response.Error(c, err, response.RoleListError)
		return
	}
	response.OK(c, result, "")
}

// SaveRole 保存角色
func SaveRole(c *gin.Context) {
	var (
		err       error
		role      models.Role
		roleCount int64
	)

	err = c.ShouldBind(&role)
	if err != nil {
		response.Error(c, err, response.InvalidParameterError)
		return
	}

	db := conn.Orm.Model(&models.Role{})

	if role.Id != 0 {
		db = db.Where("id = ?", role.Id)
	} else {
		err = conn.Orm.Model(&models.Role{}).
			Where("key = ? or name = ?", role.Key, role.Name).
			Count(&roleCount).Error
		if err != nil {
			response.Error(c, err, response.RoleExistError)
			return
		}
	}

	err = db.Save(&role).Error
	if err != nil {
		response.Error(c, err, response.SaveRoleError)
		return
	}

	response.OK(c, "", "")
}

// DeleteRole 删除角色
func DeleteRole(c *gin.Context) {
	var (
		err           error
		roleId        string
		role          models.Role
		roleMenuCount int64
	)

	roleId = c.Param("id")

	err = conn.Orm.Model(&models.Role{}).Where("id = ?", roleId).Find(&role).Error
	if err != nil {
		response.Error(c, err, response.GetRoleError)
		return
	}

	// 获取角色绑定的用户
	groups := permission.Enforcer.GetFilteredNamedGroupingPolicy("g", 1, role.Key)
	if len(groups) > 0 {
		response.Error(c, err, response.RoleUsedError)
		return
	}

	// 获取角色绑定的菜单数据
	err = conn.Orm.Model(&models.RoleMenu{}).Where("role = ?", roleId).Count(&roleMenuCount).Error
	if err != nil {
		response.Error(c, err, response.GetRoleMenuError)
		return
	}
	if roleMenuCount > 0 {
		response.Error(c, err, response.RoleBindMenuError)
		return
	}

	err = conn.Orm.Delete(&models.Role{}, roleId).Error
	if err != nil {
		response.Error(c, err, response.DeleteRoleError)
		return
	}

	response.OK(c, "", "")
}

// UpdateRolePermission 更新角色权限
func UpdateRolePermission(c *gin.Context) {
	var (
		err                error
		newRoleMenu        []*models.RoleMenu
		createRoleMenu     []models.RoleMenu
		roleMenu           []*models.RoleMenu
		roleId             string
		oldRoleMenuMap     map[string]int
		newRoleMenuMap     map[string]struct{}
		deleteRoleMenuList []int
	)

	roleId = c.Param("id")
	createRoleMenu = make([]models.RoleMenu, 0)

	err = c.ShouldBind(&newRoleMenu)
	if err != nil {
		response.Error(c, err, response.InvalidParameterError)
		return
	}

	// 查询角色对应的菜单列表
	err = conn.Orm.Model(&models.RoleMenu{}).Where("role = ?", roleId).Find(&roleMenu).Error
	if err != nil {
		response.Error(c, err, response.GetRoleMenuError)
		return
	}

	if len(roleMenu) > 0 {
		oldRoleMenuMap = make(map[string]int)
		newRoleMenuMap = make(map[string]struct{})

		for _, oldRoleMenu := range roleMenu {
			oldRoleMenuMap[fmt.Sprintf("%d-%d-%d", oldRoleMenu.Role, oldRoleMenu.Menu, oldRoleMenu.Type)] = oldRoleMenu.Id
		}

		for _, newRoleMenu := range newRoleMenu {
			roleMenuKey := fmt.Sprintf("%d-%d-%d", newRoleMenu.Role, newRoleMenu.Menu, newRoleMenu.Type)
			if _, ok := oldRoleMenuMap[roleMenuKey]; ok {
				delete(oldRoleMenuMap, roleMenuKey)
			} else {
				roleMenuSlice := strings.Split(roleMenuKey, "-")
				roleIdTmp, _ := strconv.Atoi(roleMenuSlice[0])
				menuIdTmp, _ := strconv.Atoi(roleMenuSlice[1])
				typeTmp, _ := strconv.Atoi(roleMenuSlice[2])
				createRoleMenu = append(createRoleMenu, models.RoleMenu{
					Role: roleIdTmp,
					Menu: menuIdTmp,
					Type: typeTmp,
				})
			}
			newRoleMenuMap[roleMenuKey] = struct{}{}
		}

		if len(createRoleMenu) > 0 {
			err = conn.Orm.Model(&models.RoleMenu{}).Create(&createRoleMenu).Error
			if err != nil {
				response.Error(c, err, response.CreateRoleMenuError)
				return
			}
		}

		// 获取需要删除的节点
		deleteRoleMenuList = make([]int, 0, len(oldRoleMenuMap))
		for _, v := range oldRoleMenuMap {
			deleteRoleMenuList = append(deleteRoleMenuList, v)
		}
		if len(deleteRoleMenuList) > 0 {
			err = conn.Orm.Delete(&models.RoleMenu{}, deleteRoleMenuList).Error
			if err != nil {
				response.Error(c, err, response.DeleteRoleMenuError)
				return
			}
		}
	} else {
		err = conn.Orm.Model(&models.RoleMenu{}).Create(&newRoleMenu).Error
		if err != nil {
			response.Error(c, err, response.CreateRoleMenuError)
			return
		}
	}

	response.OK(c, "", "")
}

// GetRolePermission 获取角色权限
func GetRolePermission(c *gin.Context) {
	var (
		err            error
		roleMenus      []int
		roleMenuButton []struct {
			models.RoleMenu
			Parent int `json:"parent"`
		}
		roleButtons   map[int][]int
		roleId        string
		parentMenuIds []int
	)

	roleId = c.Param("id")

	// 查询所有的父级别ID
	err = conn.Orm.Model(&models.Menu{}).
		Select("distinct parent").
		Where("parent != 0 and type = 2").
		Pluck("parent", &parentMenuIds).Error
	if err != nil {
		response.Error(c, err, response.GetMenuParentError)
		return
	}

	// 当前菜单权限
	err = conn.Orm.Model(&models.RoleMenu{}).Where("role = ? and type = 1 and menu not in (?)", roleId, parentMenuIds).Pluck("menu", &roleMenus).Error
	if err != nil {
		response.Error(c, err, response.GetRolePermissionError)
		return
	}

	// 查询按钮权限
	err = conn.Orm.Model(&models.RoleMenu{}).
		Joins("left join system_menu on system_menu.id = system_role_menu.menu").
		Select("system_role_menu.*, system_menu.parent").
		Where("system_role_menu.role = ? and system_role_menu.type = 2", roleId).Find(&roleMenuButton).Error
	if err != nil {
		response.Error(c, err, response.GetRoleButtonError)
		return
	}
	roleButtons = make(map[int][]int)
	for _, v := range roleMenuButton {
		if _, ok := roleButtons[v.Parent]; ok {
			roleButtons[v.Parent] = append(roleButtons[v.Parent], v.Menu)
		} else {
			roleButtons[v.Parent] = []int{v.Menu}
		}
	}

	response.OK(c, map[string]interface{}{
		"menu":   roleMenus,
		"button": roleButtons,
	}, "")
}

// RoleApiPermission 角色的接口权限
func RoleApiPermission(c *gin.Context) {
	var (
		err    error
		role   models.Role
		roleId string
		apis   []models.Api
		groups [][]string
		params struct {
			Type int   `json:"type"` // 1 绑定api权限，0 解除api权限
			Api  []int `json:"api"`
		}
	)

	roleId = c.Param("id")

	err = c.ShouldBind(&params)
	if err != nil {
		response.Error(c, err, response.InvalidParameterError)
		return
	}

	// 查询角色信息
	err = conn.Orm.Model(&models.Role{}).Where("id = ?", roleId).Find(&role).Error
	if err != nil {
		response.Error(c, err, response.GetRoleError)
		return
	}

	// 查询 API 接口
	err = conn.Orm.Model(&models.Api{}).Where("id in (?)", params.Api).Find(&apis).Error
	if err != nil {
		response.Error(c, err, response.GetApiError)
		return
	}

	for _, api := range apis {
		groups = append(groups, []string{role.Key, api.URL, api.Method})
	}

	if params.Type == 1 {
		_, err = permission.Enforcer.AddNamedPolicies("p", groups)
		if err != nil {
			response.Error(c, err, response.RoleBindApiError)
			return
		}
	} else if params.Type == 0 {
		_, err = permission.Enforcer.RemoveNamedPolicies("p", groups)
		if err != nil {
			response.Error(c, err, response.RoleUnBindApiError)
			return
		}
	}

	response.OK(c, "", "")
}

// GetRoleApi 获取角色绑定的API接口
func GetRoleApi(c *gin.Context) {
	var (
		err      error
		role     models.Role
		roleId   string
		menuId   string
		apis     []int
		roleApis [][]string
	)

	roleId = c.Param("id")

	menuId = c.DefaultQuery("menu", "")
	if menuId == "" {
		response.Error(c, nil, response.InvalidParameterError)
		return
	}

	err = conn.Orm.Model(&models.Role{}).Where("id = ?", roleId).Find(&role).Error
	if err != nil {
		response.Error(c, err, response.GetRoleError)
		return
	}

	roleApis = permission.Enforcer.GetFilteredNamedPolicy("p", 0, role.Key)

	if len(roleApis) > 0 {

		db := conn.Orm.Debug().Model(&models.Api{}).
			Select("system_api.id").
			Joins("left join system_menu_api on system_menu_api.api = system_api.id")

		for _, p := range roleApis {
			db = db.Or("system_api.url = ? and system_api.method = ?", p[1], p[2])
		}
		err = db.Where("system_menu_api.menu = ?", menuId).Pluck("system_api.id", &apis).Error
		if err != nil {
			response.Error(c, err, response.GetApiError)
			return
		}
	}

	response.OK(c, apis, "")
}
