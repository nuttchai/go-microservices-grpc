package main

import (
	"log"
	"net"

	pb "github.com/nuttchai/go-microservices-grpc/average/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.AverageServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	srv := grpc.NewServer()
	pb.RegisterAverageServiceServer(srv, &Server{})

	if err = srv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
