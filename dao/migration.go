package dao

import "demoProject4mall/model"

func migration() {
	err := _db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&model.Address{},
		&model.Admin{},
		&model.BasePage{},
		&model.Category{},
		&model.Cart{},
		&model.Carousel{},
		&model.Favorite{},
		&model.Notice{},
		&model.Order{},
		&model.Product{},
		&model.ProductImg{},
		&model.User{})
	if err != nil {
		panic(err)

	}
	return
}
