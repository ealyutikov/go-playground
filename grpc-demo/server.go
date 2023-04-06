package main

import (
	"context"
	"log"
	"net"
	"time"

	demo_service "github.com/elyutikov/grpc-demo/proto"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)

	demo_service.RegisterDemoServiceServer(server, &demoServer{})

	err = server.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type demoServer struct {
	demo_service.UnimplementedDemoServiceServer
}

func (s *demoServer) Do(ctx context.Context, req *demo_service.Request) (*demo_service.Response, error) {
	log.Printf("start Do(%s)\n", req.String())
	defer log.Printf("end Do(%s)\n", req.String())

	time.Sleep(5 * time.Second)
	return &demo_service.Response{Message: "success"}, nil
}
