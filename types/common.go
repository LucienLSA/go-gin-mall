package types

type BasePage struct {
	PageNum  int `form:"page_num"`  // 页码
	PageSize int `form:"page_size"` // 每页数量
}

// 带总数的Data结构
type DataListResp struct {
	Item  interface{} `json:"item"`
	Total int64       `json:"total"`
}
