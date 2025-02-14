package e

var MsgFlags = map[int]string{
	Success:                    "ok",
	Error:                      "fail",
	InvalidParams:              "参数错误",
	ErrorExistUser:             "用户已存在",
	ErrorFailEncryption:        "密码加密失败",
	ErrorExistUserNotFound:     "用户不存在",
	ErrorNotcompare:            "密码错误",
	ErrorAuthToken:             "token认证失败",
	ErrorAuthCheckTokenTimeout: "token过期",
	ErrorUploadfail:            "图片上传文件失败",
	ErrorSendEmail:             "邮件发送失败",
	ErrorProductImgUploadFail:  "图片上传错误",
	ErrorFavoriteExist:         "收藏夹不存在",
}

// 获取状态码对应信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[Error]
	}
	return msg
}
