package e

const (
	SUCCESS               = 200 // 成功
	UpdatePasswordSuccess = 201 // 修改密码成功
	NotExistIdentifier    = 202 // 资源不存在
	ERROR                 = 500 // 服务器内部错误
	InvalidParams         = 400 // 参数错误

	ErrorExistNickName      = 1001 // 昵称已存在
	ErrorExistUser          = 1002 // 用户已存在
	ErrorNotExistUser       = 1003 // 用户不存在
	ErrorNotCompare         = 1004 // 比较不一致
	ErrorNotComparePassword = 1005 // 两次输入密码不一致
	ErrorFailEncryption     = 1006 // 加密失败
	ErrorNotExistProduct    = 1007 // 商品不存在
	ErrorNotExistAddress    = 1008 // 地址不存在
	ErrorExistFavorite      = 1009 // 收藏已存在
	ErrorUserNotFound       = 1010 // 用户不存在

	//店家错误
	ErrorBossCheckTokenFail        = 2001 // 验证token失败
	ErrorBossCheckTokenTimeout     = 2002 // token过期
	ErrorBossToken                 = 2003 // token错误
	ErrorBoss                      = 2004 // 店家错误
	ErrorBossInsufficientAuthority = 2005 // 权限不足
	ErrorBossProduct               = 2006 // 店家商品错误

	// 购物车
	ErrorProductExistCart = 3007 // 商品已存在购物车
	ErrorProductMoreCart  = 3008 // 商品超过购物车数量限制

	//管理员错误
	ErrorAuthCheckTokenFail        = 4001 //管理员检查token 错误
	ErrorAuthCheckTokenTimeout     = 4002 //管理员token 过期
	ErrorAuthToken                 = 4003 // 管理员token 错误
	ErrorAuth                      = 4004 // 管理员错误
	ErrorAuthInsufficientAuthority = 4005 // 权限不足
	ErrorReadFile                  = 4006 // 读取文件错误
	ErrorSendEmail                 = 4007 // 发送邮件错误
	ErrorCallApi                   = 4008 // 调用api错误
	ErrorUnmarshalJson             = 4009 // 反序列化json错误
	ErrorAdminFindUser             = 4010 // 管理员查找用户错误

	//数据库错误
	ErrorDatabase = 5001

	//对象存储错误
	ErrorOss        = 6001
	ErrorUploadFile = 6002
)
