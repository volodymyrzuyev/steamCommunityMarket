package steamcommunityapi

import "fmt"

type MarketListingsParams struct {
	AppID          string
	MarketHashName string
	Country        string
	Currency       string
	Start          string
	Count          string
	Language       string
}

func (c *Controller) MarketListings(params MarketListingsParams) ([]byte, error) {
	path := fmt.Sprintf("/market/listings/%s/%s/render", params.AppID, params.MarketHashName)
	url, err := constructPath(path)
	if err != nil {
		return []byte{}, err
	}

	values := url.Query()
	values.Add("country", params.Country)
	values.Add("currency", params.Currency)
	values.Add("start", params.Start)
	values.Add("count", params.Count)
	values.Add("language", params.Language)

	url.RawQuery = values.Encode()

	return c.runQuerry(url.String())
}
