//go:build integration

package repo_test

import (
	"github.com/ZergsLaw/back-template/cmd/twit/internal/app"
	"time"

	"github.com/gofrs/uuid"
	"testing"
)

func TestRepo_Smoke(t *testing.T) {
	t.Parallel()

	ctx, r, assert := start(t)

	twit := app.Twit{
		ID:        uuid.Nil,
		AuthorID:  uuid.Must(uuid.NewV4()),
		Text:      "Hello World",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	twit2 := twit

	res, err := r.Save(ctx, twit)
	assert.NoError(err)
	assert.NotNil(res.ID)

	twit.ID = res.ID

	res, err = r.ByID(ctx, twit.ID)
	assert.NoError(err)
	res.CreatedAt = time.Time{}
	res.UpdatedAt = time.Time{}
	twit.CreatedAt = time.Time{}
	twit.UpdatedAt = time.Time{}
	assert.Equal(*res, twit)

	twit.Text = "Goodbye World"
	res, err = r.Update(ctx, twit)
	assert.NoError(err)
	assert.Equal("Goodbye World", res.Text)

	res, err = r.Save(ctx, twit2)
	assert.NoError(err)
	assert.NotNil(res.ID)

	reslice, total, err := r.Search(ctx, res.AuthorID, 10, 0)
	assert.NoError(err)
	assert.Equal(2, total)
	assert.Len(reslice, 2)

	err = r.Delete(ctx, *res)
	assert.NoError(err)

	err = r.Delete(ctx, twit)
	assert.NoError(err)
}
