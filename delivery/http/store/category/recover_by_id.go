package category

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s *categoryHandler) RecoverCategoryByID() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		categoryID, err := strconv.Atoi(cc.Param("category_id"))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		err = s.categoryUC.RecoverCategoryByID(newCtx, uint(categoryID))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Delete category success")
	}
}
func (s *categoryHandler) RecoverCategoryNodeByID() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		categoryID, err := strconv.Atoi(cc.Param("category_id"))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		err = s.categoryUC.RecoverCategoryNodeByID(newCtx, uint(categoryID))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Delete category success")
	}
}
