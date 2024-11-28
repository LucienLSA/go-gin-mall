package server

import (
	"ginmall/api"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	// 中间件, 顺序不能改
	// r.Use(middleware.Session("something-very-secret"))
	// r.Use(middleware.Cors())
	// r.StaticFS("/static", http.Dir("./static"))
	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("/ping", api.Ping)

		// 用户登录
		// v1.POST("user/register", api.UserRegister)

		// 用户登录
		// v1.POST("user/login", api.UserLogin)

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
