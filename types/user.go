package types

type UserServiceReq struct {
	NickName string `form:"nick_name" json:"nick_name"`
	UserName string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
	Key      string `form:"key" json:"key"` // 前端进行判断
}

type UserRegisterReq struct {
	NickName string `form:"nick_name" json:"nick_name"`
	UserName string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
	Key      string `form:"key" json:"key"` // 前端进行判断
}

type UserLoginReq struct {
	UserName string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
}

type UserInfoResp struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	NickName  string `json:"nickname"`
	Type      int    `json:"type"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}

type UserTokenData struct {
	User         interface{} `json:"user"`
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
}

type UserUpdateInfoReq struct {
	NickName string `form:"nick_name" json:"nick_name"`
	UserName string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
	Key      string `form:"key" json:"key"` // 前端进行判断
}

type UserInfoShowReq struct {
}

type BindEmailServiceReq struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	// OpertionType 1:绑定邮箱 2：解绑邮箱 3：改密码
	OpeartionType uint `form:"operation_type" json:"operation_type"`
}

type VerifyEmailServiceReq struct {
	Token string `json:"token" form:"token"`
}

type UserFollowingReq struct {
	Id uint `json:"id" form:"id"`
}

type UserUnFollowingReq struct {
	Id uint `json:"id" form:"id"`
}
