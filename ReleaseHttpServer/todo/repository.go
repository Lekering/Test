package todo

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	conn *pgx.Conn
}

func NewRepository(conn *pgx.Conn) *Repository {
	return &Repository{
		conn: conn,
	}
}

func (r *Repository) AddTask(ctx context.Context, task Task) error {
	// Сначала проверяем, существует ли задача
	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM tasks WHERE title = $1)`
	err := r.conn.QueryRow(ctx, checkQuery, task.Title).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		return ErrorTaskAlreadyExist
	}

	query := `
		INSERT INTO tasks (title, description, is_done, created_at, done_at)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err = r.conn.Exec(ctx, query, task.Title, task.Description, task.IsDone, task.CreateAt, task.DoneAt)
	return err
}

func (r *Repository) GetTask(ctx context.Context, title string) (Task, error) {
	query := `
		SELECT title, description, is_done, created_at, done_at
		FROM tasks
		WHERE title = $1
	`

	var task Task
	var doneAt *time.Time

	err := r.conn.QueryRow(ctx, query, title).Scan(
		&task.Title,
		&task.Description,
		&task.IsDone,
		&task.CreateAt,
		&doneAt,
	)

	if err == pgx.ErrNoRows {
		return Task{}, ErrorTaskNotFound
	}
	if err != nil {
		return Task{}, err
	}

	task.DoneAt = doneAt
	return task, nil
}

func (r *Repository) ListTasks(ctx context.Context) (map[string]Task, error) {
	query := `
		SELECT title, description, is_done, created_at, done_at
		FROM tasks
		ORDER BY created_at DESC
	`

	rows, err := r.conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make(map[string]Task)

	for rows.Next() {
		var task Task
		var doneAt *time.Time

		err := rows.Scan(
			&task.Title,
			&task.Description,
			&task.IsDone,
			&task.CreateAt,
			&doneAt,
		)
		if err != nil {
			return nil, err
		}

		task.DoneAt = doneAt
		tasks[task.Title] = task
	}

	return tasks, rows.Err()
}

func (r *Repository) NotDoneTasks(ctx context.Context) (map[string]Task, error) {
	query := `
		SELECT title, description, is_done, created_at, done_at
		FROM tasks
		WHERE is_done = FALSE
		ORDER BY created_at DESC
	`

	rows, err := r.conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make(map[string]Task)

	for rows.Next() {
		var task Task
		var doneAt *time.Time

		err := rows.Scan(
			&task.Title,
			&task.Description,
			&task.IsDone,
			&task.CreateAt,
			&doneAt,
		)
		if err != nil {
			return nil, err
		}

		task.DoneAt = doneAt
		tasks[task.Title] = task
	}

	return tasks, rows.Err()
}

func (r *Repository) UpdateTaskDone(ctx context.Context, title string, isDone bool) (Task, error) {
	var query string
	var doneAt *time.Time

	if isDone {
		now := time.Now()
		doneAt = &now
		query = `
			UPDATE tasks
			SET is_done = TRUE, done_at = $1
			WHERE title = $2
			RETURNING title, description, is_done, created_at, done_at
		`
	} else {
		query = `
			UPDATE tasks
			SET is_done = FALSE, done_at = NULL
			WHERE title = $1
			RETURNING title, description, is_done, created_at, done_at
		`
	}

	var task Task
	var scannedDoneAt *time.Time

	var err error
	if isDone {
		err = r.conn.QueryRow(ctx, query, doneAt, title).Scan(
			&task.Title,
			&task.Description,
			&task.IsDone,
			&task.CreateAt,
			&scannedDoneAt,
		)
	} else {
		err = r.conn.QueryRow(ctx, query, title).Scan(
			&task.Title,
			&task.Description,
			&task.IsDone,
			&task.CreateAt,
			&scannedDoneAt,
		)
	}

	if err == pgx.ErrNoRows {
		return Task{}, ErrorTaskNotFound
	}
	if err != nil {
		return Task{}, err
	}

	task.DoneAt = scannedDoneAt
	return task, nil
}

func (r *Repository) DeleteTask(ctx context.Context, title string) error {
	query := `DELETE FROM tasks WHERE title = $1`

	result, err := r.conn.Exec(ctx, query, title)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrorTaskNotFound
	}

	return nil
}
