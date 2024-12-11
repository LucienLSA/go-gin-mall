package v1

import (
	"net/http"

	"github.com/LucienLSA/go-gin-mall/pkg/util/ctl"
	"github.com/LucienLSA/go-gin-mall/pkg/util/logging"
	"github.com/LucienLSA/go-gin-mall/service"
	"github.com/LucienLSA/go-gin-mall/types"
	"github.com/gin-gonic/gin"
)

// 商品分类的列表
func ListCategoryHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.ListCategoryReq
		if err := c.ShouldBind(&req); err != nil {
			// 参数校验
			logging.LogrusObj.Infoln(err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		l := service.GetCategorySrv()
		resp, err := l.CategoryList(c.Request.Context(), &req)
		if err != nil {
			logging.LogrusObj.Infoln(err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}
