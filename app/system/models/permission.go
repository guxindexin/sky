package models

import "sky/common/models"

type Permission struct {
	Path        string `gorm:"column:path;type:varchar(200);comment:路径" json:"path"`
	Name        string `gorm:"column:name;type:varchar(100);comment:路由名称" json:"name"`
	Component   string `gorm:"column:component;type:varchar(200);comment:组件路径" json:"component"`
	Title       string `gorm:"column:title;type:varchar(200);comment:标题" json:"title"`
	IsLink      string `gorm:"column:is_link;type:varchar(200);comment:是否超链接" json:"isLink"`
	IsHide      string `gorm:"column:is_hide;type:varchar(200);comment:是否隐藏" json:"isHide"`
	IsKeepAlive string `gorm:"column:is_keep_alive;type:varchar(200);comment:是否缓存组件状态" json:"isKeepAlive"`
	IsAffix     string `gorm:"column:is_affix;type:varchar(200);comment:是否固定在 tagsView 栏上" json:"isAffix"`
	IsIframe    string `gorm:"column:is_iframe;type:varchar(200);comment:是否内嵌窗口" json:"isIframe"`
	Auth        string `gorm:"column:auth;type:varchar(200);comment:权限标识" json:"auth"`
	Icon        string `gorm:"column:icon;type:varchar(200);comment:图标" json:"icon"`
	models.BaseModel
}

func (Permission) TableName() string {
	return "system_permission"
}
