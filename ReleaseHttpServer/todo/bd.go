package todo

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Bd struct {
	conn *pgx.Conn
}

func NewBd(conn *pgx.Conn) *Bd {
	return &Bd{
		conn: conn,
	}
}

func (bd *Bd) AddTaskBd(ctx context.Context, task Task) {
	sqlQuery := `
		INSERT INTO tasks (title, description, isdone, createat, doneat)
		VALUES ($1, $2, $3, $4, $5)
	`

}
