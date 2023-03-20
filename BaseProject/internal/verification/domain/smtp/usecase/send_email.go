package usecase

import (
	"context"
	"fmt"
	"github.com/eNViDAT0001/Thesis/Backend/external/mail_template"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/smtp/usecase/io"
	"github.com/jordan-wright/email"
	"net/smtp"
)

const (
	smtpAuthAddress   = "smtp.gmail.com"
	smtpServerAddress = "smtp.gmail.com:587"
)

func (s *smtpUseCase) SendEmail(ctx context.Context, input io.EmailForm) error {
	sender := mail_template.GetGmailSender()

	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", sender.Name, sender.FromEmailAddress)
	e.Subject = input.Subject
	e.HTML = []byte(input.Content)
	e.To = input.To
	e.Cc = input.Cc
	e.Bcc = input.Bcc

	for _, f := range input.AttachFiles {
		_, err := e.AttachFile(f)
		if err != nil {
			return fmt.Errorf("failed to attach file %s: %w", f, err)
		}
	}

	smtpAuth := smtp.PlainAuth("", sender.FromEmailAddress, sender.FromEmailPassword, smtpAuthAddress)
	return e.Send(smtpServerAddress, smtpAuth)
}
