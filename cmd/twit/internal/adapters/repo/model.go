package repo

import (
	"github.com/ZergsLaw/back-template/cmd/twit/internal/app"
	"github.com/gofrs/uuid"
	"time"
)

type dbTwit struct {
	ID        uuid.UUID
	AuthorID  uuid.UUID
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (db dbTwit) convert() *app.Twit {
	return &app.Twit{
		ID:       db.ID,
		AuthorID: db.AuthorID,
		Text:     db.Text,
	}
}
