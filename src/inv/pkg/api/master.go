package api

import (
	"context"
	"database/sql"

	"github.com/Posrabi/shopify-backend-project/src/inv/pkg/entity"
	"github.com/Posrabi/shopify-backend-project/src/inv/pkg/postgres"
	"github.com/Posrabi/shopify-backend-project/src/inv/pkg/repository"
)

type masterRepo struct {
	item repository.Item
}

func NewMasterRepository(db *sql.DB, openWeatherToken string) repository.Master {
	return &masterRepo{
		item: postgres.NewItemRepository(db, openWeatherToken),
	}
}

func (m *masterRepo) CreateItem(ctx context.Context, item *entity.Item) error {
	return m.item.CreateItem(ctx, item)
}

func (m *masterRepo) DeleteItem(ctx context.Context, itemID string) error {
	return m.item.DeleteItem(ctx, itemID)
}

func (m *masterRepo) EditItem(ctx context.Context, item *entity.Item) error {
	return m.item.EditItem(ctx, item)
}

func (m *masterRepo) ListItems(ctx context.Context) ([]*entity.Item, error) {
	return m.item.ListItems(ctx)
}

func (m *masterRepo) Ship(ctx context.Context, shipment *entity.Shipment) error {
	// This is a good place to update the shipments table if we have one.
	return m.item.ShipItem(ctx, shipment.ItemInstance)
}
