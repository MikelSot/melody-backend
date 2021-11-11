package interfaces

import "github.com/MikelSot/melody-backend/domain"

type DataPersistence interface {
	AddUser(user *domain.User)
	DeleteUser(user *domain.User)
	AddMessageQueue(message domain.Message)
}

type UniqueName interface {
	UniqueName(name string) bool
}

type ReadAndWriteWS interface {
	Read()
	Write()
}