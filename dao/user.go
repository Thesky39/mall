package dao

import (
	"context"
	"demoProject4mall/model"
	"fmt"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}
func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

// 根据username判断是否存在该名字
func (dao *UserDao) ExistOrNotByUserName(userName string) (user *model.User, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.User{}).Where("user_name = ?", userName).Find(&user).Count(&count).Error

	//if user == nil || err == gorm.ErrRecordNotFound {
	//	return nil, false, err
	//}
	if count == 0 {
		return nil, false, nil
	}
	return user, true, nil

	//err = dao.DB.Model(&model.User{}).Where("username = ?", userName).Count(&count).Error
	//if err != nil {
	//	return nil, false, err
	//}
	//
	//// 如果记录数为 0，返回不存在
	//if count == 0 {
	//	return nil, false, nil
	//}
	//
	//// 使用 First 查找用户的详细信息
	//user = &model.User{}
	//err = dao.DB.Model(&model.User{}).Where("username = ?", userName).First(&user).Error
	//if err != nil {
	//	return nil, false, err
	//}
	//return user, true, nil
}
func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.DB.Model(&model.User{}).Create(&user).Error
}

// GetUserById 根据id获取user
func (dao *UserDao) GetUserById(id uint) (user *model.User, err error) {
	//err = dao.DB.Model(&model.User{}).Where("id=?", id).First(&user).Error
	//return
	err = dao.DB.Model(&model.User{}).Where("id=?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user with id %d not found", id) // 如果没有找到用户，返回明确的错误信息
		}
		return nil, err // 如果是其他错误，直接返回
	}
	return user, nil
}

// UpdateUserById 通过id更新user信息
func (dao *UserDao) UpdateUserById(id uint, user *model.User) error {
	err := dao.DB.Model(&model.User{}).Where("id=?", id).Updates(&user).Error
	return err
}
