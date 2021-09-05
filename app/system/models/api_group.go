package models

import (
	"sky/common/models"
)

type ApiGroup struct {
	App    string `gorm:"column:app;type:varchar(200);comment:应用" json:"app"`
	Remark string `gorm:"column:remark;type:text;comment:备注" json:"remark"`
	models.BaseModel
}

func (ApiGroup) TableName() string {
	return "system_api_group"
}
