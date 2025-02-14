package serializer

import (
	"context"
	"demoProject4mall/conf"
	"demoProject4mall/dao"
	"demoProject4mall/model"
)

type Favorite struct {
	UserId        uint   `json:"user_id"`
	ProductId     uint   `json:"product_id"`
	CreatedAt     int64  `json:"created_at"`
	Name          string `json:"name"`
	CreatedId     uint   `json:"created_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	BossId        uint   `json:"boss_id"`
	Num           int    `json:"num"`
	OnSale        bool   `json:"on_sale"`
}

func BuildFavorite(favorite *model.Favorite, product *model.Product, boss *model.User) Favorite {
	return Favorite{
		UserId:        favorite.UserID,
		ProductId:     favorite.ProductID,
		CreatedAt:     favorite.CreatedAt.Unix(),
		Name:          product.Name,
		CreatedId:     product.CategoryId,
		Title:         product.Title,
		Info:          product.Info,
		ImgPath:       conf.Host + conf.HttpPort + conf.ProductPath + product.ImgPath,
		Price:         product.Price,
		DiscountPrice: product.DiscountPrice,
		BossId:        boss.ID,
		Num:           product.Num,
		OnSale:        product.OnSale,
	}
}

func BuildFavorites(c context.Context, items []*model.Favorite) (favorites []Favorite) {
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
		favorite := BuildFavorite(item, product, boss)
		favorites = append(favorites, favorite)
	}
	return favorites
}
