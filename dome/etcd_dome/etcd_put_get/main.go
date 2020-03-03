package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)
// etcd client put/get demo
// use etcd/clientv3
func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})

	//watch操作
	//watch 用来获取未来更改的通知。
	if err != nil {
		// handle errors!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()
	// put
	value := `[{"path":"c:/tmp/nginx.log","topic":"web_log"},{"path":"c:/tmp/redis.log","topic":"web_log"},{"path":"d:/tmps/mysql.log","topic":"mysql_log"}]`
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "/logagent/collect_config", value)
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
	// get
	//ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	//resp, err := cli.Get(ctx, "q1mi")
	//cancel()
	//if err != nil {
	//	fmt.Printf("get from etcd failed, err:%v\n", err)
	//	return
	//}
	//for _, ev := range resp.Kvs {
	//	fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	//}
}







//etcd 续期

//func main() {
//	cli, err := clientv3.New(clientv3.Config{
//		Endpoints:   []string{"127.0.0.1:2379"},
//		DialTimeout: time.Second,
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer cli.Close()
//	//设置续期5秒
//	resp, err := cli.Grant(context.TODO(), 5)
//	if err != nil {
//		log.Fatal(err)
//	}
//	// 将 k-v 设置到etcd
//	_, err = cli.Put(context.TODO(), "root", "admin", clientv3.WithLease(resp.ID))
//	if err != nil {
//		log.Fatal(err)
//	}
//	// 若想一直有效，设置自动续期
//	ch, err := cli.KeepAlive(context.TODO(), resp.ID)
//	if err != nil {
//		log.Fatal(err)
//	}
//	for {
//		c := <-ch
//		fmt.Println("c:", c)
//	}
//}