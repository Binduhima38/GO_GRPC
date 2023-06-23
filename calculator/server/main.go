package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/emohankrishna/go-grpc/calculator/proto"
)

type Server struct {
	pb.CalculatorServiceServer
}

var address string = "0.0.0.0:50051"

func main() {
	listner, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on : %v\n", err)
	}
	log.Printf("listening on %s\n", address)
	s := grpc.NewServer()

	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err := s.Serve(listner); err != nil {
		log.Fatalf("Failed to serve %v\n", err)
	}

}
