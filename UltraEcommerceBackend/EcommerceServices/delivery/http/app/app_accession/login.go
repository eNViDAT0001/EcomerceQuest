package app_accession

import (
	"context"
	"strconv"

	"github.com/eNViDAT0001/Thesis/Ecommerce/delivery/http/app/app_accession/io"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/request"
	ioJwtSto "github.com/eNViDAT0001/Thesis/Ecommerce/internal/verification/domain/jwt/storage/io"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (a appAccessionHandler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input io.LoginForm
		if err := cc.BindJSON(&input); err != nil {
			cc.BadRequest(err)
			return
		}
		user, err := a.appUC.Login(newCtx, input.Username, input.Password)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				cc.ResponseError(request.NewUnauthorizedError("User", "", "Username Or Password not match"))
				return
			}
			cc.ResponseError(err)
			return
		}
		tokenForm := ioJwtSto.GenerateTokenInput{
			UserID:   strconv.Itoa(int(user.ID)),
			Username: *user.Username,
			Salt:     *user.Salt,
		}
		token, err := a.jwtUC.GenerateToken(tokenForm)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		result := map[string]interface{}{
			"UserID": user.ID,
			"Token":  token,
		}

		cc.Ok(result)
	}
}
