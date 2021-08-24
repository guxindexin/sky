package models

import "sky/common/models"

type RoleMenu struct {
	Role int `gorm:"column:role;type:integer;comment:角色" json:"role"`
	Menu int `gorm:"column:menu;type:integer;comment:菜单" json:"menu"`
	models.BaseModel
}

func (RoleMenu) TableName() string {
	return "system_role_menu"
}
