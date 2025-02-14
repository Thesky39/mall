package model

import (
	"demoProject4mall/cache"
	"gorm.io/gorm"
	"strconv"
)

type Product struct {
	gorm.Model
	Name          string `json:"name"`
	CategoryId    uint   `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	Num           int    `json:"num"`
	OnSale        bool   `json:"on_sale" gorm:"default:false"`
	BossID        uint   `json:"boss_id"`
	BossName      string `json:"boss_name"`
	BossAvatar    string `json:"boss_avatar"`
}

func (product *Product) View() uint64 {
	countStr, _ := cache.RedisClint.Get(cache.ProductViewKey(product.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}
func (product *Product) AddView() {
	//增加商品点击数
	cache.RedisClint.Incr(cache.ProductViewKey(product.ID))
	cache.RedisClint.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(product.ID)))

}
