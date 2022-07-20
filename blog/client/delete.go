package main

import (
	"context"
	"log"

	pb "github.com/gafriputra/grpc-udemy/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})

	if err != nil {
		log.Printf("Error deleted blog: %v", err)
		return
	}

	log.Printf("Success deleted blog")
}
