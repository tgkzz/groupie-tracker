package handler

import (
	"fmt"
	"groupie-tracker/internal/elk"
	"groupie-tracker/internal/models"
	"html/template"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, code int, st string) {
	logger := elk.GetLogger()
	w.WriteHeader(code)

	tmpl, err := template.ParseFiles("templates/html/error.html")
	if err != nil {
		text := fmt.Sprintf("Error 500\n Oppss! %s", http.StatusText(http.StatusInternalServerError))
		logger.Println("ERROR: " + st + ": " + http.StatusText(http.StatusInternalServerError))
		http.Error(w, text, http.StatusInternalServerError)
		return
	}

	res := &models.Err{Text_err: http.StatusText(code), Code_err: code}
	err = tmpl.Execute(w, &res)
	if err != nil {
		text := fmt.Sprintf("Error 500\n Oppss! %s", http.StatusText(http.StatusInternalServerError))
		logger.Error(st + ": " + http.StatusText(http.StatusInternalServerError))
		http.Error(w, text, http.StatusInternalServerError)
		return
	}
}
