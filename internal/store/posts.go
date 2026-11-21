package store

import (
	"context"
	"database/sql"
)

type Post struct {
	PostID    int64    `json:"post_id"`
	Content   string   `json:"content"`
	Title     string   `json:"tittle"`
	UserId    int64    `json:"user_id"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

type PostStore struct {
	db *sql.DB
}

func (s *PostStore) Create(ctx context.Context, post *Post) error {
	query := `	
		INSERT INTO posts(content,title,user_id,tags)
		VALUES($1,$2,$3,$4) RETURNING post_id,created_at,updated_at
	`
	err := s.db.QueryRowContext(
		ctx,
		query,
		post.Content,
		post.Title,
		post.UserId,
		post.Tags,
	).Scan(
		&post.PostID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil

}
