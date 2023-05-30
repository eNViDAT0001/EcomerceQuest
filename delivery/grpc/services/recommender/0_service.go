package recommender

import (
	"context"
	proto "github.com/eNViDAT0001/Thesis/Backend/thesis_proto"
)

type GRPCService interface {
	LisRecommendedProductIDsByUserID(ctx context.Context, req *proto.RecommentReq) ([]uint, error)
	AddComment(ctx context.Context, req *proto.CommentReq) (*proto.NonQueryResponse, error)
}
type recommenderService struct {
}

func NewRecommenderService() GRPCService {
	return &recommenderService{}
}
