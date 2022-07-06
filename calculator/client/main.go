package main

import (
	"log"

	pb "github.com/gafriputra/grpc-udemy/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "0.0.0.0:3213"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	log.Printf("listening on %v", addr)
	defer conn.Close()
	c := pb.NewCalculatorServiceClient(conn)
	// doSum(c)
	// doPrimes(c)
	// doAvg(c)
	// doMax(c)
	doSqrt(c, 10)
	doSqrt(c, -5)
}
