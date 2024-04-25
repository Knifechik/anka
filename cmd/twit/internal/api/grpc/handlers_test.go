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
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
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
		id       = uuid.Must(uuid.NewV4())
		authorId = uuid.Must(uuid.NewV4())
		text     = "ya ne hochu etogo govna"
		now      = time.Now()
		appRes   = &app.Twit{
			ID:        id,
			AuthorID:  authorId,
			Text:      text,
			CreatedAt: now,
			UpdatedAt: now,
		}
		res = &twit_pb.TwitPostResponse{
			Twit: &twit_pb.Twit{
				Id:        id.String(),
				AuthorId:  authorId.String(),
				Text:      text,
				CreatedAt: timestamppb.New(now),
				UpdatedAt: timestamppb.New(now),
			},
		}
		errInternal = status.Error(codes.Internal, fmt.Sprintf("a.app.TwitPost: %s", errAny))
	)

	testCases := map[string]struct {
		text    string
		appRes  *app.Twit
		appErr  error
		wantRes *twit_pb.TwitPostResponse
		wantErr error
	}{
		"success":        {text, appRes, nil, res, nil},
		"a.app.TwitPost": {text, nil, errAny, nil, errInternal},
	}

	for name, tc := range testCases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx, c, mockApp, assert := start(t, dom.UserStatusDefault)

			mockApp.EXPECT().TwitPost(gomock.Any(), session, tc.text).Return(tc.appRes, tc.appErr)

			resp, err := c.TwitPost(auth(ctx), &twit_pb.TwitPostRequest{
				Text: tc.text,
			})
			assert.ErrorIs(err, tc.wantErr)
			assert.True(proto.Equal(resp, tc.wantRes))
		})
	}
}

func TestApi_TwitUpdate(t *testing.T) {
	t.Parallel()

	var (
		ID       = uuid.Must(uuid.NewV4())
		authorId = uuid.Must(uuid.NewV4())
		text     = "ya ne hochu etogo govna"
		now      = time.Now()
		appRes   = &app.Twit{
			ID:        ID,
			AuthorID:  authorId,
			Text:      text,
			CreatedAt: now,
			UpdatedAt: now,
		}
		res = &twit_pb.TwitUpdateResponse{
			Twit: &twit_pb.Twit{
				Id:        ID.String(),
				AuthorId:  authorId.String(),
				Text:      text,
				CreatedAt: timestamppb.New(now),
				UpdatedAt: timestamppb.New(now),
			}}
		errInternal = status.Error(codes.Internal, fmt.Sprintf("a.app.TwitUpdate: %s", errAny))
	)

	testCases := map[string]struct {
		appRes  *app.Twit
		appErr  error
		wantRes *twit_pb.TwitUpdateResponse
		wantErr error
	}{
		"success":          {appRes, nil, res, nil},
		"a.app.TwitUpdate": {appRes, errAny, nil, errInternal},
	}

	for name, tc := range testCases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx, c, mockApp, assert := start(t, dom.UserStatusDefault)

			mockApp.EXPECT().TwitUpdate(gomock.Any(), session, ID, text).Return(tc.appRes, tc.appErr)

			resp, err := c.TwitUpdate(auth(ctx), &twit_pb.TwitUpdateRequest{
				Id:   ID.String(),
				Text: text,
			})
			assert.ErrorIs(err, tc.wantErr)
			assert.True(proto.Equal(resp, tc.wantRes))
		})
	}
}

func TestApi_TwitGet(t *testing.T) {
	t.Parallel()

	var (
		id       = uuid.Must(uuid.NewV4())
		authorId = uuid.Must(uuid.NewV4())
		text     = "ya ne hochu etogo govna"
		limit    = 10
		offset   = 0
		now      = time.Now()

		appTwitRes = []app.Twit{
			{
				ID:        id,
				AuthorID:  authorId,
				Text:      text,
				CreatedAt: now,
				UpdatedAt: now,
			},
		}
		appTotalRes = 1745
		wantRes     = twit_pb.TwitGetResponse{Twit: []*twit_pb.Twit{
			{
				Id:        id.String(),
				AuthorId:  authorId.String(),
				Text:      text,
				CreatedAt: timestamppb.New(now),
				UpdatedAt: timestamppb.New(now),
			},
		},
			Total: 1745}

		errInternal = status.Error(codes.Internal, fmt.Sprintf(
			"a.app.TwitGet: %s", errAny,
		))
	)

	testCases := map[string]struct {
		appTwitRes  []app.Twit
		appTotalRes int
		appErr      error
		wantRes     *twit_pb.TwitGetResponse
		wantErr     error
	}{
		"success":        {appTwitRes, appTotalRes, nil, &wantRes, nil},
		"a.app.TwitPost": {appTwitRes, appTotalRes, errAny, nil, errInternal},
	}

	for name, tc := range testCases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx, c, mockApp, assert := start(t, dom.UserStatusDefault)

			mockApp.EXPECT().TwitGet(gomock.Any(), dom.Session{}, authorId, limit, offset).Return(tc.appTwitRes, tc.appTotalRes, tc.appErr)

			resp, err := c.TwitGet(auth(ctx), &twit_pb.TwitGetRequest{
				AuthorId: authorId.String(),
				Limit:    int32(limit),
				Offset:   int32(offset),
			})
			assert.ErrorIs(err, tc.wantErr)
			assert.True(proto.Equal(resp, tc.wantRes))
		})
	}
}
