package fileee

import (
	"fmt"
	"ginmall/pkg/util/logging"
	"log"
	"os"
)

// 判断文件夹存在
func DirExistOrNot(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		log.Println("DirExistOrNot error:", err)
		return false
	}
	return s.IsDir()
}

// 创建文件夹
func CreateDir(dirName string) bool {
	err := os.MkdirAll(dirName, 0755)
	if err != nil {
		log.Println("CreateDir error:", err)
		return false
	}
	return true
}

// 检查文件的权限
func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

// 判断文件存在
func CheckNotExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

// 打开文件
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// 打开文件夹
func MustOpen(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		logging.LogrusObj.Info("os.Getwd err: ", err)
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath
	perm := CheckPermission(src)
	if perm {
		logging.LogrusObj.Info("file.CheckPermission Permission denied src: ", src)
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}
	if notExist := CheckNotExist(src); notExist {
		CreateDir(src)
	}
	f, err := Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		logging.LogrusObj.Info("Fail to OpenFile :", err)
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}
	return f, nil

}
