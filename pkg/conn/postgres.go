package conn

import (
	"database/sql"
	"sky/pkg/logger"
	"time"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Orm *gorm.DB
)

func Setup() {
	var (
		err   error
		sqlDB *sql.DB
	)
	sqlDB, err = sql.Open("postgres", viper.GetString("db.dsn"))
	if err != nil {
		logger.Fatalf("数据库连接失败，错误: %v", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		logger.Fatalf("无法连接到数据库，错误: %v", err)
	}

	Orm, err = gorm.Open(postgres.New(postgres.Config{
		Conn:                 sqlDB,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(viper.GetInt("db.maxIdleConn"))

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(viper.GetInt("db.maxOpenConn"))

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Duration(viper.GetInt("db.connMaxLifetime")) * time.Minute)
}
