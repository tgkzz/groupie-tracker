package handler

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/internal/elk"
	"groupie-tracker/internal/models"
	"groupie-tracker/internal/service/api"
	"groupie-tracker/internal/service/filter"
	"groupie-tracker/internal/service/search"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"strings"
)

var (
	logger      elk.Logger = elk.GetLogger()
	ArtistURL   string     = "https://groupietrackers.herokuapp.com/api/artists"
	LocationURL string     = "https://groupietrackers.herokuapp.com/api/locations/"
	RelationURL string     = "https://groupietrackers.herokuapp.com/api/relation/"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	logger.Info(w, "application is running...")
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// FIX IT
	// url handling currently work unproperly due to adding metrics
	nameFunction := "IndexHandler"
	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusNotFound, nameFunction)
		return
	}
	switch r.Method {
	case "GET":
		//parsing html
		tmpl, err := template.ParseFiles("templates/html/index.html")
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError, nameFunction)
			return
		}

		// marshaling all artist url
		var groups models.Groups
		groups, err = api.GroupsJsonMarshalling(ArtistURL)
		if err != nil {
			ErrorHandler(w, http.StatusServiceUnavailable, nameFunction)
			return
		}

		// executing template
		if err = tmpl.Execute(w, groups); err != nil {
			ErrorHandler(w, http.StatusInternalServerError, nameFunction)
			return
		}
		logger.Info("IndexHandler is succesfully used")
	default:
		ErrorHandler(w, http.StatusMethodNotAllowed, nameFunction)
		return
	}
}

func GroupHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/groups/")
	nameFunction := "GroupHandler"
	// parsing html
	tmpl, err := template.ParseFiles("templates/html/group.html")
	if err != nil {
		ErrorHandler(w, http.StatusNotFound, nameFunction)
		return
	}

	// path validation
	if err = api.PathValidation(r.URL.Path); err != nil {
		ErrorHandler(w, http.StatusNotFound, nameFunction)
		return
	}

	// marshalling artist url
	var group models.Group
	group, err = api.GroupJsonMarshalling(ArtistURL + "/" + id)
	if err != nil {
		ErrorHandler(w, http.StatusServiceUnavailable, nameFunction)
		return
	}

	// marshalling locations url
	var locations models.Locations
	locations, err = api.LocationJsonMarshalling(LocationURL + id)
	if err != nil {
		ErrorHandler(w, http.StatusServiceUnavailable, nameFunction)
		return
	}

	// marshalling relations url
	var dateLocation models.DateLocation
	dateLocation, err = api.RelationJsonMarshalling(RelationURL + id)
	if err != nil {
		ErrorHandler(w, http.StatusServiceUnavailable, nameFunction)
		return
	}

	// formulating result group request
	ResultGroup := api.FormResultGroup(group, locations, dateLocation)

	// executing template
	err = tmpl.Execute(w, ResultGroup)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, nameFunction)
		return
	}
	logger.Info("GroupHandler is succesfully used")
}

func FilterHandler(w http.ResponseWriter, r *http.Request) {
	nameFunction := "FilterHandler"
	if r.URL.Path != "/filter" {
		ErrorHandler(w, http.StatusNotFound, nameFunction)
		return
	}

	switch r.Method {
	case "POST":
		var err error

		if err := r.ParseForm(); err != nil {
			ErrorHandler(w, http.StatusInternalServerError, nameFunction)
			return
		}

		CreationDateFrom := r.FormValue("creation_date_from")
		CreationDateTo := r.FormValue("creation_date_to")
		FirstAlbumFrom := r.FormValue("firstAlbum_from")
		FirstAlbumTo := r.FormValue("firstAlbum_to")
		members := r.Form["members[]"]
		location := r.FormValue("location")
		fmt.Println(location)

		var filters models.Filter
		filters, err = filter.DataHandling(CreationDateFrom, CreationDateTo, FirstAlbumFrom, FirstAlbumTo, members, location)
		if err != nil {
			ErrorHandler(w, http.StatusBadRequest, nameFunction)
			return
		}

		var ResultGroup []models.ResultGroup

		ResultGroup, err = filter.ProcessData(filters)
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError, nameFunction)
			return
		}

		for _, group := range ResultGroup {
			fmt.Print("group id ")
			fmt.Println(group.Id)
			fmt.Print("group nameFunction ")
			fmt.Println(group.Name)
			fmt.Print("group creation ")
			fmt.Println(group.CreationDate)
			fmt.Print("group members len ")
			fmt.Println(len(group.Members))
			fmt.Println("------------------------------------------------------")
		}

		result := models.Groups{
			Groups: models.ConvertToGroup(ResultGroup),
		}

		tmpl, err := template.ParseFiles("templates/html/index.html")
		if err != nil {
			ErrorHandler(w, http.StatusInternalServerError, nameFunction)
			return
		}

		if err = tmpl.Execute(w, result); err != nil {
			ErrorHandler(w, http.StatusInternalServerError, nameFunction)
			return
		}
		logger.Info("FilterHandler is succesfully used")
	default:
		ErrorHandler(w, http.StatusMethodNotAllowed, nameFunction)
		return
	}
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	nameFunction := "SearchHandler"
	if r.URL.Path != "/search" {
		ErrorHandler(w, http.StatusNotFound, nameFunction)
		return
	}
	if r.Method != "GET" {
		ErrorHandler(w, http.StatusMethodNotAllowed, nameFunction)
		return
	}
	searchText := r.FormValue("searchText")
	if searchText == "" {
		ErrorHandler(w, http.StatusBadRequest, nameFunction)
		return
	}

	body, err := api.ReturnResponseBody(ArtistURL)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, nameFunction)
		return
	}

	responseMap, err := search.SearchByTxt(searchText, body)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, nameFunction)
		return
	}

	response, err := json.Marshal(responseMap)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, nameFunction)
		return
	}
	logger.Info("SearchHandler is succesfully used")
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func LocationHandler(w http.ResponseWriter, r *http.Request) {
	// referrer := r.Header.Get("Referrer")
	// log.Print(referrer)
	// if referrer == "" || !strings.Contains(referrer, "localhost:4000") {
	// 	log.Print("path forbidden")
	// 	ErrorHandler(w, http.StatusForbidden)
	// 	return
	// }
	nameFunction := "LocationHandler"
	id := r.URL.Path[len("/location/"):]

	idNum, err := strconv.Atoi(id)
	if err != nil {
		ErrorHandler(w, http.StatusNotFound, nameFunction)
		return
	}

	if r.URL.Path != "/location/"+id || idNum < 1 || idNum > 52 || r.URL.Path == "/location/" {
		ErrorHandler(w, http.StatusNotFound, nameFunction)
		return
	}

	resp, err := http.Get(LocationURL + id)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, nameFunction)
		return
	}
	defer resp.Body.Close()
	logger.Info("LocationHandler is succesfully used")
	w.Header().Set("Content-Type", "application/json")
	io.Copy(w, resp.Body)
}
