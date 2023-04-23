package base_websocket

import (
	"errors"
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type SocketServices struct {
	services map[string]Hub
}

func (s *SocketServices) AddService(name string, hub Hub) error {
	_, ok := s.services[name]
	if ok {
		return errors.New("Service already exists")
	}
	s.services[name] = hub
	return nil
}
func (s *SocketServices) RemoveService(name string) {
	delete(s.services, name)
}
func (s *SocketServices) Start() {
	for _, hub := range s.services {
		go hub.run()
	}
}

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	route := gin.Default()

	hub := newHub()
	go hub.run()

	route.GET("/ws", func(c *gin.Context) {
		serveWs(hub, c.Writer, c.Request)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
