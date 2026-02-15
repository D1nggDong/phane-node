package core

import (
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Hub struct {
	clients map[*websocket.Conn]bool
	mu      sync.Mutex
}

var GlobalHub = &Hub{clients: make(map[*websocket.Conn]bool)}

func (h *Hub) HandleWS(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)
	h.mu.Lock()
	h.clients[conn] = true
	h.mu.Unlock()
}

func (h *Hub) Broadcast(msg interface{}) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for client := range h.clients {
		client.WriteJSON(msg)
	}
}
