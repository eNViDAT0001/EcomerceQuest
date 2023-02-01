package provider

import (
	"context"
	"errors"
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

		userID, _ := strconv.Atoi(cc.Param("user_id"))
		if uint(userID) != input.UserID {
			cc.BadRequest(errors.New("user id does not match"))
			return
		}

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
