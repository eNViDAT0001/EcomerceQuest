package app_accession

import (
	"context"
	"strconv"

	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/user/convert"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/user/io"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	ioJwtSto "github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/jwt/storage/io"
	"github.com/gin-gonic/gin"
)

func (a appAccessionHandler) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input io.CreateUserReq
		if err := cc.BindJSON(&input); err != nil {
			cc.BadRequest(err)
			return
		}

		inputSto, err := convert.CreateReqToCreateUserInput(&input)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		userID, err := a.userUC.CreateUser(newCtx, inputSto)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		tokenForm := ioJwtSto.GenerateTokenInput{
			UserID:   strconv.Itoa(int(userID)),
			Username: inputSto.Username,
			Salt:     inputSto.Salt,
		}

		token, err := a.jwtUC.GenerateToken(nil, tokenForm)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		result := map[string]interface{}{
			"UserID": userID,
			"Token":  token,
		}
		cc.Ok(result)
	}
}
