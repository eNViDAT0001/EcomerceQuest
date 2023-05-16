package convert

import (
	ioHandler "github.com/eNViDAT0001/Thesis/Backend/delivery/http/verification/smtp/io"
	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/smtp/storage/io"
	ioUC "github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/smtp/usecase/io"
	"github.com/jinzhu/copier"
)

func CreateEmailReqToCreateEmailInput(input *ioHandler.EmailReq) (ioSto.CreateEmail, error) {
	var result ioSto.CreateEmail
	err := copier.Copy(&result, input)
	if err != nil {
		return result, err
	}
	return result, nil
}
func CreateEmailReqToSendEmailInput(input *ioHandler.EmailReq) (ioUC.EmailForm, error) {
	var result ioUC.EmailForm
	result = ioUC.EmailForm{
		Subject:     input.Subject,
		Content:     input.Descriptions,
		To:          []string{input.Email},
		Cc:          nil,
		Bcc:         nil,
		AttachFiles: nil,
	}
	return result, nil
}
