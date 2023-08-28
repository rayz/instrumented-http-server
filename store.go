package main

import (
	"context"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type ToDoStore struct {
	Tasks       map[int]*Task
	IDCounter   int // (ID given to a task before incrementing by 1)
	Uncompleted int // number of tasks not yet completed
	Completed   int // number of tasks that have been completed
	Deleted     int // number of tasks deleted
}

func NewToDoStore() *ToDoStore {
	return &ToDoStore{Tasks: map[int]*Task{}}
}

func (store *ToDoStore) Add(ctx context.Context, task *Task) {
	span, ctx := tracer.StartSpanFromContext(ctx, "add_task")
	defer span.Finish()
	store.IDCounter += 1
	store.Uncompleted += 1
	task.ID = store.IDCounter
	store.Tasks[task.ID] = task
}

func (store *ToDoStore) CompleteTask(ctx context.Context, taskID int) bool {
	span, ctx := tracer.StartSpanFromContext(ctx, "complete_task")
	defer span.Finish()
	if task, ok := store.Tasks[taskID]; ok {
		task.Completed = true
		store.Completed += 1
		store.Uncompleted -= 1
		return true
	}
	return false
}

func (store *ToDoStore) DeleteTask(ctx context.Context, taskID int) {
	span, ctx := tracer.StartSpanFromContext(ctx, "delete_task")
	defer span.Finish()
	if _, ok := store.Tasks[taskID]; ok {
		delete(store.Tasks, taskID)
		store.Deleted += 1
	}
}
