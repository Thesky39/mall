package service

import (
	"context"
	"demoProject4mall/dao"
	"demoProject4mall/model"
	"demoProject4mall/pkg/e"
	"demoProject4mall/serializer"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type OrderService struct {
	ProductId uint    `json:"product_id" form:"product_id"`
	Num       int     `json:"num" form:"num"`
	AddressId uint    `json:"address_id" form:"address_id"`
	Money     float64 `json:"money" form:"money"`
	BossId    uint    `json:"boss_id" form:"boss_id"`
	UserId    uint    `json:"user_id" form:"user_id"`
	OrderNum  uint    `json:"order_num" form:"order_num"`
	Type      int     `json:"type" form:"type"`
	model.BasePage
}

func (service *OrderService) Create(c context.Context, uId uint) serializer.Response {
	var order *model.Order
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
	orderDao := dao.NewOrderDao(c)
	order = &model.Order{
		UserId:    uId,
		ProductId: service.ProductId,
		BossId:    service.BossId,
		Num:       service.Num,
		Money:     service.Money,
		Type:      1, //未支付
	}

	addressDao := dao.NewAddressDao(c)
	address, err := addressDao.GetAddressByAid(service.AddressId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	order.AddressID = address.ID
	// 订单好创建，自动生成随机number+唯一表示的product id+用户的id
	number := fmt.Sprintf("%09v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	productNum := strconv.Itoa(int(service.ProductId))
	userNum := strconv.Itoa(int(service.UserId))
	number = number + productNum + userNum
	orderNum, _ := strconv.ParseInt(number, 10, 64)
	order.OrderNum = uint64(orderNum)
	err = orderDao.CreateOrder(order)
	if err != nil {

		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	//userDao := dao.NewUserDao(c)
	//boss, err := userDao.GetUserById(service.BossId)
	//if err != nil {
	//	code = e.Error
	//	return serializer.Response{
	//		Status: code,
	//		Msg:    e.GetMsg(code),
	//		Error:  err.Error(),
	//	}
	//}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildOrder(order, product, address),
	}
}

func (service *OrderService) Show(c context.Context, uId uint, oId string) serializer.Response {
	orderId, _ := strconv.Atoi(oId)
	code := e.Success
	orderDao := dao.NewOrderDao(c)
	orders, err := orderDao.GetOrderByAid(uint(orderId), uId)
	if err != nil {

		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	addressDao := dao.NewAddressDao(c)
	address, err := addressDao.GetAddressByAid(service.AddressId)
	if err != nil {

		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
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
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildOrder(orders, product, address),
	}
}

func (service *OrderService) List(c context.Context, uId uint) serializer.Response {
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}

	orderDao := dao.NewOrderDao(c)
	condition := make(map[string]interface{})
	if service.Type != 0 {
		condition["type"] = service.Type
	}
	condition["user_id"] = uId
	orderList, err := orderDao.ListOrderByCondition(condition, service.BasePage)
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
		Data:   serializer.BuildOrders(c, orderList),
	}
}

func (service *OrderService) Delete(c context.Context, uId uint, cId string) serializer.Response {
	orderId, _ := strconv.Atoi(cId)
	code := e.Success
	orderDao := dao.NewOrderDao(c)
	err := orderDao.DeleteOrderByOrderId(uint(orderId), uId)
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
