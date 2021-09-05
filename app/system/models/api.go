package models

import (
	"sky/common/models"
)

type Api struct {
	Title  string `gorm:"column:title;type:varchar(200);comment:标题" json:"title" binding:"required"`
	URL    string `gorm:"column:url;type:varchar(512);comment:接口地址" json:"url" binding:"required"`
	Method string `gorm:"column:method;type:varchar(45);comment:方法" json:"method" binding:"required"`
	Group  int    `gorm:"column:group;type:integer;comment:分组" json:"group" binding:"required"`
	Remark string `gorm:"column:remark;type:text;comment:备注" json:"remark"`
	models.BaseModel
}

func (Api) TableName() string {
	return "system_api"
}
