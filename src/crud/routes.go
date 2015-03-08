package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func setupRoutesAndServe() {
	r := mux.NewRouter()
	r.HandleFunc("/Create", create).Methods("POST")
	r.HandleFunc("/Update/{id}", update).Methods("PUT")
	r.HandleFunc("/FindByName/{name}", getListing).Methods("GET")
	r.HandleFunc("/Find/{id}", getListingById).Methods("GET")
	r.HandleFunc("/Delete/{id}", deleteListing).Methods("DELETE")
	http.Handle("/", r)
	http.ListenAndServe(":1234", nil)
}
