package main

import (
	"log"
	"net"

	pb "github.com/nuttchai/go-microservices-grpc/calculator/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.CalculatorServiceServer
}

var addr string = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	srv := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(srv, &Server{})

	if err = srv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
