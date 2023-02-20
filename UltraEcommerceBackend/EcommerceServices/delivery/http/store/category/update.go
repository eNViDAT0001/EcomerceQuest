package category

import (
	"context"
	"strconv"

	"github.com/eNViDAT0001/Thesis/Ecommerce/delivery/http/store/category/convert"
	"github.com/eNViDAT0001/Thesis/Ecommerce/delivery/http/store/category/io"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/request"
	"github.com/gin-gonic/gin"
)

func (s *categoryHandler) UpdateCategory() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input io.UpdateCategoryReq
		if err := cc.BindJSON(&input); err != nil {
			cc.BadRequest(err)
			return
		}

		inputSto, err := convert.UpdateCategoryReqToCategoryForm(&input)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		categoryID, err := strconv.Atoi(cc.Param("category_id"))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		err = s.categoryUC.UpdateCategory(newCtx, uint(categoryID), inputSto)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Update Category success")
	}
}
