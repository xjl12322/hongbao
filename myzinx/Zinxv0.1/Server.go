package main

import "hongbao/zinx/znet"
//基于zinx框架来开发服务器端应用程序
func main()  {
	//创建server句柄 使用zinx
	s := znet.NewServer("[zinx v0.1]")
	s.Serve()

}





