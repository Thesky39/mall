package v1

import (
	"demoProject4mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListCarousel(c *gin.Context) {
	var listCarousel service.CarouselService
	if err := c.ShouldBind(&listCarousel); err == nil {
		res := listCarousel.List(c.Request.Context())
		c.JSON(200, res)
	} else {

		c.JSON(http.StatusBadRequest, err)
	}
}
