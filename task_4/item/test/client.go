//go:build api

package integration

import (
	"testing"

	pb "github.com/KolesnikDmitriy/item/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const target = "localhost:50051"

func newItemClient(t *testing.T) pb.ItemClient {
	t.Helper()

	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("failed to create grpc client: %v", err)
	}
	t.Cleanup(func() { _ = conn.Close() })

	return pb.NewItemClient(conn)
}
