package services

import (
	"hongbao/choujiang/dao"
	"hongbao/choujiang/datasource"
	"hongbao/choujiang/models"
)
//奖品结构体
type GiftService interface {
	GetAll() []models.LtGift
	CountAll() int64
	Get(id int) *models.LtGift
	Delete(id int) error
	Update(user *models.LtGift, columns []string) error
	Create(user *models.LtGift) error
}

type giftService struct {
	dao *dao.GiftDao
}

func NewGiftService() GiftService {
	return &giftService{
		dao:dao.NewGiftDao(datasource.InstanceDBMaster()),
	}

}
//获取全部商品列表

func (s *giftService)GetAll()[]models.LtGift  {

	return s.dao.GetAll()
}
//获取全部奖品数量
func (s *giftService)CountAll()int64 {
	return s.dao.CountAll()
}
//获取全部商品
func (s *giftService)Get(id int) *models.LtGift{
	return s.dao.Get(id)
}
//删除全部奖品
func (s *giftService)Delete(id int)error  {
	return s.dao.Delete(id)
}
//更新奖品列表
func (s *giftService)Update(data *models.LtGift, columns []string) error {
	return s.dao.Update(data,columns)
}
//添加奖品列表
func (s *giftService)Create(data *models.LtGift)error  {
	return s.dao.Create(data)
}







