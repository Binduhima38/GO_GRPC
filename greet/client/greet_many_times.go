package main

import (
	"context"
	"io"
	"log"

	pb "github.com/emohankrishna/go-grpc/greet/proto"
)

// Server Streaming
func doGreetManyTimes(client pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")
	req := &pb.GreetRequest{
		FirstName: "Mohan",
	}
	res, err := client.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling GreetManyTimes %v\n", err)
	}

	for {
		msg, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading the stream %v\n", err)
		}
		log.Printf("GreetManyTimes : %v", msg.Result)
	}
}
