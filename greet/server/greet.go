package main

import (
	"context"
	"log"

	pb "github.com/gafriputra/grpc-udemy/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked")
	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
