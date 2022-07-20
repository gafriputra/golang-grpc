package main

import (
	"context"

	pb "github.com/gafriputra/grpc-udemy/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) UpdateBlog(ctx context.Context, in *pb.Blog) (response *emptypb.Empty, err error) {
	response = nil
	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		err = status.Errorf(codes.InvalidArgument, "Cannot parse ID")
		return
	}

	data := &BlogItem{
		AuthorId: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	res, err := collection.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": data})

	if err != nil {
		err = status.Errorf(codes.Internal, "Cannot update")
		return
	} else if res.MatchedCount == 0 {
		err = status.Errorf(codes.NotFound, "Cannot find blog with id")
		return
	}

	response = &emptypb.Empty{}
	return
}
