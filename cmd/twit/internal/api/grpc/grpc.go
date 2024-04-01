package grpc

import (
	"context"
	"github.com/ZergsLaw/back-template/cmd/twit/internal/app"
	"github.com/ZergsLaw/back-template/internal/dom"
	"github.com/gofrs/uuid"
)

type application interface {
	TwitPost(ctx context.Context, session dom.Session, text string) (*app.Twit, error)
	TwitGet(ctx context.Context, session dom.Session, id uuid.UUID, limit, offset int) ([]app.Twit, int, error)
	TwitUpdate(ctx context.Context, session dom.Session, id uuid.UUID, text string) (*app.Twit, error)
	TwitDelete(ctx context.Context, session dom.Session, id uuid.UUID) error
}

type api struct {
	app application
}
