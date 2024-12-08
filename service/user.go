package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/LucienLSA/go-gin-mall/conf"
	"github.com/LucienLSA/go-gin-mall/pkg/util/ctl"
	"github.com/LucienLSA/go-gin-mall/pkg/util/logging"

	"mime/multipart"
	"sync"

	"github.com/LucienLSA/go-gin-mall/consts"
	"github.com/LucienLSA/go-gin-mall/dao"
	"github.com/LucienLSA/go-gin-mall/model"
	"github.com/LucienLSA/go-gin-mall/pkg/util/email"
	"github.com/LucienLSA/go-gin-mall/pkg/util/jwt"
	"github.com/LucienLSA/go-gin-mall/pkg/util/upload"
	"github.com/LucienLSA/go-gin-mall/types"
)

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once // 单例模式

type UserSrv struct {
}

func GetUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{} // 检查是否已经实例化，没有则实例化
	})
	return UserSrvIns
}

func (s *UserSrv) UserRegister(c context.Context, req *types.UserRegisterReq) (resp interface{}, err error) {
	userDao := dao.NewUserDao(c)
	_, exist, err := userDao.ExistOrNotByUserName(req.UserName)
	if err != nil {
		logging.LogrusObj.Error(err)
		return
	}
	if exist {
		err = errors.New("用户名已存在")
		logging.LogrusObj.Error(err)
		return
	}
	user := &model.User{
		NickName: req.NickName,
		UserName: req.UserName,
		Status:   model.Active,
		Money:    consts.UserInitMoney,
	}
	// 加密密码
	if err = user.SetPassword(req.Password); err != nil {
		logging.LogrusObj.Error(err)
		return
	}
	// 加密money
	money, err := user.EncryptMoney(req.Key)
	if err != nil {
		logging.LogrusObj.Error(err)
		return
	}
	user.Money = money
	// 默认头像为local
	user.Avatar = consts.UserDefaultAvatarLocal
	if conf.Config.System.UploadModel == consts.UploadModeOSS {
		user.Avatar = consts.UserDefaultAvatarOss
	}

	// 创建用户
	err = userDao.CreateUser(user)
	if err != nil {
		logging.LogrusObj.Error(err)
		return
	}
	return
}

// 用户登录
func (s *UserSrv) UserLogin(c context.Context, req *types.UserLoginReq) (resp interface{}, err error) {
	var user *model.User
	userDao := dao.NewUserDao(c)
	// 查询用户是否存在
	user, exist, err := userDao.ExistOrNotByUserName(req.UserName)
	if err != nil {
		logging.LogrusObj.Error(err)
		return
	}
	if !exist {
		logging.LogrusObj.Error(err)
		return nil, errors.New("用户名不存在")
	}
	if !user.CheckPassword(req.Password) {
		return nil, errors.New("密码/账号不正确")
	}

	accessToken, refreshToken, err := jwt.GenerateToken(user.ID, req.UserName)
	if err != nil {
		logging.LogrusObj.Error(err)
		return nil, err
	}
	userResp := &types.UserInfoResp{
		ID:        user.ID,
		UserName:  user.UserName,
		NickName:  user.NickName,
		Avatar:    user.AvatarURL(),
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Unix(),
	}

	resp = &types.UserTokenData{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         userResp,
	}
	return
}

// 上传更新头像
func (s *UserSrv) UserAvatarUpload(c context.Context, req *types.UserServiceReq, file multipart.File, fileSize int64) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(c)
	if err != nil {
		logging.LogrusObj.Error(err)
		return nil, err
	}
	uId := u.Id
	userDao := dao.NewUserDao(c)
	// 根据id查询用户是否存在
	user, err := userDao.GetUserById(uId)
	if err != nil {
		logging.LogrusObj.Error(err)
		return nil, err
	}

	// 上传头像
	var path string
	if conf.Config.System.UploadModel == consts.UploadModeLocal {
		path, err = upload.AvatarUploadToLocalStatic(file, uId, user.UserName)
	} else {
		// path, err := upload.UploadToQiNiu()
	}
	if err != nil {
		logging.LogrusObj.Error(err)
		return nil, err
	}

	// 获取头像
	user.Avatar = path
	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		logging.LogrusObj.Error(err)
		return nil, err
	}
	return
}

// 用户昵称修改
func (s *UserSrv) UserInfoUpdate(c context.Context, req *types.UserUpdateInfoReq) (resp interface{}, err error) {
	// 找用户
	u, _ := ctl.GetUserInfo(c)
	userDao := dao.NewUserDao(c)
	// 根据id查询用户是否存在
	user, err := userDao.GetUserById(u.Id)
	if err != nil {
		logging.LogrusObj.Error(err)
		return nil, err
	}
	if req.NickName != "" {
		user.NickName = req.NickName
	}
	err = userDao.UpdateUserById(u.Id, user)
	if err != nil {
		logging.LogrusObj.Error(err)
		return nil, err
	}
	return
}

// UserInfoShow 用户信息展示
func (s *UserSrv) UserInfoShow(c context.Context, req *types.UserInfoShowReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(c)
	if err != nil {
		logging.LogrusObj.Error(err)
		return nil, err
	}
	user, err := dao.NewUserDao(c).GetUserById(u.Id)
	if err != nil {
		logging.LogrusObj.Error(err)
		return nil, err
	}
	resp = &types.UserInfoResp{
		ID:        user.ID,
		UserName:  user.UserName,
		NickName:  user.NickName,
		Avatar:    user.AvatarURL(),
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Unix(),
	}
	return
}

// 绑定邮箱服务
func (s *UserSrv) BindEmail(c context.Context, req *types.BindEmailServiceReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(c)
	if err != nil {
		logging.LogrusObj.Error(err)
		return nil, err
	}
	var address string
	token, err := jwt.GenerateEmailToken(u.Id, req.OpeartionType, req.Email, req.Password)
	if err != nil {
		logging.LogrusObj.Error(err)
		return nil, err
	}
	binder := email.NewEmailBinder()
	address = "http://" + conf.Config.System.Host + conf.Config.System.HttpPort + conf.Config.Email.Address + "?token=" + token
	mailText := fmt.Sprintf(consts.EmailOperationMap[req.OpeartionType], address)
	if err = binder.Bind(mailText, req.Email, "LucienMall"); err != nil {
		logging.LogrusObj.Error(err)
		return
	}
	return
}

// 验证邮箱
func (s *UserSrv) VerifyEmail(c context.Context, req *types.VerifyEmailServiceReq) (resp interface{}, err error) {
	var userId uint
	var email string
	var password string
	var operationType uint
	// 验证Token
	if req.Token == "" {
		err = errors.New("token为空")
		logging.LogrusObj.Error(err)
		return
	}
	claims, err := jwt.ParseEmailToken(req.Token)
	if err != nil {
		logging.LogrusObj.Error(err)
		return
	} else {
		userId = claims.UserID
		email = claims.Email
		password = claims.Password
		operationType = claims.OperationType
	}
	// 获取用户信息
	userDao := dao.NewUserDao(c)
	user, err := userDao.GetUserById(userId)
	if err != nil {
		logging.LogrusObj.Error(err)
		return
	}
	// 根据OPerationType更新用户信息
	switch operationType {
	// 绑定邮箱
	case consts.EmailOperationBinding:
		user.Email = email
	// 解绑邮箱
	case consts.EmailOperationNoBinding:
		user.Email = ""
		// 更新密码
	case consts.EmailOperationUpdatePassword:
		err = user.SetPassword(password)
		if err != nil {
			logging.LogrusObj.Error(err)
			return
		}
	default:
		return nil, errors.New("操作不符合")
	}
	// 更新用户信息
	err = userDao.UpdateUserById(userId, user)
	if err != nil {
		logging.LogrusObj.Error(err)
		return
	}
	resp = &types.UserInfoResp{
		ID:        user.ID,
		UserName:  user.UserName,
		NickName:  user.NickName,
		Email:     user.Email,
		Status:    user.Status,
		Avatar:    user.AvatarURL(),
		CreatedAt: user.CreatedAt.Unix(),
	}
	return
}

// 关注用户
func (s *UserSrv) UserFollow(c context.Context, req *types.UserFollowingReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(c)
	if err != nil {
		logging.LogrusObj.Error(err)
		return nil, err
	}
	err = dao.NewUserDao(c).FollowUser(u.Id, req.Id)
	return
}

// 取消关注用户
func (s *UserSrv) UserUnFollow(c context.Context, req *types.UserUnFollowingReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(c)
	if err != nil {
		logging.LogrusObj.Error(err)
		return nil, err
	}
	err = dao.NewUserDao(c).UnFollowUser(u.Id, req.Id)
	return
}
