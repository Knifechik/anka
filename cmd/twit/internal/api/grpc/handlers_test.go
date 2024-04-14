package grpc_test

import (
	"fmt"
	twit_pb "github.com/ZergsLaw/back-template/api/twit/v1"
	"github.com/ZergsLaw/back-template/cmd/twit/internal/app"
	"github.com/ZergsLaw/back-template/internal/dom"
	"github.com/gofrs/uuid"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestApi_TwitDelete(t *testing.T) {
	t.Parallel()

	var (
		id          = uuid.Must(uuid.NewV4())
		errInternal = status.Error(codes.Internal, fmt.Sprintf("a.app.TwitDelete: %s", errAny))
	)

	testCases := map[string]struct {
		id      uuid.UUID
		appErr  error
		wantErr error
	}{
		"success":          {id, nil, nil},
		"a.app.TwitDelete": {id, errAny, errInternal},
	}

	for name, tc := range testCases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx, c, mockApp, assert := start(t, dom.UserStatusDefault)

			mockApp.EXPECT().TwitDelete(gomock.Any(), session, tc.id).Return(tc.appErr)

			_, err := c.TwitDelete(auth(ctx), &twit_pb.TwitDeleteRequest{
				Id: tc.id.String(),
			})
			assert.ErrorIs(err, tc.wantErr)
		})
	}
}

func TestApi_TwitPost(t *testing.T) {
	t.Parallel()

	var (
		text        = "ya ne hochu etogo govna"
		id          = uuid.Must(uuid.NewV4())
		res         = &app.Twit{ID: id, Text: text}
		errInternal = status.Error(codes.Internal, fmt.Sprintf("a.app.TwitPost: %s", errAny))
	)

	testCases := map[string]struct {
		res     *app.Twit
		text    string
		appErr  error
		wantErr error
	}{
		"success":        {res, text, nil, nil},
		"a.app.TwitPost": {res, text, errAny, errInternal},
	}

	for name, tc := range testCases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx, c, mockApp, assert := start(t, dom.UserStatusDefault)

			mockApp.EXPECT().TwitPost(gomock.Any(), session, tc.text).Return(tc.res, tc.appErr)

			_, err := c.TwitPost(auth(ctx), &twit_pb.TwitPostRequest{
				Text: tc.text,
			})
			assert.ErrorIs(err, tc.wantErr)
			assert.Equal(text, tc.res.Text)
		})
	}
}
