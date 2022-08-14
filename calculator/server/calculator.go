package main

import (
	"context"
	"log"

	pb "github.com/nuttchai/go-microservices-grpc/calculator/proto"
)

func (s *Server) Calculator(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Calculator function was invoked with %v\n", in)
	return &pb.CalculatorResponse{
		Summation: in.FirstInteger + in.SecondInteger,
	}, nil
}
