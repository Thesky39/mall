package model

import "gorm.io/gorm"

// 购物车模型
type Cart struct {
	gorm.Model
	UserID    uint `gorm:"not null"` //用户id
	ProductID uint `gorm:"not null"` //商品id
	BossId    uint `gorm:"not null"` //商家id
	Num       int  `gorm:"not null"` //数量
	MaxNum    int  `gorm:"not null"` //购买限制
	Check     bool `gorm:"not null"` //是否支付
}
