package main

import (
	"context"
	"log"
	"net"

	pb "github.com/gafriputra/grpc-udemy/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:3213"
var collection *mongo.Collection

type Server struct {
	pb.BlogServiceServer
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/"))
	if err != nil {
		log.Fatalf("error new client: %v", err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatalf("error connecting to MongoDB: %v", err)
	}

	collection = client.Database("blogdb").Collection("blog")

	listener, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listener on %s: %v\n", addr, err)
	}
	defer listener.Close()

	log.Printf("Listening on %s", addr)

	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Server{})
	defer s.Stop()

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to start server: %v\n", err)
	}
}
