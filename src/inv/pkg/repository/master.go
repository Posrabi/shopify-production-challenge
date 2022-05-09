package repository

import (
	"context"

	"github.com/Posrabi/shopify-backend-project/src/inv/pkg/entity"
)

type Master interface {
	// Item
	CreateItem(ctx context.Context, item *entity.Item) error
	DeleteItem(ctx context.Context, itemID string) error
	EditItem(ctx context.Context, item *entity.Item) error
	ListItems(ctx context.Context) ([]*entity.Item, error)
	// Shipments
	Ship(ctx context.Context, itemInstance *entity.Shipment) error
}
