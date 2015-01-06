package main

import (
	"github.com/anweiss/uber-api-golang/uber"
)

func main() {
	pl := &uber.Products{}
	if err := pl.Get("38.907473", "-77.0297238"); err != nil {
		uber.LogError(&err)
	}

	pl.List()

}
