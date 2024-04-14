package grpc_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"net"
	"testing"

	pb "github.com/ZergsLaw/back-template/api/twit/v1"
	grpcapi "github.com/ZergsLaw/back-template/cmd/twit/internal/api/grpc"
	"github.com/ZergsLaw/back-template/internal/dom"
	"github.com/ZergsLaw/back-template/internal/metrics"
	"github.com/ZergsLaw/back-template/internal/testhelper"
)

var (
	token   = dom.Token{Value: "token"}
	session = dom.Session{
		ID:     uuid.Must(uuid.NewV4()),
		UserID: uuid.Must(uuid.NewV4()),
		Status: dom.UserStatusDefault,
	}
	adminSession = dom.Session{
		ID:     uuid.Must(uuid.NewV4()),
		UserID: uuid.Must(uuid.NewV4()),
		Status: dom.UserStatusAdmin,
	}

	errAny = errors.New("any err")
	origin = dom.Origin{
		IP:        net.ParseIP("127.0.0.1"),
		UserAgent: "grpc-go/1.59.0",
	}
)

type clients struct {
	pb.TwitAPIClient
}

func start(t *testing.T, userStatus dom.UserStatus) (context.Context, *clients, *Mockapplication, *require.Assertions) {
	t.Helper()
	assert := require.New(t)
	ctx := testhelper.Context(t)

	ctrl := gomock.NewController(t)
	mockApp := NewMockapplication(ctrl)

	namespace := testhelper.Namespace(t)

	reg := prometheus.NewPedanticRegistry()
	m := metrics.New(reg, namespace)

	server := grpcapi.New(ctx, m, mockApp, reg, namespace)
	addr := testhelper.Addr(t, assert)
	ln, err := net.Listen("tcp", addr.String())
	assert.NoError(err)

	go func() {
		err := server.Serve(ln)
		assert.NoError(err)
	}()

	switch userStatus {
	case dom.UserStatusAdmin:
		mockApp.EXPECT().Auth(gomock.Any(), token.Value).Return(&adminSession, nil).AnyTimes()
	default:
		mockApp.EXPECT().Auth(gomock.Any(), token.Value).Return(&session, nil).AnyTimes()
	}

	conn, err := grpc.DialContext(ctx, addr.String(),
		grpc.WithTransportCredentials(insecure.NewCredentials()), // TODO Add TLS and remove this.
		grpc.WithBlock(),
	)
	assert.NoError(err)

	t.Cleanup(func() {
		err := conn.Close()
		assert.NoError(err)
		assert.NoError(err)
		server.GracefulStop()
	})

	return ctx, &clients{TwitAPIClient: pb.NewTwitAPIClient(conn)}, mockApp, assert
}

func auth(ctx context.Context) context.Context {
	return metadata.NewOutgoingContext(ctx, metadata.MD{
		"authorization": {fmt.Sprintf("Bearer %s", token.Value)},
	})
}
