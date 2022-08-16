package main

import (
	"log"
	"time"

	pb "github.com/nuttchai/go-microservices-grpc/calculator/proto"
)

func (*Server) Primes(in *pb.PrimesRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Primes function was invoked with %v\n", in)

	var k int64 = 2
	var current int64 = in.Integer

	for current >= k {
		if current%k == 0 {
			current = current / k
			stream.Send(&pb.PrimesResponse{
				Result: k,
			})
			time.Sleep(1 * time.Second)
		} else {
			k++
		}
	}

	return nil
}
