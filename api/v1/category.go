package v1

import (
	"demoProject4mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListCategory(c *gin.Context) {
	var listCategory service.CategoryService
	if err := c.ShouldBind(&listCategory); err == nil {
		res := listCategory.List(c.Request.Context())
		c.JSON(200, res)
	} else {

		c.JSON(http.StatusBadRequest, err)
	}
}
