package service

import (
	"context"
	"demoProject4mall/dao"
	"demoProject4mall/pkg/e"
	"demoProject4mall/pkg/util"
	"demoProject4mall/serializer"
)

type CategoryService struct {
}

func (service *CategoryService) List(c context.Context) serializer.Response {
	categoryDao := dao.NewCategoryDao(c)
	code := e.Success
	category, err := categoryDao.ListCategory()
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "类型搜索错误",
		}
	}
	return serializer.BuildListResponse(serializer.BuildCategorys(category), uint(len(category)))
}
