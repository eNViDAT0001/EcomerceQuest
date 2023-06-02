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


		toEmails := strings.Split(email.ToEmails, ",")
		if len(toEmails) == 1 && toEmails[0] == "" {
			toEmails = nil
		}
		bcc := strings.Split(email.Bcc, ",")
		if len(bcc) == 1 && bcc[0] == "" {
			bcc = nil
		}
		cc := strings.Split(email.Cc, ",")
		if len(cc) == 1 && cc[0] == "" {
			cc = nil
		}
		attachFiles := strings.Split(email.AttachFiles, ",")
		if len(attachFiles) == 1 && attachFiles[0] == "" {
			attachFiles = nil
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
