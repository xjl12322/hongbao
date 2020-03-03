package main
import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)
// watch demo
func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	fmt.Println("2222")
	defer cli.Close()

	// watch key:q1mi change
	rch := cli.Watch(context.Background(), "q1mi2") // <-chan WatchResponse

	//将上面的代码保存编译执行，此时程序就会等待etcd中 q1mi 这个key的变化。
	//例如：我们打开终端执行以下命令修改、删除、设置 q1mi 这个key。
	//上面的程序都能收到如下通知。
	//其他操作
	//其他操作请查看etcd/clientv3官方文档。
	//参考链接：
	//https://etcd.io/docs/v3.3.12/demo/
	//https://www.infoq.cn/article/etcd-interpretation-application-scenario-implement-principle/
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Println("1111")
			fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key,
				ev.Kv.Value)
		}
	}
	fmt.Println("3333")
}