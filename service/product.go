package service

import (
	"context"
	"strconv"

	"github.com/LucienLSA/go-gin-mall/conf"
	"github.com/LucienLSA/go-gin-mall/pkg/util/ctl"
	"github.com/LucienLSA/go-gin-mall/pkg/util/logging"

	"mime/multipart"
	"sync"

	"github.com/LucienLSA/go-gin-mall/consts"
	"github.com/LucienLSA/go-gin-mall/dao"
	"github.com/LucienLSA/go-gin-mall/model"
	"github.com/LucienLSA/go-gin-mall/pkg/util/upload"
	"github.com/LucienLSA/go-gin-mall/types"
)

var ProductSrvIns *ProductSrv
var ProductSrvOnce sync.Once // 保证初始化只执行一次 单例模式

type ProductSrv struct {
}

// 获取商品服务实例
func GetProductSrv() *ProductSrv {
	ProductSrvOnce.Do(func() {
		ProductSrvIns = &ProductSrv{}
	})
	return ProductSrvIns
}

// 创建商品
func (s *ProductSrv) ProductCreate(c context.Context, files []*multipart.FileHeader, req *types.ProductCreateReq) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(c)
	if err != nil {
		logging.LogrusObj.Error(err)
		return nil, err
	}
	uId := u.Id
	boss, _ := dao.NewUserDao(c).GetUserById(uId) // 获取当前登录用户 作为商家
	// 第一张作为封面图
	tmp, _ := files[0].Open()
	var path string
	// 上传封面图
	if conf.Config.System.UploadModel == consts.UploadModeLocal {
		path, err = upload.ProductUploadToLocalStatic(tmp, uId, req.Name)
	} else {
		// path, err = upload.ProductUploadToOss(tmp, files[0].Size)
	}
	if err != nil {
		logging.LogrusObj.Error(err)
		return nil, err
	}
	product := &model.Product{
		Name:          req.Name,
		CategoryID:    req.CategoryID,
		Title:         req.Title,
		Info:          req.Info,
		ImgPath:       path,
		Price:         req.Price,
		DiscountPrice: req.DiscountPrice,
		Num:           req.Num,
		OnSale:        true,
		BossID:        uId,
		BossName:      boss.UserName,
		BossAvatar:    boss.Avatar,
	}
	productDao := dao.NewProductDao(c)
	err = productDao.CreateProduct(product)
	if err != nil {
		logging.LogrusObj.Error(err)
		return
	}

	// 并发处理文件上传(商品图片)
	// 将上传的文件信息存储到数据库
	wg := new(sync.WaitGroup)
	// 添加待处理的文件数量到waitgroup
	wg.Add(len(files))
	for index, file := range files {
		// 生成文件名
		num := strconv.Itoa(index)
		// 打开文件
		tmp, _ = file.Open()
		// 上传文件
		if conf.Config.System.UploadModel == consts.UploadModeLocal {
			path, err = upload.ProductUploadToLocalStatic(tmp, uId, req.Name+num)
		} else {
			// path, err = upload.ProductUploadToOss(tmp, file.Size)
		}
		if err != nil {
			logging.LogrusObj.Error(err)
			return
		}
		productImg := &model.ProductImg{
			ProductID: product.ID,
			ImgPath:   path,
		}
		err = dao.NewProductImgDaoByDB(productDao.DB).CreateProductImg(productImg)
		if err != nil {
			logging.LogrusObj.Error(err)
			return
		}
		wg.Done()
	}
	wg.Wait()
	return
}

// 列出商品
func (s *ProductSrv) ProductList(c context.Context, req *types.ProductListReq) (resp interface{}, err error) {
	var total int64
	condition := make(map[string]interface{})
	if req.CategoryID != 0 {
		condition["category_id"] = req.CategoryID
	}
	productDao := dao.NewProductDao(c)
	products, _ := productDao.ListProductByCondition(condition, req.BasePage)
	total, err = productDao.CountProductByCondition(condition)
	if err != nil {
		logging.LogrusObj.Error(err)
		return
	}
	pRespList := make([]*types.ProductResp, 0)
	for _, p := range products {
		pResp := &types.ProductResp{
			ID:            p.ID,
			Name:          p.Name,
			CategoryID:    p.CategoryID,
			Title:         p.Title,
			Info:          p.Info,
			ImgPath:       p.ImgPath,
			Price:         p.Price,
			DiscountPrice: p.DiscountPrice,
			View:          p.View(),
			CreatedAt:     p.CreatedAt.Unix(),
			Num:           p.Num,
			OnSale:        p.OnSale,
			BossID:        p.BossID,
			BossName:      p.BossName,
			BossAvatar:    p.BossAvatar,
		}
		if conf.Config.System.UploadModel == consts.UploadModeLocal {
			pResp.BossAvatar = conf.Config.PhotoPath.PhotoHost + conf.Config.System.HttpPort + conf.Config.PhotoPath.AvatarPath + pResp.BossAvatar
			pResp.ImgPath = conf.Config.PhotoPath.PhotoHost + conf.Config.System.HttpPort + conf.Config.PhotoPath.ProductPath + pResp.ImgPath
		}
		pRespList = append(pRespList, pResp)
	}
	resp = &types.DataListResp{
		Item:  pRespList,
		Total: total,
	}
	return
}

// 显示详细商品信息
func (s *ProductSrv) ProductShow(ctx context.Context, req *types.ProductShowReq) (resp interface{}, err error) {
	p, err := dao.NewProductDao(ctx).ShowProductById(req.ID)
	if err != nil {
		logging.LogrusObj.Error(err)
		return
	}
	pResp := &types.ProductResp{
		ID:            p.ID,
		Name:          p.Name,
		CategoryID:    p.CategoryID,
		Title:         p.Title,
		Info:          p.Info,
		ImgPath:       p.ImgPath,
		Price:         p.Price,
		DiscountPrice: p.DiscountPrice,
		View:          p.View(),
		CreatedAt:     p.CreatedAt.Unix(),
		Num:           p.Num,
		OnSale:        p.OnSale,
		BossID:        p.BossID,
		BossName:      p.BossName,
		BossAvatar:    p.BossAvatar,
	}
	if conf.Config.System.UploadModel == consts.UploadModeLocal {
		pResp.BossAvatar = conf.Config.PhotoPath.PhotoHost + conf.Config.System.HttpPort + conf.Config.PhotoPath.AvatarPath + pResp.BossAvatar
		pResp.ImgPath = conf.Config.PhotoPath.PhotoHost + conf.Config.System.HttpPort + conf.Config.PhotoPath.ProductPath + pResp.ImgPath
	}
	resp = pResp
	return
}

// 搜索商品 TODO后续通过脚本同步数据MySQL到ES ES搜索
func (s *ProductSrv) ProductSearch(ctx context.Context, req *types.ProductSearchReq) (resp interface{}, err error) {
	products, count, err := dao.NewProductDao(ctx).SearchProduct(req.Info, req.BasePage)
	if err != nil {
		logging.LogrusObj.Error(err)
		return
	}
	pRespList := make([]*types.ProductResp, 0)
	for _, p := range products {
		pResp := &types.ProductResp{
			ID:            p.ID,
			Name:          p.Name,
			CategoryID:    p.CategoryID,
			Title:         p.Title,
			Info:          p.Info,
			ImgPath:       p.ImgPath,
			Price:         p.Price,
			DiscountPrice: p.DiscountPrice,
			View:          p.View(),
			CreatedAt:     p.CreatedAt.Unix(),
			Num:           p.Num,
			OnSale:        p.OnSale,
			BossID:        p.BossID,
			BossName:      p.BossName,
			BossAvatar:    p.BossAvatar,
		}
		if conf.Config.System.UploadModel == consts.UploadModeLocal {
			pResp.BossAvatar = conf.Config.PhotoPath.PhotoHost + conf.Config.System.HttpPort + conf.Config.PhotoPath.AvatarPath + pResp.BossAvatar
			pResp.ImgPath = conf.Config.PhotoPath.PhotoHost + conf.Config.System.HttpPort + conf.Config.PhotoPath.ProductPath + pResp.ImgPath
		}
		pRespList = append(pRespList, pResp)
	}
	resp = &types.DataListResp{
		Item:  pRespList,
		Total: count,
	}
	return
}

// 获取商品列表图片
func (s *ProductSrv) ProductImgList(ctx context.Context, req *types.ListProductImgReq) (resp interface{}, err error) {
	productImgs, _ := dao.NewProductImgDao(ctx).ListProductImgByProductId(req.ID)
	for i := range productImgs {
		if conf.Config.System.UploadModel == consts.UploadModeLocal {
			productImgs[i].ImgPath = conf.Config.PhotoPath.PhotoHost + conf.Config.System.HttpPort + conf.Config.PhotoPath.ProductPath + productImgs[i].ImgPath
		}
	}
	resp = &types.DataListResp{
		Item:  productImgs,
		Total: int64(len(productImgs)),
	}
	return
}

// 删除商品
func (s *ProductSrv) ProductDelete(ctx context.Context, req *types.ProductDeleteReq) (resp interface{}, err error) {
	u, _ := ctl.GetUserInfo(ctx)
	err = dao.NewProductDao(ctx).DeleteProduct(req.ID, u.Id)
	if err != nil {
		logging.LogrusObj.Error(err)
		return
	}
	return
}

// 更新商品信息
func (s *ProductSrv) ProductUpdate(ctx context.Context, req *types.ProductUpdateReq) (resp interface{}, err error) {
	product := &model.Product{
		Name:          req.Name,
		CategoryID:    req.CategoryID,
		Title:         req.Title,
		Info:          req.Info,
		Price:         req.Price,
		DiscountPrice: req.DiscountPrice,
		Num:           req.Num,
		OnSale:        req.OnSale,
	}
	err = dao.NewProductDao(ctx).UpdateProduct(req.ID, product)
	if err != nil {
		logging.LogrusObj.Error(err)
		return
	}
	return
}
