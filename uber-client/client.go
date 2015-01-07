package main

import (
	"encoding/json"
	"fmt"
	"github.com/anweiss/uber-api-golang/uber"
	"io/ioutil"
	"log"
)

func main() {
	// Read API auth options
	var options uber.RequestOptions
	fileContents, err := ioutil.ReadFile("./uber-client/options.json")
	if err != nil {
		log.Fatal(err)
	}

	if e := json.Unmarshal(fileContents, &options); e != nil {
		log.Fatal(err)
	}

	// Create a Client for executing API operations
	client := uber.Create(&options)

	// Retrieve products based on lat/long coords
	pl := &uber.Products{}
	pl.Latitude = 38.897939
	pl.Longitude = -77.036541
	if e := client.Get(pl); e != nil {
		log.Fatal(e)
	}

	fmt.Println("Here are the Uber options available for your area: \n")
	for _, product := range pl.Products {
		fmt.Println(product.DisplayName + ": " + product.Description)
	}

	// Retrieve price estimates based on start and end lat/long coords
	pe := &uber.PriceEstimates{}
	pe.StartLatitude = 38.897939
	pe.StartLongitude = -77.036541
	pe.EndLatitude = 38.890152
	pe.EndLongitude = -77.009096
	if e := client.Get(pe); e != nil {
		log.Fatal(e)
	}

	fmt.Println("\nHere are the Uber price estimates from The White House to the United States Capitol: \n")
	for n, price := range pe.Prices {
		fmt.Println(price.DisplayName + ": " + price.Estimate)
		if n == len(pe.Prices)-1 {
			fmt.Print("\n")
		}
	}
}
