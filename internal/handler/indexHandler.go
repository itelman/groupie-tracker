package handler

import (
	"functions/internal/models"
	"net/http"
	"text/template"
)

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

	artistsMap, err := models.GetArtistsMap()
	if err != nil {
		ErrorPageHandler(w, http.StatusInternalServerError)
		return
	}

	err = mainTmpl.Execute(w, artistsMap)
	if err != nil {
		ErrorPageHandler(w, http.StatusInternalServerError)
		return
	}
}
