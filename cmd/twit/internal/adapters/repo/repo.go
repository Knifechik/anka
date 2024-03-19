package repo

import (
	"database/sql"
	"fmt"
	"github.com/ZergsLaw/back-template/cmd/twit/internal/app"
	_ "github.com/lib/pq"
	"golang.org/x/net/context"
)

var _ app.Repo = &Repo{}

type (
	// Repo provided data from and to database.
	Repo struct {
		sql *sql.DB
	}
)

func New(name string, info string) (*Repo, error) {
	db, err := sql.Open(name, info)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &Repo{db}, err
}

func (r *Repo) Close() error {
	return r.sql.Close()
}

func (r *Repo) Save(ctx context.Context, t app.Twit) (*app.Twit, error) {
	const query = `INSERT INTO twit_table (authorid, text) VALUES ($1, $2) RETURNING *`
	twit := dbTwit{}
	err := r.sql.QueryRowContext(ctx, query, t.AuthorID, t.Text).Scan(&twit.ID, &twit.AuthorID, &twit.Text)
	if err != nil {
		return nil, fmt.Errorf("sql.QueryRowContext: %w", err)
	}

	return twit.convert(), nil
}

func (r *Repo) ByID(ctx context.Context, t app.Twit) (*app.Twit, error) {
	twit := dbTwit{}
	const query = `SELECT * FROM twit_table WHERE id = $1`
	err := r.sql.QueryRowContext(ctx, query, t.ID).Scan(&twit.ID, &twit.AuthorID, &twit.Text)
	if err != nil {
		return nil, fmt.Errorf("sql.QueryRowContext: %w", err)
	}

	return twit.convert(), nil
}

func (r *Repo) Update(ctx context.Context, t app.Twit) (*app.Twit, error) {
	twit := dbTwit{}
	const query = `UPDATE twit_table SET text = $1 WHERE id = $2`
	err := r.sql.QueryRowContext(ctx, query, t.ID).Scan(&twit.ID, &twit.AuthorID, &twit.Text)
	if err != nil {
		return nil, fmt.Errorf("sql.QueryRowContext: %w", err)
	}

	return twit.convert(), nil
}

func (r *Repo) Delete(ctx context.Context, t app.Twit) error {
	const query = `DELETE FROM twit_table WHERE id = $1`
	_, err := r.sql.ExecContext(ctx, query, t.ID)
	if err != nil {
		return fmt.Errorf("db.ExecContext: %w", err)
	}

	return nil
}
