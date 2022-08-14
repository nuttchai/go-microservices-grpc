package main

import (
	"log"

	pb "github.com/nuttchai/go-microservices-grpc/primes/proto"
)

func (s *Server) Primes(in *pb.PrimesRequest, stream pb.PrimesService_PrimesServer) error {
	log.Printf("Primes function was invoked with %v\n", in)

	var k int64 = 2
	var current int64 = in.Number

	for current >= k {
		if current%k == 0 {
			current = current / k
			stream.Send(&pb.PrimesResponse{
				Result: k,
			})
		} else {
			k++
		}
	}

	return nil
}
