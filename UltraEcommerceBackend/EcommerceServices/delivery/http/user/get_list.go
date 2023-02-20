package user

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/delivery/http/user/convert"
	"github.com/eNViDAT0001/Thesis/Ecommerce/delivery/http/user/io"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/paging/paging_params"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/request"
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/user/entities"
	"github.com/gin-gonic/gin"
)

func (s userHandler) GetUserList() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var paginator io.GetUserListParams
		err := cc.BindQuery(&paginator)
		if err != nil {
			cc.BadRequest(err)
			return
		}

		search := cc.QueryArray("search[]")
		fields := cc.QueryArray("fields[]")
		sort := cc.QueryArray("sorts[]")

		paginator.Filter = paging_params.NewFilterBuilder().
			WithSearch(search).
			WithFields(fields).
			WithSorts(sort).
			Build()

		inValidField, val := paging_params.ValidateFilter(paginator.Filter, entities.User{})
		if len(inValidField) > 0 {
			cc.ResponseError(request.NewBadRequestError(inValidField, val, "invalid key and value"))
			return
		}

		inputUC, err := convert.ParamsToGetListInput(&paginator)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		users, err := s.userUC.GetUserList(newCtx, inputUC)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		result, err := convert.ArrUserEntityToArrUserRes(users)
		if err != nil {
			cc.ResponseError(err)
			return
		}
		if len(result) == 0 {
			cc.NoContent()
			return
		}

		cc.OkPaging(inputUC, result)
	}
}
