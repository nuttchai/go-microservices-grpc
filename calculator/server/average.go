package main

import (
	"io"
	"log"

	pb "github.com/nuttchai/go-microservices-grpc/calculator/proto"
)

func (*Server) Average(stream pb.CalculatorService_AverageServer) error {
	total := 0
	count := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			if count == 0 {
				return stream.SendAndClose(&pb.AverageResponse{
					Average: float64(total),
				})
			}

			return stream.SendAndClose(&pb.AverageResponse{
				Average: float64(total) / float64(count),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		total += int(req.Integer)
		count++
	}
}
