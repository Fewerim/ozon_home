//go:build api

package test

import (
	"testing"

	pb "github.com/Fewerim/ozon_home/task_2/product-service/pkg/api"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewProductClient(t *testing.T) pb.ProductServiceClient {
	t.Helper()

	conn, err := grpc.NewClient(
		"127.0.0.1:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, conn.Close()) })

	return pb.NewProductServiceClient(conn)
}
