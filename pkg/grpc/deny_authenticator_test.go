package grpc_test

import (
	"context"
	"testing"

	bb_grpc "github.com/buildbarn/bb-storage/pkg/grpc"
	"github.com/buildbarn/bb-storage/pkg/testutil"
	"github.com/stretchr/testify/require"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestDenyAuthenticator(t *testing.T) {
	authenticator := bb_grpc.NewDenyAuthenticator("This service has been disabled")
	newCtx, err := authenticator.Authenticate(context.Background())
	testutil.RequireEqualStatus(
		t,
		status.Error(codes.Unauthenticated, "This service has been disabled"),
		err)
	require.Equal(t, nil, newCtx)
}
