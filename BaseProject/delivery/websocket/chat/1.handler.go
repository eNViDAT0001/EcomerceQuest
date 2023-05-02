package chat

//import (
//	"github.com/eNViDAT0001/Thesis/Backend/delivery/websocket"
//	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/chat"
//)
//
//type socketChatHandler struct {
//	chatUC chat.UseCase
//}
//
//func (s socketChatHandler) OnConnect(event websocket.Event, client *websocket.Client) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (s socketChatHandler) OnNewMessage(event websocket.Event, client *websocket.Client) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (s socketChatHandler) OnSeenMessage(event websocket.Event, client *websocket.Client) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (s socketChatHandler) OnSendMessage(event websocket.Event, client *websocket.Client) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func NewSocketChatHandler(chatUC chat.UseCase) chat.WebSocketHandler {
//	return &socketChatHandler{chatUC: chatUC}
//}
