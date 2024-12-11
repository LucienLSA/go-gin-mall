package service

import (
	"context"
	"errors"
	"sync"

	"github.com/LucienLSA/go-gin-mall/dao"
	"github.com/LucienLSA/go-gin-mall/model"
	"github.com/LucienLSA/go-gin-mall/pkg/util/ctl"
	"github.com/LucienLSA/go-gin-mall/pkg/util/logging"
	"github.com/LucienLSA/go-gin-mall/types"
)

var FavoriteSrvIns *FavoriteSrv
var FavoriteSrvOnce sync.Once

type FavoriteSrv struct{}

func GetFavoriteSrv() *FavoriteSrv {
	FavoriteSrvOnce.Do(func() {
		FavoriteSrvIns = &FavoriteSrv{}
	})
	return FavoriteSrvIns
}

// 创建收藏夹
func (s *FavoriteSrv) FavoriteCreate(ctx context.Context, req *types.FavoriteCreateReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		logging.LogrusObj.Error(err)
		return nil, err
	}
	// 判断收藏夹是否已存在
	fDao := dao.NewFavoriteDao(ctx)
	exist, _ := fDao.FavoriteExistOrNot(req.ProductId, u.Id)
	if exist {
		err = errors.New("收藏夹已存在")
		logging.LogrusObj.Error(err)
		return
	}
	// 获取用户
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(u.Id)
	if err != nil {
		logging.LogrusObj.Error(err)
		return
	}

	// 获取商家
	bossDao := dao.NewUserDaoByDB(userDao.DB)
	boss, err := bossDao.GetUserById(req.BossId)
	if err != nil {
		logging.LogrusObj.Error(err)
		return
	}

	// 获取商品
	product, err := dao.NewProductDao(ctx).ShowProductById(req.ProductId)
	if err != nil {
		logging.LogrusObj.Error(err)
		return
	}

	favorite := &model.Favorite{
		UserID:    u.Id,
		User:      *user,
		ProductID: req.ProductId,
		Product:   *product,
		BossID:    req.BossId,
		Boss:      *boss,
	}
	err = fDao.CreateFavorite(favorite)
	if err != nil {
		logging.LogrusObj.Error(err)
		return
	}
	return
}

// 删除收藏夹
func (s *FavoriteSrv) FavoriteDelete(ctx context.Context, req *types.FavoriteDeleteReq) (resp interface{}, err error) {
	favoriteDao := dao.NewFavoriteDao(ctx)
	err = favoriteDao.DB.Delete(&model.Favorite{}, req.Id).Error
	if err != nil {
		logging.LogrusObj.Error(err)
		return
	}
	return
}
