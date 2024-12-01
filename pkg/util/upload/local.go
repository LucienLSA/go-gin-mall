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
