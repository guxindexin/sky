package models

import (
	"sky/common/models"
)

type ApiGroup struct {
	Name   string `gorm:"column:name;type:varchar(200);comment:名称" json:"name"`
	Remark string `gorm:"column:remark;type:text;comment:备注" json:"remark"`
	models.BaseModel
}

func (ApiGroup) TableName() string {
	return "system_api_group"
}
