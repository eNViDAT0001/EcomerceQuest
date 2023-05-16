package usecase

import "github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/smtp"

type smtpUseCase struct {
	smtpStorage smtp.Storage
}

func NewSmtpUseCase(smtpStorage smtp.Storage,
) smtp.UseCase {
	return &smtpUseCase{
		smtpStorage: smtpStorage,
	}
}
