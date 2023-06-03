package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_viper"
	ioRequest "github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/request_contact/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/smtp/storage/io"
)

func (s *smtpUseCase) CreateEmail(ctx context.Context, email io.CreateEmail, request ioRequest.CreateRequestContact) (id uint, err error) {
	sEmail := io.EmailForm{
		Subject:     email.Subject,
		Content:     email.Descriptions,
		To:          []string{wrap_viper.GetViper().GetString("MAIL_SHOP")},
		Cc:          nil,
		Bcc:         nil,
		AttachFiles: nil,
	}
	_, err = s.requestStorage.CreateRequest(ctx, request)
	if err != nil {
		return 0, err
	}

	err = s.SendEmail(ctx, sEmail)
	if err != nil {
		return 0, err
	}

	return s.smtpStorage.CreateEmail(ctx, email)
}
