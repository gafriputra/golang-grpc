package main

import (
	"context"
	"log"

	pb "github.com/gafriputra/grpc-udemy/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	blog := &pb.Blog{
		AuthorId: "Gafri",
		Title:    "Buku KU",
		Content:  "Content of this blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error creating blog: %v", err)
	}

	log.Printf("Created blog: %v", res.Id)
	return res.Id
}
