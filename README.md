# Steam Community Market
A very light Golang wrapper around steam community methods outlined in [Revadike/InternalSteamWebAPI](https://github.com/Revadike/InternalSteamWebAPI)

## Getting started

Pulling the library
```bash
go mod init steamApiTest
go get https://github.com/volodymyrzuyev/steamCommunityMarket
touch main.go
nvim main.go
```
in main.go
```go
package main

import (
	"fmt"
	mApi "github.com/volodymyrzuyev/steamCommunityMarket"
	"time"
)

func main() {
	requestInterval := time.Duration(10 * time.Second)
	api := mApi.NewApiController(requestInterval)

	params := mApi.MarketRecentParams{
		Country:  mApi.UnitedStates,
		Language: mApi.English,
		Currency: mApi.USD,
	}

	resp, err := api.MarketRecent(params)
	if err != nil {
		panic("Gaben does not like you")
	}

	fmt.Println(string(resp))
}
```
`resp` will be a []byte that contains a json response from steam. You can 
implement any logic that you might like.

## Methods
### MarketListings
Get's listings for an item you want. [Example](https://github.com/Revadike/InternalSteamWebAPI/wiki/Get-Market-Listing)
```go
func MarketListings(params MarketListingsParams) ([]byte, error)
```
```go
type MarketListingsParams struct {
	AppID          string	`Appid of the game you are interested in`
	MarketHashName string	`Name that shows up in url bar when you open item in browser`
	Country        country	`One of the contries from the list defined in consts`
	Currency       currency	`One of courrenties defined in consts.go`
	Start          string	`First item to show`
	Count          string	`Number of items that will be fetched`
	Language       language	`One of languages defined in consts.go`
}
```

### MarketPriceOverview
Get's price overview of an item. [Example](https://github.com/Revadike/InternalSteamWebAPI/wiki/Get-Market-Price-Overview)
```go
func MarketPriceOverview(params MarketPriceOverviewParams) ([]byte, error)
```
```go
type MarketPriceOverviewParams struct {
	AppID          string	`Appid of the game you are interested in`
	MarketHashName string	`Name that shows up in url bar when you open item in browser`
	Country        country	`One of courrenties defined in consts.go`
	Currency       currency	`One of courrenties defined in consts.go`
}
```

### MarketRecent
Get's newly posted listings to market. [Example](https://github.com/Revadike/InternalSteamWebAPI/wiki/Get-Recent-Market-Listings)
```go
func MarketRecent(params MarketRecentParams) ([]byte, error) {
```
```go
type MarketRecentParams struct {
	Country  country		`One of courrenties defined in consts.go`
	Language language		`One of languages defined in consts.go`
	Currency currency		`One of courrenties defined in consts.go`
}
```

### MarketSearch
Find's items with similar name to the query. [Example](https://github.com/Revadike/InternalSteamWebAPI/wiki/Search-Market)
```go
func MarketSearch(params MarketSearchParams) ([]byte, error) {
```
```go
type MarketSearchParams struct {
	AppID string
	Start string
	Count string
	Query string
}
```
