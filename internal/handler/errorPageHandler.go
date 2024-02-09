package handler

import (
	"functions/internal/models"
	"net/http"
	"text/template"
)

func ErrorPageHandler(w http.ResponseWriter, statCode int) {
	errors := map[int]models.Error{
		http.StatusBadRequest:          {Code: http.StatusBadRequest, Name: "Bad Request", Content: "The server cannot process the request due to something that is perceived to be a client error."},
		http.StatusNotFound:            {Code: http.StatusNotFound, Name: "Resource not found", Content: "The requested resource could not be found but may be available again in the future."},
		http.StatusMethodNotAllowed:    {Code: http.StatusMethodNotAllowed, Name: "Method Not Allowed", Content: "A request method is not supported for the requested resource."},
		http.StatusInternalServerError: {Code: http.StatusInternalServerError, Name: "Webservice currently unavailable", Content: "An unexpected condition was encountered.<br />Our service team has been dispatched to bring it back online."},
	}

	errTmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statCode)
	err = errTmpl.Execute(w, errors[statCode])
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
