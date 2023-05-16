package provider

import (
	"context"
	"strconv"

	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/store/provider/convert"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/store/provider/io"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
)

func (s providerHandler) CreateProvider() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input io.CreateProviderReq
		if err := cc.BindJSON(&input); err != nil {
			cc.BadRequest(err)
			return
		}

		userID, err := strconv.Atoi(cc.Param("user_id"))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		input.UserID = uint(userID)

		inputSto, err := convert.ProviderCreateReqToProviderForm(&input)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		result, err := s.providerUC.CreateProvider(newCtx, inputSto)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok(result)
	}
}
