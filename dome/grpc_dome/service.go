package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pd"hongbao/dome/grpc_dome/proto"
	"net"
)

//1开启监听
//2实例化grpc服务端
//3grpc中注册服务
//4启动服务





//定义空接口
type UserInfoService struct {}
var u = UserInfoService{}

//实现方法
func (s *UserInfoService)GetUserInfo(ctx context.Context, req *pd.UserRequest)(resp *pd.UserResponse,err error)  {
	name := req.Name
	fmt.Println(name)
	if name == "zs" {
		resp = &pd.UserResponse{
			Id:1,
			Name:name,
			Age:22,
			Hobby:[]string{"shangge","paobu"},

		}
	}
	return
}

func main()  {
	//地址
	addr := "127.0.0.1:8089"
	listener, err := net.Listen("tcp",addr)
	if err != nil{
		fmt.Println("监听端口")
	}
	fmt.Println("jianting")
	s := grpc.NewServer()
	pd.RegisterUserInfoServiceServer(s,&u)
	s.Serve(listener)

}

