package routers

import (
	api "ginmall/api/v1"
	"ginmall/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Session("something-very-secret"))
	r.Use(middleware.Cors())
	// 静态文件读取
	r.StaticFS("/runtime", http.Dir("./runtime"))

	v1 := r.Group("/api/v1")
	{
		// 测试连通
		v1.GET("/ping", api.Ping)

		// 用户操作
		v1.POST("user/register", api.UserRegisterHandler())
		v1.POST("user/login", api.UserLoginHandler())

		// 需要登录保护的
		auth := v1.Group("/")
		auth.Use(middleware.AuthRequired())
		{
			// 用户操作
			// 更新用户头像
			auth.POST("user/avatar", api.UserAvatarHandler())
		}
	}
	return r
}
