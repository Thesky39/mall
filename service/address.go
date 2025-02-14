package service

import (
	"context"
	"demoProject4mall/dao"
	"demoProject4mall/model"
	"demoProject4mall/pkg/e"
	"demoProject4mall/serializer"
	"strconv"
)

type AddressService struct {
	Name    string `json:"name" form:"name"`
	Phone   string `json:"phone" form:"phone"`
	Address string `json:"address" form:"address"`
}

func (service *AddressService) Create(c context.Context, uId uint) serializer.Response {
	var address *model.Address
	code := e.Success
	addressDao := dao.NewAddressDao(c)
	address = &model.Address{
		UserID:  uId,
		Phone:   service.Phone,
		Address: service.Address,
		Name:    service.Name,
	}
	err := addressDao.CreateAddress(address)
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
	}
}

func (service *AddressService) Show(c context.Context, aId string) serializer.Response {
	addressId, _ := strconv.Atoi(aId)
	var address *model.Address
	code := e.Success
	addressDao := dao.NewAddressDao(c)
	address, err := addressDao.GetAddressByAid(uint(addressId))

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
		Data:   serializer.BuildAddress(address),
	}
}

func (service *AddressService) List(c context.Context, uId uint) serializer.Response {

	code := e.Success
	addressDao := dao.NewAddressDao(c)
	addressList, err := addressDao.ListAddressByUserId(uId)
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
		Data:   serializer.BuildAddresses(addressList),
	}
}

func (service *AddressService) Update(c context.Context, uId uint, aId string) serializer.Response {
	var address *model.Address
	code := e.Success
	addressDao := dao.NewAddressDao(c)
	address = &model.Address{
		UserID:  uId,
		Phone:   service.Phone,
		Address: service.Address,
		Name:    service.Name,
	}
	addressId, _ := strconv.Atoi(aId)
	err := addressDao.UpdateAddressByUserId(uint(addressId), address)
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
	}
}

func (service *AddressService) Delete(c context.Context, uId uint, aId string) serializer.Response {
	addressId, _ := strconv.Atoi(aId)

	code := e.Success
	addressDao := dao.NewAddressDao(c)
	err := addressDao.DeleteAddressByAddressId(uint(addressId), uId)

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
	}
}
