package types

type ProductCreateReq struct {
	ID            uint   `form:"id" json:"id"`
	Name          string `form:"name" json:"name"`
	CategoryID    uint   `form:"category_id" json:"category_id"`
	Title         string `form:"title" json:"title" `
	Info          string `form:"info" json:"info" `
	ImgPath       string `form:"img_path" json:"img_path"`
	Price         string `form:"price" json:"price"`
	DiscountPrice string `form:"discount_price" json:"discount_price"`
	OnSale        bool   `form:"on_sale" json:"on_sale"` // 是否上架
	Num           int    `form:"num" json:"num"`
}
type ProductListReq struct {
	CategoryID uint `form:"category_id" json:"category_id"` // 分类ID
	BasePage
}
type ProductShowReq struct {
	ID uint `form:"id" json:"id"` // 商品ID
}

type ProductResp struct {
	ID            uint   `json:"id"` // 商品ID
	Name          string `json:"name"`
	CategoryID    uint   `json:"category_id"` // 分类ID
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`       // 图片路径
	Price         string `json:"price"`          // 原价
	DiscountPrice string `json:"discount_price"` // 打折价格
	View          uint64 `json:"view"`           // 浏览量
	CreatedAt     int64  `json:"created_at"`
	Num           int    `json:"num"`
	OnSale        bool   `json:"on_sale"`     // 是否上架
	BossID        uint   `json:"boss_id"`     // 商家ID
	BossName      string `json:"boss_name"`   // 商家名称
	BossAvatar    string `json:"boss_avatar"` // 商家头像
}

type ProductSearchReq struct {
	ID            uint   `form:"id" json:"id"`
	Name          string `form:"name" json:"name"`
	CategoryID    int    `form:"category_id" json:"category_id"`
	Title         string `form:"title" json:"title" `
	Info          string `form:"info" json:"info" `
	Price         string `form:"price" json:"price"`
	DiscountPrice string `form:"discount_price" json:"discount_price"`
	OnSale        bool   `form:"on_sale" json:"on_sale"`
	BasePage
}
