package websocket

//import (
//	"github.com/gorilla/websocket"
//	"time"
//)
//
//type Client interface {
//	WriteMessage()
//	ReadMessage()
//	GetConnection() *websocket.Conn
//	Pong() error
//}
//
//var (
//	// pongWait is how long we will await a pong response from client
//	PongWait = 10 * time.Second
//	// pingInterval has to be less than pongWait, We cant multiply by 0.9 to get 90% of time
//	// Because that can make decimals, so instead *9 / 10 to get 90%
//	// The reason why it has to be less than PingRequency is becuase otherwise it will send a new Ping before getting response
//	PingInterval = (PongWait * 9) / 10
//)
