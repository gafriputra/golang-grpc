package main

import (
	"context"
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
