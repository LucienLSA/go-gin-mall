package dao

import (
	"context"
	"ginmall/model"

	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(c context.Context) *UserDao {
	return &UserDao{NewDBClient(c)}
}

// 根据username判断是否存在该名字
func (dao *UserDao) ExistOrNotByUserName(username string) (user *model.User, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.User{}).Where("user_name =?", username).Count(&count).Error
	if count == 0 {
		return user, false, err
	}
	err = dao.DB.Model(&model.User{}).Where("user_name =?", username).First(&user).Error
	if err != nil {
		return user, false, err
	}
	return user, true, nil
}

// 创建用户
func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.DB.Model(&model.User{}).Create(&user).Error
}

// 根据id获取用户信息
func (dao *UserDao) GetUserById(uid uint) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("id =?", uid).First(&user).Error
	return user, err
}

// 根据ID更新用户信息
func (dao *UserDao) UpdateUserById(uId uint, user *model.User) error {
	return dao.DB.Model(&model.User{}).Where("id =?", uId).Updates(user).Error
}
