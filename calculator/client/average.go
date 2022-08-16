package main

import (
	"context"
	"log"
	"time"

	pb "github.com/nuttchai/go-microservices-grpc/calculator/proto"
)

func doAverage(c pb.CalculatorServiceClient) {
	log.Println("doAverage was invoked")

	reqs := []*pb.AverageRequest{
		{Integer: 1},
		{Integer: 2},
		{Integer: 3},
		{Integer: 4},
	}
	stream, err := c.Average(context.Background())

	if err != nil {
		log.Fatalf("Error while calling Average: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from Average: %v\n", err)
	}

	log.Printf("Average: %f\n", res.Average)
}
