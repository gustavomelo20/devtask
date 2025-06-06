package tasks

import (
	"context"
	"database/sql"
)

type ListTasksCommand struct {
}

type Task struct {
	ID          int
	Title       string
	Description string
	Done        bool
	CreatedAt   string
}

type ListTasksHandler struct {
	db *sql.DB
}

func NewListTasksHandler(db *sql.DB) *ListTasksHandler {
	return &ListTasksHandler{db: db}
}

func (h *ListTasksHandler) Handle(ctx context.Context, cmd ListTasksCommand) ([]Task, error) {
	rows, err := h.db.QueryContext(ctx, "SELECT id, title, description, done, created_at FROM tasks ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var t Task
		var description sql.NullString

		if err := rows.Scan(&t.ID, &t.Title, &description, &t.Done, &t.CreatedAt); err != nil {
			return nil, err
		}

		if description.Valid {
			t.Description = description.String
		}

		tasks = append(tasks, t)
	}

	return tasks, rows.Err()
}
