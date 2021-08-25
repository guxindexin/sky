package migrate

import (
	"sky/app/system/models"
	"sky/pkg/conn"
	"sky/pkg/logger"
)

func AutoMigrate() {
	// 同步数据结构
	logger.Info("开始同步数据结构...")
	err := conn.Orm.AutoMigrate(
		models.SystemModels...,
	)
	if err != nil {
		logger.Fatal("同步数据结构失败，错误：%v", err)
	}
	logger.Info("数据结构同步完成")

	// 同步初始数据
	logger.Info("开始同步基础数据...")
	err = InitDb(conn.Orm, "cmd/migrate/sql/init.sql")
	if err != nil {
		logger.Fatal("同步初始数据失败，错误：%v", err)
	}
	logger.Info("基础数据同步完成")
	return
}
