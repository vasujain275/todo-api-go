package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Tasks interface {
		Create(context.Context, *Task) error
	}

	Users interface {
		Create(context.Context, *User) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Tasks: &TaskStore{db},
		Users: &UserStore{db},
	}
}
