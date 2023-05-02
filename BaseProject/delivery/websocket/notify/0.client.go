package notify

//import (
//	deliverysocket "github.com/eNViDAT0001/Thesis/Backend/delivery/websocket"
//	"github.com/gorilla/websocket"
//	"log"
//	"time"
//)
//
//type NotifySocket struct {
//	// the websocket connection
//	connection *websocket.Conn
//
//	// manager is the manager used to manage the client
//	manager *deliverysocket.Manager
//	// egress is used to avoid concurrent writes on the WebSocket
//	egress chan deliverysocket.Event
//}
//
//func (s *NotifySocket) GetConnection() *websocket.Conn {
//	return s.connection
//}
//
//func (s *NotifySocket) WriteMessage() {
//	ticker := time.NewTicker(deliverysocket.PingInterval)
//	defer func() {
//		ticker.Stop()
//		s.manager.RemoveClient(s)
//	}()
//}
//
//func (s *NotifySocket) ReadMessage() {
//	defer func() {
//		s.manager.RemoveClient(s)
//	}()
//
//}
//
//func (s *NotifySocket) Pong() error {
//	log.Println("pong")
//	return s.connection.SetReadDeadline(time.Now().Add(deliverysocket.PongWait))
//}
//
//func NewNotifySocket(conn *websocket.Conn, manager *deliverysocket.Manager) deliverysocket.Client {
//	return &NotifySocket{
//		connection: conn,
//		manager:    manager,
//		egress:     make(chan deliverysocket.Event),
//	}
//}
