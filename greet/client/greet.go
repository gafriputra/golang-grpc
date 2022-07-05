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

func doLongGreet(c pb.GreetServiceClient) {
	reqs := []*pb.GreetRequest{
		{FirstName: "Gafri"},
		{FirstName: "Putra"},
		{FirstName: "Aliffansah"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling long_greet: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from long greet: %v", err)
	}

	log.Printf("LongGreet : %s\n", res.Result)
}
