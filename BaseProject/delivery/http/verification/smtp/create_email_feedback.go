package smtp

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/verification/smtp/convert"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/verification/smtp/io"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
)

func (s *smtpHandler) CreateEmailFeedback() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input io.EmailReq
		if err := cc.BindJSON(&input); err != nil {
			cc.ResponseError(err)
			return
		}

		inputUC, err := convert.CreateEmailReqToCreateEmailInput(&input)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		id, err := s.smtpUC.CreateEmail(newCtx, inputUC)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok(id)
	}
}
