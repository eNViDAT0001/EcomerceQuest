package convert

import (
	ioHandler "github.com/eNViDAT0001/Thesis/Backend/delivery/http/verification/smtp/io"
	ioSto "github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/smtp/storage/io"
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
func CreateEmailReqToSendEmailInput(input *ioHandler.EmailReq) (ioSto.EmailForm, error) {
	var result ioSto.EmailForm
	result = ioSto.EmailForm{
		Subject:     input.Subject,
		Content:     input.Descriptions,
		To:          []string{input.Email},
		Cc:          nil,
		Bcc:         nil,
		AttachFiles: nil,
	}
	return result, nil
}
