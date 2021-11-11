package infrastructure

import (
	"github.com/MikelSot/melody-backend/domain"
	"net/http"
)

type router struct {
	serveMux   *http.ServeMux
	hanbler    *handler
	controller *domain.Controller
}

func NewRouter(
	s *http.ServeMux,
	h *handler,
	c *domain.Controller ) *router {
	return &router{s, h, c}
}

func (ro router) Router() {
	ro.serveMux.HandleFunc("/ws", ro.Request)
}

func (ro router) Request(w http.ResponseWriter, r *http.Request) {
	ro.hanbler.handler(ro.controller, w, r)
}
