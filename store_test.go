package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddToStore(t *testing.T) {
	store := NewToDoStore()
	for i := 1; i <= 10; i++ {
		// each task gets id + 1 of previous task id
		task := new(Task)
		store.Add(task)
		assert.Equal(t, task.ID, i, "task given unique id")

		_, task_exists := store.Tasks[i]
		assert.True(t, task_exists, "task stored in TaskStore")
	}
}

func TestCompleteTask(t *testing.T) {
	store := NewToDoStore()
	for i := 1; i <= 10; i++ {
		task := new(Task)
		store.Add(task)
		assert.False(t, task.Completed, "task completed false on creation")

		store.CompleteTask(i)
		assert.True(t, task.Completed, "task completed true after update")
	}
}
