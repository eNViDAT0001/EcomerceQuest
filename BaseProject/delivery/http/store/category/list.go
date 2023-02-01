package category

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
)

func (s *categoryHandler) ListCategories() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		categories, err := s.categoryUC.ListCategories(newCtx)
		if err != nil {
			cc.ResponseError(err)
		}
		cc.Ok(categories)
	}
}
