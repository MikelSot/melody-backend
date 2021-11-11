package interfaces

import "github.com/MikelSot/melody-backend/domain"

type Runner interface {
	Run()
}

type DataPersistence interface {
	AddUser(user *domain.User)
	DeleteUser(user *domain.User)
	AddMessageQueue(message domain.Message)
}
