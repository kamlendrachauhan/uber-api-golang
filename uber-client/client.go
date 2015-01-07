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

	// Create a Client retrieve Products based on lat/long coords
	client := uber.Create(&options)
	pl := &uber.Products{}
	pl.Latitude = ""
	pl.Longitude = ""
	if e := client.Get(pl); e != nil {
		log.Fatal(e)
	}

	for _, product := range pl.Products {
		fmt.Println(product.DisplayName + ": " + product.Description)
	}
}
