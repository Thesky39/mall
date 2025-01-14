package service

import (
	"demoProject4mall/conf"
	"demoProject4mall/dao"
	"demoProject4mall/model"
	"demoProject4mall/pkg/e"
	"demoProject4mall/pkg/util"
	"demoProject4mall/serializer"
	"fmt"
	"golang.org/x/net/context"
	"gopkg.in/mail.v2"
	"mime/multipart"
	"strings"
	"time"
)

// 数据传输对象
type UserService struct {
	NickName string `json:"nick_Name" form:"nick_name"`
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
	Key      string `json:"key" form:"key"` //前端验证
}

type SendEmailService struct {
	Email         string `json:"email" form:"email"`
	Password      string `json:"password" form:"password"`
	OperationType uint   `json:"operation_type" form:"operation_type"` //1.绑定邮箱 2.解绑邮箱 3.改密码

}
type ValidEmailService struct {
	Email         string `json:"email" form:"email"`
	Password      string `json:"password" form:"password"`
	OperationType uint   `json:"operation_type" form:"operation_type"` //1.绑定邮箱 2.解绑邮箱 3.改密码
}
type ShowMoneyService struct {
	Key string `json:"key" form:"key"`
}

func (service *UserService) Register(c context.Context) serializer.Response {
	var user model.User
	code := e.Success
	if service.Key == "" || len(service.Password) <= 8 {
		code = e.Error
		fmt.Println(len(service.Password))
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "密码长度不足",
		}
	}
	//密文存储
	util.Encrypt.SetKey(service.Key)

	userDao := dao.NewUserDao(c)
	_, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if exist {
		code = e.ErrorExistUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user = model.User{
		UserName: service.UserName,

		NickName: "",
		Status:   model.Active,
		Avatar:   "avatar.JPG",
		Money:    util.Encrypt.AesEncoding("10000"),
	}
	//密码加密
	if err = user.SetPassword(service.Password); err != nil {
		code = e.ErrorFailEncryption
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//创建用户
	err = userDao.CreateUser(&user)
	if err != nil {
		code = e.Error

	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// 用户登录
func (service *UserService) Login(c context.Context) serializer.Response {
	var user *model.User
	code := e.Success
	userDao := dao.NewUserDao(c)
	//判断用户是否存在
	user, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if !exist || err != nil {
		code = e.ErrorExistUserNotFound
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "用户不存在",
		}
	}
	//校验密码
	if user.CheckPassword(service.Password) == false {
		code = e.ErrorNotcompare
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "密码错误，请重新登录",
		}
	}
	//http 无状态（认证，token） 签发token
	token, err := util.GenerateToken(user.ID, service.UserName, 0)
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "token 认证失败",
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    e.GetMsg(code),
	}
}

// 用户修改信息
func (service *UserService) Update(c context.Context, uId uint) serializer.Response {
	var user *model.User
	var err error
	code := e.Success
	//找到这个用户
	userDao := dao.NewUserDao(c)
	user, err = userDao.GetUserById(uId)

	if err != nil {
		// 如果是用户未找到的错误
		if err.Error() == fmt.Sprintf("user with id %d not found", uId) {
			code = e.ErrorExistUserNotFound // 假设你在 `e` 中定义了这个错误码
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}

		// 其他错误
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	//修改昵称
	if service.NickName != "" {
		user.NickName = service.NickName
	}
	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}

	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}

// post 头像更新
func (service *UserService) Post(c context.Context, uId uint, file multipart.File, fileSize int64) serializer.Response {
	code := e.Success
	var user *model.User
	var err error
	userDao := dao.NewUserDao(c)
	user, err = userDao.GetUserById(uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//保存图片到本地
	path, err := UploadAvatarToLocalStatic(file, uId, user.UserName)
	if err != nil {
		code = e.ErrorUploadfail
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	user.Avatar = path
	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}

// 发送邮箱
func (service *SendEmailService) Send(c context.Context, uId uint) serializer.Response {
	code := e.Success
	var address string
	var notice *model.Notice //绑定邮箱的模板
	token, err := util.GenerateEmailToken(uId, service.OperationType, service.Email, service.Password)
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	noticeDao := dao.NewNoticeDao(c)
	notice, err = noticeDao.GetNoticeById(uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	address = conf.ValidEmail + token //发送方
	mailStr := notice.Text
	mailTex := strings.Replace(mailStr, "Email", address, -1)
	m := mail.NewMessage()
	m.SetHeader("From", conf.SmtpEmail)
	m.SetHeader("To", service.Email)
	m.SetHeader("Subject", "FanOne")
	m.SetBody("text/html", mailTex)
	d := mail.NewDialer(conf.SmtpHost, 465, conf.SmtpEmail, conf.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err = d.DialAndSend(m); err != nil {
		code = e.ErrorSendEmail
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// 验证邮箱
func (service *ValidEmailService) Valid(c context.Context, token string) serializer.Response {
	var userId uint
	var email string
	var password string
	var operationType uint
	code := e.Success
	fmt.Println("hahahahah", token)
	if token == "" {
		code = e.InvalidParams
	} else {
		claims, err := util.ParseEmailToken(token)
		fmt.Println("claims", claims)
		if err != nil {
			code = e.ErrorAuthToken
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ErrorAuthCheckTokenTimeout
		} else {
			userId = claims.UserID
			email = claims.Email
			password = claims.Password
			operationType = claims.OperationType
		}
	}
	if code != e.Success {
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//换取该用户的信息
	userDao := dao.NewUserDao(c)
	user, err := userDao.GetUserById(userId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if operationType == 1 {
		//绑定邮箱
		fmt.Println(email)
		user.Email = email
	} else if operationType == 2 {
		//解绑邮箱
		user.Email = ""
	} else if operationType == 3 {
		err = user.SetPassword(password)
		if err != nil {
			code = e.Error
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
	}
	err = userDao.UpdateUserById(userId, user)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}

// 展示用户金额
func (service *ShowMoneyService) ShowMony(c context.Context, uId uint) serializer.Response {
	code := e.Success
	userDao := dao.NewUserDao(c)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   serializer.BuildUser(user),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildMoney(user, service.Key),
		Msg:    e.GetMsg(code),
	}
}
