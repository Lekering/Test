package todo

import (
	"errors"
	"maps"
	"sync"
)

type List struct {
	tasks map[string]Task
	mtx   sync.RWMutex
}

func NewList() *List {
	return &List{
		make(map[string]Task),
		sync.RWMutex{},
	}
}

func (l *List) AddTask(taks Task) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	_, ok := l.tasks[taks.Title]
	if ok {
		return ErrorTaskAlreadyExist
	}
	l.tasks[taks.Title] = taks
	return nil
}

func (l *List) GetTasks(title string) (Task, error) {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	if task, ok := l.tasks[title]; !ok {
		return Task{}, errors.New("Task not found")
	} else {
		return task, nil
	}
}

func (l *List) ListTasks() map[string]Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	fakeList := make(map[string]Task, len(l.tasks))

	maps.Copy(fakeList, l.tasks)
	return fakeList
}

// NotDoneTasks возвращает все задачи, которые еще не выполнены (IsDone == false).
func (l *List) NotDoneTasks() map[string]Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	notDoneTasks := make(map[string]Task)

	for title, task := range l.tasks {
		if !task.IsDone {
			notDoneTasks[title] = task
		}
	}
	return notDoneTasks
}

func (l *List) DoneTasks(title string) (Task, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	task, ok := l.tasks[title]
	if !ok {
		return Task{}, ErrorTaskNotFound
	}

	task.Done()
	l.tasks[title] = task
	return task, nil
}

func (l *List) UnDoneTasks(title string) (Task, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	task, ok := l.tasks[title]
	if !ok {
		return Task{}, ErrorTaskNotFound
	}

	task.UnDone()
	l.tasks[title] = task
	return task, nil
}

func (l *List) DeleteTask(title string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	_, ok := l.tasks[title]
	if !ok {
		return ErrorTaskNotFound
	}

	delete(l.tasks, title)

	return nil
}
