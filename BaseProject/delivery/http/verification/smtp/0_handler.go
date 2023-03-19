package smtp

import (
	"github.com/eNViDAT0001/Thesis/Backend/internal/user/domain/user"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/jwt"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/smtp"
	"github.com/gin-gonic/gin"
)

type smtpHandler struct {
	jwtUC  jwt.UseCase
	userUC user.UseCase
	smtpUC smtp.UseCase
}

func (s *smtpHandler) SendEmail() func(*gin.Context) {
	//TODO implement me
	panic("implement me")
}

func NewSmtpHandler(
	jwtUC jwt.UseCase,
	userUC user.UseCase,
	smtpUC smtp.UseCase,
) smtp.HttpHandler {
	return &smtpHandler{
		jwtUC:  jwtUC,
		userUC: userUC,
		smtpUC: smtpUC,
	}
}
