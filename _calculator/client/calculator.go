package main

import (
	"context"
	"log"

	pb "github.com/nuttchai/go-microservices-grpc/_calculator/proto"
)

func doCalculator(c pb.CalculatorServiceClient) {
	log.Println("doCalculator was invoked")
	res, err := c.Calculator(context.Background(), &pb.CalculatorRequest{
		FirstInteger:  10,
		SecondInteger: 2,
	})

	if err != nil {
		log.Fatalf("Could not calculate: %v\n", err)
	}

	log.Printf("The summation: %d\n", res.Summation)
}
