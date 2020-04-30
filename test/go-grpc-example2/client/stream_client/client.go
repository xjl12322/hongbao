package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	pb "hongbao/test/go-grpc-example/proto"
	"io"
	"log"
	"net/http"
)
const (
	PORT = "9002"
)
var client pb.StreamServiceClient
func init()  {
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	if err != nil{
		log.Fatalf("grpc.Dial err: %v", err)
	}
	client = pb.NewStreamServiceClient(conn)
}

func main() {
	app := gin.Default()
	app.GET("/printLists",printLists)
	app.GET("/printRecord",printRecord)
	app.GET("/printRoute",printRoute)


	app.Run("127.0.0.1:8080")

	//err = printLists(client,&pb.StreamRequest{Pt: &pb.StreamPoint{Name: "gRPC Stream Client: List", Value: 2018}})
	//if err != nil {
	//	log.Fatalf("printLists.err: %v", err)
	//}
	//err = printRecord(client, &pb.StreamRequest{Pt: &pb.StreamPoint{Name: "gRPC Stream Client: Record", Value: 1000}})
	//if err != nil {
	//	log.Fatalf("printRecord.err: %v", err)
	//}
	//err = printRoute(client, &pb.StreamRequest{Pt: &pb.StreamPoint{Name: "gRPC Stream Client: Route", Value: 2018}})
	//if err != nil {
	//	log.Fatalf("printRoute.err: %v", err)
	//}
}
func printLists(ctx *gin.Context)  {
	//服务器端流式
	stream,err := client.List(context.Background(),&pb.StreamRequest{
		Pt:&pb.StreamPoint{
			Name:                 "gRPC Stream Client: List",
			Value:                2018,
		},
	})
	if err != nil {
		ctx.JSON(http.StatusOK,err)
		return
	}

	for {
		response, err := stream.Recv()
		if err == io.EOF {
			ctx.JSON(http.StatusOK,gin.H{
				"err":err,
				"msg":"服务器端流式结束",
			})
			return
		}
		if err != nil {
			ctx.JSON(http.StatusOK,err)
			return
		}
		ctx.JSON(http.StatusOK,gin.H{"name":response.Pt.Name,"value":response.Pt.Value})
	}
}


func printRecord(ctx *gin.Context)  {
	//客户端流式  连续发送多次到服务器端  服务器端返回一次响应
	stream, err := client.Record(context.Background())
	if err != nil {
		ctx.JSON(http.StatusOK,err)
		return
	}
	for n := 0; n < 6; n++ {
		err := stream.Send(&pb.StreamRequest{
			Pt:&pb.StreamPoint{
				Name:                 "gRPC Stream Client: Record",
				Value:                11,
			},
		})
		if err != nil {
			ctx.JSON(http.StatusOK,err)
			return
		}
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		ctx.JSON(http.StatusOK,err)
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"name":resp.Pt.Name,"value":resp.Pt.Value})

}
func printRoute(ctx *gin.Context) {
	//双向流式 RPC，顾名思义是双向流。由客户端以流式的方式发起多次请求，
	// 服务端同样以流式的方式响应多次请求
	stream,err := client.Route(context.Background())
	if err != nil {
		ctx.JSON(http.StatusOK,err)
		return
	}

	response := map[string]interface{}{
		"errno":rsp.Error,
		"errmsg:":rsp.Errmsg,
		"data":area_list,
	}
	for n := 0;n<=6;n++{
		err = stream.Send(&pb.StreamRequest{
			Pt:&pb.StreamPoint{
				Name:                 "gRPC 双向 printRoute",
				Value:                22,
			},
		})
		if err != nil {
			ctx.JSON(http.StatusOK,err)
			return
		}
		resp,err := stream.Recv()
		if err == io.EOF {
			ctx.JSON(http.StatusOK,gin.H{
				"err":err,
				"msg":"服务器端流式结束",
			})
			break
		}
		if err != nil{
			ctx.JSON(http.StatusOK,err)
			return
		}
		//ctx.JSON(http.StatusOK,gin.H{"name":resp.Pt.Name,"value":resp.Pt.Value})

	}
	stream.CloseSend()
	//ctx.JSON(http.StatusOK,gin.H{"name":resp.Pt.Name,"value":resp.Pt.Value})
	ctx.JSON(http.StatusOK,gin.H{"1":1})


}