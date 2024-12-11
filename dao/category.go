package dao

import (
	"context"

	"github.com/LucienLSA/go-gin-mall/model"
	"gorm.io/gorm"
)

type CategoryDao struct {
	*gorm.DB
}

func NewCategoryDao(ctx context.Context) *CategoryDao {
	return &CategoryDao{NewDBClient(ctx)}
}

func (dao *CategoryDao) ListCategory() (r []*model.Category, err error) {
	err = dao.DB.Model(&model.Category{}).Find(&r).Error
	return
}
