package service

import (
	"context"

	"sync"

	"github.com/LucienLSA/go-gin-mall/dao"
	"github.com/LucienLSA/go-gin-mall/pkg/util/logging"
	"github.com/LucienLSA/go-gin-mall/types"
)

var CarouselSrvIns *CarouselSrv
var CarouselSrvOnce sync.Once

type CarouselSrv struct {
}

func GetCarouselSrv() *CarouselSrv {
	CarouselSrvOnce.Do(func() {
		CarouselSrvIns = &CarouselSrv{}
	})
	return CarouselSrvIns
}

// 列出轮播图的服务
func (s *CarouselSrv) ListCarousel(ctx context.Context, req *types.ListCarouselReq) (resp interface{}, err error) {
	carousels, err := dao.NewCarouselDao(ctx).ListCarousels()
	if err != nil {
		logging.LogrusObj.Error(err)
		return
	}

	resp = &types.DataListResp{
		Item:  carousels,
		Total: int64(len(carousels)),
	}
	return
}
