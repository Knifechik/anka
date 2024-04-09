package grpc

import (
	"context"
	"fmt"
	twit_pb "github.com/ZergsLaw/back-template/api/twit/v1"
	"github.com/ZergsLaw/back-template/internal/adapters/session"
	"github.com/ZergsLaw/back-template/internal/dom"
	"github.com/gofrs/uuid"
)

func (a *api) TwitPost(ctx context.Context, request *twit_pb.TwitPostRequest) (*twit_pb.TwitPostResponse, error) {
	twit, err := a.app.TwitPost(ctx, dom.Session{}, request.Text)
	if err != nil {
		return nil, fmt.Errorf("a.app.TwitPost: %w", err)
	}

	return &twit_pb.TwitPostResponse{Twit: toTwit(*twit)}, nil
}

func (a *api) TwitGet(ctx context.Context, request *twit_pb.TwitGetRequest) (*twit_pb.TwitGetResponse, error) {
	userSession := session.FromContext(ctx)
	twit, total, err := a.app.TwitGet(ctx, *userSession, uuid.FromStringOrNil(request.AuthorId), int(request.Limit), int(request.Offset))
	if err != nil {
		return nil, fmt.Errorf("a.app.TwitGet: %w", err)
	}

	return &twit_pb.TwitGetResponse{Twit: toTwits(twit), Total: int32(total)}, nil
}

func (a *api) TwitUpdate(ctx context.Context, request *twit_pb.TwitUpdateRequest) (*twit_pb.TwitUpdateResponse, error) {
	userSession := session.FromContext(ctx)
	twit, err := a.app.TwitUpdate(ctx, *userSession, uuid.FromStringOrNil(request.Id), request.Text)
	if err != nil {
		return nil, fmt.Errorf("a.app.TwitUpdate: %w", err)
	}

	return &twit_pb.TwitUpdateResponse{Twit: toTwit(*twit)}, nil
}

func (a *api) TwitDelete(ctx context.Context, request *twit_pb.TwitDeleteRequest) (*twit_pb.TwitDeleteResponse, error) {
	userSession := session.FromContext(ctx)
	err := a.app.TwitDelete(ctx, *userSession, uuid.FromStringOrNil(request.Id))
	if err != nil {
		return nil, fmt.Errorf("a.app.TwitDelete: %w", err)
	}

	return &twit_pb.TwitDeleteResponse{}, nil
}
