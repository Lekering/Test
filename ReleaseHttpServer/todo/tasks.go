package todo

import "time"

type Task struct {
	Title       string
	Description string
	IsDone      bool

	CreateAt time.Time
	DoneAt   *time.Time
}

func NewTask(titel string, descrition string) Task {
	return Task{
		Title:       titel,
		Description: descrition,
		IsDone:      false,

		CreateAt: time.Now(),
		DoneAt:   nil,
	}
}
