package model

import "gorm.io/gorm"

type ProductImg struct {
	gorm.Model
	ProductID uint   `json:"product_id" form:"product_id"`
	ImgPath   string `json:"img_path" form:"img_path"`
}
