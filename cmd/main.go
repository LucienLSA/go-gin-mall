package main

import (
	"fmt"
	"ginmall/conf"
	"ginmall/dao"
	"ginmall/pkg/util/logging"
	"ginmall/server"
)

func main() {
	// 从配置文件读取配置
	InitLoading()

	// 装载路由
	// gin.SetMode(os.Getenv("GIN_MODE"))
	r := server.NewRouter()
	_ = r.Run(conf.Config.System.HttpPort)
	fmt.Println("启动配成功...")
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
func scriptStarting() {
	// 启动一些脚本
}
