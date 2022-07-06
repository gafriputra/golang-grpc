package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"

	pb "github.com/gafriputra/grpc-udemy/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoked")
	return &pb.SumResponse{
		Result: in.FirstNumber + in.SecondNumber,
	}, nil
}

func (s *Server) Primes(in *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {
	number := in.Number
	divisor := int64(2)
	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: divisor,
			})
			number /= divisor
		} else {
			divisor++
		}
	}
	return nil
}

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	var sum float32 = 0
	count := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: sum / float32(count),
			})
		}

		if err != nil {
			log.Printf("error while reading client stream: %v", err)
		}

		sum += req.Number
		count++
	}
}

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	var max int32 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.Send(&pb.MaxResponse{
				Result: max,
			})
		}

		if err != nil {
			log.Printf("error while reading client stream: %v", err)
		}

		if req.Number > max {
			max = req.Number
		}
	}
}

func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	number := in.Number

	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received negative number: %v", in.Number),
		)
	}

	return &pb.SqrtResponse{
		Result: math.Sqrt(float64(number)),
	}, nil
}
