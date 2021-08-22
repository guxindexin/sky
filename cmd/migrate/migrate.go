package migrate

import (
	"sky/app/system/models"
	"sky/pkg/conn"
	"sky/pkg/logger"
)

func AutoMigrate() {
	err := conn.Orm.AutoMigrate(
		models.SystemModels...,
	)
	if err != nil {
		logger.Fatal("同步数据结构失败，错误：%v", err)
	}
	return
}
