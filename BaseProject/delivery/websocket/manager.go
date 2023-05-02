package websocket

//import (
//	"context"
//	"errors"
//	"github.com/eNViDAT0001/Thesis/Backend/delivery/websocket/chat"
//	"github.com/eNViDAT0001/Thesis/Backend/delivery/websocket/notify"
//	"github.com/eNViDAT0001/Thesis/Backend/external/request"
//	chatStoPKG "github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/chat/storage"
//	chatUC "github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/chat/usecase"
//	"github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/notification/storage"
//	notifyUC "github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/notification/usecase"
//	"github.com/gin-gonic/gin"
//	"log"
//	"net/http"
//	"sync"
//	"time"
//
//	"github.com/gorilla/websocket"
//)
//
//var (
//	/**
//	websocketUpgrader is used to upgrade incomming HTTP requests into a persitent websocket connection
//	*/
//	websocketUpgrader = websocket.Upgrader{
//		// Apply the Origin Checker
//		CheckOrigin:     checkOrigin,
//		ReadBufferSize:  1024,
//		WriteBufferSize: 1024,
//	}
//)
//
//var (
//	ErrEventNotSupported = errors.New("this event type is not supported")
//)
//
//// checkOrigin will check origin and return true if its allowed
//func checkOrigin(r *http.Request) bool {
//
//	// Grab the request origin
//	origin := r.Header.Get("Origin")
//
//	switch origin {
//	// Update this to HTTPS
//	case "https://localhost:8080":
//		return true
//	default:
//		return false
//	}
//}
//
//// Manager is used to hold references to all Clients Registered, and Broadcasting etc
//type Manager struct {
//	clients map[Client]bool
//	// Using a syncMutex here to be able to lcok state before editing clients
//	// Could also use Channels to block
//	sync.RWMutex
//	// handlers are functions that are used to handle Events
//	handlers map[string]EventHandler
//	// otps is a map of allowed OTP to accept connections from
//	otps RetentionMap
//}
//
//// NewManager is used to initalize all the values inside the manager
//func NewManager(ctx context.Context) *Manager {
//	m := &Manager{
//		clients:  make(map[Client]bool),
//		handlers: make(map[string]EventHandler),
//		// Create a new retentionMap that removes Otps older than 5 seconds
//		otps: NewRetentionMap(ctx, 5*time.Second),
//	}
//	m.setupEventHandlers()
//	return m
//}
//
//// setupEventHandlers configures and adds all handlers
//func (m *Manager) setupEventHandlers() {
//	chatSto := chatStoPKG.NewChatStorage()
//	chatUseCase := chatUC.NewChatUseCase(chatSto)
//	chatHandler := chat.NewSocketChatHandler(chatUseCase)
//	notifySto := storage.NewNotificationStorage()
//	notifyUseCase := notifyUC.NewNotificationUseCase(notifySto)
//	notifyHandler := notify.NewSocketNotificationHandler(notifyUseCase)
//
//	m.handlers["OnChatConnect"] = chatHandler.OnConnect
//	m.handlers["OnChatSeenMessage"] = chatHandler.OnSeenMessage
//	m.handlers["OnChatNewMessage"] = chatHandler.OnNewMessage
//	m.handlers["OnChatSendMessage"] = chatHandler.OnSendMessage
//	m.handlers["OnNotifyConnect"] = notifyHandler.OnConnect
//	m.handlers["OnNewNotification"] = notifyHandler.OnNewNotification
//	m.handlers["OnSeenNotification"] = notifyHandler.OnSeenNotification
//}
//
//// routeEvent is used to make sure the correct event goes into the correct handler
//func (m *Manager) routeEvent(event Event, c *Client) error {
//	if handler, ok := m.handlers[event.Type]; ok {
//		if err := handler(event, c); err != nil {
//			return err
//		}
//		return nil
//	} else {
//		return ErrEventNotSupported
//	}
//}
//
//func (m *Manager) ConnectNotifyWS() func(ctx *gin.Context) {
//	return func(ctx *gin.Context) {
//		cc := request.FromContext(ctx)
//		// Begin by upgrading the HTTP request
//		conn, err := websocketUpgrader.Upgrade(cc.Writer, cc.Request, nil)
//		if err != nil {
//			log.Println(err)
//			return
//		}
//		log.Println("New Notify connection")
//
//		// Create New Client
//		client := notify.NewNotifySocket(conn, m)
//		// Add the newly created client to the manager
//		m.AddClient(client)
//
//		go client.ReadMessage()
//		go client.WriteMessage()
//	}
//}
//func (m *Manager) ConnectChatWS() func(ctx *gin.Context) {
//	return func(ctx *gin.Context) {
//		cc := request.FromContext(ctx)
//		// Begin by upgrading the HTTP request
//		conn, err := websocketUpgrader.Upgrade(cc.Writer, cc.Request, nil)
//		if err != nil {
//			log.Println(err)
//			return
//		}
//		log.Println("New Notify connection")
//
//		// Create New Client
//		client := chat.NewChatSocket(conn, m)
//		// Add the newly created client to the manager
//		m.AddClient(client)
//
//		go client.ReadMessage()
//		go client.WriteMessage()
//	}
//}
//
//func (m *Manager) AddClient(client Client) {
//	// Lock so we can manipulate
//	m.Lock()
//	defer m.Unlock()
//
//	// Add Client
//	m.clients[client] = true
//}
//
//func (m *Manager) RemoveClient(client Client) {
//	m.Lock()
//	defer m.Unlock()
//
//	// Check if Client exists, then delete it
//	if _, ok := m.clients[client]; ok {
//		// close connection
//		client.GetConnection().Close()
//		// remove
//		delete(m.clients, client)
//	}
//}
