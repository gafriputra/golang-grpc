package main

import (
	"context"

	pb "github.com/gafriputra/grpc-udemy/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ReadBlog(ctx context.Context, in *pb.BlogId) (data *pb.Blog, err error) {
	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		err = status.Errorf(codes.InvalidArgument, "Cannot parse ID")
		return
	}

	blogData := &BlogItem{}
	filter := bson.M{"_id": oid}

	res := collection.FindOne(ctx, filter)
	if err = res.Decode(blogData); err != nil {
		err = status.Errorf(codes.NotFound, "Cannot find blog with ID %s", oid)
		return
	}

	data = documentToBlog(blogData)
	return
}
