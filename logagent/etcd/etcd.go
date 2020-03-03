package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var (
	cli *clientv3.Client
	err error
)

type LogEntry struct {
	Path string `json"path"`  //日志存放的信息
	Topic string `json:"topic"` //日志要发往kafka中的topic
	
}
func Init(addr string)  {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints: []string{addr},
		DialTimeout: 6*time.Second,
	})
	//watch操作
	//watch 用来获取未来更改的通知。
	if err != nil {
		// handle errors!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")


}

//从etcd 中获取配置项
func GetCont(key string)(logEnTryConf []*LogEntry, err error)  {
	// get
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		err = json.Unmarshal(ev.Value,&logEnTryConf)
		//fmt.Printf("%s:%s\n", ev.Key, ev.Value)
		if err != nil{
			fmt.Printf("unmarshal etcd value failed,err:%v\n",err)
			return
		}
	}
	return
}

//etcd watch<- chan []*etcd.LogEntr
func WatchConf(key string,newConfCh chan []*LogEntry)  {
	ch := cli.Watch(context.Background(),key)
	//从通道中取值 监视的信息
	for wresp := range ch {
		for _, evt := range wresp.Events {
			fmt.Printf("Type: %s Key:%s Value:%s\n", evt.Type, evt.Kv.Key,
				evt.Kv.Value)
			//通知别人 通知taillog.tskMgr
			//1.先判断操作类型 get delete ....
			var newConf []*LogEntry
			if evt.Type == clientv3.EventTypeDelete{
				//如果删除操作
			}
			err := json.Unmarshal(evt.Kv.Value,&newConf)
			if err != nil{
				fmt.Printf("unmarshal failed,err :%v\n",err)
				continue
			}
			fmt.Printf("get new conf : %v\n",newConf)
			newConfCh <- newConf
		}
	}

}


