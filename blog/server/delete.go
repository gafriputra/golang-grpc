package main

import (
	"context"
	"fmt"

	pb "github.com/gafriputra/grpc-udemy/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteBlog(ctx context.Context, in *pb.BlogId) (response *emptypb.Empty, err error) {
	response = nil
	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		err = status.Errorf(codes.InvalidArgument, "Cannot parse ID")
		return
	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": oid})

	if err != nil {
		err = status.Errorf(codes.Internal, fmt.Sprintf("Cannot delete blog %v", err))
		return
	} else if res.DeletedCount == 0 {
		err = status.Errorf(codes.NotFound, "Cannot find blog with id")
		return
	}

	response = &emptypb.Empty{}
	return
}
