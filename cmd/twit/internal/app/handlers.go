package app

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
)

func (a *App) TwitPost(ctx context.Context, text string) (*Twit, error) {
	twit := Twit{Text: text}
	res, err := a.repo.Save(ctx, twit)
	if err != nil {
		return res, fmt.Errorf("repo.Save: %w", err)
	}

	return res, nil
}

func (a *App) TwitGet(ctx context.Context, id uuid.UUID) (*Twit, error) {
	twit := Twit{ID: id}
	res, err := a.repo.ByID(ctx, twit)
	if err != nil {
		return res, fmt.Errorf("repo.ByID: %w", err)
	}

	return res, nil
}

func (a *App) TwitUpdate(ctx context.Context, id uuid.UUID, text string) (*Twit, error) {
	twit := Twit{ID: id, Text: text}
	res, err := a.repo.Update(ctx, twit)
	if err != nil {
		return res, fmt.Errorf("repo.Update: %w", err)
	}
	return res, nil
}
func (a *App) TwitDelete(ctx context.Context, id uuid.UUID) error {
	twit := Twit{ID: id}
	err := a.repo.Delete(ctx, twit)
	if err != nil {
		return fmt.Errorf("repo.Delete: %w", err)
	}
	return nil
}
