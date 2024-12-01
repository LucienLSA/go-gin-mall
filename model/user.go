package model

import (
	"ginmall/conf"
	"ginmall/consts"

	"github.com/CocaineCong/secret"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	UserName       string `gorm:"unique"` // 用户名
	Email          string // 邮箱
	PasswordDigest string // 密码散列值
	NickName       string // 昵称
	Status         string // 状态
	Avatar         string `gorm:"size:1000"` // 头像
	Money          string // 余额
	Relations      []User `gorm:"many2many:relation;"` // 关系
}

const (
	PassWordCost        = 12       // 密码加密难度
	Active       string = "active" // 激活用户
)

// EncryptMoney 加密金额 (参考 https://github.com/CocaineCong/secret)
func (u *User) EncryptMoney(key string) (money string, err error) {
	aesObj, err := secret.NewAesEncrypt(conf.Config.EncryptSecret.MoneySecret, key, "", secret.AesEncrypt128, secret.AesModeTypeCBC)
	if err != nil {
		return
	}
	money = aesObj.SecretEncrypt(u.Money)

	return
}

// SetPassword 设置密码
func (u *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	u.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(password))
	return err == nil
}

// AvatarURL 获取头像URL
func (u *User) AvatarURL() string {
	// OSS上传模式
	if conf.Config.System.UploadModel == consts.UploadModeOSS {
		return u.Avatar
	}
	// 本地上传模式
	pConfig := conf.Config.PhotoPath
	return pConfig.PhotoHost + conf.Config.System.HttpPort + pConfig.AvatarPath + u.Avatar

}

// // GetUser 用ID获取用户
// func GetUser(ID interface{}) (User, error) {
// 	var user User
// 	result := dao.DB.First(&user, ID)
// 	return user, result.Error
// }

// // SetPassword 设置密码
// func (user *User) SetPassword(password string) error {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
// 	if err != nil {
// 		return err
// 	}
// 	user.PasswordDigest = string(bytes)
// 	return nil
// }

// // CheckPassword 校验密码
// func (user *User) CheckPassword(password string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
// 	return err == nil
// }
