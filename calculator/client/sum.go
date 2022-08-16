package main

import (
	"context"
	"log"

	pb "github.com/nuttchai/go-microservices-grpc/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		Integer_1: 3,
		Integer_2: 5,
	})

	if err != nil {
		log.Fatalf("Could not calculate: %v\n", err)
	}

	log.Printf("The summation: %d\n", res.Summation)
}
