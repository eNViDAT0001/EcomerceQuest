package user

import (
	"context"
	"strconv"

	"github.com/eNViDAT0001/Thesis/Ecommerce/external/request"
	"github.com/gin-gonic/gin"
)

func (s userHandler) DeleteUserByID() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		userID, _ := strconv.Atoi(cc.Param("user_id"))

		IDs := []uint{uint(userID)}
		err := s.userUC.DeleteUserByIDs(newCtx, IDs)

		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Delete User success")
	}
}
