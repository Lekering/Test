package todo

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type List struct {
	repo *Repository
	ctx  context.Context
}

func NewList(ctx context.Context, conn *pgx.Conn) *List {
	return &List{
		repo: NewRepository(conn),
		ctx:  ctx,
	}
}

func (l *List) AddTask(task Task) error {
	return l.repo.AddTask(l.ctx, task)
}

func (l *List) GetTasks(title string) (Task, error) {
	return l.repo.GetTask(l.ctx, title)
}

func (l *List) ListTasks() map[string]Task {
	tasks, err := l.repo.ListTasks(l.ctx)
	if err != nil {
		return make(map[string]Task)
	}
	return tasks
}

// NotDoneTasks возвращает все задачи, которые еще не выполнены (IsDone == false).
func (l *List) NotDoneTasks() map[string]Task {
	tasks, err := l.repo.NotDoneTasks(l.ctx)
	if err != nil {
		return make(map[string]Task)
	}
	return tasks
}

func (l *List) DoneTasks(title string) (Task, error) {
	return l.repo.UpdateTaskDone(l.ctx, title, true)
}

func (l *List) UnDoneTasks(title string) (Task, error) {
	return l.repo.UpdateTaskDone(l.ctx, title, false)
}

func (l *List) DeleteTask(title string) error {
	return l.repo.DeleteTask(l.ctx, title)
}
