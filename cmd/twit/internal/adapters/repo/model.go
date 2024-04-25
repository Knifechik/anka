package repo

import (
	"github.com/ZergsLaw/back-template/cmd/twit/internal/app"
	"github.com/gofrs/uuid"
	"time"
)

type dbTwit struct {
	ID        uuid.UUID `db:"id"`
	AuthorID  uuid.UUID `db:"author_id"`
	Text      string    `db:"text"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (db dbTwit) convert() *app.Twit {
	return &app.Twit{
		ID:        db.ID,
		AuthorID:  db.AuthorID,
		Text:      db.Text,
		CreatedAt: db.CreatedAt,
		UpdatedAt: db.UpdatedAt,
	}
}
