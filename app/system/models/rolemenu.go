package models

import "sky/common/models"

type RoleMenu struct {
	Role int `gorm:"column:role;type:integer;comment:角色" json:"role"`
	Menu int `gorm:"column:menu;type:integer;comment:菜单" json:"menu"`
	Type int `gorm:"column:type;type:smallint;comment:类型" json:"type"` // 1 菜单， 2 按钮
	models.BaseModel
}

func (RoleMenu) TableName() string {
	return "system_role_menu"
}
