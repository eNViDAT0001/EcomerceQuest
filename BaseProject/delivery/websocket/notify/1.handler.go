package notify

//import (
//	"github.com/eNViDAT0001/Thesis/Backend/delivery/websocket"
//	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/notification"
//)
//
//type socketNotificationHandler struct {
//	notifyUC notification.UseCase
//}
//
//func (s socketNotificationHandler) OnConnect(event websocket.Event, client *websocket.Client) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (s socketNotificationHandler) OnNewNotification(event websocket.Event, client *websocket.Client) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (s socketNotificationHandler) OnSeenNotification(event websocket.Event, client *websocket.Client) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func NewSocketNotificationHandler(notifyUC notification.UseCase) notification.WebSocketHandler {
//	return &socketNotificationHandler{notifyUC: notifyUC}
//}
