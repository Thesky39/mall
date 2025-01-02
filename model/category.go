package model

import "gorm.io/gorm"

// 商品分类
type Category struct {
	gorm.Model
	CategoryName string //商品类型名称
}
