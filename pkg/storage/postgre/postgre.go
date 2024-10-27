package postgres

import (
	"GoNews/pkg/storage"
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Postgres struct {
	db *pgxpool.Pool
}

func New(connection string) (*Postgres, error) {
	db, err := pgxpool.Connect(context.Background(), connection)
	if err != nil {
		return nil, err
	}
	s := Postgres{
		db: db,
	}
	return &s, nil
}

func (s *Postgres) Posts() ([]storage.Post, error) {
	rows, err := s.db.Query(context.Background(), `SELECT * FROM posts;`)
	if err != nil {
		return nil, err
	}
	var posts []storage.Post
	for rows.Next() {
		var post storage.Post
		err = rows.Scan(
			&post.ID,
			&post.AuthorID,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)

	}
	return posts, rows.Err()
}

func (s *Postgres) AddPost(post storage.Post) error {
	_, err := s.db.Exec(context.Background(), `
		INSERT INTO posts (author_id, title, content, created_at)
		VALUES ($1, $2, $3, $4);
		`,
		post.AuthorID,
		post.Title,
		post.Content,
		post.CreatedAt,
	)
	return err
}
func (s *Postgres) UpdatePost(post storage.Post) error {
	_, err := s.db.Exec(context.Background(), `
	UPDATE posts
	SET 
		author_id = $2,
		title = $3,
		content = $4,
		created_at = $5
	WHERE
   		id = $1;
	`,
		post.ID,
		post.AuthorID,
		post.Title,
		post.Content,
		post.CreatedAt,
	)

	return err
}
func (s *Postgres) DeletePost(post storage.Post) error {
	_, err := s.db.Exec(context.Background(), `
	DELETE FROM posts WHERE id = $1;
	`,
		post.ID,
	)

	return err
}
