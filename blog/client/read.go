package main

import (
	"context"
	"log"

	pb "github.com/gafriputra/grpc-udemy/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	req := &pb.BlogId{Id: id}
	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		log.Printf("Error reading blog: %v", err)
	}

	log.Printf("Got blog: %v", res)
	return res
}
