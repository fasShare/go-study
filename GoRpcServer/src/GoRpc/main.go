package main

import (
	"context"
	"echo"
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
)

type myecho string

func (s *myecho) Echo(ctx context.Context, req *echo.EchoRequest) (*echo.EchoResponse, error) {
	var res echo.EchoResponse
	str := req.GetContent()
	res.Content = &str
	log.Printf("recv:", res.GetContent())
	return &res, nil
}

func main() {
	ip := flag.String("ip", "127.0.0.1", "The default ip of server!")
	port := flag.String("port", "8080", "The defaut port of server!")
	flag.Parse()
	log.Printf("%s\n", *ip+":"+*port)

	lis, err := net.Listen("tcp", *ip+":"+*port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server
	s := grpc.NewServer()
	echo.RegisterEchoServiceServer(s, new(myecho))
	s.Serve(lis)
}
