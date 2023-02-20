package category

import (
	"context"
	"strconv"

	"github.com/eNViDAT0001/Thesis/Ecommerce/external/request"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (s *categoryHandler) GetCategoryParentsTreeWithCategoryID() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		categoryID, err := strconv.Atoi(cc.Param("category_id"))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		categoryTree, err := s.categoryUC.GetCategoryParentsTreeWithCategoryID(newCtx, uint(categoryID))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				cc.NoContent()
				return
			}
			cc.ResponseError(err)
			return
		}

		cc.Ok(categoryTree)
	}
}
