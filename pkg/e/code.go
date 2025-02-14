package e

const (
	Success       = 200
	Error         = 500
	InvalidParams = 400
	//user模块错误
	ErrorExistUser             = 30001
	ErrorFailEncryption        = 30002
	ErrorExistUserNotFound     = 30003
	ErrorNotcompare            = 30004
	ErrorAuthToken             = 30005
	ErrorAuthCheckTokenTimeout = 30006
	ErrorUploadfail            = 30007
	ErrorSendEmail             = 30008
	//商品模块错误
	ErrorProductImgUploadFail = 40001
	//收藏夹模块错误
	ErrorFavoriteExist = 50001
)
