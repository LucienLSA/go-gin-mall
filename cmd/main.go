package main

import (
	"ginmall/conf"
	"ginmall/server"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// 从配置文件读取配置
	conf.InitLodding()

	// 装载路由
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := server.NewRouter()
	r.Run(":3000")
}
