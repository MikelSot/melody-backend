package infrastructure

import (
	"github.com/MikelSot/melody-backend/domain"
	"github.com/MikelSot/melody-backend/middleware"
	"net/http"
)

type router struct {
	serveMux   *http.ServeMux
	handler    *handler
	controller *domain.Controller
}

func NewRouter(
	s *http.ServeMux,
	h *handler,
	c *domain.Controller ) *router {
	return &router{s, h, c}
}

func (ro router) Router() {
	ro.serveMux.HandleFunc("/ws", middleware.Log(ro.Request))
}

func (ro router) Request(w http.ResponseWriter, r *http.Request) {
	ro.handler.handler(ro.controller, w, r)
}
