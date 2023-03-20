package user

import (
	"github.com/eNViDAT0001/Thesis/Backend/internal/user/domain/user"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/jwt"
)

type userHandler struct {
	userUC user.UseCase
	jwtUC  jwt.UseCase
}

func NewUserHandler(userUC user.UseCase, jwtUC jwt.UseCase) user.HttpHandler {
	return &userHandler{
		userUC: userUC,
		jwtUC:  jwtUC,
	}
}
