package banner

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s favoriteHandler) DeleteFavorite() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()
		userID, _ := strconv.Atoi(cc.Param("user_id"))
		providerID, _ := strconv.Atoi(cc.Param("provider_id"))
		err := s.favoriteUC.DeleteFavorite(newCtx, uint(userID), uint(providerID))
		if err != nil {
			cc.ResponseError(err)
		}
		cc.Ok("Remove Favorite Success")
	}
}
