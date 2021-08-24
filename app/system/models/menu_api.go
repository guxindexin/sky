package models

import (
	"sky/common/models"
)

type MenuApi struct {
	Menu int    `gorm:"column:path;type:varchar(200);comment:菜单" json:"path"`
	Api  string `gorm:"column:name;type:varchar(100);comment:接口" json:"name"`
	models.BaseModel
}

func (MenuApi) TableName() string {
	return "system_menu_api"
}
