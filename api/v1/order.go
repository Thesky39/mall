package v1

import (
	"demoProject4mall/pkg/util"
	"demoProject4mall/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOrder(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	createOrderService := service.OrderService{}
	fmt.Println(createOrderService.ProductId, "---------------")
	if err := c.ShouldBind(&createOrderService); err == nil {
		res := createOrderService.Create(c.Request.Context(), claim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func ShowOrder(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	showOrderService := service.OrderService{}
	if err := c.ShouldBind(&showOrderService); err == nil {
		res := showOrderService.Show(c.Request.Context(), claim.ID, c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func DeleteOrder(c *gin.Context) {
	deleteOrderService := service.OrderService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteOrderService); err == nil {
		res := deleteOrderService.Delete(c.Request.Context(), claim.ID, c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}

}

func ListOrder(c *gin.Context) {
	listOrderService := service.OrderService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listOrderService); err == nil {
		res := listOrderService.List(c.Request.Context(), claim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}

}
