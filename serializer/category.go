package serializer

import "demoProject4mall/model"

type Category struct {
	Id           uint   `json:"id"`
	CategoryName string `json:"category_name"`
	CreatedAt    int64  `json:"created_at"`
}

func BuildCategory(item *model.Category) Category {
	return Category{
		Id:           item.ID,
		CategoryName: item.CategoryName,
		CreatedAt:    item.CreatedAt.Unix(),
	}
}

func BuildCategorys(items []model.Category) (carousels []Category) {
	for _, item := range items {
		carousel := BuildCategory(&item)
		carousels = append(carousels, carousel)
	}
	return carousels
}
