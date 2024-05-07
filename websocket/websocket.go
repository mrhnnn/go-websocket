package websocket

import (
	"log"
	"sync"

	"golang.org/x/net/websocket"
)

type WSServer struct {
	connections     map[string]map[*websocket.Conn]struct{}
	connectionMutex sync.Mutex
}

func NewWsServer() *WSServer {
	return &WSServer{
		connections: make(map[string]map[*websocket.Conn]struct{}),
	}
}

type WsResponseData struct {
	Action   string `json:"action"`
	Message  string `json:"message"`
	SenderId string `json:"senderid"`
	Name     string `json:"name"`
	CreateAt string `json:"createat"`
}

func (w *WSServer) Broadcast(chatRoomId, msg string) {
	w.connectionMutex.Lock()
	defer w.connectionMutex.Unlock()

	if groupConnections, ok := w.connections[chatRoomId]; ok {
		for conn := range groupConnections {
			if err := websocket.Message.Send(conn, msg); err != nil {
				log.Println(err)
			}
		}
	}
}
