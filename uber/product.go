package uber

import (
	"encoding/json"
)

// List of Uber products
type Products struct {
	Latitude  string
	Longitude string
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

// Products list method that implements the Lister interface
func (pl *Products) get(c *Client) error {
	productParams := map[string]string{
		"server_token": c.Options.ServerToken,
		"latitude":     pl.Latitude,
		"longitude":    pl.Longitude,
	}

	data := getRequest(&productParams)
	if e := json.Unmarshal(data, &pl); e != nil {
		return e
	}
	return nil
}
