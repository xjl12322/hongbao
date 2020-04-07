package datasource

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/prometheus/common/log"
	"hongbao/choujiang/conf"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)




var dbLock sync.Mutex
var masterInstance *xorm.Engine

func InstanceDbMaster() *xorm.Engine  {
	if masterInstance != nil{
		return masterInstance
	}
	dbLock.Lock()
	defer dbLock.Unlock()
	if masterInstance != nil{
		return masterInstance
	}
	return NewDbMaster()
}
func NewDbMaster() *xorm.Engine {
	sourcename := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		conf.DbMaster.User,
		conf.DbMaster.Pwd,
		conf.DbMaster.Host,
		conf.DbMaster.Port,
		conf.DbMaster.Database)

	instance,err := xorm.NewEngine(conf.DriverName,sourcename)
	if err != nil{
		log.Fatal("dbhelper.newdbmaster newengine error",err)
		return nil
	}
	instance.ShowSQL(true)  //显示sql执行时间默认false
	masterInstance = instance
	return instance
}




