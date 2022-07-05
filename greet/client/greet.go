package main

import (
	"context"
	"io"
	"log"

	pb "github.com/gafriputra/grpc-udemy/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Gafri",
	})

	if err != nil {
		log.Fatalf("Couldn't get greeting: %v", err)
	}

	log.Printf("Got greeting: %v", res)
}

func doGreetManyTimes(c pb.GreetServiceClient) {
	req := &pb.GreetRequest{
		FirstName: "Gafri",
	}
	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Couldn't get greeting: %v", err)
	}
	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the streaming %v\n", err)
		}
		log.Printf("GreetManyTimes: %v", msg)
	}

}
