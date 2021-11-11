package infrastructure

import (
	"encoding/json"
	"net/http"
)

const(
	ERROR      = "error"
	MESSAGE    = "message"
)

type response struct {
	MessageType string      `json:"message_type"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}

func newResponse(messageType string, message string, data interface{}) response {
	return response{
		messageType,
		message,
		data,
	}
}

func responseJson(w http.ResponseWriter, statusCode int,r response){
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(&r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

