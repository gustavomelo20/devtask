package tasks

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type AddTaskCommand struct {
	Title       string
	Description string
}

type AddTaskHandler struct {
	DB *sql.DB
}

func NewAddTaskHandler(db *sql.DB) *AddTaskHandler {
	return &AddTaskHandler{DB: db}
}

func (h *AddTaskHandler) Handle(ctx context.Context, cmd AddTaskCommand) error {
	if cmd.Title == "" {
		return errors.New("título não pode ser vazio")
	}

	query := `INSERT INTO tasks (title, description, done, created_at) VALUES (?, ?, ?, ?)`
	_, err := h.DB.ExecContext(ctx, query, cmd.Title, cmd.Description, false, time.Now())
	if err != nil {
		return err
	}

	return nil
}
