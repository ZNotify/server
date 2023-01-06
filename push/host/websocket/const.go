package websocket

import (
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait = 5 * time.Second

	pongTimeout = 60 * time.Second

	pingPeriod = (pongTimeout * 7) / 10

	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
