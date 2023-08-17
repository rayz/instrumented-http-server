package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/DataDog/datadog-go/v5/statsd"
	"github.com/gorilla/mux"

	log "github.com/sirupsen/logrus"
)

type ToDoServer struct {
	ToDoStore *ToDoStore
	Statsd    *statsd.Client
}

func NewServer(statsd *statsd.Client) *ToDoServer {
	return &ToDoServer{
		ToDoStore: NewToDoStore(),
		Statsd:    statsd,
	}
}

func (server *ToDoServer) GetToDos(w http.ResponseWriter, req *http.Request) {
	j, _ := json.Marshal(server.ToDoStore.Tasks)
	fmt.Fprintf(w, string(j))
	server.Statsd.Incr("todo_get", nil, 1)
	log.WithFields(log.Fields{"total_tasks": len(server.ToDoStore.Tasks)}).Info("tasks retrieved")
}

func (server *ToDoServer) AddTask(w http.ResponseWriter, req *http.Request) {
	var task Task
	err := json.NewDecoder(req.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	server.ToDoStore.Add(&task)
	server.Statsd.Incr("todo_tasks", nil, 1)
	log.WithFields(log.Fields{"id": task.ID, "description": task.Description, "completed": task.Completed}).Info("task added")
}

func (server *ToDoServer) CompleteTask(w http.ResponseWriter, req *http.Request) {
	target, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !server.ToDoStore.CompleteTask(target) {
		http.Error(w, "Task id not found", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Task id %d completed\n", target)
	server.Statsd.Incr("todo_tasks_completed", nil, 1)
	log.WithFields(log.Fields{"id": target}).Info("task completed")
}

func (server *ToDoServer) DeleteTask(w http.ResponseWriter, req *http.Request) {
	target, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	server.ToDoStore.DeleteTask(target)
	log.WithFields(log.Fields{"id": target}).Info("task deleted")
	server.Statsd.Decr("todo_tasks", nil, 1)
}
