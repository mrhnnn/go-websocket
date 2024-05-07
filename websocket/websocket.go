package websocket

import (
	"sync"

	"golang.org/x/net/websocket"
)

type WSServer struct {
	connection      map[string]map[*websocket.Conn]struct{}
	connectionMutex sync.Mutex
}

func NewWsServer() *WSServer {
	return &WSServer{
		connection: make(map[string]map[*websocket.Conn]struct{}),
	}
}

type WsResponseData struct {
	Action   string `json:"action"`
	Message  string `json:"message"`
	SenderId string `json:"senderid"`
	Name     string `json:"name"`
	CreateAt string `json:"createat"`
}
