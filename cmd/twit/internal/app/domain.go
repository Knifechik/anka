package app

import (
	"github.com/gofrs/uuid"
	"time"
)

type Twit struct {
	ID        uuid.UUID
	AuthorID  uuid.UUID
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
