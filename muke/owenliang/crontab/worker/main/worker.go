package main

import (
	"flag"
	"fmt"
	"owenliang/crontab/worker"
	"runtime"
	"time"
)

var (
	confFile string // 配置文件路径
)

// 解析命令行参数
func initArgs() {
	// worker -config ./master.json -xxx 123 -yyy ddd
	// worker -h
	flag.StringVar(&confFile, "config", "./worker.json", "指定worker.json")
	flag.Parse()
}

// 初始化线程数量
func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var (
		err error
	)

	// 初始化命令行参数
	initArgs()

	// 初始化线程
	initEnv()

	// 加载配置
	if err = worker.InitConfig(confFile); err != nil {
		goto ERR
	}
	//
	// 初始化服务发现模块
	//if err = master.InitWorkerMgr(); err != nil {
	//	goto ERR
	//}
	//
	//// 日志管理器
	//if err =master.InitLogMgr(); err != nil {
	//	goto ERR
	//}
	//
	//  任务管理器
	if err = worker.InitJobMgr(); err != nil {
		goto ERR
	}

	//// 启动Api HTTP服务
	//if err = master.InitApiServer(); err != nil {
	//	goto ERR
	//}

	// 正常退出
	for {
		time.Sleep(1 * time.Second)
	}

	return

ERR:
	fmt.Println(err)
}























