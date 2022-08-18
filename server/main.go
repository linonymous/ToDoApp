package main

import (
	"ToDoGRPC/todo"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal("port listen failed ", err)
	}

	service := todo.Service{}
	grpcServer := grpc.NewServer()
	service.t
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("couldnt start grpc server", err)
	}
}
