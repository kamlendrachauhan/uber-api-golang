package main

import (
	"encoding/json"
	"fmt"
	"github.com/anweiss/uber-api-golang/uber"
	"io/ioutil"
	"log"
	"strconv"
)

const (
	WhiteHouseLat  float64 = 38.897939
	WhiteHouseLong float64 = -77.036541
	USCapitolLat   float64 = 38.890152
	USCapitolLong  float64 = -77.009096
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
	pl.Latitude = WhiteHouseLat
	pl.Longitude = WhiteHouseLong
	if e := client.Get(pl); e != nil {
		log.Fatal(e)
	}

	fmt.Println("Here are the Uber options available for your area: \n")
	for _, product := range pl.Products {
		fmt.Println(product.DisplayName + ": " + product.Description)
	}

	// Retrieve price estimates based on start and end lat/long coords
	pe := &uber.PriceEstimates{}
	pe.StartLatitude = WhiteHouseLat
	pe.StartLongitude = WhiteHouseLong
	pe.EndLatitude = USCapitolLat
	pe.EndLongitude = USCapitolLong
	if e := client.Get(pe); e != nil {
		log.Fatal(e)
	}

	fmt.Println("\nHere are the Uber price estimates from The White House to the United States Capitol: \n")
	for _, price := range pe.Prices {
		fmt.Println(price.DisplayName + ": " + price.Estimate + "; Surge: " + strconv.FormatFloat(price.SurgeMultiplier, 'f', 2, 32))
	}

	// Retrieve ETA estimates based on start lat/long coords
	te := &uber.TimeEstimates{}
	te.StartLatitude = WhiteHouseLat
	te.StartLongitude = WhiteHouseLong
	if e := client.Get(te); e != nil {
		log.Fatal(e)
	}

	fmt.Println("\nHere are the Uber ETA estimates if leaving from The White House: \n")
	for n, eta := range te.Times {
		fmt.Println(eta.DisplayName + ": " + strconv.Itoa(eta.Estimate/60))
		if n == len(te.Times)-1 {
			fmt.Print("\n")
		}
	}
}
