package config

const (
	MysqlDSN     = "root:mysqlxjl12322@163.com@(152.136.43.225:3306)/movies"
	Namespace    = "com.movies."
	LogPath      = "D:\\go_work\\src\\my-micro\\logdata"
	Num          = 20 // 分页每次取多少
	TickingNow   = 1  // 正在上映
	TickingWill  = 2  // 即将上映
	ActorType    = 1  // 演员
	DirectorType = 2  // 导演
)

const (
	ServiceNameUser    = "user"
	ServiceNameFilm    = "film"
	ServiceNameComment = "comment"
	ServiceNameCinema  = "cinema"
	ServiceNameOrder   = "order"
	ServiceNameCMS     = "cms"
)
