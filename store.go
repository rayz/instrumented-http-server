package main

type ToDoStore struct {
	Tasks map[int]*Task
	Count int // total number of tasks (used as ID for individual task)
}

func NewToDoStore() *ToDoStore {
	return &ToDoStore{Tasks: map[int]*Task{}}
}

func (store *ToDoStore) Add(task *Task) {
	store.Count += 1
	task.ID = store.Count
	store.Tasks[task.ID] = task
}

func (store *ToDoStore) CompleteTask(taskID int) bool {
	if task, ok := store.Tasks[taskID]; ok {
		task.Completed = true
		return true
	}
	return false
}

func (store *ToDoStore) DeleteTask(taskID int) {
	delete(store.Tasks, taskID)
}
