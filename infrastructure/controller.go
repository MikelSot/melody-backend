package infrastructure

import (
	"github.com/MikelSot/melody-backend/domain"
	"github.com/MikelSot/melody-backend/interfaces"
	"log"
)

type controller struct {
	data       interfaces.DataPersistence
	controller *domain.Controller
}

func newController(
	data interfaces.DataPersistence,
	c *domain.Controller ) *controller {
	return &controller{data, c}
}

func (c controller) run() {
	for {
		select {
		case user := <- c.controller.Login:
			c.data.AddUser(user)
		case user := <- c.controller.Logout:
			c.data.DeleteUser(user)
		case message := <- c.controller.Broadcast:
			c.data.AddMessageQueue(message)
		default:
			log.Println("F")
		}
	}
}
