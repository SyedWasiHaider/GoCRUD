package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func setupRoutesAndServe() {
	r := mux.NewRouter()
	r.HandleFunc("/User/Create", createUser)
	r.HandleFunc("/ProductListing/Create", createListing).Methods("POST")
	r.HandleFunc("/ProductListing/Update/{id}", updateListing).Methods("PUT")
	r.HandleFunc("/ProductListing/FindByName/{name}", getListing).Methods("GET")
	r.HandleFunc("/ProductListing/Find/{id}", getListingById).Methods("GET")
	r.HandleFunc("/ProductListing/Delete/{id}", deleteListing).Methods("DELETE")
	http.Handle("/", r)
	http.ListenAndServe(":1234", nil)
}
