package main

import (
	"log"
	"time"

	pb "github.com/gafriputra/grpc-udemy/greet/proto"
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
	c := pb.NewGreetServiceClient(conn)
	// doGreet(c)
	// doGreetManyTimes(c)
	// doLongGreet(c)
	// doGreetEveryone(c)
	doGreetWithDeadline(c, 5*time.Second)
	doGreetWithDeadline(c, 3*time.Second)

}
