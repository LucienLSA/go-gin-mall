package v1

import (
	"net/http"

	"github.com/LucienLSA/go-gin-mall/pkg/util/ctl"
	"github.com/LucienLSA/go-gin-mall/pkg/util/logging"

	"github.com/LucienLSA/go-gin-mall/service"
	"github.com/LucienLSA/go-gin-mall/types"
	"github.com/gin-gonic/gin"
)

// 显示商品列表
func ListProductHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// 显示商品详情
func ShowProductHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

// 创建商品
func CreateProductHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.ProductCreateReq
		if err := c.ShouldBind(&req); err != nil {
			// 参数校验
			logging.LogrusObj.Infoln(err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		form, _ := c.MultipartForm()
		files := form.File["image"] // 上传商品图片
		l := service.GetProductSrv()
		resp, err := l.ProductCreate(c.Request.Context(), files, &req)
		if err != nil {
			logging.LogrusObj.Infoln(err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}

}
