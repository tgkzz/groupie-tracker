package handler

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/internal/models"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func ErrorHandler(w http.ResponseWriter, code int) {
	w.WriteHeader(code)

	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		text := fmt.Sprintf("Error 500\n Oppss! %s", http.StatusText(http.StatusInternalServerError))
		http.Error(w, text, http.StatusInternalServerError)
		return
	}

	res := &models.Err{Text_err: http.StatusText(code), Code_err: code}
	err = tmpl.Execute(w, &res)
	if err != nil {
		text := fmt.Sprintf("Error 500\n Oppss! %s", http.StatusText(http.StatusInternalServerError))
		http.Error(w, text, http.StatusInternalServerError)
		return
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}

	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	var groups models.Groups
	json.Unmarshal(body, &groups.Groups)

	err = tmpl.Execute(w, groups)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}

func GroupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}

	tmpl, err := template.ParseFiles("templates/group.html")
	if err != nil {
		ErrorHandler(w, http.StatusNotFound)
		return
	}

	path := r.URL.Path

	id := strings.TrimPrefix(path, "/groups/")

	idNum, err := strconv.Atoi(id)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	if id == "" || idNum > 52 {
		ErrorHandler(w, http.StatusNotFound)
		return
	}

	// masrshaling artist url
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + id)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	var group models.Group
	json.Unmarshal(body, &group)

	// marshalling locations url
	respLocations, err := http.Get("https://groupietrackers.herokuapp.com/api/locations/" + id)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	bodyLocations, err := io.ReadAll(respLocations.Body)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	var locations models.Locations
	json.Unmarshal(bodyLocations, &locations)

	// marshaling relations url
	respRelation, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + id)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	bodyRelation, err := io.ReadAll(respRelation.Body)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	var dateLocation models.DateLocation
	json.Unmarshal(bodyRelation, &dateLocation)

	var ResultGroup models.ResultGroup

	ResultGroup.Id = group.Id
	ResultGroup.Image = group.Image
	ResultGroup.Name = group.Name
	ResultGroup.CreationDate = group.CreationDate
	ResultGroup.Members = group.Members
	ResultGroup.FirstAlbum = group.FirstAlbum

	for _, loc := range locations.Locate {
		ResultGroup.ConcertData = append(ResultGroup.ConcertData, models.Concerts{Location: loc, Dates: dateLocation.DateLoc[loc]})
	}

	err = tmpl.Execute(w, ResultGroup)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}
