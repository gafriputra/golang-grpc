package main

import (
	"context"
	"log"

	pb "github.com/gafriputra/grpc-udemy/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  10,
		SecondNumber: 20,
	})

	if err != nil {
		log.Fatalf("Couldn't get greeting: %v", err)
	}

	log.Printf("Got Sum : %v", res)
}
