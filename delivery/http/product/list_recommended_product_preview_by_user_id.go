package product

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/grpc/grpc_base"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	proto "github.com/eNViDAT0001/Thesis/Backend/thesis_proto"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s *productHandler) ListRecommendedProductsPreview() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		userID, err := strconv.Atoi(cc.Param("user_id"))
		if err != nil {
			cc.ResponseError(err)
			return
		}
		productIDs, err := grpc_base.GetServices().RecommenderService.
			LisRecommendedProductIDsByUserID(newCtx, &proto.RecommendReq{UserId: int32(userID)})
		if err != nil {
			cc.ResponseError(err)
			return
		}

		result := map[string][]uint{
			"recommended": productIDs,
		}
		cc.Ok(result)
	}
}
