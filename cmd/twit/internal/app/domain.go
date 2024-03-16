package app

import "github.com/gofrs/uuid"

type Twit struct {
	ID       uuid.UUID
	AuthorID uuid.UUID
	Text     string
}
