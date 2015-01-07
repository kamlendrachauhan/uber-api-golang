package uber

import (
	"encoding/json"
	"strconv"
)

// List of Uber products with given lat/long coords
type Products struct {
	Latitude  float64
	Longitude float64
	Products  []Product `json:"products"`
}

// Uber product
type Product struct {
	ProductId   string `json:"product_id"`
	Description string `json:"description"`
	DisplayName string `json:"display_name"`
	Capacity    int    `json:"capacity"`
	Image       string `json:"image"`
}

// Internal method that implements the getter interface
func (pl *Products) get(c *Client) error {
	productParams := map[string]string{
		"latitude":  strconv.FormatFloat(pl.Latitude, 'f', 2, 32),
		"longitude": strconv.FormatFloat(pl.Longitude, 'f', 2, 32),
	}

	data := c.getRequest("products", productParams)
	if e := json.Unmarshal(data, &pl); e != nil {
		return e
	}
	return nil
}
