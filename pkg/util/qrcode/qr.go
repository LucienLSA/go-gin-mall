package qrcode

import (
	"image/jpeg"

	"github.com/LucienLSA/go-gin-mall/conf"
	"github.com/LucienLSA/go-gin-mall/pkg/util/encryption"
	"github.com/LucienLSA/go-gin-mall/pkg/util/fileee"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

type QrCode struct {
	URL    string
	Width  int
	Height int
	Ext    string
	Level  qr.ErrorCorrectionLevel
	Mode   qr.Encoding
}

const (
	EXT_JPG = ".jpg"
)

// 生成一个二维码
func NewQrCode(url string, width, height int, level qr.ErrorCorrectionLevel, mode qr.Encoding) *QrCode {
	return &QrCode{
		URL:    url,
		Width:  width,
		Height: height,
		Level:  level,
		Mode:   mode,
		Ext:    EXT_JPG,
	}
}

// 获取二维码的路径
func GetQrCodePath() string {
	return conf.Config.PhotoPath.QrcodePath
}

// 获取二维码的所有url
func GetQrCodeFullUrl(name string) string {
	return conf.Config.PhotoPath.PhotoHost + conf.Config.System.HttpPort + GetQrCodePath() + name
}

// 获取二维码的文件名
func GetQrCodeFileName(value string) string {
	return encryption.EncodeMD5(value)
}

// 获取二维码的后缀名
func (q *QrCode) GetQrCodeExt() string {
	return q.Ext
}

func (q *QrCode) Encode(path string) (string, string, error) {
	name := GetQrCodeFileName(q.URL) + q.GetQrCodeExt()
	src := path + name
	if fileee.CheckNotExist(src) {
		code, err := qr.Encode(q.URL, q.Level, q.Mode)
		if err != nil {
			return "", "", err
		}

		code, err = barcode.Scale(code, q.Width, q.Height)
		if err != nil {
			return "", "", err
		}

		f, err := fileee.MustOpen(name, path)
		if err != nil {
			return "", "", err
		}
		defer f.Close()

		err = jpeg.Encode(f, code, nil)
		if err != nil {
			return "", "", err
		}
	}

	return name, path, nil
}
