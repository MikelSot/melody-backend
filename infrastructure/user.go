package infrastructure

import (
	"github.com/MikelSot/melody-backend/domain"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type user struct {}

func NewUser() *user{
	return &user{}
}

// leemos el mensaje del frontend
func (u user) Read(user *domain.User) {
	defer func() {
		user.Controller.Logout <- user
		user.Conn.Close()
	}()

	user.Conn.SetReadLimit(maxMessageSize)

	for {
		message := domain.Message{}
		err := user.Conn.ReadJSON(&message)
		if err != nil {
			if websocket.IsUnexpectedCloseError(
				err,
				websocket.CloseGoingAway,
				websocket.CloseAbnormalClosure,
				){
				log.Println("No se puede leer el mensaje", err)
			}
			return
		}

		user.Controller.Broadcast <- message
	}
}

// leemos el mensaje y lo enviamos a frontend
func (u user) Write(user *domain.User) {
	defer func() {
		user.Conn.Close()
	}()

	for {
		select {
		case message, status := <-user.Message:
			if !status {
				return
			}
			user.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			err := user.Conn.WriteJSON(message)
			if err != nil {
				log.Println("no se puedo escribir el mensaje en el ws")
				return
			}
		}
	}
}
