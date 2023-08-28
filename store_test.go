package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddToStore(t *testing.T) {
	store := NewToDoStore()
	for i := 1; i <= 10; i++ {
		// each task gets id + 1 of previous task id
		task := &Task{}
		store.Add(context.Background(), task)
		assert.Equal(t, task.ID, i, "task given unique id")

		_, task_exists := store.Tasks[i]
		assert.True(t, task_exists, "task stored in TaskStore")
	}
}

func TestCompleteTask(t *testing.T) {
	store := NewToDoStore()
	for i := 1; i <= 10; i++ {
		task := &Task{}
		store.Add(context.Background(), task)
		assert.False(t, task.Completed, "task completed false on creation")

		store.CompleteTask(context.Background(), i)
		assert.True(t, task.Completed, "task completed true after update")
	}
}
