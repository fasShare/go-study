package main

import (
	"context"
	"echo"
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
	lis, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server
	s := grpc.NewServer()
	echo.RegisterEchoServiceServer(s, new(myecho))
	s.Serve(lis)
}
