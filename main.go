package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DataDog/datadog-go/v5/statsd"

	"github.com/gorilla/mux"
)

func main() {
	statsd, err := statsd.New("127.0.0.1:8125")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("stats created")
	r := mux.NewRouter()
	todoserver := NewServer(statsd)
	r.HandleFunc("/", todoserver.GetToDos).Methods("GET")
	r.HandleFunc("/add", todoserver.AddTask).Methods("POST")
	r.HandleFunc("/complete/{id}", todoserver.CompleteTask).Methods("PUT")
	r.HandleFunc("/delete/{id}", todoserver.DeleteTask).Methods("POST")
	http.ListenAndServe(":8080", r)
}
