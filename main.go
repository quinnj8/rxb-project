package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/quinnj8/rxb-project/handlers"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/Welcome", handlers.WelcomeHandler).Methods("GET")
	r.HandleFunc("/Films", handlers.GetFilms).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8080", r)
}
