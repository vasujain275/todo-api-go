package store

import (
	"context"
	"database/sql"
	"time"
)

type User struct {
	ID        int64     `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"-" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Create(ctx context.Context, user *User) error {
	query := `INSERT INTO users (username, password)
	VALUES ($1, $2) RETURNING id, created_at, updated_at`

	err := s.db.QueryRowContext(ctx,
		query,
		user.Username,
		user.Password).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	return err
}
