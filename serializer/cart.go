package serializer

import (
	"context"
	"demoProject4mall/conf"
	"demoProject4mall/dao"
	"demoProject4mall/model"
)

type Cart struct {
	Id             uint   `json:"id"`
	UserId         uint   `json:"user_id"`
	ProductId      uint   `json:"product_id"`
	CreatedAt      int64  `json:"created_at"`
	Num            int    `json:"num"`
	Name           string `json:"name"`
	MaxNum         int    `json:"max_num"`
	ImgPath        string `json:"img_path"`
	Check          bool   `json:"check"`
	DiscountParice string `json:"discount_parice"`
	BossId         uint   `json:"boss_id"`
	BossName       string `json:"boss_name"`
}

func BuildCart(cart *model.Cart, product *model.Product, boss *model.User) Cart {
	return Cart{
		Id:             cart.ID,
		UserId:         cart.UserID,
		ProductId:      cart.ProductID,
		CreatedAt:      cart.CreatedAt.Unix(),
		Num:            int(cart.Num),
		MaxNum:         int(cart.MaxNum),
		Check:          cart.Check,
		Name:           product.Name,
		ImgPath:        conf.Host + conf.HttpPort + conf.ProductPath + product.ImgPath,
		DiscountParice: product.DiscountPrice,
		BossId:         boss.ID,
		BossName:       boss.UserName,
	}
}

func BuildCarts(c context.Context, items []*model.Cart) (carts []Cart) {
	productDao := dao.NewProductDao(c)
	bossDao := dao.NewUserDao(c)
	for _, item := range items {
		product, err := productDao.GetProductById(item.ProductID)
		if err != nil {
			continue
		}
		boss, err := bossDao.GetUserById(item.UserID)
		if err != nil {
			continue
		}
		favorite := BuildCart(item, product, boss)
		carts = append(carts, favorite)
	}
	return carts
}
