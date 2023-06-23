package main

import (
	"fmt"
	"log"

	pb "github.com/emohankrishna/go-grpc/greet/proto"
)

// Server Streaming
func (s *Server) GreetManyTimes(r *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked %v\n", r)
	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d\n", r.FirstName, i)
		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}
	return nil
}
