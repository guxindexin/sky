package models

import "sky/common/models"

type Dept struct {
	Name   string `gorm:"column:name;type:varchar(100);comment:名称" json:"name"`
	Leader int    `gorm:"column:leader;type:integer;comment:负责人" json:"leader"`
	Status bool   `gorm:"column:status;type:boolean;comment:状态" json:"status"`
	Remark string `gorm:"column:remark;type:text;comment:备注" json:"remark"`
	models.BaseModel
}

func (Dept) TableName() string {
	return "system_dept"
}
