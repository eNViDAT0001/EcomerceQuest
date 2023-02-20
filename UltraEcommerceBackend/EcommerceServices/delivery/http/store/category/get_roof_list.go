package category

import (
	"context"

	"github.com/eNViDAT0001/Thesis/Ecommerce/external/request"
	"github.com/gin-gonic/gin"
)

func (s *categoryHandler) GetCategoryRoofList() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		categories, err := s.categoryUC.GetCategoryRoofList(newCtx)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok(categories)
	}
}
