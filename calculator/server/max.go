package main

import (
	"io"
	"log"

	pb "github.com/nuttchai/go-microservices-grpc/calculator/proto"
)

func (*Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max was invoked")

	var maxPointer *int64

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		current := req.Integer
		if maxPointer == nil || current > *maxPointer {
			maxPointer = &current
			err = stream.Send(&pb.MaxResponse{
				MaxInteger: current,
			})

			if err != nil {
				log.Fatalf("Error while sending data to client %v\n", err)
			}
		}
	}
}
