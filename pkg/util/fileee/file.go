package fileee

import (
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
