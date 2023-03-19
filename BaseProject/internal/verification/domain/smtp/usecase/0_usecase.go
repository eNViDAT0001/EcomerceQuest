package usecase

import "github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/smtp"

type smtpUseCase struct {
}

func NewSmtpUseCase() smtp.UseCase {
	return &smtpUseCase{}
}
