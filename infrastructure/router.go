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
	c *domain.Controller) *router {
	return &router{s, h, c}
}

//START OMIT
func (ro router) Router() {
	ro.webPage()
	ro.serveMux.HandleFunc("/ws", middleware.Log(ro.request)) // HL
}

func (ro router) request(w http.ResponseWriter, r *http.Request) {
	ro.handler.handler(ro.controller, w, r)
}

func (ro router) webPage() {
	ro.serveMux.Handle("/", http.FileServer(http.Dir("public"))) // HL
}

//END OMIT
