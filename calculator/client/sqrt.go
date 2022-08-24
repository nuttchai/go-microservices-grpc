package main

import (
	"context"
	"log"

	pb "github.com/nuttchai/go-microservices-grpc/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) {
	log.Println("doSqrt was invoked")
	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{Number: n})

	if err != nil {
		e, ok := status.FromError(err)

		// ok tells whether the error is from grpc or not
		if ok {
			log.Printf("Error message from server: %s\n", e.Message())
			log.Printf("Error code from server: %s\n", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Printf("You probably sent a negative input!")
				return
			}
		} else {
			log.Fatalf("A non gRPC error: %v\n", err)
		}
	}

	log.Printf("The result of square root %d is %f", n, res.Result)
}
