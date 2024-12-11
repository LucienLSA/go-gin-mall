package dao

import (
	"context"

	"github.com/LucienLSA/go-gin-mall/model"
	"github.com/LucienLSA/go-gin-mall/types"
	"gorm.io/gorm"
)

type CarouselDao struct {
	*gorm.DB
}

func NewCarouselDao(ctx context.Context) *CarouselDao {
	return &CarouselDao{NewDBClient(ctx)}
}

func (dao *CarouselDao) ListCarousels() (r []*types.ListCarouselResp, err error) {
	err = dao.DB.Model(&model.Carousel{}).
		Select("id, img_path, product_id, UNIX_TIMESTAMP(created_at)").
		Find(&r).Error
	return
}
