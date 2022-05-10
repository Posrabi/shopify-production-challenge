package entity

import (
	"reflect"
	"strconv"
)

// Generic item entity
type Item struct {
	ID          string `json:"item_id" db:"item_id"` // UUID
	Name        string `json:"item_name" db:"item_name"`
	Brand       string `json:"brand" db:"brand"`
	Quantity    int    `json:"item_quantity" db:"item_quantity"`
	StorageCity City   `json:"storage_city" db:"storage_city"`
	Weather     string `json:"weather"`
}

func (i *Item) ToCSVRecord() []string {
	v := reflect.ValueOf(*i)

	res := []string{}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.CanInt() {
			res = append(res, strconv.FormatInt(v.Field(i).Int(), 10))
		} else {
			res = append(res, v.Field(i).String())
		}
	}

	return res
}

func (i *Item) ToCSVHeader() []string {
	v := reflect.ValueOf(*i)

	res := []string{}

	for i := 0; i < v.NumField(); i++ {
		res = append(res, v.Type().Field(i).Tag.Get("json"))
	}

	return res
}

// An instance of Item for shipment, is not present on the DB.
type ItemInstance struct {
	ID       string `json:"item_id"`
	Quantity int    `json:"item_quantity"`
}

type City string

const (
	CITY_1 City = "Toronto"
	CITY_2 City = "New York"
	CITY_3 City = "San Francisco"
	CITY_4 City = "Boston"
	CITY_5 City = "Atlanta"
)

type CreateItemRequest struct {
	Item Item `json:"item"`
	City int  `json:"city"`
}

// Not a very good idea IRL but in this case, it works!
func (r *CreateItemRequest) AssignCityToItem() *CreateItemRequest {
	switch r.City {
	case 2:
		r.Item.StorageCity = CITY_2
	case 3:
		r.Item.StorageCity = CITY_3
	case 4:
		r.Item.StorageCity = CITY_4
	case 5:
		r.Item.StorageCity = CITY_5
	default:
		r.Item.StorageCity = CITY_1
	}
	return r
}
