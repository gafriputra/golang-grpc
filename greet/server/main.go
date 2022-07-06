package main

import (
	"log"
	"net"

	pb "github.com/gafriputra/grpc-udemy/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "0.0.0.0:3213"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	listener, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listener on %s: %v\n", addr, err)
	}

	defer listener.Close()

	log.Printf("Listening on %s", addr)

	opts := []grpc.ServerOption{}
	tls := true

	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			log.Fatalf("Failed loading certificate: %v\n", err)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(s, &Server{})
	defer s.Stop()

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to start server: %v\n", err)
	}
}
