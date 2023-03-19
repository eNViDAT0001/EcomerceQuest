package app_accession

import (
	"github.com/eNViDAT0001/Thesis/Backend/internal/app/domain/app_accession"
	"github.com/eNViDAT0001/Thesis/Backend/internal/user/domain/user"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/jwt"
)

type appAccessionHandler struct {
	jwtUC  jwt.UseCase
	userUC user.UseCase
	appUC  app_accession.UseCase
}

func NewAppAccessionHandler(jwtUC jwt.UseCase, userUC user.UseCase, appUC app_accession.UseCase) app_accession.HttpHandler {
	return &appAccessionHandler{jwtUC: jwtUC, userUC: userUC, appUC: appUC}
}
