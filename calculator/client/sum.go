package main

import (
	"context"
	"io"
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

func doPrimes(c pb.CalculatorServiceClient) {
	req := &pb.PrimeRequest{
		Number: 12345678,
	}
	stream, err := c.Primes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling primes %v\n", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}

		log.Printf("Primes received : %v", res)
	}

}

func doAvg(c pb.CalculatorServiceClient) {
	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Fatalf("error while oppening avg: %v", err)
	}

	numbers := []float32{3.0, 1.0, 3.12, 3.19}

	for _, number := range numbers {
		stream.Send((&pb.AvgRequest{
			Number: number,
		}))
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receive result: %v", err)
	}

	log.Printf("result received : %v", res.Result)
}
