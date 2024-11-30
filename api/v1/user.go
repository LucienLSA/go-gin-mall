package v1

import (
	"errors"
	"ginmall/consts"
	"ginmall/pkg/util/ctl"
	"ginmall/pkg/util/logging"
	"ginmall/service"
	"ginmall/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 用户注册接口
func UserRegisterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.UserRegisterReq
		if err := c.ShouldBind(&req); err != nil {
			logging.LogrusObj.Infoln(err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}

		// 参数检验
		if req.Key == "" || len(req.Key) != consts.EncryptMoneyKeyLength {
			err := errors.New("key长度错误,必须是6位数")
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}

		l := service.GetUserSrv()
		resp, err := l.UserRegister(c.Request.Context(), &req)
		if err != nil {
			logging.LogrusObj.Infoln(err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

// func UserRegister(c *gin.Context) {
// 	var service service.UserRegisterService
// 	if err := c.ShouldBind(&service); err == nil {
// 		res := service.Register()
// 		c.JSON(200, res)
// 	} else {
// 		c.JSON(200, ErrorResponse(err))
// 	}
// }

// // UserLogin 用户登录接口
// func UserLogin(c *gin.Context) {
// 	var service service.UserLoginService
// 	if err := c.ShouldBind(&service); err == nil {
// 		res := service.Login(c)
// 		c.JSON(200, res)
// 	} else {
// 		c.JSON(200, ErrorResponse(err))
// 	}
// }

// // UserMe 用户详情
// func UserMe(c *gin.Context) {
// 	user := CurrentUser(c)
// 	res := serializer.BuildUserResponse(*user)
// 	c.JSON(200, res)
// }

// // UserLogout 用户登出
// func UserLogout(c *gin.Context) {
// 	s := sessions.Default(c)
// 	s.Clear()
// 	s.Save()
// 	c.JSON(200, serializer.Response{
// 		Code: 0,
// 		Msg:  "登出成功",
// 	})
// }
