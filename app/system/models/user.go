package models

import (
	"sky/common/models"
)

type User struct {
	Username string `gorm:"column:username;type:varchar(100);comment:用户名" json:"username"`
	Password string `gorm:"column:password;type:varchar(100);comment:密码" json:"-"`
	Nickname string `gorm:"column:nickname;type:varchar(100);comment:姓名" json:"nickname"`
	Tel      string `gorm:"column:tel;type:varchar(100);comment:手机号" json:"tel"`
	Email    string `gorm:"column:email;type:varchar(100);comment:邮箱" json:"email"`
	Sex      int    `gorm:"column:sex;type:smallint;comment:性别" json:"sex"`
	Status   bool   `gorm:"column:status;type:boolean;comment:状态" json:"status"`
	Remark   string `gorm:"column:remark;type:text;comment:备注" json:"remark"`
	Role     []int  `gorm:"column:role;type:integer[];comment:角色" json:"role"`
	Dept     []int  `gorm:"column:dept;type:integer[];comment:部门" json:"dept"`
	models.BaseModel
}

func (User) TableName() string {
	return "system_user"
}
