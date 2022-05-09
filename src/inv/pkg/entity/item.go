package entity

// Generic item entity
type Item struct {
	ID       string `json:"item_id" db:"item_id"` // UUID
	Name     string `json:"item_name" db:"item_name"`
	Brand    string `json:"brand" db:"brand"`
	Quantity int    `json:"item_quantity" db:"item_quantity"`
}

// An instance of Item for shipment, is not present on the DB.
type ItemInstance struct {
	ID       string `json:"item_id"`
	Quantity int    `json:"item_quantity"`
}
