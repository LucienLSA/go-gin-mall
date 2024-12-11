package v1

import (
	"net/http"

	"github.com/LucienLSA/go-gin-mall/pkg/util/ctl"
	"github.com/LucienLSA/go-gin-mall/pkg/util/logging"
	"github.com/LucienLSA/go-gin-mall/service"
	"github.com/LucienLSA/go-gin-mall/types"
	"github.com/gin-gonic/gin"
)

// 新建收藏
func CreateFavoritesHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.FavoriteCreateReq
		if err := c.ShouldBindJSON(&req); err != nil {
			// 参数校验
			logging.LogrusObj.Infoln(err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		l := service.GetFavoriteSrv()
		resp, err := l.FavoriteCreate(c.Request.Context(), &req)
		if err != nil {
			logging.LogrusObj.Infoln(err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

// 删除收藏
func DeleteFavoritesHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req types.FavoriteDeleteReq
		if err := c.ShouldBindJSON(&req); err != nil {
			// 参数校验
			logging.LogrusObj.Infoln(err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		l := service.GetFavoriteSrv()
		resp, err := l.FavoriteDelete(c.Request.Context(), &req)
		if err != nil {
			logging.LogrusObj.Infoln(err)
			c.JSON(http.StatusOK, ErrorResponse(c, err))
			return
		}
		c.JSON(http.StatusOK, ctl.RespSuccess(c, resp))
	}
}

// // 列出收藏内容
// func ListFavoritesHandler() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var req types.FavoriteListReq
// 	}
// }
