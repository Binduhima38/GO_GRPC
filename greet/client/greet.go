package main

import (
	"context"
	"log"

	pb "github.com/emohankrishna/go-grpc/greet/proto"
)

func doGreet(client pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := client.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Mohan",
	})

	if err != nil {
		log.Fatalf("couldn't greet %v\n", err)
		return
	}

	log.Printf("Greeting : %v", res.Result)
}
