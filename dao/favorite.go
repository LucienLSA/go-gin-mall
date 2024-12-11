package dao

import (
	"context"

	"github.com/LucienLSA/go-gin-mall/model"
	"gorm.io/gorm"
)

type FavoriteDao struct {
	*gorm.DB
}

func NewFavoriteDao(ctx context.Context) *FavoriteDao {
	return &FavoriteDao{NewDBClient(ctx)}
}

// 判断收藏是否存在
func (dao *FavoriteDao) FavoriteExistOrNot(pId, uId uint) (exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.Favorite{}).
		Where("product_id =? AND user_id =?", pId, uId).
		Count(&count).Error
	if count == 0 || err != nil {
		return false, err
	}
	return true, nil
}

// 创建收藏 收藏商品
func (dao *FavoriteDao) CreateFavorite(favorite *model.Favorite) error {
	err := dao.DB.Create(&favorite).Error
	return err
}
