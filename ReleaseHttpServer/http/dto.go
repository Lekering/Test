package http

import (
	"errors"
	"time"
)

type doneDto struct {
	Done bool `json:"done"`
}

type taskDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ErorDto struct {
	Massege string
	Time    time.Time
}

func NewErorDto(err string) *ErorDto {
	return &ErorDto{
		Massege: err,
		Time:    time.Now(),
	}
}

func NewTaskDto(title, description string) *taskDto {
	return &taskDto{
		Title:       title,
		Description: description,
	}
}

func (t *taskDto) Validate() error {
	if t.Title == "" {
		return errors.New("title is required")
	}
	if t.Description == "" {
		return errors.New("description is required")
	}
	return nil
}
