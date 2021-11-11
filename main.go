package main

import (
	"github.com/MikelSot/melody-backend/data"
	"github.com/MikelSot/melody-backend/domain"
	"github.com/MikelSot/melody-backend/infrastructure"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	serveMux := http.NewServeMux()
	controller := domain.NewController()
	user := &domain.User{}

	data := data.NewUser(user, controller)
	rw := infrastructure.NewUser()
	cont := infrastructure.NewController(data, controller)
	go cont.Run()

	handler := infrastructure.NewHandler(data, rw)
	router := infrastructure.NewRouter(serveMux, handler, controller)
	router.Router()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet},
		AllowCredentials: true,
	}).Handler(serveMux)

	log.Println("Servidor iniciado en el puerto :3000")
	log.Println(http.ListenAndServe(":3000", cors))
}