package dao

import (
	"context"
	"demoProject4mall/model"
	"gorm.io/gorm"
)

type OrderDao struct {
	*gorm.DB
}

func NewOrderDao(ctx context.Context) *OrderDao {
	return &OrderDao{NewDBClient(ctx)}
}

func (dao *OrderDao) CreateOrder(in *model.Order) error {
	return dao.DB.Model(&model.Order{}).Create(in).Error
}

func (dao *OrderDao) GetOrderByAid(aId uint, userId uint) (order *model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Where("id = ? AND user_id=?", aId, userId).First(&order).Error
	return
}

func (dao *OrderDao) ListOrderByUserId(uId uint) (order []*model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Where("user_id = ?", uId).Find(&order).Error
	return
}

func (dao *OrderDao) UpdateOrderByUserId(aId uint, order *model.Order) error {
	return dao.DB.Model(&model.Order{}).Where("id = ?", aId).Updates(order).Error
}

func (dao *OrderDao) DeleteOrderByOrderId(aId uint, uId uint) error {
	return dao.DB.Model(&model.Order{}).Where("id = ? AND user_id =? ", aId, uId).Delete(&model.Order{}).Error
}

func (dao *OrderDao) ListOrderByCondition(condition map[string]interface{}, page model.BasePage) (order []*model.Order, err error) {

	err = dao.DB.Model(&model.Order{}).Where(condition).
		Offset((page.PageNum - 1) * (page.PageSize)).
		Limit(page.PageSize).Find(&order).Error
	return
}
