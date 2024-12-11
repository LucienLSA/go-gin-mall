package routers

import (
	"net/http"

	api "github.com/LucienLSA/go-gin-mall/api/v1"
	"github.com/LucienLSA/go-gin-mall/middleware"
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
		v1.GET("/ping")
		// 生成二维码
		v1.POST("/qrcode", api.GenerateQrcode)

		// 用户操作
		v1.POST("user/register", api.UserRegisterHandler())
		v1.POST("user/login", api.UserLoginHandler())

		// 商品操作
		v1.GET("product/list", api.ListProductHandler())
		v1.GET("product/show", api.ShowProductHandler())
		v1.POST("product/search", api.SearchProductsHandler())
		// v1.GET("product/imgs/list", api.ListProductImgHandler()) // 商品图片列表
		// v1.GET("category/list", api.ListCategoryHandler())       // 商品分类列表
		// v1.GET("carousels", api.ListCarouselsHandler())          // 首页轮播图

		// 需要登录保护的
		auth := v1.Group("/")
		auth.Use(middleware.AuthRequired())
		{
			// 用户操作
			// 更新用户头像
			auth.POST("user/avatar", api.UserAvatarHandler())
			// 更新用户信息
			auth.POST("user/update", api.UserUpdateHandler())
			// 获取用户信息
			auth.GET("user/show_info", api.ShowUserInfoHandler())
			// 关注用户
			auth.POST("user/following", api.UserFollowingHandler())
			// 取关用户
			auth.POST("user/unfollowing", api.UserUnFollowingHandler())
			// 邮箱验证
			auth.GET("user/verify_email", api.VerifyEmailHandler())
			// 绑定邮箱
			auth.POST("user/bind_email", api.BindEmailHandler())

			// 商品操作
			auth.POST("product/create", api.CreateProductHandler())
		}
	}
	return r
}
