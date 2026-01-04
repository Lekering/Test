package simpletable

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL UNIQUE,
		description TEXT NOT NULL,
		is_done BOOLEAN NOT NULL DEFAULT FALSE,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		done_at TIMESTAMP NULL
	);
	
	CREATE INDEX IF NOT EXISTS idx_tasks_title ON tasks(title);
	CREATE INDEX IF NOT EXISTS idx_tasks_is_done ON tasks(is_done);
	`

	_, err := conn.Exec(ctx, query)
	return err
}
