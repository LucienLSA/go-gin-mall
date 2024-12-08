package v1

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/LucienLSA/go-gin-mall/conf"
	"github.com/LucienLSA/go-gin-mall/pkg/e"
	"github.com/LucienLSA/go-gin-mall/pkg/util/ctl"
	"github.com/LucienLSA/go-gin-mall/pkg/util/qrcode"
	"github.com/boombuler/barcode/qr"
	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
)

func ErrorResponse(ctx *gin.Context, err error) *ctl.TrackedErrorResponse {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range ve {
			field := conf.T(fmt.Sprintf("Field.%s", fieldError.Field))
			tag := conf.T(fmt.Sprintf("Tag.Valid.%s", fieldError.Tag))
			return ctl.RespError(ctx, err, fmt.Sprintf("%s%s", field, tag))
		}
	}

	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return ctl.RespError(ctx, err, "JSON类型不匹配")
	}

	return ctl.RespError(ctx, err, err.Error(), e.ERROR)
}

const (
	QRCODE_URL = "https://lucienlsa.github.io/docsify-blog.github.io/#/README"
)

// 生成二维码
func GenerateQrcode(ctx *gin.Context) {
	qr := qrcode.NewQrCode(QRCODE_URL, 300, 300, qr.M, qr.Auto)
	path := qrcode.GetQrCodePath()
	fmt.Println(path)
	name := qrcode.GetQrCodeFileName(qr.URL) + qr.GetQrCodeExt()
	_, _, err := qr.Encode(path)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, map[string]string{
		"poster_url":      qrcode.GetQrCodeFullUrl(name),
		"poster_save_url": path + name,
	}))
}
