package service

import (
	"context"
	"demoProject4mall/dao"
	"demoProject4mall/serializer"
	"strconv"
)

type ListProductImg struct {
}

func (service *ListProductImg) List(c context.Context, pId string) serializer.Response {
	productImgDao := dao.NewProductImgDao(c)
	productId, _ := strconv.Atoi(pId)
	prouctImgs, _ := productImgDao.ListProductImg(uint(productId))
	return serializer.BuildListResponse(serializer.BuildProductImgs(prouctImgs), uint(len(prouctImgs)))
}
