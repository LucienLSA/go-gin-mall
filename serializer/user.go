package serializer

import (
	"ginmall/conf"
	"ginmall/model"
)

// User 用户序列化器
type User struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Type     int    `json:"type"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	Avatar   string `json:"avatar"`
	CreateAt int64  `json:"create_at"`
}

// BuildUser 序列化用户
func BuildUser(user *model.User) *User {
	return &User{
		ID:       user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
		Email:    user.Email,
		Status:   user.Status,
		Avatar:   conf.Config.System.Host + conf.Config.System.HttpPort + conf.Config.PhotoPath.AvatarPath + user.Avatar,
		CreateAt: user.CreatedAt.Unix(),
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user *model.User) *Response {
	return &Response{
		Data: BuildUser(user),
	}
}
