package dao

import (
	"context"

	"github.com/LucienLSA/go-gin-mall/model"
	"github.com/LucienLSA/go-gin-mall/types"
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

// 根据条件列出商品
func (dao *ProductDao) ListProductByCondition(condition map[string]interface{}, page types.BasePage) (products []*model.Product, err error) {
	err = dao.DB.Where(condition).
		Offset((page.PageNum - 1) * page.PageSize). // 偏移量
		Limit(page.PageSize).                       // 限制查询结果的数量
		Find(&products).Error
	return products, err
}

// 根据条件获取商品的数量
func (dao *ProductDao) CountProductByCondition(condition map[string]interface{}) (count int64, err error) {
	err = dao.DB.Model(&model.Product{}).
		Where(condition).
		Count(&count).Error
	return count, err
}

// 根据id 获取商品
func (dao *ProductDao) ShowProductById(id uint) (product *model.Product, err error) {
	err = dao.DB.Model(&model.Product{}).
		Where("id = ?", id).
		First(&product).Error
	return product, err
}

// 根据信息 搜索商品
func (dao *ProductDao) SearchProduct(info string, page types.BasePage) (products []*model.Product, count int64, err error) {
	err = dao.DB.Model(&model.Product{}).
		Where("name LIKE ? OR info LIKE ?", "%"+info+"%", "%"+info+"%"). // 条件筛选 搜素词info使用LIKE实现模糊查询
		Offset((page.PageNum - 1) * page.PageSize).                      // 偏移量
		Limit(page.PageSize).Find(&products).Error

	if err != nil {
		return nil, 0, err
	}

	err = dao.DB.Model(&model.Product{}).
		Where("name LIKE ? OR info LIKE ?", "%"+info+"%", "%"+info+"%").
		Count(&count).Error
	return products, count, err
}

// 删除商品
func (dao *ProductDao) DeleteProduct(pId, uId uint) error {
	return dao.DB.Model(&model.Product{}).
		Where("id = ? AND boss_id = ?", pId, uId).
		Delete(&model.Product{}).Error
}

// 更新商品信息
func (dao *ProductDao) UpdateProduct(pId uint, product *model.Product) error {
	return dao.DB.Model(&model.Product{}).
		Where("id = ?", pId).
		Updates(product).Error
}
