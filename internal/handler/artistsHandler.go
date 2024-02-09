package handler

import (
	"functions/internal/api"
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

	if mapErr != nil {
		ErrorPageHandler(w, http.StatusInternalServerError)
		return
	}

	if _, ok := artistsMap[id]; !ok {
		ErrorPageHandler(w, http.StatusNotFound)
		return
	}

	artistFullInfo, err := api.GetArtistFullInfo(artistsMap[id])
	if err != nil {
		ErrorPageHandler(w, http.StatusInternalServerError)
		return
	}

	mainTmpl, err := template.ParseFiles("templates/artists.html")
	if err != nil {
		ErrorPageHandler(w, http.StatusInternalServerError)
		return
	}

	err = mainTmpl.Execute(w, artistFullInfo)
	if err != nil {
		ErrorPageHandler(w, http.StatusInternalServerError)
		return
	}
}
