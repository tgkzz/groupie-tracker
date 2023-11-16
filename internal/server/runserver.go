package server

import (
	"groupie-tracker/internal/handler"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Runserver() {
	mux := http.NewServeMux()

	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./templates/css"))))
	mux.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./templates/js"))))

	mux.HandleFunc("/", handler.IndexHandler)
	mux.HandleFunc("/groups/", handler.GroupHandler)
	mux.HandleFunc("/filter", handler.FilterHandler)
	mux.HandleFunc("/location/", handler.LocationHandler)
	mux.HandleFunc("/search", handler.SearchHandler)

	//adding metrics
	mux.Handle("/metrics", promhttp.Handler())

	log.Println("Listening on: http://localhost:4000/")
	http.ListenAndServe(":4000", mux)
}
