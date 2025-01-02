package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	UserName       string `gorm:"type:varchar(20);not null"`
	PasswordDigest string `gorm:"type:varchar(20);not null"`
	Avatar         string `gorm:"type:varchar(255);not null"`
}
