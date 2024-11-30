package service

import (
	"context"
	"ginmall/conf"
	"ginmall/consts"
	"ginmall/dao"
	"ginmall/model"
	"ginmall/pkg/util/logging"
	"ginmall/types"
	"sync"

	"github.com/pkg/errors"
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
