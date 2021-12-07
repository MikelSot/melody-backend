package infrastructure

import (
	"github.com/MikelSot/melody-backend/domain"
	"github.com/MikelSot/melody-backend/interfaces"
)

type controller struct {
	data       interfaces.DataPersistence
	controller *domain.Controller
}

func NewController(
	data interfaces.DataPersistence,
	c *domain.Controller) *controller {
	return &controller{data, c}
}

//START OMIT
func (c controller) Run() {
	for { // HL
		select {
		case user := <-c.controller.Login: // HL
			c.data.AddUser(user)
		case user := <-c.controller.Logout: // HL
			c.data.DeleteUser(user)
		case message := <-c.controller.Broadcast: // HL
			c.data.AddMessageQueue(message)
		}
	} // HL
}

//END OMIT
