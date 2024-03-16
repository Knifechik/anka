package grpc

import (
	"context"
	"fmt"
	twit_pb "github.com/ZergsLaw/back-template/api/twit/v1"
)

func (a *api) TwitPost(ctx context.Context, request *twit_pb.TwitPostRequest) (*twit_pb.TwitPostResponse, error) {
	twit, err := a.app.TwitPost(ctx, request.Text)
	if err != nil {
		return nil, fmt.Errorf("a.app.TwitPost: %w", err)
	}
	return &twit_pb.TwitPostResponse{Twit: toUser(twit)}, nil
}

func (a *api) TwitGet(ctx context.Context, request *twit_pb.TwitGetRequest) (*twit_pb.TwitGetResponse, error) {
	twit, err := a.app.TwitGet(ctx, request.Id)
	if err != nil {
		return nil, fmt.Errorf("a.app.TwitGet: %w", err)
	}
	return &twit_pb.TwitGetResponse{Twit: toUser(twit)}, nil
}

func (a *api) TwitUpdate(ctx context.Context, request *twit_pb.TwitUpdateRequest) (*twit_pb.TwitUpdateResponse, error) {
	twit, err := a.app.TwitUpdate(ctx, request.Id, request.Text)
	if err != nil {
		return nil, fmt.Errorf("a.app.TwitUpdate: %w", err)
	}
	return &twit_pb.TwitUpdateResponse{Twit: toUser(twit)}, nil
}

func (a *api) TwitDelete(ctx context.Context, request *twit_pb.TwitDeleteRequest) (*twit_pb.TwitDeleteResponse, error) {
	err := a.app.TwitDelete(ctx, request.Id)
	if err != nil {
		return nil, fmt.Errorf("a.app.TwitDelete: %w", err)
	}
	return &twit_pb.TwitDeleteResponse{}, nil
}
