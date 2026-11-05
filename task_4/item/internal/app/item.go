package app

import (
	"context"
	"time"

	"github.com/KolesnikDmitriy/item/internal/domain"
	pb "github.com/KolesnikDmitriy/item/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ItemService) GetItem(ctx context.Context, in *pb.GetItemRequest) (*pb.GetItemResponse, error) {
	start := time.Now()
	defer func() { s.analytics.Log("GET", time.Since(start)) }()

	if in.Id <= 0 {
		return nil, status.Error(codes.InvalidArgument, "wrong id")
	}

	item, err := s.storage.Get(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetItemResponse{Title: item.Title, Description: item.Description}, nil
}
func (s *ItemService) PostItem(ctx context.Context, in *pb.PostItemRequest) (*pb.PostItemResponse, error) {
	start := time.Now()
	defer func() { s.analytics.Log("POST", time.Since(start)) }()

	if in.Title == "" {
		return nil, status.Error(codes.InvalidArgument, "wrong title")
	}
	if in.Description == "" {
		return nil, status.Error(codes.InvalidArgument, "wrong description")
	}

	id, err := s.storage.Add(ctx, domain.Item{
		Title:       in.Title,
		Description: in.Description,
	})
	if err != nil {
		return nil, err
	}
	return &pb.PostItemResponse{Id: id}, nil
}
