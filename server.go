package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/DataDog/datadog-go/v5/statsd"
	"github.com/gorilla/mux"
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
	server.Statsd.Count("todo_tasks_uncompleted.count", int64(server.ToDoStore.Uncompleted), []string{"environment:dev"}, 1)
}

func (server *ToDoServer) AddTask(w http.ResponseWriter, req *http.Request) {
	var task Task
	err := json.NewDecoder(req.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	server.ToDoStore.Add(&task)
	server.Statsd.Count("todo_tasks_uncompleted.count", int64(server.ToDoStore.Uncompleted), []string{"environment:dev"}, 1)
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
	server.Statsd.Count("todo_tasks_uncompleted.count", int64(server.ToDoStore.Uncompleted), []string{"environment:dev"}, 1)
	server.Statsd.Count("todo_tasks_completed.count", int64(server.ToDoStore.Completed), []string{"environment:dev"}, 1)
}

func (server *ToDoServer) DeleteTask(w http.ResponseWriter, req *http.Request) {
	target, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	server.ToDoStore.DeleteTask(target)
	fmt.Fprintf(w, "Task id %d deleted\n", target)
	server.Statsd.Count("todo_tasks_deleted.count", int64(server.ToDoStore.Deleted), []string{"environment:dev"}, 1)
}
