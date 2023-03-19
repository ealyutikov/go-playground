package main

import (
	"context"
	"log"
	"time"

	"github.com/elyutikov/goblockchain/crypto"
	"github.com/elyutikov/goblockchain/node"
	"github.com/elyutikov/goblockchain/proto"
	"github.com/elyutikov/goblockchain/util"
	"google.golang.org/grpc"
)

func main() {
	makeNode(":3000", []string{}, true)
	time.Sleep(time.Second)
	makeNode(":4000", []string{":3000"}, false)
	time.Sleep(time.Second)
	makeNode(":5000", []string{":4000"}, false)

	for {
		time.Sleep(100 * time.Millisecond)
		makeTransaction()
	}
	select {}
}

func makeNode(address string, bootstrapNodes []string, isValidator bool) *node.Node {
	config := node.ServerConfig{
		Version:    "glockchain-1",
		ListenAddr: address,
	}

	if isValidator {
		config.PrivateKey = crypto.GeneratePrivateKey()
	}

	n := node.NewNode(config)
	go n.Start(address, bootstrapNodes)
	return n
}

func makeTransaction() {
	client, err := grpc.Dial(":3000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := proto.NewNodeClient(client)
	privKey := crypto.GeneratePrivateKey()
	tx := &proto.Transaction{
		Version: 1,
		Inputs: []*proto.TxInput{
			{
				PrevTxHash:   util.RandomHash(),
				PrevOutIndex: 0,
				PublicKey:    privKey.Public().Bytes(),
			},
		},
		Outputs: []*proto.TxOutput{
			{
				Amount:  99,
				Address: privKey.Public().Address().Bytes(),
			},
		},
	}

	_, err = c.HandleTransaction(context.TODO(), tx)
	if err != nil {
		log.Fatal(err)
	}
}
