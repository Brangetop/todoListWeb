package http

import (
	"encoding/json"
	"errors"
	"time"
)

type TaskDTO struct {
	Title       string
	Description string
}

func (t *TaskDTO) ValidateToCreate() error {
	if t.Title == "" {
		return errors.New("empty title")
	}
	if t.Description == "" {
		return errors.New("empty description")
	}
	return nil
}

type CompleteTaskDTO struct {
	Complete bool
}

type ErrorDTO struct {
	Message string
	Time    time.Time
}

func (e *ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}

	return string(b)
}
