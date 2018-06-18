package main

import (
	"context"
	"echo"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connect to rpc server error:%v", err)
	} else {
		log.Println("Connection is build!")
	}

	defer conn.Close()

	client := echo.NewEchoServiceClient(conn)

	req_str := "Hello Go grpc!"
	res, err := client.Echo(context.Background(), &echo.EchoRequest{Content: &req_str})
	if err != nil {
		log.Fatalf("Client echo rpc failed:%v", err)
	} else {
		log.Println("after Echo!")
	}

	log.Printf("res:%s", res.GetContent())
}
