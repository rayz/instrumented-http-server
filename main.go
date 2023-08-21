package main

import (
	"net/http"

	"github.com/DataDog/datadog-go/v5/statsd"

	"os"

	log "github.com/sirupsen/logrus"

	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

var logFile = "/tmp/instrumentedhttpserver.log"

func main() {

	// logs
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(f)
	log.WithFields(log.Fields{"string": "foo", "int": 1, "float": 1.1}).Info("instrumented http server started...")

	// statsd
	statsd, err := statsd.New("127.0.0.1:8125")
	if err != nil {
		log.Fatal(err)
	}

	// tracer
	tracer.Start(
		tracer.WithEnv("dev"),
		tracer.WithService("instrumented-http-server"),
		tracer.WithServiceVersion("1.0"),
	)
	defer tracer.Stop()

	r := muxtrace.NewRouter()
	todoserver := NewServer(statsd)
	r.HandleFunc("/", todoserver.GetToDos).Methods("GET")
	r.HandleFunc("/add", todoserver.AddTask).Methods("POST")
	r.HandleFunc("/complete/{id}", todoserver.CompleteTask).Methods("PUT")
	r.HandleFunc("/delete/{id}", todoserver.DeleteTask).Methods("POST")
	http.ListenAndServe(":8080", r)
}
