package main

import (
	"google.golang.org/grpc"
	pb "hongbao/test/go-grpc-example/proto"
	"io"
	"log"
	"net"
)
type StreamService struct {}
const PORT  = "9002"


func (s *StreamService) List(r *pb.StreamRequest, stream pb.StreamService_ListServer) error {
	/*
	服务器端流式 RPC
	简单来讲就是客户端发起一次普通的RPC请求，服务端通过流式响应多次发送数据集，客户端 Recv 接收数据集
	 */
	for n := 0;n<=6;n++{
		err := stream.Send(&pb.StreamResponse{
			Pt:&pb.StreamPoint{
				Name:r.Pt.Name,
				Value: r.Pt.Value + int32(n),
			},
		})
		if err != nil {
			return err
		}

	}
	return nil
}
func (s *StreamService) Record(stream pb.StreamService_RecordServer) error {
	/*
		客户端流式 RPC
	单向流，客户端通过流式发起多次 RPC 请求给服务端，服务端发起一次响应给客户端
	*/
	for {
		r,err := stream.Recv()
		//在这段程序中，我们对每一个 Recv 都进行了处理，当发现 io.EOF (流关闭) 后，
		// 需要将最终的响应结果发送给客户端，同时关闭正在另外一侧等待的 Rec   在client.go 文件里关闭另一端的recv方法也就是这个里的方法
		if err == io.EOF{
			return stream.SendAndClose(&pb.StreamResponse{Pt:&pb.StreamPoint{Name:"gggggg",Value:1}})
		}
		if err != nil {
			return err
		}
		log.Printf("stream.Recv pt.name: %s, pt.value: %d", r.Pt.Name, r.Pt.Value)
	}
	return nil
}
func (s *StreamService) Route(stream pb.StreamService_RouteServer) error {
	/*
		双向流式 RPC
		双向流式 RPC，顾名思义是双向流。由客户端以流式的方式发起多次请求，
		服务端同样以流式的方式响应多次请求
	*/
	n := 0
	for {
		err := stream.Send(&pb.StreamResponse{
			Pt:&pb.StreamPoint{
				Name: "双向流",
				Value: 12,
			},
		})
		if err != nil{
			return err
		}
		r,err := stream.Recv()
		if err == io.EOF{
			return nil
		}
		if err != nil {
			return err
		}
		n++
		log.Printf("stream.Recv pt.name: %s, pt.value: %d", r.Pt.Name, r.Pt.Value)
	}
	return nil
}



func main() {
	server := grpc.NewServer()
	pb.RegisterStreamServiceServer(server,&StreamService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	server.Serve(lis)


}





