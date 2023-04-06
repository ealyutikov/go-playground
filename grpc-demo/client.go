package main

import (
	"context"
	"fmt"
	"log"
	"time"

	demo_service "github.com/elyutikov/grpc-demo/proto"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
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

	callOpts := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffExponential(1 * time.Second)),
		grpc_retry.WithMax(10),
		grpc_retry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded, codes.Unavailable),
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(retryPolicy),
	}

	opts = append(opts, grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(callOpts...)))

	conn, err := grpc.Dial("127.0.0.1:50051", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := demo_service.NewDemoServiceClient(conn)
	response, err := client.Do(context.Background(), &demo_service.Request{Message: "hello"})

	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	fmt.Println(response.Message)
}
