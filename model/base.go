package model

// 翻页
type BasePage struct {
	PageNum  int `form:"pageNum"`
	PageSize int `form:"pageSize"`
}
