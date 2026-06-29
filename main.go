package main

import "time"

type Task struct {
	Title        string
	Description  string
	isDone       bool
	initTime     time.Time
	finishedTime time.Time
}

func newTask(title string, description string) *Task {

	return &Task{
		Title:       title,
		Description: description,
		isDone:      false,
		initTime:    time.Now(),
	}
}

func main() {

}
