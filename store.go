package main

type ToDoStore struct {
	Tasks []Task
	Count int // total number of tasks (used as ID for individual task)
}

func NewToDoStore() *ToDoStore {
	return &ToDoStore{}
}

func (store *ToDoStore) Add(task *Task) {
	store.Count += 1
	task.ID = store.Count
	store.Tasks = append(store.Tasks, *task)
}

func (store *ToDoStore) CompleteTask(taskID int) bool {
	for i := range store.Tasks {
		task := &store.Tasks[i]
		if task.ID == taskID {
			task.Completed = true
			return true
		}
	}
	return false
}
