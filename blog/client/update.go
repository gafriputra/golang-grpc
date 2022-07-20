package main

import (
	"context"
	"log"

	pb "github.com/gafriputra/grpc-udemy/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	updatedBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Gafri Putra ALiffansah",
		Title:    "New Title",
		Content:  "Contents of the first blog",
	}
	_, err := c.UpdateBlog(context.Background(), updatedBlog)

	if err != nil {
		log.Printf("Error updated blog: %v", err)
	}

	log.Printf("Success updated blog")
}
