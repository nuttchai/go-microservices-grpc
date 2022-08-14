package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/nuttchai/go-microservices-grpc/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Nutt"},
		{FirstName: "Chai"},
		{FirstName: "World"},
	}

	// creating a channel
	waitc := make(chan struct{})

	// this goroutine is for sending requests
	go func() {
		for _, req := range reqs {
			log.Printf("Send request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	// this goroutine is for receiving responses
	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}

			log.Printf("Received: %v\n", res.Result)
		}
		// close channel
		close(waitc)
	}()

	// make waitc wait for "close(waitc)"
	<-waitc
}
