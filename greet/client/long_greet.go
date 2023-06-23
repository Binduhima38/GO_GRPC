package main

import (
	"context"
	"log"
	"time"

	pb "github.com/emohankrishna/go-grpc/greet/proto"
)

// Client Streaming
func doLongGreet(client pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")
	stream, err := client.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("couldn't long greet %v\n", err)
		return
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Mohan"},
		{FirstName: "Hima"},
		{FirstName: "Krishna"},
		{FirstName: "Bindu"},
		{FirstName: "Karthik"},
		{FirstName: "Mounika"},
		{FirstName: "keerthana"},
		{FirstName: "Arnab"},
	}

	for _, req := range reqs {
		log.Printf("Sending greeting request %v\n", req)
		time.Sleep(time.Second * 1)
		stream.Send(req)
	}
	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("error while receiving response from long greet %v\n", err)
		return
	}
	log.Printf("LongGreet : %v", res.Result)
}
