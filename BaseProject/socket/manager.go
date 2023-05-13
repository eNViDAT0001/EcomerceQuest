package socket

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	storage2 "github.com/eNViDAT0001/Thesis/Backend/internal/user/domain/user/storage"
	"github.com/eNViDAT0001/Thesis/Backend/internal/user/entities"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/jwt/storage"
	"github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/jwt/usecase"
	"github.com/eNViDAT0001/Thesis/Backend/socket/io"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var websocketUpgrader *websocket.Upgrader

func GetWsServer() *websocket.Upgrader {
	if websocketUpgrader == nil {
		websocketUpgrader = &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(r *http.Request) bool { return true },
		}
	}

	return websocketUpgrader
}

var (
	ErrEventNotSupported = errors.New("this event type is not supported")
)

// Manager is used to hold references to all Clients Registered, and Broadcasting etc
type Manager struct {
	Clients map[string]io.Client
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
		Clients:  make(map[string]io.Client),
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
		newCtx := context.Background()

		tokenString := cc.Param("token")

		jwtSto := storage.NewJwtStorage()
		userSto := storage2.NewUserStorage()
		jwtUC := usecase.NewJwtUseCase(userSto, jwtSto)
		token, err := jwtUC.VerifyToken(newCtx, tokenString)
		cancel := func() {
			cc.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid access_token"})
		}
		if err != nil {
			cancel()
			return
		}
		claims, _ := token.Claims.(jwt.MapClaims)

		userID, _ := strconv.Atoi(claims["user_id"].(string))
		user, err := userSto.GetUserDetailByID(newCtx, uint(userID))

		id, err := strconv.Atoi(cc.Param("user_id"))
		if err != nil {
			cc.ResponseError(err)
			return
		}
		if *user.Type != entities.Admin {
			if id != int(user.ID) {
				cc.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid access_token"})
				return
			}
		}

		// Begin by upgrading the HTTP request
		conn, err := GetWsServer().Upgrade(cc.Writer, cc.Request, nil)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("New Client connect")
		oldClient, ok := m.Clients[cc.Param("user_id")]
		if ok {
			m.RemoveClient(oldClient)
		}

		client := NewSocketClient(conn, m, strconv.Itoa(userID))
		m.AddClient(client)

		go client.ReadMessage()
		go client.WriteMessage()
	}
}

func (m *Manager) EmitNotify(event string, data interface{}, targetID string) error {
	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	m.Lock()
	if _, ok := m.Clients[targetID]; ok {
		m.Clients[targetID].AddEvent(io.Event{
			Type:    event,
			Payload: payload,
		})
	}
	m.Unlock()

	return nil
}
