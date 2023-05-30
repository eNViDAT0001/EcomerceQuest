package recommender

import (
	"context"
	proto "github.com/eNViDAT0001/Thesis/Backend/thesis_proto"
)

func (r recommenderService) AddComment(ctx context.Context, req *proto.CommentReq) (*proto.NonQueryResponse, error) {
	conn, client, err := GRPCClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	return client.AddComment(ctx, req)
}
