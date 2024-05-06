package handler

import (
	"functions/internal/models"
	"net/http"
	"strconv"
	"text/template"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artists" {
		ErrorPageHandler(w, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		ErrorPageHandler(w, http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("id")

	if _, err := strconv.Atoi(id); err != nil {
		ErrorPageHandler(w, http.StatusBadRequest)
		return
	}

	artistsMap, err := models.GetArtistsMap()
	if err != nil {
		ErrorPageHandler(w, http.StatusInternalServerError)
		return
	}

	if _, ok := artistsMap[id]; !ok {
		ErrorPageHandler(w, http.StatusNotFound)
		return
	}

	artistInfo := artistsMap[id]
	err = artistInfo.GetFullInfo()
	if err != nil {
		ErrorPageHandler(w, http.StatusInternalServerError)
		return
	}

	mainTmpl, err := template.ParseFiles("templates/artists.html")
	if err != nil {
		ErrorPageHandler(w, http.StatusInternalServerError)
		return
	}

	err = mainTmpl.Execute(w, artistInfo)
	if err != nil {
		ErrorPageHandler(w, http.StatusInternalServerError)
		return
	}
}
