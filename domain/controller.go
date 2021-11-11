package domain

type Controller struct {
	Users     map[string]*User
	Login     chan *User
	Logout    chan *User
	Broadcast chan Message
}

func NewController() *Controller{
	return &Controller{
		Users:     make(map[string]*User),
		Login:     make(chan *User),
		Logout:    make(chan *User),
		Broadcast: make(chan Message),
	}
}
