package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ToDoServer struct {
	ToDoStore *ToDoStore
}

func NewServer() *ToDoServer {
	return &ToDoServer{
		ToDoStore: NewToDoStore(),
	}
}

func (server *ToDoServer) GetToDos(w http.ResponseWriter, req *http.Request) {
	j, _ := json.Marshal(server.ToDoStore.Tasks)
	fmt.Fprintf(w, string(j))

}

func (server *ToDoServer) AddTask(w http.ResponseWriter, req *http.Request) {
	var task Task
	err := json.NewDecoder(req.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	server.ToDoStore.Add(&task)
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
}

func (server *ToDoServer) DeleteTask(w http.ResponseWriter, req *http.Request) {
	target, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	server.ToDoStore.DeleteTask(target)
	fmt.Fprintf(w, "Task id %d deleted\n", target)
}
