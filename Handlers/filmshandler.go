package Handlers

import (
	"encoding/json"
	"net/http"

	"github.com/quinnj8/rxb-project/Models"
)

type Response struct {
	FilmsHandler Models.Films `json:"filmshandler"`
}

func FilmsHandler(w http.ResponseWriter, r *http.Request) {

	var response Response

	films := db.film.getAllFilms()

	response.FilmsHandler = films

	res, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}
