//go:build integration

package repo_test

import (
	"net"
	"testing"
	"time"

	"github.com/gofrs/uuid"

	"github.com/ZergsLaw/back-template/cmd/session/internal/app"
	"github.com/ZergsLaw/back-template/internal/dom"
)

func TestRepo_Smoke(t *testing.T) {
	t.Parallel()

	ctx, r, assert := start(t)

	session := app.Session{
		ID: uuid.Must(uuid.NewV4()),
		Origin: app.Origin{
			IP:        net.ParseIP("192.100.10.4"),
			UserAgent: "Mozilla/5.0",
		},
		Token: app.Token{
			Value: "token",
		},
		Status:    dom.UserStatusDefault,
		UserID:    uuid.Must(uuid.NewV4()),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := r.Save(ctx, session)
	assert.NoError(err)

	res, err := r.ByID(ctx, session.ID)
	assert.NoError(err)
	session.CreatedAt = res.CreatedAt
	session.UpdatedAt = res.UpdatedAt
	if session.Origin.IP.Equal(res.Origin.IP) {
		session.Origin.IP = res.Origin.IP
	}
	assert.Equal(session, *res)

	upStatusID := uuid.Must(uuid.NewV4())

	err = r.UpdateStatus(ctx, upStatusID, session.UserID, dom.UserStatusDefault)
	assert.NoError(err)

	err = r.UpdateStatus(ctx, upStatusID, session.UserID, dom.UserStatusDefault)
	assert.ErrorIs(err, app.ErrDuplicate)

	res, err = r.ByID(ctx, session.ID)
	assert.NoError(err)
	assert.Equal(res.Status, dom.UserStatusDefault)

	err = r.Delete(ctx, session.ID)
	assert.NoError(err)

	res, err = r.ByID(ctx, session.ID)
	assert.Nil(res)
	assert.ErrorIs(err, app.ErrNotFound)
}
