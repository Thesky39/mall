package serializer

import (
	"demoProject4mall/conf"
	"demoProject4mall/model"
)

type ProductImg struct {
	ProductId uint   `json:"product_id"`
	ImgPath   string `json:"img_path"`
}

func BuildProductImg(item *model.ProductImg) ProductImg {
	return ProductImg{
		ProductId: item.ProductID,
		ImgPath:   conf.Host + conf.HttpPort + conf.ProductPath + item.ImgPath,
	}
}

func BuildProductImgs(items []*model.ProductImg) (ProductImg []ProductImg) {
	for _, item := range items {
		product := BuildProductImg(item)
		ProductImg = append(ProductImg, product)
	}
	return
}
