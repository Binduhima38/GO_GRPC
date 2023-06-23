package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/emohankrishna/go-grpc/calculator/proto"
)

func doCalculate(client pb.CalculatorServiceClient) {
	log.Println("doCalculate invoked")
	res, err := client.Sum(context.Background(), &pb.SumRequest{
		Num1: 3,
		Num2: 10,
	})
	if err != nil {
		log.Printf("couldn't calculate %v\n", err)
	}
	log.Printf("Calculation of 3 + 10 is %v\n", res.Result)
}

func doCalculatePrimeDecomposition(client pb.CalculatorServiceClient) {
	log.Println("doCalculatePrimeDecomposition invoked")
	res, err := client.PrimeDecomposition(context.Background(), &pb.NumberRequest{
		Num: 120,
	})
	if err != nil {
		log.Fatalf("error while calling PrimeDecomposition %v\n", err)
	}

	for {
		msg, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading the stream %v\n", err)
		}
		log.Printf("One of the prime number is : %v", msg.Result)

	}
}

func doCalculateAverage(client pb.CalculatorServiceClient) {
	log.Println("doCalculateAverage invoked")
	stream, err := client.Average(context.Background())
	if err != nil {
		log.Fatalf("error while calling doCalculateAverage %v\n", err)
	}
	arr := []uint32{1, 2, 3, 4}

	for _, v := range arr {
		time.Sleep(time.Second * 1)
		stream.Send(&pb.AverageRequest{
			Num: v,
		})
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while reading the stream %v\n", err)
		return
	}
	log.Printf("Average of  %v : %v", arr, res.Res)
}
