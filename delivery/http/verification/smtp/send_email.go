package smtp

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/smtp/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/entities"
	"github.com/gin-gonic/gin"
)

func (s *smtpHandler) SendEmail() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input io.EmailForm
		if err := cc.BindJSON(&input); err != nil {
			cc.ResponseError(err)
			return
		}

		err := s.smtpUC.SendEmail(newCtx, input)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Send Email Success")
	}
}
func (s *smtpHandler) ReSendUnSendEmail() error {
	newCtx := context.Background()
	return s.smtpUC.ResendEmails(newCtx, entities.EmailWaiting)
}

func (s *smtpHandler) ReSendFailedEmail() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		err := s.smtpUC.ResendEmails(newCtx, entities.EMailFailed)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Send Email Success")
	}
}
