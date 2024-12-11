package types

type ListCategoryReq struct{}  // 请求参数结构体
type ListCategoryResp struct { // 返回参数结构体
	ID           uint   `json:"id"`
	CategoryName string `json:"category_name"`
	CreatedAt    int64  `json:"created_at"`
}
