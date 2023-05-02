package chat

//import (
//	deliverysocket "github.com/eNViDAT0001/Thesis/Backend/delivery/websocket"
//	"github.com/gorilla/websocket"
//	"log"
//	"time"
//)
//
//type ChatSocket struct {
//	// the websocket connection
//	connection *websocket.Conn
//
//	// manager is the manager used to manage the client
//	manager *deliverysocket.Manager
//	// egress is used to avoid concurrent writes on the WebSocket
//	egress chan deliverysocket.Event
//	// chatroom is used to know what room user is in
//	chatroom string
//}
//
//func (s *ChatSocket) GetConnection() *websocket.Conn {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (s *ChatSocket) WriteMessage() {
//	ticker := time.NewTicker(deliverysocket.PingInterval)
//	defer func() {
//		ticker.Stop()
//		s.manager.RemoveClient(s)
//	}()
//
//}
//
//func (s *ChatSocket) ReadMessage() {
//	defer func() {
//		s.manager.RemoveClient(s)
//	}()
//}
//
//func (s *ChatSocket) Pong() error{
//	log.Println("pong")
//	return s.connection.SetReadDeadline(time.Now().Add(deliverysocket.PongWait))
//}
//
//func NewChatSocket(conn *websocket.Conn, manager *deliverysocket.Manager) deliverysocket.Client {
//	return &ChatSocket{
//		connection: conn,
//		manager:    manager,
//		egress:     make(chan deliverysocket.Event),
//	}
//}
