package grpc

import (
	"context"
)

func (a *api) TwitPost(context.Context, *TwitPostRequest) (*TwitPostResponse, error)       {}
func (a *api) TwitGet(context.Context, *TwitGetRequest) (*TwitGetResponse, error)          {}
func (a *api) TwitUpdate(context.Context, *TwitUpdateRequest) (*TwitUpdateResponse, error) {}
func (a *api) TwitDelete(context.Context, *TwitDeleteRequest) (*TwitDeleteResponse, error) {}
