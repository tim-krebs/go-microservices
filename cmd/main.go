package main

import (
	"net/http"
	"registrierung"

	"github.com/gorilla/mux"
)

func main() {
	regHandler := &registrierung.RegistrierungsHandler{}
	r := mux.NewRouter()
	r.PathPrefix("/").Methods("POST").Handler(regHandler)
	http.ListenAndServe(":8080", r)
}
