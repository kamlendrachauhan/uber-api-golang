Golang SDK for the Uber API
===============

[![GoDoc](https://godoc.org/github.com/anweiss/uber-api-golang/uber?status.svg)](https://godoc.org/github.com/anweiss/uber-api-golang/uber) [![Build Status](https://ci.weisslabs.info/api/badge/github.com/anweiss/uber-api-golang/status.svg?branch=master)](https://ci.weisslabs.info/github.com/anweiss/uber-api-golang)

This package provides an SDK for use with the Uber API. Currently, the following API endpoints are supported:
- Products
- Price Estimates
- Time Estimates

OAuth support for user profile and past activity access is forthcoming.

Prior to use, you'll need to need to register an application in the Uber developer portal and obtain the appropriate tokens and keys. More information about the Uber API can be found
here: [https://developer.uber.com/](https://developer.uber.com/).

## Usage

Full usage info can be found in the [godocs](https://godoc.org/github.com/anweiss/uber-api-golang/uber).

A sample client has been included with this source that demonstrates its usage. The package import path is `github.com/anweiss/uber-api-golang/uber`. The package requires the creation of a `Client` type with a set of `RequestOptions`. The `RequestOptions` type holds the token and OAuth parameters.

```go
import "github.com/anweiss/uber-api-golang/uber"

var options uber.RequestOptions{
  ServerToken: SERVER_TOKEN
  ClientId: CLIENT_ID
  ...
}

client := uber.Create(&options)
```

### Retrieving Products

You can retrieve product endpoints by creating a variable which will hold the endpoint's type and required properties and pass the object `Client.Get()`

```go
pl := &uber.Products{}
pl.Latitude = LAT
pl.Longitude = LONG
if e := client.Get(p1); e != nil {
  log.Fatal(e)
}
```

The products are retrieved and stored in the `Products` property.

```go
fmt.Println("Here are the Uber options available for your area: \n")
for _, product := range pl.Products {
  fmt.Println(product.DisplayName + ": " + product.Description)
}
```
