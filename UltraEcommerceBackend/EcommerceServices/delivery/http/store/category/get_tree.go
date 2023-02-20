package category

import (
	"context"

	"github.com/eNViDAT0001/Thesis/Ecommerce/external/request"
	"github.com/gin-gonic/gin"
)

func (s *categoryHandler) GetCategoryTree() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		categories, err := s.categoryUC.GetCategoryTree(newCtx)
		if err != nil {
			cc.ResponseError(err)
			return
		}
		if len(categories) == 0 {
			cc.NoContent()
			return
		}
		cc.Ok(categories)
	}
}
