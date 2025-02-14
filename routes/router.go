package routes

import (
	api "demoProject4mall/api/v1"
	"demoProject4mall/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")

		})
		//用户操作
		v1.POST("/user/register", api.UserRegister)
		v1.POST("/user/login", api.UserLogin)

		//轮播图
		v1.GET("carousels", api.ListCarousel)

		//商品操作
		v1.GET("products", api.ListProduct)
		v1.GET("products/:id", api.ShowProduct)
		v1.GET("imgs/:id", api.ListProductImg)
		v1.GET("categories", api.ListCategory) //商品分类

		authed := v1.Group("/") //需要登录保护
		authed.Use(middleware.JWT())
		{
			//用户操作
			authed.PUT("user", api.UserUpdate)
			authed.POST("avatar", api.UploadAvatar)
			authed.POST("/user/sending-email", api.SendEmail)
			authed.POST("/user/valid-email", api.ValidEmail)

			//显示金额
			authed.POST("/money", api.ShowMoney)

			//商品操作
			authed.POST("/product", api.CreateProduct)
			authed.POST("/products", api.SearchProduct)

			//收藏夹操作
			authed.GET("favorites", api.ListFavorite)
			authed.POST("favorites", api.CreateFavorite)
			authed.DELETE("favorites/:id", api.DeleteFavorite)
			// 地址模块
			authed.POST("addresses", api.CreateAddress)
			authed.GET("addresses/:id", api.ShowAddress)
			authed.GET("addresses", api.ListAddress)
			authed.PUT("addresses", api.UpdateAddress)
			authed.DELETE("addresses/:id", api.DeleteAddress)

			// 购物车操作
			authed.POST("carts", api.CreateCart)
			authed.GET("carts/:id", api.ListCart)
			authed.PUT("carts/:id", api.UpdateCart)
			authed.DELETE("carts/:id", api.DeleteCart)
			// 订单操作
			authed.POST("orders", api.CreateOrder)
			authed.GET("order", api.ListOrder)
			authed.PUT("order/:id", api.ShowOrder)
			authed.DELETE("order/:id", api.DeleteOrder)
		}
	}
	return r
}
