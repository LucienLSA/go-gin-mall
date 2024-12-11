package dao

import (
	"fmt"

	"github.com/LucienLSA/go-gin-mall/model"
)

//执行数据迁移

func Migration() (err error) {
	// 自动迁移模式
	err = DB.AutoMigrate(
		&model.User{},
		&model.Product{},
		&model.Favorite{},
		&model.ProductImg{},
	)
	fmt.Println("register stable success")
	return
}
