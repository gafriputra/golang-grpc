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

func (s *Server) ListBlogs(in *emptypb.Empty, stream pb.BlogService_ListBlogsServer) (err error) {
	cur, err := collection.Find(context.Background(), primitive.D{})
	if err != nil {
		err = status.Errorf(codes.Internal, fmt.Sprintf("Internal error: %v", err))
		return
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		data := &BlogItem{}
		err = cur.Decode(data)
		if err != nil {
			err = status.Errorf(codes.Internal, fmt.Sprintf("Error while decoding data from MongoDB: %v", err))
			return
		}

		err = stream.SendMsg(documentToBlog(data))
		if err != nil {
			err = status.Errorf(codes.Internal, fmt.Sprintf("Error while send streaming: %v", err))
			return
		}
	}

	if err = cur.Err(); err != nil {
		err = status.Errorf(codes.Internal, fmt.Sprintf("Internal error: %v", err))
	}

	return
}
