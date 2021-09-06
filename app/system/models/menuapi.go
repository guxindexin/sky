package models

import (
	"sky/common/models"
)

type MenuApi struct {
	Menu int `gorm:"column:menu;type:integer;comment:菜单" json:"menu"`
	Api  int `gorm:"column:api;type:integer;comment:接口" json:"api"`
	models.BaseModel
}

func (MenuApi) TableName() string {
	return "system_menu_api"
}
