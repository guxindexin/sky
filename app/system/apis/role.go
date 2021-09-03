package apis

import (
	"sky/app/system/models"
	"sky/pkg/conn"
	"sky/pkg/pagination"
	"sky/pkg/tools/response"

	"github.com/gin-gonic/gin"
)

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

func SaveRole(c *gin.Context) {
	var (
		err  error
		role models.Role
	)

	err = c.ShouldBind(&role)
	if err != nil {
		response.Error(c, err, response.InvalidParameterError)
		return
	}

	db := conn.Orm.Model(&models.Role{})

	if role.Id != 0 {
		db = db.Where("id = ?", role.Id)
	}

	err = db.Save(&role).Error
	if err != nil {
		response.Error(c, err, response.SaveRoleError)
		return
	}

	response.OK(c, "", "")
}

func DeleteRole(c *gin.Context) {
	var (
		err       error
		roleId    string
		userCount int64
	)

	roleId = c.Param("id")

	// 查询是否有用户关联了角色
	err = conn.Orm.Model(&models.User{}).Where("(? = any(role))", roleId).Count(&userCount).Error
	if err != nil {
		response.Error(c, err, response.GetRoleUserError)
		return
	}

	if userCount > 0 {
		response.Error(c, nil, response.RoleUsedError)
		return
	}

	err = conn.Orm.Delete(&models.Role{}, roleId).Error
	if err != nil {
		response.Error(c, err, response.DeleteRoleError)
		return
	}

	response.OK(c, "", "")
}
