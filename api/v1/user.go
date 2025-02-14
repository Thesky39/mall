package v1

import (
	"demoProject4mall/pkg/util"
	"demoProject4mall/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegister(c *gin.Context) {
	var userReister service.UserService
	if err := c.ShouldBind(&userReister); err == nil {
		res := userReister.Register(c.Request.Context())
		c.JSON(200, res)
	} else {
		fmt.Println(111111)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}
func UserLogin(c *gin.Context) {
	var userLogin service.UserService
	if err := c.ShouldBind(&userLogin); err == nil {
		res := userLogin.Login(c.Request.Context())
		c.JSON(200, res)
	} else {
		//fmt.Println(111111)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func UserUpdate(c *gin.Context) {
	var UserUpdate service.UserService
	fmt.Println("hhhh")

	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	fmt.Println("-----------", claims.ID)
	if err := c.ShouldBind(&UserUpdate); err == nil {
		res := UserUpdate.Update(c.Request.Context(), claims.ID)
		c.JSON(200, res)
	} else {
		fmt.Println("----------------12222")
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}

	//var userUpdate service.UserService
	//
	//// 解析 Token 获取用户信息
	//claims, err := util.ParseToken(c.GetHeader("Authorization"))
	//if err != nil || claims == nil {
	//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
	//	return
	//}
	//
	//// 调试日志，确认 claims.ID 是否正常
	//fmt.Println("-----------", claims.ID)
	//
	//// 绑定请求中的参数到 UserService 结构体
	//if err := c.ShouldBind(&userUpdate); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	//
	//// 调用 Update 方法进行更新
	//res := userUpdate.Update(c.Request.Context(), claims.ID)
	//c.JSON(200, res)
}
func UploadAvatar(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	var uploadAvatar service.UserService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&uploadAvatar); err == nil {
		res := uploadAvatar.Post(c.Request.Context(), claims.ID, file, fileSize)
		c.JSON(200, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}
func SendEmail(c *gin.Context) {
	var sendEmail service.SendEmailService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&sendEmail); err == nil {
		res := sendEmail.Send(c.Request.Context(), claims.ID)
		c.JSON(200, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// ValidEmail
func ValidEmail(c *gin.Context) {
	var validEmail service.ValidEmailService
	fmt.Println("111111")
	if err := c.ShouldBind(&validEmail); err == nil {
		res := validEmail.Valid(c.Request.Context(), c.GetHeader("Authorization"))
		c.JSON(200, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func ShowMoney(c *gin.Context) {
	var showMoney service.ShowMoneyService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showMoney); err == nil {
		res := showMoney.ShowMony(c.Request.Context(), claims.ID)
		c.JSON(200, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}
