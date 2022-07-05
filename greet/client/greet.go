package main

import (
	"context"
	"io"
	"log"
	"time"

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

func doGreetEveryone(c pb.GreetServiceClient) {
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("error while creating stream: %v", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Gafri"},
		{FirstName: "Putra"},
		{FirstName: "Aliffansah"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send request: %v", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("error while receiving response: %v", err)
				break
			}

			log.Printf("Received: %v\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}
