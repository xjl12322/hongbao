package dao

import (
	"github.com/go-xorm/xorm"
	"hongbao/choujiang/models"
	_ "github.com/go-sql-driver/mysql"
	"log"
)


type CodeDao struct {
	engine *xorm.Engine
}

func NewCodeDao(engine *xorm.Engine)*CodeDao  {
	return &CodeDao{
		engine:engine,
	}
}


func (d *CodeDao)Get(id int)*models.LtGift  {
	data := &models.LtGift{Id:id}
	ok,err := d.engine.Get(data)
	if ok && err == nil{
		return data
	}else{
		data.Id = 0
		return data
	}
}


func (d *CodeDao)GetAll()[]models.LtGift {
	dataList := make([]models.LtGift,0)
	err := d.engine.Desc("id").Find(dataList)

	if err != nil{
		log.Println("Code_dao.getAll error=",err)
		return dataList
	}
	return dataList
}


func (d *CodeDao)CountAll()int64 {
	num,err := d.engine.Count(&models.LtGift{})
	if err != nil{
		return 0
	}else{
		return num
	}
}


//根据id删除奖品
func (d *CodeDao) Delete(id int) error {
	data := &models.LtGift{Id: id, SysStatus: 1}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}


func (d *CodeDao) Update(data *models.LtGift, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}


func (d *CodeDao) Create(data *models.LtGift) error {
	_, err := d.engine.Insert(data)
	return err
}