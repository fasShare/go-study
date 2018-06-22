package main

import (
	"EtcdResolver"
	"context"
	"echo"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"log"
	"time"
)

func main() {
	var erb EtcdResolver.EtcdResolverBuilder
	fmt.Println(erb.Scheme())
	fmt.Println(resolver.Get(erb.Scheme()).Scheme())

	conn, err := grpc.Dial("EtcdResolver://localhost/127.0.0.1:2379:serverList", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connect to rpc server error:%v", err)
	} else {
		log.Println("Connection is build!")
	}

	defer conn.Close()

	client := echo.NewEchoServiceClient(conn)

	req_str := "Hello Go grpc!"
	for i := 0; i < 100; i++ {
		res, err := client.Echo(context.Background(), &echo.EchoRequest{Content: &req_str})
		if err != nil {
			log.Fatalf("Client echo rpc failed:%v", err)
		} else {
			log.Println("after Echo!")
		}
		log.Printf("res:%s", res.GetContent())
		time.Sleep(time.Second * 10)
	}
}
