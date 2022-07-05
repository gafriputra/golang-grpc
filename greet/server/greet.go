package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/gafriputra/grpc-udemy/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked")
	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked")
	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hellom %s, number %d", in.FirstName, i)
		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}
	return nil
}

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{Result: res})
		}

		if err != nil {
			log.Fatalf("Error while client streaming: %v", err)
		}

		res += fmt.Sprintf("Hello %s!\n", req.FirstName)
	}
}

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("error while reading client streaming: %v", err)
		}

		res := "Hello " + req.FirstName + "!\n"
		err = stream.Send(&pb.GreetResponse{Result: res})

		if err != nil {
			log.Fatalf("error while sending data to client: %v", err)
		}

	}
}
