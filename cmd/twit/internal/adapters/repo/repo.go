package repo

import (
	"context"
	"fmt"
	"github.com/ZergsLaw/back-template/cmd/twit/internal/app"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sipki-tech/database"
	"github.com/sipki-tech/database/connectors"
	"github.com/sipki-tech/database/migrations"
)

var _ app.Repo = &Repo{}

type (
	// Config provide connection info for database.
	Config struct {
		Cockroach  connectors.CockroachDB
		MigrateDir string
		Driver     string
	}
	// Repo provided data from and to database.
	Repo struct {
		sql *database.SQL
	}
)

func New(ctx context.Context, reg *prometheus.Registry, namespace string, cfg Config) (*Repo, error) {
	const subsystem = "repo"
	m := database.NewMetrics(reg, namespace, subsystem, new(app.Repo))

	returnErrs := []error{ // List of app.Err… returned by Repo methods.
		app.ErrNotFound,
	}

	migrates, err := migrations.Parse(cfg.MigrateDir)
	if err != nil {
		return nil, fmt.Errorf("migrations.Parse: %w", err)
	}

	err = migrations.Run(ctx, cfg.Driver, &cfg.Cockroach, migrations.Up, migrates)
	if err != nil {
		return nil, fmt.Errorf("migrations.Run: %w", err)
	}

	conn, err := database.NewSQL(ctx, cfg.Driver, database.SQLConfig{
		Metrics:    m,
		ReturnErrs: returnErrs,
	}, &cfg.Cockroach)
	if err != nil {
		return nil, fmt.Errorf("librepo.NewCockroach: %w", err)
	}

	return &Repo{
		sql: conn,
	}, nil
}

func (r *Repo) Close() error {
	return r.sql.Close()
}

func (r *Repo) Save(ctx context.Context, t app.Twit) (twit *app.Twit, err error) {
	dbtwit := dbTwit{}
	err = r.sql.NoTx(func(db *sqlx.DB) error {
		const query = `INSERT INTO twit_table (author_id, text) VALUES ($1, $2) RETURNING *`
		err = db.GetContext(ctx, &dbtwit, query, t.AuthorID, t.Text)

		if err != nil {
			return fmt.Errorf("sql.QueryRowContext: %w", err)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return dbtwit.convert(), nil
}

func (r *Repo) ByID(ctx context.Context, id uuid.UUID) (twit *app.Twit, err error) {
	dbtwit := dbTwit{}
	err = r.sql.NoTx(func(db *sqlx.DB) error {
		const query = `SELECT * FROM twit_table WHERE id = $1`
		err = db.GetContext(ctx, &dbtwit, query, id)
		if err != nil {
			return fmt.Errorf("sql.QueryRowContext: %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	} 

	return dbtwit.convert(), nil
}

func (r *Repo) Update(ctx context.Context, t app.Twit) (twit *app.Twit, err error) {
	dbtwit := dbTwit{}
	err = r.sql.NoTx(func(db *sqlx.DB) error {
		const query = `UPDATE twit_table SET text = $1, updated_at = now() WHERE id = $2 RETURNING *`
		err = db.GetContext(ctx, &dbtwit, query, t.Text, t.ID)
		if err != nil {
			return fmt.Errorf("sql.GetContextContext: %w", err)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return dbtwit.convert(), nil
}

func (r *Repo) Delete(ctx context.Context, t app.Twit) (err error) {
	err = r.sql.NoTx(func(db *sqlx.DB) error {
		const query = `DELETE FROM twit_table WHERE id = $1`
		_, err := db.ExecContext(ctx, query, t.ID)
		if err != nil {
			return fmt.Errorf("db.ExecContext: %w", err)
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) Search(ctx context.Context, authorId uuid.UUID, limit, offset int) (twits []app.Twit, total int, err error) {
	dbtwit := []dbTwit{}
	err = r.sql.NoTx(func(db *sqlx.DB) error {
		const query = `SELECT * FROM twit_table WHERE author_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`
		err = db.SelectContext(ctx, &dbtwit, query, authorId, limit, offset)
		if err != nil {
			return fmt.Errorf("db.SelectContext: %w", err)
		}

		const getTotal = `SELECT count(*) over() AS total FROM twit_table WHERE author_id = $1`
		err = db.GetContext(ctx, &total, getTotal, authorId)
		if err != nil {
			return fmt.Errorf("db.GetContext: %w", err)
		}

		for i := range dbtwit {
			twits = append(twits, *dbtwit[i].convert())
		}

		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	return twits, total, nil
}
