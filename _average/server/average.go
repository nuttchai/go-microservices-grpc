package main

import (
	"io"
	"log"

	pb "github.com/nuttchai/go-microservices-grpc/_average/proto"
)

func (s *Server) Average(stream pb.AverageService_AverageServer) error {
	log.Println("Average function was invoked")

	summation := 0
	count := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			if count == 0 {
				return stream.SendAndClose(&pb.AverageResponse{
					Result: float64(summation),
				})
			}

			return stream.SendAndClose(&pb.AverageResponse{
				Result: float64(summation) / float64(count),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		summation += int(req.Number)
		count++
	}
}
