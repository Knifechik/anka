package grpc

import (
	"context"
	"errors"
	pb "github.com/ZergsLaw/back-template/api/twit/v1"
	"github.com/ZergsLaw/back-template/cmd/twit/internal/app"
	"github.com/ZergsLaw/back-template/internal/dom"
	"github.com/ZergsLaw/back-template/internal/grpchelper"
	"github.com/ZergsLaw/back-template/internal/logger"
	"github.com/ZergsLaw/back-template/internal/metrics"
	"github.com/gofrs/uuid"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

type application interface {
	Auth(ctx context.Context, token string) (*dom.Session, error)
	TwitPost(ctx context.Context, session dom.Session, text string) (*app.Twit, error)
	TwitGet(ctx context.Context, session dom.Session, authorID uuid.UUID, limit, offset int) ([]app.Twit, int, error)
	TwitUpdate(ctx context.Context, session dom.Session, id uuid.UUID, text string) (*app.Twit, error)
	TwitDelete(ctx context.Context, session dom.Session, id uuid.UUID) error
}

type api struct {
	pb.UnimplementedTwitAPIServer
	app  application
	auth map[string]bool
}

// New creates and returns gRPC server.
func New(ctx context.Context, m metrics.Metrics, applications application, reg *prometheus.Registry, namespace string) *grpc.Server {
	log := logger.FromContext(ctx)
	subsystem := "api"
	grpcMetrics := grpchelper.NewServerMetrics(reg, namespace, subsystem)

	srv, health := grpchelper.NewServer(m, log, grpcMetrics, apiError,
		[]grpc.UnaryServerInterceptor{grpc_auth.UnaryServerInterceptor(nil)},   // Nil because we are using override.
		[]grpc.StreamServerInterceptor{grpc_auth.StreamServerInterceptor(nil)}, // Nil because we are using override.
	)
	health.SetServingStatus(pb.TwitAPI_ServiceDesc.ServiceName, healthpb.HealthCheckResponse_SERVING)

	pb.RegisterTwitAPIServer(srv, &api{
		app: applications,
		auth: map[string]bool{
			"TwitPost":   true,
			"TwitGet":    false,
			"TwitUpdate": true,
			"TwitDelete": true,
		},
	})

	return srv
}

func apiError(err error) *status.Status {
	if err == nil {
		return nil
	}
	code := codes.Internal
	switch {
	case errors.Is(err, app.ErrAccessDenied):
		code = codes.PermissionDenied
	case errors.Is(err, app.ErrNotFound):
		code = codes.NotFound
	case errors.Is(err, context.DeadlineExceeded):
		code = codes.DeadlineExceeded
	case errors.Is(err, context.Canceled):
		code = codes.Canceled
	}

	return status.New(code, err.Error())
}
