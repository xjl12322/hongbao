package main

import (
	"fmt"
	"net"
	"time"
)

/*
模拟客户端
 */

func main() {
	//1 链接远程服务器获得 conn链接
	conn,err := net.Dial("tcp","127.0.0.1:8999")
	if err != nil{
		fmt.Println("client start err,exit")
	}

	for {
		//2 链接调用write写数据
		_,err := conn.Write([]byte("ni hao0.2"))
		if err != nil{
			fmt.Println("write conn err",err)
			return
		}
		buf := make([]byte,512)
		cnt,err := conn.Read(buf)
		if err != nil{
			fmt.Println("read buf error",err)
			return
		}
		fmt.Printf("server call back: %s, cnt=%d\n",buf,cnt)

		time.Sleep(time.Second*2)

	}

}





