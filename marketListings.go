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

	query := make([]string, 5)
	query[0] = "currency=" + params.Currency
	query[1] = "coutry=" + params.Country
	query[2] = "start=" + params.Start
	query[3] = "count=" + params.Count
	query[4] = "language=" + params.Language

	url.RawQuery = constructQuery(query)

	return c.runQuerry(url.String())
}
