func (h *handler) handler(controller *domain.Controller, w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query()["name"] // HL

	conn, err := upgrader.Upgrade(w, r, nil) // HL
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
