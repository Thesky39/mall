package dao

import (
	"context"
	"demoProject4mall/model"
	//"fmt"
	"gorm.io/gorm"
)

type CarouselDao struct {
	*gorm.DB
}

func NewCarouselDao(ctx context.Context) *CarouselDao {
	return &CarouselDao{NewDBClient(ctx)}
}
func NewCarouselDaoByDB(db *gorm.DB) *CarouselDao {
	return &CarouselDao{db}
}

// GetCarouselById 根据id获取Carousel
func (dao *CarouselDao) ListCarousel() (carousel []model.Carousel, err error) {

	err = dao.DB.Model(&model.Carousel{}).Find(&carousel).Error
	return
}
