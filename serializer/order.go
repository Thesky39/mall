package serializer

import (
	"context"
	"demoProject4mall/conf"
	"demoProject4mall/dao"
	"demoProject4mall/model"
)

type Order struct {
	Id           uint    `json:"id"`
	OrderNum     uint64  `json:"order_num"`
	CreateAt     int64   `json:"create_at"`
	UpdateAt     int64   `json:"update_at"`
	UserId       uint    `json:"user_id"`
	ProductId    uint    `json:"product_id"`
	BossId       uint    `json:"boss_id"`
	Num          int     `json:"num"`
	AddressName  string  `json:"address_name"`
	AddressPhone string  `json:"address_phone"`
	Address      string  `json:"address"`
	Type         uint    `json:"type"`
	ProductName  string  `json:"product_name"`
	ImgPath      string  `json:"img_path"`
	Money        float64 `json:"discount_price"`
}

func BuildOrder(order *model.Order, product *model.Product, address *model.Address) Order {
	return Order{
		Id:           order.ID,
		OrderNum:     order.OrderNum,
		CreateAt:     order.CreatedAt.Unix(),
		UpdateAt:     order.UpdatedAt.Unix(),
		UserId:       order.UserId,
		ProductId:    product.ID,
		BossId:       order.ID,
		Num:          order.Num,
		AddressName:  address.Name,
		AddressPhone: address.Phone,
		Address:      address.Address,
		Type:         order.Type,
		ProductName:  product.Name,
		ImgPath:      conf.Host + conf.HttpPort + conf.ProductPath + product.ImgPath,
		Money:        order.Money,
	}
}

func BuildOrders(c context.Context, items []*model.Order) (orders []Order) {
	productDao := dao.NewProductDao(c)
	addressDao := dao.NewAddressDao(c)
	for _, item := range items {
		product, err := productDao.GetProductById(item.ProductId)
		if err != nil {
			continue
		}
		address, err := addressDao.GetAddressByAid(item.AddressID)
		if err != nil {
			continue
		}
		order := BuildOrder(item, product, address)
		orders = append(orders, order)
	}
	return orders
}
