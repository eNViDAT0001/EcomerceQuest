package websocket

import (
	address "github.com/eNViDAT0001/Thesis/Backend/internal/address/domain/address"
)

type SocketHttpHandler interface {
}
type webSocketHandler struct {
	addressUC address.UseCase
}

func NewWebSocketHandler(addressUC address.UseCase) SocketHttpHandler {
	return &webSocketHandler{}
}
