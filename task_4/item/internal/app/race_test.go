//go:build race

package app

import (
	"context"
	"sync"
	"testing"

	pb "github.com/KolesnikDmitriy/item/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestConcurrentPostItem(t *testing.T) {
	t.Parallel()

	svc := NewItemService(nil)
	ctx := context.Background()

	var wg sync.WaitGroup
	for range 50 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := svc.PostItem(ctx, &pb.PostItemRequest{Title: "Book", Description: "Very Great Book"})
			if err != nil && status.Code(err) != codes.OK {
				t.Errorf("PostItem failed: %v", err)
			}
		}()
	}
	wg.Wait()
}
