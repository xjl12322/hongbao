package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"hongbao/movies/src/share/pb"

)
func testregister()  {
	services := micro.NewService(
		micro.Name("test"),
	)
	services.Init()
	cli := pb.NewUserServiceExtService("go.micro.srv.user",services.Client())
	rep,err := cli.RegistAccount(context.TODO(),&pb.RegistAccountReq{Email:"xjl2@163.com",UserName:"xjl12322",Password:"12322"})
	if err != nil{
		//t.Error(err)
		fmt.Println(err)
		return
	}
	//t.Log(rep.String())
	fmt.Println(rep.String())
}

func testlgin()  {
	services := micro.NewService(
		micro.Name("test"),
	)
	services.Init()
	cli := pb.NewUserServiceExtService("go.micro.srv.user",services.Client())
	rep,err := cli.LoginAccount(context.TODO(),&pb.LoginAccountReq{Email:"xjl2@163.com",Password:"12322"})
	if err != nil{
		//t.Error(err)
		fmt.Println(err)
		return
	}
	//t.Log(rep.String())
	fmt.Println(rep.String())
}

func testupdatauser()  {
	services := micro.NewService(
		micro.Name("test"),
	)
	services.Init()
	cli := pb.NewUserServiceExtService("go.micro.srv.user",services.Client())
	rep,err := cli.UpdateUserProfile(context.TODO(),&pb.UpdateUserProfileReq{UserPhone:"12355555",UserEmail:"xjl2@123.com",UserID:5})
	if err != nil{
		//t.Error(err)
		fmt.Println(err)
		return
	}
	//t.Log(rep.String())
	fmt.Println(rep.String())
}
func main() {
	testupdatauser()
}

