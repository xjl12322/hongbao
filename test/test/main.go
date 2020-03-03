package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	//"github.com/micro/go-micro/registry"
	//"github.com/micro/go-plugins/registry/etcdv3"
	"hongbao/movies/src/share/pb"
)


func main() {
	//reg := etcdv3.NewRegistry(func(options *registry.Options) {
	//	options.Addrs = []string{
	//		"http://127.0.0.1:2379",
	//	}
	//})

	server := micro.NewService(micro.Name("test-user"))
	server.Init()
	userservice := pb.NewUserServiceExtService("go.micro.srv.user",server.Client())

	rep,err := userservice.RegistAccount(context.TODO(),&pb.RegistAccountReq{
		Email:"xjl12322@163.com",
		UserName:"xinjialei",
		Password:"1366313",

	})

	if err != nil{
		fmt.Println("err",err)
		return
	}
	fmt.Println(rep)


}