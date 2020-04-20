package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)
func main() {
	lis,err := net.Listen("tcp",fmt.Sprintf(":%d", 8888))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()




}
