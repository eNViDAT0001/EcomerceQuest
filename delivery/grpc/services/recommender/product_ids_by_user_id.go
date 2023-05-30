package recommender

import (
	"context"
	proto "github.com/eNViDAT0001/Thesis/Backend/thesis_proto"
)

func (r recommenderService) LisRecommendedProductIDsByUserID(ctx context.Context, req *proto.RecommentReq) ([]uint, error) {
	conn, client, err := GRPCClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	result, err := client.LisRecommendedProductIDsByUserID(ctx, req)
	if err != nil {
		return nil, err
	}

	productIDs := make([]uint, 0)
	for _, i := range result.ProductId {
		productIDs = append(productIDs, uint(i))
	}
	return productIDs, err
}
