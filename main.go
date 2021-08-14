package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/quinnj8/rxb-project/Handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/Welcome", Handlers.WelcomeHandler).Methods("GET")
	r.HandleFunc("/Films", Handlers.FilmsHandler).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8080", r)
}
