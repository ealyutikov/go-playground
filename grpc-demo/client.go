package main

import (
	"context"
	"fmt"
	"log"

	demo_service "github.com/elyutikov/grpc-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	retryPolicy := `{
	      "methodConfig": [{
	          "name": [{"service": "demo_service.DemoService"}],
			  "timeout": "5.0s",
	          "waitForReady": true,
	          "retryPolicy": {
	              "MaxAttempts": 5,
	              "InitialBackoff": "5s",
	              "MaxBackoff": "20s",
	              "BackoffMultiplier": 1.5,
	              "RetryableStatusCodes": [ "UNAVAILABLE", "DEADLINE_EXCEEDED", "UNKNOWN" ]
	          }
	      }]
	  }`

	conn, err := grpc.Dial(
		"127.0.0.1:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(retryPolicy),
	)

	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := demo_service.NewDemoServiceClient(conn)

	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("header-key-from-client", "val"))
	//ctx = metadata.AppendToOutgoingContext(ctx, "header-key-from-client", "val")

	var header, trailer metadata.MD
	response, err := client.Do(ctx, &demo_service.Request{Message: "hello"}, grpc.Header(&header), grpc.Trailer(&trailer))

	log.Printf("headers from server(%+v)\n", header)
	log.Printf("trailers from server(%+v)\n", trailer)

	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	fmt.Println(response.Message)
}
