package main

import (
	"ToDoGRPCClient/todo"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("couldnt connect ", err)
	}
	defer conn.Close()
	client := todo.NewToDoServiceClient(conn)
	response, err := client.Create(context.Background(), &todo.ToDo{
		Goal:   "to kill the mockingbird 1",
		Status: 2,
		DueDate: &timestamppb.Timestamp{
			Seconds: int64(time.Second * 10000),
			Nanos:   int32(time.Nanosecond * 100),
		},
	})
	if err != nil {
		log.Println("creating new todo failed", err)
	}
	log.Println("CreateResponse ", response)

	listResponse, err := client.List(context.Background(), &todo.ListRequest{})
	if err != nil {
		log.Println("getting list failed", err)
	}
	log.Println("ListResponse ", listResponse)
}
