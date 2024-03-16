package grpc

import (
	"github.com/ZergsLaw/back-template/cmd/twit/internal/app"
	"golang.org/x/net/context"
)

type application interface {
	TwitPost(ctx context.Context, text string) (app.Twit, error)
	TwitGet(ctx context.Context, id string) (app.Twit, error)
	TwitUpdate(ctx context.Context, id, text string) (app.Twit, error)
	TwitDelete(ctx context.Context, id string) error
}

type api struct {
	app application
}
