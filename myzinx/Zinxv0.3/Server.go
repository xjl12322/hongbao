package main

import (
	"fmt"
	"hongbao/zinx/ziface"
	"hongbao/zinx/znet"
)
//基于zinx框架来开发服务器端应用程序

//test自定义路由
type PingRouter struct {
	znet.BaseRouter
	
} 


//test prehanle 
func (this *PingRouter) PreHandle(request ziface.IRequest)  {
	fmt.Println("call before  ping err")
	_,err := request.GetConnection().GetTCPConnection().Write([]byte("before ping\n"))
	if err != nil{
		fmt.Println("call back before ping err")
	}


}

func (this *PingRouter) Handle(request ziface.IRequest)  {
	fmt.Println("call back  ping")
	_,err := request.GetConnection().GetTCPConnection().Write([]byte(" ping\n"))
	if err != nil{
		fmt.Println("call back  ping err")
	}

}

func (this *PingRouter) PostHandle(request ziface.IRequest)  {
	fmt.Println("call after ping")
	_,err := request.GetConnection().GetTCPConnection().Write([]byte("after ping\n"))
	if err != nil{
		fmt.Println("call back after ping err")
	}

}
func main()  {
	//创建server句柄 使用zinx
	s := znet.NewServer("[zinxv0.3]")

	//2 给当前zinx框架添加一个自定义root
	s.AddRouter(&PingRouter{})


	s.Serve()


}





