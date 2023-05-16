package socket

import (
	chatHandlerPKG "github.com/eNViDAT0001/Thesis/Backend/delivery/websocket/chat"
	notifyHandlerPKG "github.com/eNViDAT0001/Thesis/Backend/delivery/websocket/notify"
	userStoPKG "github.com/eNViDAT0001/Thesis/Backend/internal/user/domain/user/storage"

	chatPKG "github.com/eNViDAT0001/Thesis/Backend/internal/chat/domain/chat"
	chatStoPKG "github.com/eNViDAT0001/Thesis/Backend/internal/chat/domain/chat/storage"
	chatUCPKG "github.com/eNViDAT0001/Thesis/Backend/internal/chat/domain/chat/usecase"

	notificationPKG "github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/notification"
	notificationStoPKG "github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/notification/storage"
	notificationUCPKG "github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/notification/usecase"

	"github.com/google/wire"
)

var IteratorCollection = wire.NewSet(
	chatHandlerPKG.NewSocketChatHandler,
	chatUCPKG.NewChatUseCase,
	chatStoPKG.NewChatStorage,
	userStoPKG.NewUserStorage,

	notifyHandlerPKG.NewSocketNotificationHandler,
	notificationUCPKG.NewNotificationUseCase,
	notificationStoPKG.NewNotificationStorage,

	NewSocketCollection,
)

type WebSocketCollection struct {
	chatSocket         chatPKG.WebSocketHandler
	notificationSocket notificationPKG.WebSocketHandler
}

func NewSocketCollection(
	chatSocket chatPKG.WebSocketHandler,
	notificationSocket notificationPKG.WebSocketHandler,

) *WebSocketCollection {
	return &WebSocketCollection{
		chatSocket:         chatSocket,
		notificationSocket: notificationSocket,
	}
}
