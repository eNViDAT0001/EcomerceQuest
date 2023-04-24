package chat

import (
	"github.com/eNViDAT0001/Thesis/Backend/internal/real_time/domain/notification"
)

type notificationHandler struct {
	notifyUC notification.UseCase
}

func NewNotificationHandler(notifyUC notification.UseCase) notification.HttpHandler {
	return &notificationHandler{notifyUC: notifyUC}
}
