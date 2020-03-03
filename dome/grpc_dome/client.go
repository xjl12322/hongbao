package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pd"hongbao/dome/grpc_dome/proto"
)

//1连续服务端
//2实例化grpc客户端
//3调用
func main() {
	//1连续服务端
	conn,err := grpc.Dial("127.0.0.1:8089",grpc.WithInsecure())
	if err != nil{
		fmt.Println("yic",err)
	}

	defer conn.Close()
	//实例化grpc 客户端
	client := pd.NewUserInfoServiceClient(conn)



	response,err := client.GetUserInfo(context.Background(),&pd.UserRequest{Name:"zs"})
	if err != nil{
		fmt.Println("zs",err)
	}

	fmt.Println("yes",response)
}















