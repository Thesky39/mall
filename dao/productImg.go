package dao

import (
	"context"
	"demoProject4mall/model"
	"gorm.io/gorm"
)

type ProductImgDao struct {
	*gorm.DB
}

func NewProductImgDao(c context.Context) *ProductImgDao {
	return &ProductImgDao{NewDBClient(c)}
}

func (dao *ProductImgDao) CreateProductImg(productImg *model.ProductImg) error {
	return dao.DB.Model(&model.ProductImg{}).Create(&productImg).Error
}

func NewProductImgDapByDB(db *gorm.DB) *ProductImgDao {
	return &ProductImgDao{db}
}

func (dao *ProductImgDao) ListProductImg(id uint) (productImg []*model.ProductImg, err error) {
	err = dao.DB.Model(&model.ProductImg{}).Where("product_id = ?", id).Find(&productImg).Error
	return
}
