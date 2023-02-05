package main

import (
	"context"
	"log"

	"github.com/elyutikov/goblockchain/node"
	"github.com/elyutikov/goblockchain/proto"
	"google.golang.org/grpc"
)

func main() {
	makeNode(":3000", []string{})
	makeNode(":4000", []string{})

	select {}

	// node := node.NewNode()

	// go func() {
	// 	for {
	// 		time.Sleep(2 * time.Second)
	// 		makeTransaction()
	// 	}
	// }()

	// log.Fatal(node.Start(":3000"))
}

func makeNode(address string, bootstrapNodes []string) *node.Node {
	n := node.NewNode()
	go n.Start(address)

	if len(bootstrapNodes) > 0 {
		if err := n.BootstrapNetwork(bootstrapNodes); err != nil {
			log.Fatal(err)
		}
	}
	return n
}

func makeTransaction() {
	client, err := grpc.Dial(":3000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := proto.NewNodeClient(client)

	version := &proto.Version{
		Version:    "blocker-0.1",
		Height:     1,
		ListenAddr: ":4000",
	}

	_, err = c.Handshake(context.TODO(), version)
	if err != nil {
		log.Fatal(err)
	}
}
