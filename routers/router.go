package routers

import (
	api "ginmall/api/v1"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	// 中间件, 顺序不能改
	// r.Use(middleware.Session("something-very-secret"))
	// r.Use(middleware.Cors())
	// r.StaticFS("/runtime/static", http.Dir("./runtime/static"))

	v1 := r.Group("/api/v1")
	{
		// 测试连通
		v1.GET("/ping", api.Ping)

		// 用户操作
		v1.POST("user/register", api.UserRegisterHandler())
		// v1.POST("user/login", api.UserLoginHandler)

		// 需要登录保护的
		// auth := v1.Group("/")
		// auth.Use(middleware.AuthRequired())
		{
			// User Routing
			// auth.GET("user/me", api.UserLogin)
			// auth.DELETE("user/logout", api.UserLogout)
		}
	}
	return r
}
