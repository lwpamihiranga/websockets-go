package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	websocketUpgrade = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type Manager struct {
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) serverWS(w http.ResponseWriter, r *http.Request) {
	log.Println("new connection")

	conn, err := websocketUpgrade.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}

	conn.Close()
}
