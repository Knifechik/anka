package app

import (
	"github.com/ZergsLaw/back-template/internal/dom"
	"github.com/gofrs/uuid"
	"golang.org/x/net/context"
)

type (
	Repo interface {
		Save(ctx context.Context, twit Twit) (*Twit, error)
		ByID(ctx context.Context, id uuid.UUID) (*Twit, error)
		Update(ctx context.Context, twit Twit) (*Twit, error)
		Delete(ctx context.Context, twit Twit) error
		Search(ctx context.Context, authorId uuid.UUID, limit, offset int) ([]Twit, int, error)
	}
	// Sessions module for manager user's session.
	Sessions interface {
		// Get returns user session by his token.
		// Errors: ErrNotFound, unknown.
		Get(context.Context, string) (*dom.Session, error)
	}
)
