package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/event_background"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_viper"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/smtp/storage/io"
	io2 "github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/smtp/usecase/io"
)

func (s *smtpUseCase) CreateEmail(ctx context.Context, email io.CreateEmail) (id uint, err error) {
	event_background.GetBackGroundJobs().Group <- event_background.NewGroup(false, event_background.NewJob(func(ctx context.Context) error {
		sEmail := io2.EmailForm{
			Subject:     email.Subject,
			Content:     email.Descriptions,
			To:          []string{wrap_viper.GetViper().GetString("MAIL_SHOP")},
			Cc:          nil,
			Bcc:         nil,
			AttachFiles: nil,
		}
		return s.SendEmail(ctx, sEmail)
	}))

	return s.smtpStorage.CreateEmail(ctx, email)
}
