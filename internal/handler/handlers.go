package handler

import (
	"groupie-tracker/internal/models"
	"groupie-tracker/internal/service/api"
	"groupie-tracker/internal/service/filter"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var (
	ArtistURL   string = "https://groupietrackers.herokuapp.com/api/artists"
	LocationURL string = "https://groupietrackers.herokuapp.com/api/locations/"
	RelationURL string = "https://groupietrackers.herokuapp.com/api/relation/"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		log.Print("incorrect path")
		ErrorHandler(w, http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		//parsing html
		tmpl, err := template.ParseFiles("templates/html/index.html")
		if err != nil {
			log.Print(err)
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}

		// marshaling all artist url
		var groups models.Groups
		groups, err = api.GroupsJsonMarshalling(ArtistURL)
		if err != nil {
			log.Print(err)
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}

		// executing template
		if err = tmpl.Execute(w, groups); err != nil {
			log.Print(err)
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
	default:
		log.Print("incorrect method")
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
}

func GroupHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/groups/")

	// parsing html
	tmpl, err := template.ParseFiles("templates/html/group.html")
	if err != nil {
		log.Print(err)
		ErrorHandler(w, http.StatusNotFound)
		return
	}

	// path validation
	if err = api.PathValidation(r.URL.Path); err != nil {
		log.Print(err)
		ErrorHandler(w, http.StatusNotFound)
		return
	}

	// marshalling artist url
	var group models.Group
	group, err = api.GroupJsonMarshalling(ArtistURL + "/" + id)
	if err != nil {
		log.Print(err)
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	// marshalling locations url
	var locations models.Locations
	locations, err = api.LocationJsonMarshalling(LocationURL + id)
	if err != nil {
		log.Print(err)
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	// marshalling relations url
	var dateLocation models.DateLocation
	dateLocation, err = api.RelationJsonMarshalling(RelationURL + id)
	if err != nil {
		log.Print(err)
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	// formulating result group request
	var ResultGroup models.ResultGroup
	ResultGroup = api.FormResultGroup(group, locations, dateLocation)

	// executing template
	err = tmpl.Execute(w, ResultGroup)
	if err != nil {
		log.Print(err)
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}

func FilterHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/filter" {
		log.Print("incorrect path")
		ErrorHandler(w, http.StatusNotFound)
		return
	}

	switch r.Method {
	case "POST":
		var err error

		if err := r.ParseForm(); err != nil {
			log.Print(err)
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}

		CreationDateFrom := r.FormValue("creation_date_from")
		CreationDateTo := r.FormValue("creation_date_to")
		FirstAlbum := r.FormValue("firstAlbum")
		members := r.Form["members[]"]

		var filters models.Filter
		filters, err = filter.DataHandling(CreationDateFrom, CreationDateTo, FirstAlbum, members)
		if err != nil {
			log.Print(err)
			ErrorHandler(w, http.StatusBadRequest)
			return
		}

		var ResultGroup []models.ResultGroup

		ResultGroup, err = filter.ProcessData(filters)
		if err != nil {
			log.Print(err)
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}

		// for _, group := range ResultGroup {
		// 	fmt.Print("group id ")
		// 	fmt.Println(group.Id)
		// 	fmt.Print("group name ")
		// 	fmt.Println(group.Name)
		// 	fmt.Print("group creation ")
		// 	fmt.Println(group.CreationDate)
		// 	fmt.Print("group members len ")
		// 	fmt.Println(len(group.Members))
		// 	fmt.Println("------------------------------------------------------")
		// }

		result := models.Groups{
			Groups: models.ConvertToGroup(ResultGroup),
		}

		tmpl, err := template.ParseFiles("templates/html/index.html")
		if err != nil {
			log.Print(err)
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}

		if err = tmpl.Execute(w, result); err != nil {
			log.Print(err)
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}

	default:
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
}
