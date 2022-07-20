package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/gafriputra/grpc-udemy/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
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

func listBlog(c pb.BlogServiceClient) {
	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Error listing blog: %v", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Error reading blog: %v", err)
		}

		fmt.Println(res)
	}
}
