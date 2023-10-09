package handler

import (
	"encoding/json"
	"groupie-tracker/internal/models"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal("Error 400")
	}

	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal("Error 400")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var groups models.Groups
	json.Unmarshal(body, &groups.Groups)

	err = tmpl.Execute(w, groups)
	if err != nil {
		log.Print(err)
	}
}

func GroupHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/group.html")
	if err != nil {
		log.Fatal("Error 400")
	}

	path := r.URL.Path

	id := strings.TrimPrefix(path, "/groups/")

	if id == "" {
		log.Fatal("Error: empty id")
	}

	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + id)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var group models.Group
	json.Unmarshal(body, &group)

	err = tmpl.Execute(w, group)
}
