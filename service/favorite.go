package service

import (
	"context"
	"demoProject4mall/dao"
	"demoProject4mall/model"
	"demoProject4mall/pkg/e"
	"demoProject4mall/pkg/util"
	"demoProject4mall/serializer"
	"strconv"
)

type FavoriteService struct {
	ProductId  uint `json:"product_id" form:"product_id"`
	BossId     uint `json:"boss_id" form:"boss_id"`
	FavoriteId uint `json:"favorite_id" form:"favorite_id"`
	model.BasePage
}

func (service *FavoriteService) List(c context.Context, uId uint) serializer.Response {
	favoriteDao := dao.NewFavoriteDao(c)
	code := e.Success
	favorite, err := favoriteDao.ListFavorite(uId)
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "类型搜索错误",
		}
	}
	return serializer.BuildListResponse(serializer.BuildFavorites(c, favorite), uint(len(favorite)))
}

func (service *FavoriteService) Create(c context.Context, uId uint) serializer.Response {
	favoriteDao := dao.NewFavoriteDao(c)
	code := e.Success
	exist, _ := favoriteDao.FavoriteExistOrNot(service.ProductId, uId)
	if exist {
		code = e.ErrorFavoriteExist
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	userDao := dao.NewUserDao(c)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	bossDao := dao.NewUserDao(c)
	boss, err := bossDao.GetUserById(service.BossId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	productDao := dao.NewProductDao(c)
	product, err := productDao.GetProductById(service.ProductId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	favorite := &model.Favorite{
		User:      *user,
		UserID:    uId,
		Product:   *product,
		ProductID: service.ProductId,
		Boss:      *boss,
		BossID:    service.BossId,
	}
	err = favoriteDao.CreateFavorite(favorite)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *FavoriteService) Delete(c context.Context, uId uint, fId string) serializer.Response {
	favoriteDao := dao.NewFavoriteDao(c)
	favoriteId, _ := strconv.Atoi(fId)
	code := e.Success
	err := favoriteDao.DeleteFavorit(uId, uint(favoriteId))
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "类型搜索错误",
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
