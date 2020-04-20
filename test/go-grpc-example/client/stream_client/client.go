package main
import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "hongbao/test/go-grpc-example/proto"
	"io"
	"log"
)
const (
	PORT = "9002"
)

func main() {
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	if err != nil{
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()
	client := pb.NewStreamServiceClient(conn)
	//err = printLists(client,&pb.StreamRequest{Pt: &pb.StreamPoint{Name: "gRPC Stream Client: List", Value: 2018}})
	//if err != nil {
	//	log.Fatalf("printLists.err: %v", err)
	//}
	//err = printRecord(client, &pb.StreamRequest{Pt: &pb.StreamPoint{Name: "gRPC Stream Client: Record", Value: 1000}})
	//if err != nil {
	//	log.Fatalf("printRecord.err: %v", err)
	//}
	err = printRoute(client, &pb.StreamRequest{Pt: &pb.StreamPoint{Name: "gRPC Stream Client: Route", Value: 2018}})
	if err != nil {
		log.Fatalf("printRoute.err: %v", err)
	}
}


func printLists(client pb.StreamServiceClient, r *pb.StreamRequest) error {
	//服务器端流式
	stream,err := client.List(context.Background(),r)
	if err != nil {
		return err
	}
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("ninininininininiininni",err)
		}
		if err != nil {
			return err
		}
		log.Printf("resp: pj.name: %s, pt.value: %d", response.Pt.Name, response.Pt.Value)
	}
	return nil
}


func printRecord(client pb.StreamServiceClient, r *pb.StreamRequest) error {
	//客户端流式  连续发送多次到服务器端  服务器端返回一次响应
	stream, err := client.Record(context.Background())
	if err != nil {
		return err
	}
	for n := 0; n < 6; n++ {
		err := stream.Send(r)
		if err != nil {
			return err
		}
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}
	log.Printf("resp: pj.name: %s, pt.value: %d", resp.Pt.Name, resp.Pt.Value)
	return nil
}
func printRoute(client pb.StreamServiceClient, r *pb.StreamRequest) error {
	//双向流式 RPC，顾名思义是双向流。由客户端以流式的方式发起多次请求，
	// 服务端同样以流式的方式响应多次请求
	stream,err := client.Route(context.Background())
	if err != nil {
		return err
	}
	for n := 0;n<=6;n++{
		err = stream.Send(r)
		if err != nil {
			return err
		}
		resp,err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil{
			return err
		}
		log.Printf("resp: pj.name: %s, pt.value: %d", resp.Pt.Name, resp.Pt.Value)
	}
	stream.CloseSend()
	return nil
}