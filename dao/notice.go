package dao

import (
	"context"
	"demoProject4mall/model"
	"fmt"

	//"fmt"
	"gorm.io/gorm"
)

type NoticeDao struct {
	*gorm.DB
}

func NewNoticeDao(ctx context.Context) *NoticeDao {
	return &NoticeDao{NewDBClient(ctx)}
}
func NewNoticeDaoByDB(db *gorm.DB) *NoticeDao {
	return &NoticeDao{db}
}

// GetNoticeById 根据id获取notice
func (dao *NoticeDao) GetNoticeById(id uint) (notic *model.Notice, err error) {
	//err = dao.DB.Model(&model.User{}).Where("id=?", id).First(&user).Error
	//return
	//if err != nil {
	err = dao.DB.Model(&model.Notice{}).Where("id=?", id).First(&notic).Error
	//	if err == gorm.ErrRecordNotFound {
	//		return nil, fmt.Errorf("user with id %d not found", id) // 如果没有找到用户，返回明确的错误信息
	//	}
	//	return nil, err // 如果是其他错误，直接返回
	//}
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 返回更具体的错误信息
			return nil, fmt.Errorf("notice with id %d not found", id)
		}
		return nil, err // 处理其他错误
	}
	return
}
