package main

type Task struct {
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	ID          int    `json:"id"`
}
