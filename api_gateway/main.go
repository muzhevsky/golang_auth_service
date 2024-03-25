package main

import (
	"api_gateway/pkg/proto"
	"context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := proto.NewAuthClient(conn)
	response, err := client.Register(context.Background(), &proto.RegisterRequest{Email: "test", Password: "123"})
	if err != nil {
		log.Fatal(err)
	}
	log.Print(response.UserId)
}
