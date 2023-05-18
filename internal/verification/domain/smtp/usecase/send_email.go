package usecase

import (
	"context"
	"fmt"
	"github.com/eNViDAT0001/Thesis/Backend/external/event_background"
	"github.com/eNViDAT0001/Thesis/Backend/external/mail_template"
	io2 "github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/smtp/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/entities"
	"github.com/jordan-wright/email"
	"net/smtp"
)

const (
	smtpAuthAddress   = "smtp.gmail.com"
	smtpServerAddress = "smtp.gmail.com:587"
)

func (s *smtpUseCase) SendEmail(ctx context.Context, input io2.EmailForm) error {
	storedEmail, err := s.smtpStorage.StoreEmail(ctx, input)
	if err != nil {
		return err
	}
	sender := mail_template.GetGmailSender()

	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", sender.Name, sender.FromEmailAddress)
	e.Subject = "[No Reply] " + input.Subject
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

	event_background.AddBackgroundJobs(false, event_background.NewJob(func(ctx context.Context) error {
		err = e.Send(smtpServerAddress, smtpAuth)
		if err != nil {
			err = s.smtpStorage.UpdateSentByID(ctx, storedEmail.ID, entities.EMailFailed)
			return fmt.Errorf("failed to send email: %w", err)
		}
		return s.smtpStorage.UpdateSentByID(ctx, storedEmail.ID, entities.EmailSuccess)
	}))

	return nil
}
func (s *smtpUseCase) ReSendEmail(ctx context.Context, id uint, input io2.EmailForm) error {

	sender := mail_template.GetGmailSender()

	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", sender.Name, sender.FromEmailAddress)
	e.Subject = "[No Reply] " + input.Subject
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

	event_background.AddBackgroundJobs(false, event_background.NewJob(func(ctx context.Context) error {
		err := e.Send(smtpServerAddress, smtpAuth)
		if err != nil {
			err = s.smtpStorage.UpdateSentByID(ctx, id, entities.EMailFailed)
			return fmt.Errorf("failed to send email: %w", err)
		}
		return s.smtpStorage.UpdateSentByID(ctx, id, entities.EmailSuccess)
	}))

	return nil
}
