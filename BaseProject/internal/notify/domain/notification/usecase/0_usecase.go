package usecase

import (
	notification2 "github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/notification"
)

type notificationUseCase struct {
	notifySto notification2.Storage
}

func NewNotificationUseCase(notifySto notification2.Storage) notification2.UseCase {
	return &notificationUseCase{notifySto: notifySto}
}
