package Handlers

import (
	"encoding/json"
	"net/http"

	"github.com/quinnj8/rxb-project/Models"
)

type Response struct {
	FilmsHandler []Models.Film `json:"filmshandler"`
}

func FilmsHandler(w http.ResponseWriter, r *http.Request) {

	response := Response

	films := getFilms()

	response.FilmsHandler = films

	res, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func getFilms() []Models.Film {
	films := []Models.Film

	film := Models.Film
	film.Id = 534
	films = append(films, film)

	film.Id = 426
	films = append(films, film)

	film.Id = 698
	films = append(films, film)

	return films
}
