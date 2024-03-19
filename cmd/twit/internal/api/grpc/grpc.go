package grpc

import (
	"context"
	"github.com/ZergsLaw/back-template/cmd/twit/internal/app"
	"github.com/gofrs/uuid"
)

type application interface {
	TwitPost(ctx context.Context, text string) (*app.Twit, error)
	TwitGet(ctx context.Context, id uuid.UUID) (*app.Twit, error)
	TwitUpdate(ctx context.Context, id uuid.UUID, text string) (*app.Twit, error)
	TwitDelete(ctx context.Context, id uuid.UUID) error
}

type api struct {
	app application
}
