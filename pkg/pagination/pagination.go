package pagination

/*
 @Author : lanyulei
*/

import (
	"fmt"
	"math"
	"sky/pkg/logger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Param struct {
	C       *gin.Context
	DB      *gorm.DB
	ShowSQL bool
}

type Paginator struct {
	Total     int64       `json:"total"`
	TotalPage int         `json:"total_page"`
	List      interface{} `json:"list"`
	Size      int         `json:"size"`
	Page      int         `json:"page"`
}

type ListRequest struct {
	Page int `json:"page" form:"page"`
	Size int `json:"size" form:"size"`
	Sort int `json:"sort" form:"sort"`
}

// Paging 分页
func Paging(p *Param, result interface{}, args ...interface{}) (*Paginator, error) {
	var (
		err       error
		param     ListRequest
		paginator Paginator
		count     int64
		offset    int
		tableName string
	)

	if err := p.C.Bind(&param); err != nil {
		logger.Errorf("参数绑定失败，错误：%v", err)
		return nil, err
	}

	db := p.DB

	if p.ShowSQL {
		db = db.Debug()
	}

	if param.Page < 1 {
		param.Page = 1
	}

	if param.Size < 1 {
		param.Size = 10
	}

	if param.Sort == 0 || param.Sort == -1 {
		db = db.Order("id desc")
	}

	if len(args) > 1 {
		tableName = fmt.Sprintf("`%s`.", args[1].(string))
	}

	if len(args) > 0 {
		for paramType, paramsValue := range args[0].(map[string]map[string]interface{}) {
			if paramType == "like" {
				for key, value := range paramsValue {
					db = db.Where(fmt.Sprintf("%v%v like ?", tableName, key), fmt.Sprintf("%%%v%%", value))
				}
			} else if paramType == "equal" {
				for key, value := range paramsValue {
					db = db.Where(fmt.Sprintf("%v%v = ?", tableName, key), value)
				}
			}
		}
	}

	err = db.Count(&count).Error

	if param.Page == 1 {
		offset = 0
	} else {
		offset = (param.Page - 1) * param.Size
	}
	err = db.Limit(param.Size).Offset(offset).Scan(result).Error
	if err != nil {
		logger.Errorf("数据查询失败，错误：%v", err)
		return nil, err
	}

	paginator.Total = count
	paginator.List = result
	paginator.Page = param.Page
	paginator.Size = param.Size
	paginator.TotalPage = int(math.Ceil(float64(count) / float64(param.Size)))

	return &paginator, nil
}
