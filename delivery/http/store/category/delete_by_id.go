package category

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/store/category/io"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s *categoryHandler) DeleteByCategoryID() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input io.DeleteCategoryNodeReq
		if err := cc.BindJSON(&input); err != nil {
			cc.BadRequest(err)
			return
		}

		err := s.categoryUC.DeleteCategoryByID(newCtx, input.ID, input.CategoryParentID)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Delete category success")
	}
}
func (s *categoryHandler) DeleteNodeByCategoryID() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		categoryID, err := strconv.Atoi(cc.Param("category_id"))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		err = s.categoryUC.DeleteCategoryNodeByID(newCtx, uint(categoryID))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Delete category success")
	}
}
