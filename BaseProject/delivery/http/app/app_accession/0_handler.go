package app_accession

import (
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
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

func (a appAccessionHandler) LoginByGoogle() gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO implement me
		cc := request.FromContext(c)
		cc.Ok("Chức năng này đang làm bruh ơii")
	}
}

func (a appAccessionHandler) LoginByFacebook() gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO implement me
		cc := request.FromContext(c)
		cc.Ok("Chức năng này đang làm bruh ơii")
	}
}

func NewAppAccessionHandler(jwtUC jwt.UseCase, userUC user.UseCase, appUC app_accession.UseCase) app_accession.HttpHandler {
	return &appAccessionHandler{jwtUC: jwtUC, userUC: userUC, appUC: appUC}
}
