package banner

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/store/favorite/io"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
)

func (s favoriteHandler) AddFavorite() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()
		var input io.CreateFavoriteReq
		if err := cc.BindJSON(&input); err != nil {
			cc.ResponseError(err)
			return
		}

		err := s.favoriteUC.AddFavorite(newCtx, input.UserID, input.ProviderID)
		if err != nil {
			cc.ResponseError(err)
			return
		}
		cc.Ok("Add Favorite Success")
	}
}
