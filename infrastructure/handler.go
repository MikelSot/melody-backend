package infrastructure

import (
	"github.com/MikelSot/melody-backend/domain"
	"github.com/MikelSot/melody-backend/interfaces"
	"log"
	"net/http"
)

type handler struct {
	exists interfaces.UniqueName
	rw     interfaces.ReadAndWriteWS
}

func NewHandler(
	e interfaces.UniqueName,
	rw interfaces.ReadAndWriteWS) *handler {
	return &handler{e, rw}
}

//START OMIT
func (h *handler) handler(controller *domain.Controller, w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query()["name"]
	//START OMIT
	if len(name) != 1 {
		res := newResponse(ERROR, "Datos no validos", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}

	if h.exists.UniqueName(name[0]) {
		res := newResponse(ERROR, "Datos no validos", nil)
		responseJson(w, http.StatusBadRequest, res)
		return
	}
	//END OMIT
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("ERROR en conexion websocket", err)
		return
	}

	user := &domain.User{
		Name:       name[0],
		Controller: controller,
		Conn:       conn,
		Message:    make(chan domain.Message, 10),
	}

	user.Controller.Login <- user // HL

	go h.rw.Read(user)  // HL
	go h.rw.Write(user) // HL
}

//END OMIT
