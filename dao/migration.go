package dao

import (
	"fmt"
	"ginmall/model"
)

//执行数据迁移

func Migration() (err error) {
	// 自动迁移模式
	err = DB.AutoMigrate(
		&model.User{},
		&model.Product{},
	)
	fmt.Println("register stable success")
	return
}
