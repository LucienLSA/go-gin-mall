package v1

import (
	"errors"

	"github.com/LucienLSA/go-gin-mall/pkg/util/ctl"
	"github.com/LucienLSA/go-gin-mall/pkg/util/logging"

	"net/http"

	"github.com/LucienLSA/go-gin-mall/consts"
	"github.com/LucienLSA/go-gin-mall/pkg/e"
	"github.com/LucienLSA/go-gin-mall/service"
	"github.com/LucienLSA/go-gin-mall/types"
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

// 用户登录
func UserLoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.UserLoginReq
		if err := c.ShouldBind(&req); err != nil {
			logging.LogrusObj.Infoln(err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		l := service.GetUserSrv()
		resp, err := l.UserLogin(c.Request.Context(), &req)
		if err != nil {
			logging.LogrusObj.Infoln(err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

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

// 用户更新头像
func UserAvatarHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.UserServiceReq
		// 参数校验
		if err := c.ShouldBind(&req); err != nil {
			logging.LogrusObj.Infoln(err)
			c.JSON(http.StatusBadRequest, ErrorResponse(c, err))
			return
		}
		file, fileHeader, _ := c.Request.FormFile("file")
		if fileHeader == nil {
			err := errors.New(e.GetMsg(e.ErrorUploadFile))
			c.JSON(http.StatusBadRequest, ErrorResponse(c, err))
			logging.LogrusObj.Infoln(err)
			return
		}
		fileSize := fileHeader.Size
		l := service.GetUserSrv()
		resp, err := l.UserAvatarUpload(c.Request.Context(), &req, file, fileSize)
		if err != nil {
			logging.LogrusObj.Infoln(err)
			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

// 用户昵称修改
func UserUpdateHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.UserUpdateInfoReq
		if err := c.ShouldBind(&req); err != nil {
			// 参数校验
			logging.LogrusObj.Infoln(err)
			c.JSON(http.StatusBadRequest, ErrorResponse(c, err))
			return
		}
		l := service.GetUserSrv()
		resp, err := l.UserInfoUpdate(c.Request.Context(), &req)
		if err != nil {
			logging.LogrusObj.Infoln(err)
			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

func ShowUserInfoHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.UserInfoShowReq
		if err := c.ShouldBind(&req); err != nil {
			// 参数校验
			logging.LogrusObj.Infoln(err)
			c.JSON(http.StatusBadRequest, ErrorResponse(c, err))
			return
		}
		l := service.GetUserSrv()
		resp, err := l.UserInfoShow(c.Request.Context(), &req)
		if err != nil {
			logging.LogrusObj.Infoln(err)
			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

// 绑定邮箱
func BindEmailHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.BindEmailServiceReq
		if err := c.ShouldBind(&req); err != nil {
			// 参数校验
			logging.LogrusObj.Infoln(err)
			c.JSON(http.StatusBadRequest, ErrorResponse(c, err))
			return
		}
		l := service.GetUserSrv()
		resp, err := l.BindEmail(c.Request.Context(), &req)
		if err != nil {
			logging.LogrusObj.Infoln(err)
			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

// 验证邮箱
func VerifyEmailHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.VerifyEmailServiceReq
		if err := c.ShouldBind(&req); err != nil {
			// 参数校验
			logging.LogrusObj.Infoln(err)
			c.JSON(http.StatusBadRequest, ErrorResponse(c, err))
			return
		}
		l := service.GetUserSrv()
		resp, err := l.VerifyEmail(c.Request.Context(), &req)
		if err != nil {
			logging.LogrusObj.Infoln(err)
			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

// 关注用户
func UserFollowingHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.UserFollowingReq
		if err := c.ShouldBind(&req); err != nil {
			// 参数校验
			logging.LogrusObj.Infoln(err)
			c.JSON(http.StatusBadRequest, ErrorResponse(c, err))
			return
		}
		l := service.GetUserSrv()
		resp, err := l.UserFollow(c.Request.Context(), &req)
		if err != nil {
			logging.LogrusObj.Infoln(err)
			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

// 取关用户
func UserUnFollowingHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.UserUnFollowingReq
		if err := c.ShouldBind(&req); err != nil {
			// 参数校验
			logging.LogrusObj.Infoln(err)
			c.JSON(http.StatusBadRequest, ErrorResponse(c, err))
			return
		}
		l := service.GetUserSrv()
		resp, err := l.UserUnFollow(c.Request.Context(), &req)
		if err != nil {
			logging.LogrusObj.Infoln(err)
			c.JSON(http.StatusInternalServerError, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

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
