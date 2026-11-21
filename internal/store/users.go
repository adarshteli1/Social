package store

import (
	"context"
	"database/sql"
)

type User struct {
	UserID    string `json:"user_id"`
	Username  string `json:"username"`
	Password  string `json:"-"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

type UserStorage struct {
	db *sql.DB
}

func (s *UserStorage) Create(ctx context.Context, user *User) error {
	query := `	
		INSERT INTO posts(username,password,email)
		VALUES($1,$2,$3) RETURNING user_id,created_at	
	`
	err := s.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Password,
		user.Email,
	).Scan(
		&user.UserID,
		&user.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
