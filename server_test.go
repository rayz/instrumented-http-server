package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestAddHandler(t *testing.T) {
	server := NewServer()
	tt := []struct {
		name       string
		method     string
		data       map[string]string
		id         int
		statusCode int
	}{
		{
			name:       "add task1",
			method:     http.MethodPost,
			data:       map[string]string{"description": "task1"},
			statusCode: 200,
		},
		{
			name:       "add task2",
			method:     http.MethodPost,
			data:       map[string]string{"description": "task1"},
			statusCode: 200,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			body, _ := json.Marshal(tc.data)
			responseRecorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/add", bytes.NewReader(body))
			server.AddTask(responseRecorder, req)
			if responseRecorder.Code != tc.statusCode {
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
			}
		})
	}
}

func TestGetHandler(t *testing.T) {
	server := NewServer()

	// add task1
	task1_data := map[string]string{"description": "fake_task1"}
	body, _ := json.Marshal(task1_data)
	responseRecorder := httptest.NewRecorder()
	add_task1_req := httptest.NewRequest(http.MethodPost, "/add", bytes.NewReader(body))
	server.AddTask(responseRecorder, add_task1_req)

	expected_task_1 := &Task{
		ID:          1,
		Description: "fake_task1",
		Completed:   false,
	}

	// get tasks and compare
	get_tasks_req := httptest.NewRequest(http.MethodGet, "/", nil)
	server.GetToDos(responseRecorder, get_tasks_req)
	get_tasks_resp := responseRecorder.Result()
	b, _ := io.ReadAll(get_tasks_resp.Body)

	tasks := make(map[int]*Task)
	json.Unmarshal(b, &tasks)

	assert.Equal(t, tasks[1], expected_task_1, "should get original task after adding and retrieving")
}

func TestCompleteHandler(t *testing.T) {
	server := NewServer()
	responseRecorder := httptest.NewRecorder()

	// add some tasks
	for i := 1; i < 5; i++ {
		task_data := map[string]string{"description": "fake_task"}
		body, _ := json.Marshal(task_data)
		add_task_req := httptest.NewRequest(http.MethodPost, "/add", bytes.NewReader(body))
		server.AddTask(responseRecorder, add_task_req)

	}

	// check task completion status is false on add
	for i := 1; i < 5; i++ {
		task := server.ToDoStore.Tasks[i]
		assert.False(t, task.Completed, "task completion should be false on initial add")

	}

	// complete handler
	for i := 1; i < 5; i++ {
		complete_route := fmt.Sprintf("/complete/%d", i)
		complete_task_req := httptest.NewRequest(http.MethodPut, complete_route, nil)
		// need to fake mux vars
		vars := map[string]string{"id": fmt.Sprint(i)}
		complete_task_req = mux.SetURLVars(complete_task_req, vars)
		server.CompleteTask(responseRecorder, complete_task_req)
	}

	// check if completed is set to true
	for i := 1; i < 5; i++ {
		task := server.ToDoStore.Tasks[i]
		assert.True(t, task.Completed, "task completion should be true after update")
	}
}

func TestDeleteHandler(t *testing.T) {
	server := NewServer()
	responseRecorder := httptest.NewRecorder()

	// add 5 tasks
	for i := 1; i <= 5; i++ {
		task_data := map[string]string{"description": "fake_task"}
		body, _ := json.Marshal(task_data)
		add_task_req := httptest.NewRequest(http.MethodPost, "/add", bytes.NewReader(body))
		server.AddTask(responseRecorder, add_task_req)
	}

	assert.Equal(t, len(server.ToDoStore.Tasks), 5, "store should have 5 tasks")

	// delete task1
	delete_task_req := httptest.NewRequest(http.MethodPost, "/delete/1", nil)
	delete_task_req = mux.SetURLVars(delete_task_req, map[string]string{"id": "1"})
	server.DeleteTask(responseRecorder, delete_task_req)

	assert.Equal(t, len(server.ToDoStore.Tasks), 4, "store should have 4 tasks after deleting 1 from 5")

}
