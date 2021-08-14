package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/quinnj8/rxb-project/db"
	"github.com/quinnj8/rxb-project/models"
)

type Response struct {
	FilmsHandler []models.Film `json:"filmshandler"`
}

func FilmsHandler(w http.ResponseWriter, r *http.Request) {

	var response Response

	films := db.GetAllFilms()

	response.FilmsHandler = films

	res, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}
