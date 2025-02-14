package v1

import (
	"demoProject4mall/pkg/util"
	"demoProject4mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCart(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	createCartService := service.CartService{}
	if err := c.ShouldBind(&createCartService); err == nil {
		res := createCartService.Create(c.Request.Context(), claim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func UpdateCart(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	updateCartService := service.CartService{}
	if err := c.ShouldBind(&updateCartService); err == nil {
		res := updateCartService.Update(c.Request.Context(), claim.ID, c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func DeleteCart(c *gin.Context) {
	deleteCartService := service.CartService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteCartService); err == nil {
		res := deleteCartService.Delete(c.Request.Context(), claim.ID, c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}

}

func ListCart(c *gin.Context) {
	listCartService := service.CartService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listCartService); err == nil {
		res := listCartService.List(c.Request.Context(), claim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}

}
