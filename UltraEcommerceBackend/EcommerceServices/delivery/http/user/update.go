package user

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Ecommerce/delivery/http/user/convert"
	"github.com/eNViDAT0001/Thesis/Ecommerce/delivery/http/user/io"
	"github.com/gin-gonic/gin"
	"strconv"
)

func UpdateUser[T io.UpdateUserReq](cc *gin.Context, newCtx context.Context, s *userHandler, updateForm T) error {

	if err := cc.BindJSON(&updateForm); err != nil {
		return err
	}

	inputSto, err := convert.UpdateReqToUpdateUserRepo(&updateForm)
	if err != nil {
		return err
	}

	id, _ := strconv.Atoi(cc.Param("user_id"))

	err = s.userUC.UpdateUser(newCtx, uint32(id), inputSto)
	if err != nil {
		return err
	}

	return nil
}
