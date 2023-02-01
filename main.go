package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/elyutikov/goblockchain/node"
	"github.com/elyutikov/goblockchain/proto"
	"google.golang.org/grpc"
)

func main() {
	node := node.NewNode()

	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)

	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	proto.RegisterNodeServer(server, node)
	fmt.Println("node running on port:", ":3000")

	go func() {
		for {
			time.Sleep(2 * time.Second)
			makeTransaction()
		}
	}()

	server.Serve(ln)
}

func makeTransaction() {
	client, err := grpc.Dial(":3000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := proto.NewNodeClient(client)

	version := &proto.Version{
		Version: "blocker-0.1",
		Height:  1,
	}

	_, err = c.Handshake(context.TODO(), version)
	if err != nil {
		log.Fatal(err)
	}
}
