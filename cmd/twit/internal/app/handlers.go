package app

import (
	"context"
	"fmt"
	"github.com/ZergsLaw/back-template/internal/dom"
	"github.com/gofrs/uuid"
)

// Auth get user session by token.
func (a *App) Auth(ctx context.Context, token string) (*dom.Session, error) {
	return a.sessions.Get(ctx, token)
}

func (a *App) TwitPost(ctx context.Context, session dom.Session, text string) (*Twit, error) {
	twit := Twit{Text: text, AuthorID: session.UserID}
	res, err := a.repo.Save(ctx, twit)
	if err != nil {
		return res, fmt.Errorf("repo.Save: %w", err)
	}

	return res, nil
}

func (a *App) TwitGet(ctx context.Context, _ dom.Session, authorId uuid.UUID, limit, offset int) ([]Twit, int, error) {
	res, total, err := a.repo.Search(ctx, authorId, limit, offset)
	if err != nil {
		return res, 0, fmt.Errorf("repo.ByID: %w", err)
	}

	return res, total, nil
}

func (a *App) TwitUpdate(ctx context.Context, session dom.Session, id uuid.UUID, text string) (*Twit, error) {
	checkTwit, err := a.repo.ByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("repo.ByID: %w", err)
	}
	if checkTwit.AuthorID != session.UserID {
		return nil, ErrAccessDenied
	}
	twit := Twit{ID: id, Text: text}
	res, err := a.repo.Update(ctx, twit)
	if err != nil {
		return res, fmt.Errorf("repo.Update: %w", err)
	}
	return res, nil
}
func (a *App) TwitDelete(ctx context.Context, session dom.Session, id uuid.UUID) error {
	checkTwit, err := a.repo.ByID(ctx, id)
	if err != nil {
		return fmt.Errorf("repo.ByID: %w", err)
	}
	if checkTwit.AuthorID != session.UserID {
		return ErrAccessDenied
	}
	twit := Twit{ID: id}
	err = a.repo.Delete(ctx, twit)
	if err != nil {
		return fmt.Errorf("repo.Delete: %w", err)
	}
	return nil
}
