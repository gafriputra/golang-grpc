package main

import (
	"log"

	pb "github.com/gafriputra/grpc-udemy/blog/proto"
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
	c := pb.NewBlogServiceClient(conn)

	id := createBlog(c)
	readBlog(c, id)
	updateBlog(c, id)
	listBlog(c)
	deleteBlog(c, id)
}
