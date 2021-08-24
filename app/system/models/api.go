package models

import (
	"sky/common/models"
)

type Api struct {
	Title  string `gorm:"column:title;type:varchar(200);comment:标题" json:"title"`
	URL    string `gorm:"column:url;type:varchar(512);comment:接口地址" json:"url"`
	Method string `gorm:"column:method;type:varchar(45);comment:方法" json:"method"`
	models.BaseModel
}

func (Api) TableName() string {
	return "system_api"
}
