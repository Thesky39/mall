package v1

import (
	"demoProject4mall/pkg/util"
	"demoProject4mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateAddress(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	createAddressService := service.AddressService{}
	if err := c.ShouldBind(&createAddressService); err == nil {
		res := createAddressService.Create(c.Request.Context(), claim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func UpdateAddress(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	updateAddressService := service.AddressService{}
	if err := c.ShouldBind(&updateAddressService); err == nil {
		res := updateAddressService.Update(c.Request.Context(), claim.ID, c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// 分页查询商品
func ListAddress(c *gin.Context) {
	listAddressService := service.AddressService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listAddressService); err == nil {
		res := listAddressService.List(c.Request.Context(), claim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}

}

func DeleteAddress(c *gin.Context) {
	deleteAddressService := service.AddressService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteAddressService); err == nil {
		res := deleteAddressService.Delete(c.Request.Context(), claim.ID, c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}

}

func ShowAddress(c *gin.Context) {
	listAddressService := service.AddressService{}
	if err := c.ShouldBind(&listAddressService); err == nil {
		res := listAddressService.Show(c.Request.Context(), c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}

}
