package server

import (
	"groupie-tracker/internal/handler"
	"log"
	"net/http"
)

func Runserver() {
	mux := http.NewServeMux()

	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./templates/css"))))
	mux.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./templates/js"))))

	mux.HandleFunc("/", handler.IndexHandler)
	mux.HandleFunc("/groups/", handler.GroupHandler)
	mux.HandleFunc("/filter", handler.FilterHandler)
	mux.HandleFunc("/location/", handler.LocationHandler)

	log.Println("Listening on: http://localhost:4000/")
	http.ListenAndServe(":4000", mux)
}
