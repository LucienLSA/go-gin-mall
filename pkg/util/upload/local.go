package upload

import (
	"fmt"
	"ginmall/conf"
	"ginmall/pkg/util/fileee"
	"ginmall/pkg/util/logging"
	"io/ioutil"
	"mime/multipart"
	"strconv"
)

// AvatarUploadToLocalStatic 上传头像到本地静态目录
func AvatarUploadToLocalStatic(file multipart.File, userId uint, userName string) (filePath string, err error) {
	bId := strconv.Itoa(int(userId))
	basePath := "." + conf.Config.PhotoPath.AvatarPath + "user" + bId + "/"
	if !fileee.DirExistOrNot(basePath) {
		fileee.CreateDir(basePath)
	}
	// 头像路径
	avatarPath := fmt.Sprintf("%s%s.jpg", basePath, userName)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		logging.LogrusObj.Error(err)
		return "", err
	}
	err = ioutil.WriteFile(avatarPath, content, 0666)
	if err != nil {
		logging.LogrusObj.Error(err)
		return "", err
	}
	return fmt.Sprintf("user%s/%s.jpg", bId, userName), nil

}

// 上传商品图片到本地文件中
func ProductUploadToLocalStatic(file multipart.File, bossId uint, productName string) (filePath string, err error) {
	bId := strconv.Itoa(int(bossId))
	// 商品图片路径
	basePath := "." + conf.Config.PhotoPath.ProductPath + "boss" + bId + "/"
	// 创建商品图片目录
	if !fileee.DirExistOrNot(basePath) {
		fileee.CreateDir(basePath)
	}
	// 输出商品图片路径
	productPath := fmt.Sprintf("%s%s.jpg", basePath, productName)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		logging.LogrusObj.Error(err)
		return "", err
	}
	// 写入商品图片到本地文件中
	err = ioutil.WriteFile(productPath, content, 0666)
	if err != nil {
		logging.LogrusObj.Error(err)
		return "", err
	}
	return fmt.Sprintf("boss%s/%s.jpg", bId, productName), nil
}
