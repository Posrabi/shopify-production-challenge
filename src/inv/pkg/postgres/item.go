package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	resty "github.com/go-resty/resty/v2"

	"github.com/Posrabi/shopify-backend-project/src/common/exception"
	"github.com/Posrabi/shopify-backend-project/src/inv/pkg/entity"
	"github.com/Posrabi/shopify-backend-project/src/inv/pkg/repository"
)

const ITEM_COLUMNS = "item_id, item_name, brand, item_quantity, storage_city"
const openWeatherURL = "http://api.openweathermap.org/data/2.5/weather"

type itemRepo struct {
	db               *sql.DB
	openWeatherToken string
}

func NewItemRepository(db *sql.DB, openWeatherToken string) repository.Item {
	return &itemRepo{
		db:               db,
		openWeatherToken: openWeatherToken,
	}
}

func (i *itemRepo) CreateItem(ctx context.Context, item *entity.Item) error {
	q := `INSERT INTO items (%s) VALUES ($1, $2, $3, $4, $5)`

	args := []any{item.ID, item.Name, item.Brand, item.Quantity, item.StorageCity}
	if _, err := i.db.ExecContext(ctx, fmt.Sprintf(q, ITEM_COLUMNS), args...); err != nil {
		return exception.NewPQError(err, q, args)
	}

	return nil
}

func (i *itemRepo) DeleteItem(ctx context.Context, itemID string) error {
	q := `DELETE FROM items WHERE item_id = $1`

	if _, err := i.db.ExecContext(ctx, q, itemID); err != nil {
		return exception.NewPQError(err, q, []any{itemID})
	}

	return nil
}

func (i *itemRepo) EditItem(ctx context.Context, item *entity.Item) error {
	q := `UPDATE items
	SET item_name = $1, brand = $2, item_quantity = $3
	WHERE item_id = $4`

	args := []any{item.Name, item.Brand, item.Quantity, item.ID}
	if _, err := i.db.ExecContext(ctx, q, args...); err != nil {
		return exception.NewPQError(err, q, args)
	}

	return nil
}

func (i *itemRepo) ListItems(ctx context.Context) ([]*entity.Item, error) {
	q := `SELECT %s FROM items`
	res, err := i.db.QueryContext(ctx, fmt.Sprintf(q, ITEM_COLUMNS))
	if err != nil {
		return nil, exception.NewPQError(err, q, nil)
	}

	items, err := i.scanItems(res)
	if err != nil {
		return nil, err
	}

	client := resty.New().SetHeader("Accept", "application/json")
	for _, item := range items {
		item.Weather, err = i.getWeatherReport(ctx, client, string(item.StorageCity))
		if err != nil {
			return nil, err
		}
	}

	return items, nil
}

func (i *itemRepo) ShipItem(ctx context.Context, itemInstance *entity.ItemInstance) error {
	getCurrentQuantity := `SELECT item_quantity from items WHERE item_id = $1`
	row := i.db.QueryRowContext(ctx, getCurrentQuantity, itemInstance.ID)

	var currentQuantity int
	if err := row.Scan(&currentQuantity); err != nil {
		return exception.NewPQError(err, getCurrentQuantity, []any{itemInstance.ID})
	}

	// Prevent decrementing to negative
	if currentQuantity < itemInstance.Quantity {
		return exception.NewPQError(exception.InventoryError, getCurrentQuantity, []any{itemInstance.ID})
	}

	update := `UPDATE items SET item_quantity = item_quantity - $2 WHERE item_id = $1`
	updateArgs := []any{itemInstance.ID, itemInstance.Quantity}
	if _, err := i.db.ExecContext(ctx, update, updateArgs...); err != nil {
		return exception.NewPQError(err, update, updateArgs)
	}

	return nil
}

func (i *itemRepo) scanItems(rows *sql.Rows) ([]*entity.Item, error) {
	items := []*entity.Item{}
	for rows.Next() {
		item := entity.Item{}
		if err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Brand,
			&item.Quantity,
			&item.StorageCity,
		); err != nil {
			return nil, exception.NewError(err)
		}

		items = append(items, &item)
	}

	return items, rows.Err()
}

// getWeatherReport returns a weather report from the city name, which makes it a little less accurate than using coordinates but is
// perfectly fine in this situation.
func (i *itemRepo) getWeatherReport(ctx context.Context, client *resty.Client, city string) (string, error) {
	resp, err := client.R().SetContext(ctx).SetQueryParam("appid", i.openWeatherToken).SetQueryParam("q", city).Get(openWeatherURL)
	if err != nil {
		return "", err
	}

	var report entity.WeatherReport
	if err := json.Unmarshal(resp.Body(), &report); err != nil {
		return "", err
	}

	return fmt.Sprintf("City: %s, Weather: %s, Description: %s", report.Name, report.Weather[0].Main, report.Weather[0].Description), nil
}
