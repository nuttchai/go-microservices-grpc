package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/nuttchai/go-microservices-grpc/blog/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	client := pb.NewBlogServiceClient(conn)

	id := createBlog(client)
	readBlog(client, id) // valid Id
	// readBlog(client, "nonExistingId") // non valid Id
	updateBlog(client, id)
	listBlog(client)
	deleteBlog(client, id)
}
