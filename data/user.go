package data

import (
	"github.com/MikelSot/melody-backend/domain"
)

type User struct {
	user       *domain.User
	controller *domain.Controller
}

func NewUser(u *domain.User, c *domain.Controller) *User {
	return &User{u, c}
}

//START OMIT
func (u User) AddUser(user *domain.User) {
	u.controller.Users[user.Name] = user // HL
}

func (u User) DeleteUser(user *domain.User) {
	_, state := u.controller.Users[user.Name]
	if state {
		delete(u.controller.Users, user.Name)
		close(user.Message) // HL
	}
}

func (u User) AddMessageQueue(message domain.Message) {
	for name, user := range u.controller.Users {
		if message.Username != name {
			select { // HL
			case user.Message <- message: // HL
			default:
				delete(u.controller.Users, user.Name)
				close(user.Message) // HL
			} // HL
		}
	}
}

//END OMIT

func (u User) UniqueName(name string) bool {
	for username, _ := range u.controller.Users {
		if name == username {
			return true
		}
	}
	return false
}
