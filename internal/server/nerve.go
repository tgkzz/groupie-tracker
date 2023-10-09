package server

import (
	handler "groupie-tracker/internal/handler"
	"log"
	"net/http"
)

func RunServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.IndexHandler)
	mux.HandleFunc("/groups/", handler.GroupHandler)

	log.Println("Listening on: http://localhost:8080/")
	http.ListenAndServe(":8080", mux)
}
