package middleware

import (
	"net/http"

	"github.com/LucienLSA/go-gin-mall/consts"
	"github.com/LucienLSA/go-gin-mall/pkg/e"
	"github.com/LucienLSA/go-gin-mall/pkg/util/ctl"
	"github.com/LucienLSA/go-gin-mall/pkg/util/jwt"
	"github.com/gin-gonic/gin"
)

// token验证中间件
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = e.SUCCESS
		// 获取access_token和refresh_token
		accessToken := c.GetHeader("access_token")
		refreshToken := c.GetHeader("refresh_token")
		if accessToken == "" {
			code = e.InvalidParams
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": e.GetMsg(code),
				"data":    "Token不能为空",
			})
			c.Abort()
			return
		}
		// 解析Token，包括每次请求时刷新token，返回到新的token
		newAccessToken, newRefreshToken, err := jwt.ParseRefreshToken(accessToken, refreshToken)
		if err != nil {
			code = e.ErrorAuthCheckTokenFail
		}
		if code != e.SUCCESS {
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": e.GetMsg(code),
				"data":    "Token验证失败",
				"error":   err.Error(),
			})
			c.Abort()
			return
		}
		// 验证新的Token
		claims, err := jwt.ParseToken(newAccessToken)
		if err != nil {
			code = e.ErrorAuthCheckTokenFail
			c.JSON(http.StatusOK, gin.H{
				"status":  code,
				"message": e.GetMsg(code),
				"data":    err.Error(),
			})
			c.Abort()
			return
		}
		// 设置刷新后新的Token
		SetToken(c, newAccessToken, newRefreshToken)
		// 设置用户信息到上下文
		c.Request = c.Request.WithContext(ctl.NewContext(c.Request.Context(), &ctl.UserInfo{Id: claims.ID}))
		ctl.InitUserInfo(c.Request.Context())
		c.Next()
	}
}

// 将刷新后的token设置到cookie和header中
func SetToken(c *gin.Context, accessToken, refreshToken string) {
	secure := IsHttps(c)
	c.Header(consts.AccessTokenHeader, accessToken)
	c.Header(consts.RefreshTokenHeader, refreshToken)
	c.SetCookie(consts.AccessTokenHeader, accessToken, consts.MaxAge, "/", "", secure, true)
	c.SetCookie(consts.RefreshTokenHeader, refreshToken, consts.MaxAge, "/", "", secure, true)
}

// 判断是否https
func IsHttps(c *gin.Context) bool {
	if c.GetHeader(consts.HeaderForwardedProto) == "https" || c.Request.TLS != nil {
		return true
	}
	return false
}
