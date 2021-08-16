package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/quinnj8/rxb-project/db"
	"github.com/quinnj8/rxb-project/models"
)

type Response struct {
	FilmsHandler []models.Film `json:"filmshandler"`
}

func GetAllFilms(w http.ResponseWriter, r *http.Request) {

	var response Response

	films := db.GetAllFilms()

	response.FilmsHandler = films

	res, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func GetFilmsByTitle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	var film models.Film

	film = db.GetFilmByTitle(title)

	res, _ := json.Marshal(film)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func GetFilmsByRating(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filmrating := vars["rating"]

	film := db.GetFilmByRating(filmrating)

	res, _ := json.Marshal(film)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func GetFilmsByCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category := vars["category"]

	films := db.GetFilmByCategory(category)

	res, _ := json.Marshal(films)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func GetFilmDetails(w http.ResponseWriter, r *http.Request) {

}

func AddFilmComment(w http.ResponseWriter, r *http.Request) {

}

func GetFilmComments(w http.ResponseWriter, r *http.Request) {

}
