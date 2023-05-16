package comment

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func (s commentHandler) GetCommentDetailByID() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		comment, err := strconv.Atoi(cc.Param("comment_id"))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		result, err := s.commentUC.GetCommentDetailByID(newCtx, uint(comment))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				cc.NoContent()
				return
			}

			cc.ResponseError(err)
			return
		}

		cc.Ok(result)
	}
}
