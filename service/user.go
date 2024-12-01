package service

import (
	"context"
	"errors"
	"ginmall/conf"
	"ginmall/consts"
	"ginmall/dao"
	"ginmall/model"
	"ginmall/pkg/util/ctl"
	"ginmall/pkg/util/jwt"
	"ginmall/pkg/util/logging"
	"ginmall/pkg/util/upload"
	"ginmall/types"
	"mime/multipart"
	"sync"
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
