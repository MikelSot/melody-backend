package domain

import "github.com/gorilla/websocket"

type User struct {
	Name       string `json:"name"`
	Controller *Controller
	Conn       *websocket.Conn
	Message    chan Message
}
