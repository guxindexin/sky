package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sky/common/config"
	"sky/common/logger"
	"sky/common/router"
	"sky/common/tools"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configYml string
	StartCmd  = &cobra.Command{
		Use:          "server",
		Short:        "Start API server",
		Example:      "sky server -c config/settings.yml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/settings.yml", "Start server with provided configuration file")
}

func setup() {
	// 加载配置文件
	viper.SetConfigFile(configYml) // 指定配置文件
	err := viper.ReadInConfig()    // 读取配置信息
	if err != nil {                // 读取配置信息失败
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 日志配置
	logger.Setup()

	// 数据库配置
}

func run() (err error) {
	if viper.GetString("server.mode") == config.ModeProd.String() {
		gin.SetMode(gin.ReleaseMode)
	}

	// 路由引擎实例
	r := gin.Default()
	router.Load(r)

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", viper.GetString("server.host"), viper.GetInt("server.post")),
		Handler: r,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func() {
		// 服务连接
		if viper.GetBool("ssl.enable") {
			if err := srv.ListenAndServeTLS(viper.GetString("ssl.pem"), viper.GetString("ssl.key")); err != nil && err != http.ErrServerClosed {
				logger.Fatal("listen: ", err)
			}
		} else {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				logger.Fatal("listen: ", err)
			}
		}
	}()

	fmt.Println("Server run at:")
	fmt.Printf("-  Local:   http://localhost:%d/ \r\n", viper.GetInt("server.post"))
	fmt.Printf("-  Network: http://%s:%d/ \r\n", tools.GetLocaHonst(), viper.GetInt("server.post"))
	fmt.Printf("%s Enter Control + C Shutdown Server \r\n", time.Now().Format("2006-01-02 15:04:05.000"))
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Printf("%s Shutdown Server ... \r\n", time.Now().Format("2006-01-02 15:04:05"))

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown:", err)
	}
	logger.Info("Server exiting")

	return nil
}
