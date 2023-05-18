package usecase

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/smtp"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/smtp/storage/io"
	"strings"
)

type smtpUseCase struct {
	smtpStorage smtp.Storage
}

func (s *smtpUseCase) ResendEmails(ctx context.Context, status string) (err error) {
	emails, err := s.smtpStorage.GetEmailsByStatus(ctx, status)
	if err != nil {
		return err
	}
	for _, email := range emails {
		toEmails := make([]string, 0)
		bcc := make([]string, 0)
		cc := make([]string, 0)
		attachFiles := make([]string, 0)
		for _, to := range strings.Split(email.ToEmails, ",") {
			toEmails = append(toEmails, to)
		}

		for _, to := range strings.Split(email.Bcc, ",") {
			bcc = append(toEmails, to)
		}
		for _, to := range strings.Split(email.Cc, ",") {
			cc = append(toEmails, to)
		}
		for _, to := range strings.Split(email.AttachFiles, ",") {
			attachFiles = append(toEmails, to)
		}
		emailForm := io.EmailForm{
			ID:          email.ID,
			Subject:     email.Subject,
			Content:     email.Content,
			To:          toEmails,
			Cc:          cc,
			Bcc:         bcc,
			AttachFiles: attachFiles,
		}
		err = s.ReSendEmail(ctx, email.ID, emailForm)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewSmtpUseCase(smtpStorage smtp.Storage,
) smtp.UseCase {
	return &smtpUseCase{
		smtpStorage: smtpStorage,
	}
}
