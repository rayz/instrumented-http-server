package main

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

func (store *ToDoStore) Add(task *Task) {
	store.IDCounter += 1
	store.Uncompleted += 1
	task.ID = store.IDCounter
	store.Tasks[task.ID] = task
}

func (store *ToDoStore) CompleteTask(taskID int) bool {
	if task, ok := store.Tasks[taskID]; ok {
		task.Completed = true
		store.Completed += 1
		store.Uncompleted -= 1
		return true
	}
	return false
}

func (store *ToDoStore) DeleteTask(taskID int) {
	if _, ok := store.Tasks[taskID]; ok {
		delete(store.Tasks, taskID)
		store.Deleted += 1
	}
}
