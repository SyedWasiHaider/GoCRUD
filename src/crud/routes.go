package main

import (
"github.com/gorilla/mux"
"net/http"
)

func setupRoutesAndServe(){
	r := mux.NewRouter()
	r.HandleFunc("/Create", create).Methods("POST");
	r.HandleFunc("/Update/{id}", update).Methods("PUT");
	r.HandleFunc("/Find/{name}", getListing).Methods("GET");
	http.Handle("/", r)
	http.ListenAndServe(":1234", nil)
}