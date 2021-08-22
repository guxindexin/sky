package models

import "sky/common/models"

type RolePermission struct {
	Role       int `gorm:"column:role;type:integer;comment:角色" json:"role"`
	Permission int `gorm:"column:menu;type:integer;comment:权限" json:"permission"`
	models.BaseModel
}

func (RolePermission) TableName() string {
	return "system_role_permission"
}
