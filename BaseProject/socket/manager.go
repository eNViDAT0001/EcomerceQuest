package socket

import (
	"context"
	"errors"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/eNViDAT0001/Thesis/Backend/socket/io"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

var websocketUpgrader *websocket.Upgrader

func GetWsServer() *websocket.Upgrader {
	if websocketUpgrader == nil {
		websocketUpgrader = &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		}
	}

	return websocketUpgrader
}

var (
	ErrEventNotSupported = errors.New("this event type is not supported")
)

// Manager is used to hold references to all Clients Registered, and Broadcasting etc
type Manager struct {
	Clients map[interface{}]io.Client
	// Using a syncMutex here to be able to lcok state before editing clients
	// Could also use Channels to block
	sync.RWMutex
	// handlers are functions that are used to handle Events
	Handlers map[string]io.EventHandler
}

var manager *Manager

func GetManager() *Manager {
	if manager == nil {
		manager = NewManager(context.Background())
	}
	return manager
}

// NewManager is used to initalize all the values inside the manager
func NewManager(ctx context.Context) *Manager {
	m := &Manager{
		Clients:  make(map[interface{}]io.Client),
		Handlers: make(map[string]io.EventHandler),
	}
	m.setupEventHandlers()
	return m
}

// setupEventHandlers configures and adds all handlers
func (m *Manager) setupEventHandlers() {
	route := initSocketCollection()
	m.Handlers[io.ChatUnseenMessage] = route.chatSocket.UnseenMessages()
	m.Handlers[io.ChatSeenMessage] = route.chatSocket.SeenMessage()
	m.Handlers[io.ChatNewMessage] = route.chatSocket.NewMessage()
	m.Handlers[io.NotifyUnseenMessage] = route.notificationSocket.UnseenNotification()
	m.Handlers[io.NotificationNew] = route.notificationSocket.NewNotification()
}

// routeEvent is used to make sure the correct event goes into the correct handler
func (m *Manager) routeEvent(event io.Event, c io.Client) error {
	if handler, ok := m.Handlers[event.Type]; ok {
		outGoingEvent, targetID, err := handler(&event, c)
		if err != nil {
			return err
		}
		GetManager().Clients[targetID].AddEvent(outGoingEvent)
		return nil
	} else {
		return ErrEventNotSupported
	}
}

func (m *Manager) AddClient(client io.Client) {
	// Lock so we can manipulate
	m.Lock()
	defer m.Unlock()

	// Add Client
	m.Clients[client.GetID()] = client
}

func (m *Manager) RemoveClient(client io.Client) {
	m.Lock()
	defer m.Unlock()

	// Check if Client exists, then delete it
	if _, ok := m.Clients[client.GetID()]; ok {
		// close connection
		err := client.GetConnection().Close()
		if err != nil {
			log.Fatal("Error closing connection: ", err)
		}
		// remove
		delete(m.Clients, client.GetID())
	}
}
func (m *Manager) ConnectChatWS() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		cc := request.FromContext(ctx)
		// Begin by upgrading the HTTP request
		conn, err := GetWsServer().Upgrade(cc.Writer, cc.Request, nil)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("New Client connect")

		socketManager := GetManager()
		client := NewSocketClient(conn, socketManager)
		socketManager.AddClient(client)

		go client.ReadMessage()
		go client.WriteMessage()
	}
}
