package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"path"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"

	"github.com/ZergsLaw/back-template/cmd/twit/internal/app"
	"github.com/ZergsLaw/back-template/internal/adapters/session"
)

const (
	userAgentForward = `grpcgateway-user-agent`
	userAgent        = `user-agent`
	auth             = `authorization`
	scheme           = `Bearer`
)

var ErrUnauthenticated = status.Error(codes.Unauthenticated, "unauthenticated")

// AuthFuncOverride implements grpc_auth.ServiceAuthFuncOverride.
func (a *api) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	_, method := path.Split(fullMethodName)
	if !a.auth[method] {
		return ctx, nil
	}

	token, err := grpc_auth.AuthFromMD(ctx, scheme)
	if err != nil {
		return nil, fmt.Errorf("grpc_auth.AuthFromMD: %w", err)
	}

	userSession, err := a.app.Auth(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", app.ErrInvalidAuth, err)
	}

	return session.NewContext(ctx, userSession), nil
}
