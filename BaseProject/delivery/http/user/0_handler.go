package user

import (
	"github.com/eNViDAT0001/Thesis/Backend/internal/user/domain/user"
)

type userHandler struct {
	userUC user.UseCase
}

func NewUserHandler(userUC user.UseCase) user.HttpHandler {
	return &userHandler{
		userUC: userUC,
	}
}
