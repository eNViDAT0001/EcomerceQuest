package app_accession

import (
	"github.com/eNViDAT0001/Thesis/Backend/internal/app/domain/app_accession"
	"github.com/eNViDAT0001/Thesis/Backend/internal/user/domain/user"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/jwt"
	"github.com/gin-gonic/gin"
)

type appAccessionHandler struct {
	jwtUC  jwt.UseCase
	userUC user.UseCase
	appUC  app_accession.UseCase
}

func (a *appAccessionHandler) FacebookSSO() func(*gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a *appAccessionHandler) CallbackFacebookSSO() func(*gin.Context) {
	//TODO implement me
	panic("implement me")
}

func NewAppAccessionHandler(jwtUC jwt.UseCase, userUC user.UseCase, appUC app_accession.UseCase) app_accession.HttpHandler {
	return &appAccessionHandler{jwtUC: jwtUC, userUC: userUC, appUC: appUC}
}
