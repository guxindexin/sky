package permission

import (
	"sky/pkg/conn"
	"sky/pkg/logger"
	"sky/pkg/tools/response"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/spf13/viper"

	"github.com/casbin/casbin/v2"
	gormAdapter "github.com/casbin/gorm-adapter/v3"
)

var Enforcer *casbin.SyncedEnforcer

func CasbinSetup() *casbin.SyncedEnforcer {
	var (
		err     error
		adapter *gormAdapter.Adapter
	)
	adapter, err = gormAdapter.NewAdapterByDBWithCustomTable(conn.Orm, nil, viper.GetString("casbin.tableName"))
	if err != nil {
		logger.Fatal("创建 casbin gorm adapter 失败，错误：%v", err)
	}

	Enforcer, err = casbin.NewSyncedEnforcer(viper.GetString("casbin.rbacModel"), adapter)
	if err != nil {
		logger.Fatal("创建 casbin enforcer 失败，错误：%v", err)
	}

	err = Enforcer.LoadPolicy()
	if err != nil {
		logger.Fatal("从数据库加载策略失败，错误：%v", err)
	}

	// 定时同步策略
	if viper.GetBool("casbin.isTiming") {
		// 间隔多长时间同步一次权限策略，单位：秒
		Enforcer.StartAutoLoadPolicy(time.Second * time.Duration(viper.GetInt("casbin.intervalTime")))
	}

	return Enforcer
}

func CheckPermMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取资源
		obj := c.Request.URL.Path
		//获取方法
		act := c.Request.Method
		//获取实体
		sub := c.GetString("username")

		isAdmin := c.GetBool("isAdmin")
		if isAdmin {
			c.Next()
		} else {
			//判断策略中是否存在
			if ok, _ := Enforcer.Enforce(sub, obj, act); ok {
				c.Next()
			} else {
				response.Error(c, nil, response.NoPermissionError)
				c.Abort()
			}
		}
	}
}
