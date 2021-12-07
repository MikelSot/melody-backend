package infrastructure

import (
	"github.com/MikelSot/melody-backend/domain"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type user struct{}

func NewUser() *user {
	return &user{}
}

//START OMIT
// leemos el mensaje del frontend
func (u user) Read(user *domain.User) {
	defer func() {
		user.Controller.Logout <- user // HL
		user.Conn.Close()              // HL
	}()

	user.Conn.SetReadLimit(maxMessageSize)
	for { // HL
		message := domain.Message{}
		err := user.Conn.ReadJSON(&message) // HL
		if err != nil {
			if websocket.IsUnexpectedCloseError(
				err, websocket.CloseGoingAway,
				websocket.CloseAbnormalClosure,
			) {
				log.Println("No se puede leer el mensaje", err)
			}
			return
		}
		user.Controller.Broadcast <- message // HL
	} // HL
}

//END OMIT

//START_WRITE OMIT
// leemos el mensaje y lo enviamos a frontend
func (u user) Write(user *domain.User) {
	defer func() {
		user.Conn.Close() // HL
	}()

	for { // HL
		select {
		case message, status := <-user.Message: // HL
			if !status {
				return
			}
			user.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			err := user.Conn.WriteJSON(message) // HL
			if err != nil {
				log.Println("no se puedo escribir el mensaje en el ws")
				return
			}
		}
	} // HL
}

//END_WRITE OMIT
