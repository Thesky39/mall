package service

import (
	"context"
	"demoProject4mall/dao"
	"demoProject4mall/pkg/e"
	"demoProject4mall/pkg/util"
	"demoProject4mall/serializer"
)

type CarouselService struct {
}

func (service *CarouselService) List(c context.Context) serializer.Response {
	carouselDao := dao.NewCarouselDao(c)
	code := e.Success
	carousels, err := carouselDao.ListCarousel()
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "轮播图获取错误",
		}
	}
	return serializer.BuildListResponse(serializer.BuildCarousels(carousels), uint(len(carousels)))
}
