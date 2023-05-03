package notify

import (
	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/notification"
)

type socketNotificationHandler struct {
	notifyUC notification.UseCase
}

func NewSocketNotificationHandler(notifyUC notification.UseCase) notification.WebSocketHandler {
	return &socketNotificationHandler{notifyUC: notifyUC}
}
