package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"hongbao/logagent/etcd"
	"hongbao/logagent/kafka_go"
	"hongbao/logagent/tailog"
	"os"
)

//logagent入口
var cfg *ini.File
var err error

func main()  {
	cfg,err = ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}


	//1.初始化kafka链接
	err := kafka_go.Init([]string{cfg.Section("kafka").Key("address").String()})
	if err != nil{
		fmt.Println("init kafka failed,err:%v\n",err)
		return
	}
	fmt.Println("init kafka")

	//2初始化etcd
	addr := cfg.Section("etcd").Key("address").String()
	etcd.Init(addr)
	fmt.Println("init etcd")
	//2.1 从etcd拉取日志收集的配置信息
	collect_log_key := cfg.Section("etcd").Key("collect_log_key").String()
	logEntryConf,err :=etcd.GetCont(collect_log_key)
	if err != nil{
		fmt.Println("etcd.GetCont,err:",err)
		return
	}
	fmt.Println("get conf from etcd success,%v\n",logEntryConf)

	//2.2 派一个任务去监视日志收集项的变化（有变化及时通知我的logagent实现热加载配置）





	//for index,value := range logEntryConf{
	//	fmt.Printf("index:%v value :%v\n",index,value)
	//}
	//3.收集日志发往kafka
	//3.1 循环每一个日志收集项 创建tailobj对象
	tailog.Init(logEntryConf)
	NewConfChan := tailog.NewConfChan() //从taillog中获取对外暴漏的通道
	go etcd.WatchConf(collect_log_key,NewConfChan) //发现配置的信息通知上面的通道
	////2.打开日志文件开始收集日志
    //err = tailog.Init(cfg.Section("taillog").Key("path").String())
	//if err != nil{
	//	fmt.Println("init tail failed,err:%v\n",err)
	//	return
	//}
	//fmt.Println("init tail")
	//
	//run()
}
//
//func run()  {
//	//1.读取日志
//	for {
//		select {
//		    case line := <- tailog.ReadChan():
//		    	kafka_go.SendToKafka(cfg.Section("kafka").Key("topic").String(),line.Text)
//		default:
//			time.Sleep(time.Second)
//		}
//
//	}
//
//}

