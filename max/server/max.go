package main

import (
	"io"
	"log"

	pb "github.com/nuttchai/go-microservices-grpc/max/proto"
)

func (s *Server) Max(stream pb.MaxService_MaxServer) error {
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

		num := req.Number

		if maxPointer == nil || *maxPointer < num {
			maxPointer = &num
			err = stream.Send(&pb.MaxResponse{
				Max: num,
			})

			if err != nil {
				log.Fatalf("Error while sending data to client %v\n", err)
			}
		}

	}
}
