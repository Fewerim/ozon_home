package app

import (
	"github.com/KolesnikDmitriy/item/internal/analytics"
	"github.com/KolesnikDmitriy/item/internal/storage"
	pb "github.com/KolesnikDmitriy/item/pkg/api"
)

type ItemService struct {
	pb.ItemServer

	analytics *analytics.Analytics

	storage storage.Storage
}

func NewItemService(st storage.Storage, a *analytics.Analytics) *ItemService {
	return &ItemService{
		storage:   st,
		analytics: a,
	}
}
