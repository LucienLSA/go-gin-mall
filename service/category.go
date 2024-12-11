package service

import (
	"context"
	"sync"

	"github.com/LucienLSA/go-gin-mall/dao"
	"github.com/LucienLSA/go-gin-mall/pkg/util/logging"
	"github.com/LucienLSA/go-gin-mall/types"
)

var CategorySrvIns *CategorySrv
var CategorySrvOnce sync.Once

type CategorySrv struct {
}

func GetCategorySrv() *CategorySrv {
	CategorySrvOnce.Do(func() {
		CategorySrvIns = &CategorySrv{}
	})
	return CategorySrvIns
}

// 列出商品分类的列表
func (s *CategorySrv) CategoryList(ctx context.Context, req *types.ListCategoryReq) (resp interface{}, err error) {
	categories, err := dao.NewCategoryDao(ctx).ListCategory()
	if err != nil {
		logging.LogrusObj.Error(err)
		return
	}
	cResp := make([]*types.ListCategoryResp, 0)
	for _, v := range categories {
		cResp = append(cResp, &types.ListCategoryResp{
			ID:           v.ID,
			CategoryName: v.CategoryName,
			CreatedAt:    v.CreatedAt.Unix(),
		})
	}
	resp = types.DataListResp{
		Item:  cResp,
		Total: int64(len(cResp)),
	}
	return
}
