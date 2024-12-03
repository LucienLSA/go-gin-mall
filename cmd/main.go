package main

import (
	"context"
	"fmt"
	"ginmall/cache"
	"ginmall/conf"
	"ginmall/dao"
	"ginmall/pkg/util/logging"
	"ginmall/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func main() {
	// 从配置文件读取配置
	InitLoading()
	// 装载路由
	// gin.SetMode(os.Getenv("GIN_MODE"))
	// 初始化路由
	InitRoutes()
}

// Init 初始化配置项
func InitLoading() {
	// 从本地读取环境变量
	// godotenv.Load()
	// 读取配置文件
	conf.ReadConfig()
	// 初始化Mysql
	dao.InitMySQL()
	// 初始化Redisd
	cache.InitRedis()
	// 初始化日志
	logging.InitLog()
	// 设置日志级别
	// util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 读取翻译文件
	if err := conf.LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		logging.LogrusObj.Panic("翻译文件加载失败", err)
		// util.Log().Panic("翻译文件加载失败", err)
	}
	fmt.Println("加载配置完成...")
	go scriptStarting()
	// 连接数据库
	// model.Database(os.Getenv("MYSQL_DSN"))
	// cache.Redis()
}

func InitRoutes() {
	// 注册路由
	gin.SetMode(conf.Config.System.RunMode)
	// 正常启动
	// normalStarting()
	// 优雅启动和关闭
	gracefulStartingAndClosing()

}

// 正常启动
func normalStarting() {
	r := routers.NewRouter()
	readTimeout := conf.Config.System.ReadTimeout
	writeTimeout := conf.Config.System.WriteTimeout
	endPoint := fmt.Sprintf("%s%s", conf.Config.System.Host, conf.Config.System.HttpPort)
	maxHeaderBytes := 1 << 20
	// 正常启动服务
	// err := r.Run(conf.Config.System.HttpPort)
	server := &http.Server{
		Addr:           endPoint,
		Handler:        r,
		ReadTimeout:    readTimeout * time.Second,
		WriteTimeout:   writeTimeout * time.Second,
		MaxHeaderBytes: maxHeaderBytes,
	}
	err := server.ListenAndServe()
	if err != nil {
		logging.LogrusObj.Panic("启动服务失败", err)
	}
	fmt.Printf("启动配置成功... \n start server listening %s", endPoint)
}
func gracefulStartingAndClosing() {
	gin.SetMode(conf.Config.System.RunMode)
	endless.DefaultReadTimeOut = conf.Config.System.ReadTimeout * time.Second
	endless.DefaultWriteTimeOut = conf.Config.System.WriteTimeout * time.Second
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf("%s%s", conf.Config.System.Host, conf.Config.System.HttpPort)
	server := endless.NewServer(endPoint, routers.NewRouter())
	server.BeforeBegin = func(add string) {
		log.Println("Actual pid is ", syscall.Getpid())
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Println("server err: ", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Waiting for the system signal to gracefully shutdown.")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logging.LogrusObj.Info("Server forced to shutdown: ", err)
	}
	log.Println("Server exiting")
}

func scriptStarting() {
	// 启动一些脚本
}
