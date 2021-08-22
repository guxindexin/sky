package models

import "sky/common/models"

type Role struct {
	Name   string `gorm:"column:name;type:varchar(100);comment:名称" json:"name"`
	Status bool   `gorm:"column:status;type:boolean;comment:状态" json:"status"`
	Remark string `gorm:"column:remark;type:text;comment:备注" json:"remark"`
	models.BaseModel
}

func (Role) TableName() string {
	return "system_role"
}
