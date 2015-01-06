package uber

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	// Uber API endpoint
	APIUrl string = "https://api.uber.com/v1/%s"

	// Token used for interacting with Uber API
	ServerToken string = ""
)

type Lister interface {
	List()
}

// List of Uber products
type Products struct {
	Products []Product `json:"products"`
}

// Uber product
type Product struct {
	ProductId   string `json:"product_id"`
	Description string `json:"description"`
	DisplayName string `json:"display_name"`
	Capacity    int    `json:"capacity"`
	Image       string `json:"image"`
}

// List of price estimates
type PriceEstimates struct {
	Prices []PriceEstimate `json:"prices"`
}

// Uber price estimate
type PriceEstimate struct {
	ProductId       string  `json:"product_id"`
	CurrencyCode    string  `json:"currency_code"`
	DisplayName     string  `json:"display_name"`
	Estimate        string  `json:"estimate"`
	LowEstimate     int     `json:"low_estimate"`
	HighEstimate    int     `json:"high_estimate"`
	SurgeMultiplier float32 `json:"surge_multiplier"`
	Duration        int     `json:"duration"`
	Distance        float32 `json:"distance"`
}

// Method to retrieve list of Uber products
func (p *Products) Get(latitude string, longitude string) error {
	productParams := map[string]string{
		"server_token": ServerToken,
		"latitude":     latitude,
		"longitude":    longitude,
	}
	url := fmt.Sprintf(APIUrl, "products"+encodeUrl(&productParams))
	data := getRequest(url)

	if e := json.Unmarshal(data, &p); e != nil {
		return e
	}

	return nil
}

// Products list method that implements the Lister interface
func (pl *Products) List() {
	for _, product := range pl.Products {
		fmt.Println(product.DisplayName + ": " + product.Description)
	}
}

// Method to retrieve list of Uber price estimates given
// start and end latitude/longitude values
/*func (pe *PriceEstimates) Get(startLat string, startLong string, endLat string, endLong string) error {


}
*/

// Retrieve HTTP response from Uber API
func getRequest(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	return data
}

// Logging function
func LogError(err *error) {
	log.Fatal(err)
}

// Encodes a URL given a map of query parameters
func encodeUrl(params *map[string]string) string {
	urlParams := "?"
	for k, v := range *params {
		if len(urlParams) > 1 {
			urlParams += "&"
		}
		urlParams += fmt.Sprintf("%s=%s", k, v)
	}

	return urlParams
}
