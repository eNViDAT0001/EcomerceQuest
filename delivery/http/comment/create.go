package comment

import (
	"context"
	"strconv"

	"github.com/eNViDAT0001/Thesis/Backend/delivery/grpc/grpc_base"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/comment/convert"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/comment/io"
	"github.com/eNViDAT0001/Thesis/Backend/external/event_background"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	proto "github.com/eNViDAT0001/Thesis/Backend/thesis_proto"
	"github.com/gin-gonic/gin"
)

func (s commentHandler) CreateComment() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input io.CreateCommentReq
		if err := cc.ShouldBind(&input); err != nil {
			cc.BadRequest(err)
			return
		}

		productID, _ := strconv.Atoi(cc.Param("product_id"))
		userID, _ := strconv.Atoi(cc.Param("user_id"))

		input.UserID = uint(userID)
		input.ProductID = uint(productID)

		inputSto, err := convert.CreateReqToCreateCommentInput(&input)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		commentID, err := s.commentUC.CreateComment(newCtx, inputSto, input.Media)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		event_background.AddBackgroundJobs(false, event_background.NewJob(func(ctx context.Context) error {
			_, err = grpc_base.GetServices().RecommenderService.AddComment(newCtx, &proto.CommentReq{
				UserId:    int32(userID),
				ProductId: int32(productID),
				Rating:    int32(input.Rating),
			})

			return err
		}))


		result := map[string]interface{}{
			"CommentID": commentID,
		}
		cc.Ok(result)
	}
}
