package v1

import (
	"demoProject4mall/pkg/util"
	"demoProject4mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 创建商品
func CreateProduct(c *gin.Context) {
	from, _ := c.MultipartForm()
	files := from.File["file"]
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	createProuctService := service.ProductService{}
	if err := c.ShouldBind(&createProuctService); err == nil {
		res := createProuctService.Create(c.Request.Context(), claim.ID, files)
		c.JSON(200, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func SearchProduct(c *gin.Context) {
	searchProductService := service.ProductService{}
	if err := c.ShouldBind(&searchProductService); err == nil {
		res := searchProductService.Search(c.Request.Context())
		c.JSON(200, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}

}

// 分页查询商品
func ListProduct(c *gin.Context) {
	listProductService := service.ProductService{}
	if err := c.ShouldBind(&listProductService); err == nil {
		res := listProductService.List(c.Request.Context())
		c.JSON(200, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}

}

// 获取商品详细信息
func ShowProduct(c *gin.Context) {
	listProductService := service.ProductService{}
	if err := c.ShouldBind(&listProductService); err == nil {
		res := listProductService.Show(c.Request.Context(), c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}

}
