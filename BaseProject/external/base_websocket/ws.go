package base_websocket

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type SocketServices struct {
	services map[string]Hub
}

func (s *SocketServices) AddService(name string, hub Hub) {
	s.services[name] = hub
}
func (s *SocketServices) RemoveService(name string) {
	//TODO: Remove
}
func (s *SocketServices) Start() {

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
