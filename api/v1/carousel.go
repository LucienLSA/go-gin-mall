package v1

import (
	"net/http"

	"github.com/LucienLSA/go-gin-mall/pkg/util/ctl"
	"github.com/LucienLSA/go-gin-mall/pkg/util/logging"
	"github.com/LucienLSA/go-gin-mall/service"
	"github.com/LucienLSA/go-gin-mall/types"
	"github.com/gin-gonic/gin"
)

// 显示轮播图
func ListCarouselsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.ListCarouselReq
		if err := c.ShouldBind(&req); err != nil {
			// 参数校验
			logging.LogrusObj.Infoln(err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		l := service.GetCarouselSrv()
		resp, err := l.ListCarousel(c.Request.Context(), &req)
		if err != nil {
			logging.LogrusObj.Infoln(err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}
