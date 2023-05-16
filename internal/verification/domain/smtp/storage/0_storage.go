package storage

import (
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_viper"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/smtp"
)

var wViper = wrap_viper.GetViper()

type smtpStorage struct {
}

func NewSmtpStorage() smtp.Storage {
	return &smtpStorage{}
}
