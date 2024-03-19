package app

import "golang.org/x/net/context"

type (
	Repo interface {
		Save(ctx context.Context, twit Twit) (*Twit, error)
		ByID(ctx context.Context, twit Twit) (*Twit, error)
		Update(ctx context.Context, twit Twit) (*Twit, error)
		Delete(ctx context.Context, twit Twit) error
	}
)
