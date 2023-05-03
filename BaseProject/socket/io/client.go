package io

import (
	"github.com/gorilla/websocket"
	"time"
)

type Client interface {
	GetID() interface{}
	AddEvent(event Event)
	WriteMessage()
	ReadMessage()
	GetConnection() *websocket.Conn
	Pong(pongMsg string) error
}

var (
	PongWait     = 10 * time.Second
	PingInterval = (PongWait * 9) / 10
)
