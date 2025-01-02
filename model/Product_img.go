package model

import "gorm.io/gorm"

type ProductImg struct {
	gorm.Model
	Text string `gorm:"type:text"`
}
