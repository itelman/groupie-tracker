package handler

import (
	"functions/internal/api"
	"net/http"
	"text/template"
)

var artistsMap, mapErr = api.GetArtistsMap()

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorPageHandler(w, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		ErrorPageHandler(w, http.StatusMethodNotAllowed)
		return
	}

	mainTmpl, err := template.ParseFiles("templates/main.html")
	if err != nil {
		ErrorPageHandler(w, http.StatusInternalServerError)
		return
	}

	if mapErr != nil {
		ErrorPageHandler(w, http.StatusInternalServerError)
		return
	}

	err = mainTmpl.Execute(w, artistsMap)
	if err != nil {
		ErrorPageHandler(w, http.StatusInternalServerError)
		return
	}
}
