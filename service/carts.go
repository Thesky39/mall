package service

import (
	"context"
	"demoProject4mall/dao"
	"demoProject4mall/model"
	"demoProject4mall/pkg/e"
	"demoProject4mall/serializer"
	"strconv"
)

type CartService struct {
	Id        uint `json:"id" form:"id"`
	BossId    uint `json:"boss_id" form:"boss_id"`
	ProductId uint `json:"product_id" form:"product_id"`
	Num       int  `json:"num" form:"num"`
}

func (service *CartService) Create(c context.Context, uId uint) serializer.Response {
	var cart *model.Cart
	code := e.Success
	//判断商品是否存在
	productDao := dao.NewProductDao(c)
	product, err := productDao.GetProductById(service.ProductId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	cartDao := dao.NewCartDao(c)
	cart = &model.Cart{
		UserID:    uId,
		ProductID: service.ProductId,
		BossId:    service.BossId,
		Num:       service.Num,
	}
	err = cartDao.CreateCart(cart)
	if err != nil {

		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	userDao := dao.NewUserDao(c)
	boss, err := userDao.GetUserById(service.BossId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildCart(cart, product, boss),
	}
}

func (service *CartService) List(c context.Context, uId uint) serializer.Response {
	code := e.Success
	cartDao := dao.NewCartDao(c)
	carts, err := cartDao.ListCartByUserId(uId)
	if err != nil {

		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildCarts(c, carts),
	}
}

func (service *CartService) Update(c context.Context, uId uint, cId string) serializer.Response {
	code := e.Success
	cartDao := dao.NewCartDao(c)
	cartId, _ := strconv.Atoi(cId)
	err := cartDao.UpdateCartById(uint(cartId), service.Num)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *CartService) Delete(c context.Context, uId uint, cId string) serializer.Response {
	cartId, _ := strconv.Atoi(cId)
	code := e.Success
	cartDao := dao.NewCartDao(c)
	err := cartDao.DeleteCartByCartId(uint(cartId), uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
