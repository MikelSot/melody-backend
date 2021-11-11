package middleware

import (
	"log"
	"net/http"
)

type handler func(w http.ResponseWriter, r *http.Request)

func Log(fun handler) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Peticion ->  %q, metodo -> %q", r.URL.Path, r.Method)
		fun(w,r)
	}
}
