package category

import (
	"context"

	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/store/category/convert"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/store/category/io"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
)

func (s *categoryHandler) CreateCategory() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input io.CreateCategoryReq
		if err := cc.BindJSON(&input); err != nil {
			cc.BadRequest(err)
			return
		}

		inputSto, err := convert.CreateCategoryReqToCategoryForm(&input)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		err = s.categoryUC.CreateCategory(newCtx, inputSto)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Create Category success")
	}
}
