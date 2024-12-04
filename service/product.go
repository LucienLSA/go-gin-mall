package service

import (
	"context"
	"ginmall/conf"
	"ginmall/consts"
	"ginmall/dao"
	"ginmall/model"
	"ginmall/pkg/util/ctl"
	"ginmall/pkg/util/logging"
	"ginmall/pkg/util/upload"
	"ginmall/types"
	"mime/multipart"
	"sync"
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

	// // 并发处理文件上传(商品图片)
	// // 将上传的文件信息存储到数据库
	// wg := new(sync.WaitGroup)
	// // 添加待处理的文件数量到waitgroup
	// wg.Add(len(files))
	// for index, file := range files {
	// 	// 生成文件名
	// 	num := strconv.Itoa(index)
	// 	// 打开文件
	// 	tmp, _ = file.Open()
	// 	// 上传文件
	// 	if conf.Config.System.UploadModel == consts.UploadModeLocal {
	// 		path, err = upload.ProductUploadToLocalStatic(tmp, uId, req.Name+num)
	// 	} else {
	// 		path, err = upload.ProductUploadToOss(tmp, file.Size)
	// 	}
	// 	if err != nil {
	// 		logging.LogrusObj.Error(err)
	// 		return
	// 	}
	// 	productImg := &model.ProductImg{
	// 		ProductID: product.ID,
	// 		ImgPath:   path,
	// 	}

	// }
	return
}
