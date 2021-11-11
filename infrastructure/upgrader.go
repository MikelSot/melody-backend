package infrastructure

import (
	"github.com/gorilla/websocket"
	"time"
)

const(
	maxMessageSize = 512  // en bytes (tamaño)
	writeWait = 10 * time.Second
)


var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
