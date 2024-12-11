package types

type ListProductImgReq struct {
	ID uint `json:"id" form:"id"`
}

type ProductImgResp struct {
	ProductID uint   `json:"product_id" form:"product_id"`
	ImgPath   string `json:"img_path" form:"img_path"`
}
