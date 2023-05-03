//go:build wireinject
// +build wireinject

package socket

import "github.com/google/wire"

func initSocketCollection() *WebSocketCollection {

	wire.Build(IteratorCollection)

	return &WebSocketCollection{}
}
