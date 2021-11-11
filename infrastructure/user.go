package infrastructure

import (
	"github.com/MikelSot/melody-backend/domain"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type user struct {
	user *domain.User
}

func newUser(u *domain.User) *user{
	return &user{u}
}

// leemos el mensaje del frontend
func (u user) Read() {
	defer func() {
		u.user.Controller.Logout <- u.user
		u.user.Conn.Close()
	}()

	u.user.Conn.SetReadLimit(maxMessageSize)

	for {
		message := domain.Message{}
		err := u.user.Conn.ReadJSON(&message)
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

		u.user.Controller.Broadcast <- message
	}
}

// leemos el mensaje y lo enviamos a frontend
func (u user) Write() {
	defer func() {
		u.user.Conn.Close()
	}()

	for {
		select {
		case message, status := <-u.user.Message:
			if !status {
				return
			}
			u.user.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			err := u.user.Conn.WriteJSON(message)
			if err != nil {
				log.Println("no se puedo escribir el mensaje en el ws")
				return
			}
		}
	}
}
