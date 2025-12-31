package todo

import "maps"

type List struct {
	tasks map[string]Task
}

func NewList() *List {
	return &List{
		make(map[string]Task),
	}
}

func (l *List) AddTask(taks Task) error {
	_, ok := l.tasks[taks.Title]
	if ok {
		return ErrorTaskAlreadyExist
	}
	l.tasks[taks.Title] = taks
	return nil
}

func (l *List) ListTasks() map[string]Task {
	fakeList := make(map[string]Task, len(l.tasks))

	maps.Copy(fakeList, l.tasks)
	return fakeList
}

func (l *List) NotDoneTasks() map[string]Task {
	notDoneTasks := make(map[string]Task)

	for title, task := range l.tasks {
		if !task.IsDone {
			notDoneTasks[title] = task
		}
	}
	return notDoneTasks
}

func (l *List) DoneTasks(title string) error {
	task, ok := l.tasks[title]
	if !ok {
		return ErrorTaskNotFound
	}

	task.Done()
	l.tasks[title] = task
	return nil
}

func (l *List) DeleteTask(title string) error {
	_, ok := l.tasks[title]
	if !ok {
		return ErrorTaskNotFound
	}

	delete(l.tasks, title)

	return nil
}
