package apis

import (
	"fmt"
	"sky/app/system/models"
	"sky/pkg/conn"
	"sky/pkg/pagination"
	"sky/pkg/tools/response"

	"github.com/gin-gonic/gin"
)

// ApiList 接口列表
func ApiList(c *gin.Context) {
	var (
		err     error
		apiList []*models.Api
		result  interface{}
	)

	db := conn.Orm.Debug().Model(&models.Api{})

	title := c.DefaultQuery("title", "")
	if title != "" {
		db = db.Where("title like ?", fmt.Sprintf("%%%s%%", title))
	}

	group := c.DefaultQuery("group", "")
	if group != "" {
		db = db.Where("\"group\" = ?", group)
	}

	result, err = pagination.Paging(&pagination.Param{
		C:  c,
		DB: db,
	}, &apiList)
	if err != nil {
		response.Error(c, err, response.ApiListError)
		return
	}
	response.OK(c, result, "")
}

// SaveApi 保存接口
func SaveApi(c *gin.Context) {
	var (
		err error
		api models.Api
	)

	err = c.ShouldBind(&api)
	if err != nil {
		response.Error(c, err, response.InvalidParameterError)
		return
	}

	db := conn.Orm.Model(&models.Api{})

	if api.Id != 0 {
		db = db.Where("id = ?", api.Id)
	}

	err = db.Save(&api).Error
	if err != nil {
		response.Error(c, err, response.SaveApiError)
		return
	}

	response.OK(c, "", "")
}

// DeleteApi 删除接口
func DeleteApi(c *gin.Context) {
	var (
		err          error
		apiId        string
		meunApiCount int64
	)

	apiId = c.Param("id")

	// 查询是否有菜单绑定了接口
	err = conn.Orm.Model(&models.MenuApi{}).Where("api = ?", apiId).Count(&meunApiCount).Error
	if err != nil {
		response.Error(c, nil, response.GetApiMenuError)
		return
	}

	if meunApiCount > 0 {
		response.Error(c, nil, response.ApiUsedError)
		return
	}

	err = conn.Orm.Delete(&models.Api{}, apiId).Error
	if err != nil {
		response.Error(c, err, response.DeleteApiError)
		return
	}

	response.OK(c, "", "")
}
