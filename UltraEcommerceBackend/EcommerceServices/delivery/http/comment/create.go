package comment

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/delivery/http/comment/convert"
	"github.com/eNViDAT0001/Thesis/Ecommerce/delivery/http/comment/io"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/request"
	"github.com/gin-gonic/gin"
	"strconv"
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

		commentID, err := s.commentUC.CreateComment(newCtx, inputSto, input.Files)
		if err != nil {
			cc.ResponseError(err)
			return
		}
		result := map[string]interface{}{
			"CommentID": commentID,
		}
		cc.Ok(result)
	}
}
