package conf

const DriverName = "mysql"

type DbConfig struct {
	Host      string
	Port      int
	User      string
	Pwd       string
	Database  string
	IsRunning bool // 是否正常运行
}

// 系统中所有mysql主库 root:root@tcp(127.0.0.1:3306)/lottery?charset=utf-8
var DbMasterList = []DbConfig{
	{
		Host:      "152.136.43.225",
		Port:      3306,
		User:      "root",
		Pwd:       "mysqlxjl12322@163.com",
		Database:  "choujiang",
		IsRunning: true,
	},
}
var DbMaster DbConfig = DbMasterList[0]
