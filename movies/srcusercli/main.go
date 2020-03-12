package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"hongbao/movies/src/share/pb"
)
func main() {
	services := micro.NewService(
		micro.Name("test"),
		)

	services.Init()
	cli := pb.NewUserServiceExtService("go.micro.srv.user",services.Client())
	rep,err := cli.RegistAccount(context.TODO(),&pb.RegistAccountReq{Email:"xjl1@163.com",UserName:"xjl12322",Password:"12322"})
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(rep.String())


}
