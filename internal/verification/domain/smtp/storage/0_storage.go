package storage

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_gorm"
	"github.com/eNViDAT0001/Thesis/Backend/external/wrap_viper"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/smtp"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/smtp/storage/io"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/entities"
	"strings"
)

var wViper = wrap_viper.GetViper()

type smtpStorage struct {
}

func (s smtpStorage) StoreEmail(ctx context.Context, email io.EmailForm) (io.EmailForm, error) {
	toEmails := ""
	for _, v := range email.To {
		toEmails += v + ","
	}
	bcc := ""
	for _, v := range email.Bcc {
		bcc += v + ","
	}
	cc := ""
	for _, v := range email.Cc {
		cc += v + ","
	}
	attachFiles := ""
	for _, v := range email.AttachFiles {
		attachFiles += v + ","
	}

	createSmtp := io.CreateSmtp{
		ID:          0,
		Subject:     email.Subject,
		Content:     email.Content,
		ToEmails:    strings.TrimSuffix(toEmails, ","),
		Cc:          strings.TrimSuffix(cc, ","),
		Bcc:         strings.TrimSuffix(bcc, ","),
		AttachFiles: strings.TrimSuffix(attachFiles, ","),
	}
	err := wrap_gorm.GetDB().Table(entities.Smtp{}.TableName()).Create(&createSmtp).Error
	email.ID = createSmtp.ID
	return email, err
}

func (s smtpStorage) UpdateSentByID(ctx context.Context, id uint, status string) (err error) {
	return wrap_gorm.GetDB().Model(&entities.Smtp{}).Where("id = ?", id).Update("status", status).Error
}

func (s smtpStorage) GetEmailsByStatus(ctx context.Context, status string) (emails []entities.Smtp, err error) {
	err = wrap_gorm.GetDB().Model(&entities.Smtp{}).Where("status = ?", status).Find(&emails).Error
	return emails, err
}

func (s smtpStorage) DeleteUnsentEmails(ctx context.Context) (err error) {
	return wrap_gorm.GetDB().Model(&entities.Smtp{}).Where("status = ?", entities.EmailSuccess).Delete(&entities.Email{}).Error
}

func NewSmtpStorage() smtp.Storage {
	return &smtpStorage{}
}
