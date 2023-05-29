package grpc

import (
	"context"
	proto "github.com/eNViDAT0001/Thesis/Backend/ThesisProto"
	"github.com/eNViDAT0001/Thesis/Backend/external/grpc"
)

func AddComment() (*proto.NonQueryResponse, error) {
	conn := grpc.RecommenderClient()
	defer conn.Close()
	client := proto.NewRecommenderBaseCommentClient(conn)
	request := &proto.CommentReq{
		UserId:    1,
		ProductId: 2,
		Rating:    3,
	}
	// Call the Add RPC
	response, err := client.AddComment(context.Background(), request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
