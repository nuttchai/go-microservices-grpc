package main

import (
	"context"

	pb "github.com/nuttchai/go-microservices-grpc/calculator/proto"
)

func (*Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	return &pb.SumResponse{
		Summation: in.Integer_1 + in.Integer_2,
	}, nil
}
