package smtp

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/verification/smtp/convert"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/verification/smtp/io"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	io2 "github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/request_contact/storage/io"
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
		requestContact := io2.CreateRequestContact{
			Subject:       input.Subject,
			Content:       input.Descriptions,
			AttachedFile:  input.AttachedFiles,
			Type:          input.Type,
			Seen:          false,
		}

		id, err := s.smtpUC.CreateEmail(newCtx, inputUC, requestContact)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok(id)
	}
}
