package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	todoserver := NewServer()
	r.HandleFunc("/", todoserver.GetToDos).Methods("GET")
	r.HandleFunc("/add", todoserver.AddTask).Methods("POST")
	r.HandleFunc("/complete/{id}", todoserver.CompleteTask).Methods("PUT")
	http.ListenAndServe(":8080", r)
}