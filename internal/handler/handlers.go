package handler

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/internal/models"
	"groupie-tracker/internal/resp"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

var (
	ArtistURL   string = "https://groupietrackers.herokuapp.com/api/artists"
	LocationURL string = "https://groupietrackers.herokuapp.com/api/locations/"
	RelationURL string = "https://groupietrackers.herokuapp.com/api/relation/"
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
		ErrorHandler(w, http.StatusMethodNotAllowed)
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

	// marshaling all artist url
	body, err := resp.ReturnResponseBody(ArtistURL)
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
		ErrorHandler(w, http.StatusMethodNotAllowed)
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

	// marshaling artist url
	bodyArtist, err := resp.ReturnResponseBody(ArtistURL + "/" + id)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	var group models.Group
	json.Unmarshal(bodyArtist, &group)

	// marshalling locations url
	bodyLocation, err := resp.ReturnResponseBody(LocationURL + id)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	var locations models.Locations
	json.Unmarshal(bodyLocation, &locations)

	// marshaling relations url
	bodyRelation, err := resp.ReturnResponseBody(RelationURL + id)
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
