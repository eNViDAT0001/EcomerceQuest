package socket

import (
	"encoding/json"
	"github.com/eNViDAT0001/Thesis/Backend/socket/io"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type socketClient struct {
	ID string

	// the websocket connection
	connection *websocket.Conn

	// manager is the manager used to manage the client
	manager *Manager
	// egress is used to avoid concurrent writes on the WebSocket
	egress chan io.Event
	// chatroom is used to know what room user is in
}

func (s *socketClient) AddEvent(event io.Event) {
	s.egress <- event
}
func (s *socketClient) GetID() interface{} {
	return s.ID
}

func (s *socketClient) GetConnection() *websocket.Conn {
	return s.connection
}

func (s *socketClient) WriteMessage() {
	ticker := time.NewTicker(io.PingInterval)
	defer func() {
		ticker.Stop()
		manager.RemoveClient(s)
	}()

	for {
		select {
		case message, ok := <-s.egress:
			if !ok {
				// Manager has closed this connection channel, so communicate that to frontend
				if err := s.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					// Log that the connection is closed and the reason
					log.Println("connection closed: ", err)
				}
				// Return to close the goroutine
				return
			}

			data, err := json.Marshal(message)
			if err != nil {
				log.Println(err)
				return // closes the connection, should we really
			}
			// Write a Regular text message to the connection
			if err = s.connection.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Println(err)
			}
			log.Println("sent message")
		case <-ticker.C:
			// Send the Ping
			if err := s.connection.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				log.Println("writemsg: ", err)
				return // return to break this goroutine triggeing cleanup
			}
		}
	}
}

func (s *socketClient) ReadMessage() {
	defer func() {
		manager.RemoveClient(s)
	}()

	// Set Max Size of Messages in Bytes
	s.connection.SetReadLimit(512)
	// Configure Wait time for Pong response, use Current time + pongWait
	// This has to be done here to set the first initial timer.
	if err := s.connection.SetReadDeadline(time.Now().Add(io.PongWait)); err != nil {
		log.Println(err)
		return
	}
	// Configure how to handle Pong responses
	s.connection.SetPongHandler(s.Pong)

	// Loop Forever
	for {
		// ReadMessage is used to read the next message in queue
		// in the connection
		_, payload, err := s.connection.ReadMessage()

		if err != nil {
			// If Connection is closed, we will Recieve an error here
			// We only want to log Strange errors, but simple Disconnection
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error reading message: %v", err)
			}
			break // Break the loop to close conn & Cleanup
		}
		// Marshal incoming data into a Event struct
		var request io.Event
		if err := json.Unmarshal(payload, &request); err != nil {
			log.Printf("error marshalling message: %v", err)
			break // Breaking the connection here might be harsh xD
		}
		// Route the Event
		if err := s.manager.routeEvent(request, s); err != nil {
			log.Println("Error handeling Message: ", err)
		}
	}
}

func (s *socketClient) Pong(pongMsg string) error {
	return s.connection.SetReadDeadline(time.Now().Add(io.PongWait))
}

func NewSocketClient(conn *websocket.Conn, manager *Manager) io.Client {
	return &socketClient{
		connection: conn,
		manager:    manager,
		egress:     make(chan io.Event),
	}
}
