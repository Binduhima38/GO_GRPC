package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/emohankrishna/go-grpc/calculator/proto"
)

// Unary
func (s *Server) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Calculate function was invoked %v\n", req)
	return &pb.SumResponse{
		Result: req.Num1 + req.Num2,
	}, nil
}

// Server streaming
func (s *Server) PrimeDecomposition(req *pb.NumberRequest, stream pb.CalculatorService_PrimeDecompositionServer) error {
	log.Printf("PrimeDecomposition function was invoked %v\n", req)
	var k uint32 = 2
	N := req.Num
	for N > 1 {
		if N%k == 0 {
			stream.Send(&pb.NumberResponse{
				Result: k,
			})
			N = N / k
		} else {
			k = k + 1
		}
	}
	return nil
}
func calculateAverage(arr []uint32) float32 {
	total := uint32(0)
	for _, v := range arr {
		total += v
	}
	return float32(total) / float32(len(arr))
}

// Client Streaming
func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {
	fmt.Println("Average function was invoked")
	arr := []uint32{}
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AverageResponse{
				Res: calculateAverage(arr),
			})
		}
		if err != nil {
			log.Fatalf("error while reading Client stream %v\n", err)
		}
		arr = append(arr, req.Num)
	}
}
