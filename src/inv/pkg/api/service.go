package api

import (
	"context"

	"github.com/Posrabi/shopify-backend-project/src/inv/pkg/entity"
	"github.com/Posrabi/shopify-backend-project/src/inv/pkg/repository"
)

type Service struct {
	repo repository.Master
}

func NewService(repo repository.Master) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateItem(ctx context.Context, item *entity.Item) error {
	return s.repo.CreateItem(ctx, item)
}

func (s *Service) EditItem(ctx context.Context, item *entity.Item) error {
	return s.repo.EditItem(ctx, item)
}

func (s *Service) DeleteItem(ctx context.Context, itemID string) error {
	return s.repo.DeleteItem(ctx, itemID)
}

func (s *Service) ListItems(ctx context.Context) ([]*entity.Item, error) {
	return s.repo.ListItems(ctx)
}

func (s *Service) ShipItem(ctx context.Context, shipment *entity.Shipment) error {
	return s.repo.Ship(ctx, shipment)
}
