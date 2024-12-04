package dao

import (
	"context"
	"ginmall/model"

	"gorm.io/gorm"
)

type ProductDao struct {
	*gorm.DB
}

// 调用NewProductDao方法创建ProductDao实例
func NewProductDao(ctx context.Context) *ProductDao {
	return &ProductDao{NewDBClient(ctx)}
}

// CreateProduct 创建商品
func (dao *ProductDao) CreateProduct(product *model.Product) error {
	return dao.DB.Model(&model.Product{}).Create(product).Error
}
