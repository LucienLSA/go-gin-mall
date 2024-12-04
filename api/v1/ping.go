package v1

import (
	"ginmall/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping 状态检查页面
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// CurrentUser 获取当前用户
func CurrentUser(c *gin.Context) *model.User {
	if user, _ := c.Get("user"); user != nil {
		if u, ok := user.(*model.User); ok {
			return u
		}
	}
	return nil
}

// // ErrorResponse 返回错误消息
// func ErrorResponse(err error) serializer.Response {
// 	if ve, ok := err.(validator.ValidationErrors); ok {
// 		for _, e := range ve {
// 			field := conf.T(fmt.Sprintf("Field.%s", e.Field()))
// 			tag := conf.T(fmt.Sprintf("Tag.Valid.%s", e.Tag()))
// 			return serializer.ParamErr(
// 				fmt.Sprintf("%s%s", field, tag),
// 				err,
// 			)
// 		}
// 	}
// 	if _, ok := err.(*json.UnmarshalTypeError); ok {
// 		return serializer.ParamErr("JSON类型不匹配", err)
// 	}

// 	return serializer.ParamErr("参数错误", err)
// }
