package entity

import "time"

// A shipment entity, this is not present on the database for now.
type Shipment struct {
	ItemInstance *ItemInstance `json:"item_instance"`
	Origin       string        `json:"origin"` // Some metadata that is not in used right now.
	Destination  string        `json:"destination"`
	Location     string        `json:"location"`
	ETA          time.Time     `json:"eta"`
}
