package conf



type RdsConfig struct {
	Host      string
	Port      int
	User      string
	Pwd       string
	IsRunning bool // 是否正常运行
}

// 系统中用到的所有redis缓存资源
var RdsCacheList = []RdsConfig{
	{
		Host:      "152.136.43.225",
		Port:      6379,
		User:      "",
		Pwd:       "mysqlxjl12322@163.com",
		IsRunning: true,
	},
}

var RdsCache RdsConfig = RdsCacheList[0]
