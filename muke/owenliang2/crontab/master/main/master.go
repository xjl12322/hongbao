package main

import (
	"flag"
	"fmt"
	"owenliang/crontab/master"
	"runtime"
	"time"
)
var (
	confFile string
)
// 解析命令行参数
func initArgs() {
	// master -config ./master.json -xxx 123 -yyy ddd
	// master -h
	flag.StringVar(&confFile, "config", "./master.json", "指定master.json")
	flag.Parse()
}

func initEnv()  {
	runtime.GOMAXPROCS(runtime.NumCPU())
}


func main()  {
	var (
		err error
	)


	// 初始化命令行参数
	initArgs()
	// 加载配置
	if err = master.InitConfig(confFile); err != nil {
		goto ERR
	}
	//初始化etcd
	if err = master.InitJobMgr(); err!=nil{
		fmt.Println(err)
	}
	initEnv()
	// 启动Api HTTP服务
	if err = master.InitApiServer(); err != nil {
		goto ERR
	}
	for {
		time.Sleep(1*time.Second)
	}

ERR:
	fmt.Println(err)


}









