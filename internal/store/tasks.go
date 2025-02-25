package store

import (
	"context"
	"database/sql"
	"time"
)

type Task struct {
	ID          int64     `json:"id" db:"id"`
	UserID      int64     `json:"user_id" db:"user_id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Completed   bool      `json:"completed" db:"completed"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	DueAt       time.Time `json:"due_at" db:"due_at"`
}

type TaskStore struct {
	db *sql.DB
}

func (s *TaskStore) Create(ctx context.Context, task *Task) error {
}
