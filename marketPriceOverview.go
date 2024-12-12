package steamcommunityapi

type MarketPriceOverviewParams struct {
	AppID          string
	MarketHashName string
	Country        string
	Currency       string
}

func (c *Controller) MarketPriceOverview(params MarketPriceOverviewParams) ([]byte, error) {
	url, err := constructPath("/market/priceoverview/")
	if err != nil {
		return []byte{}, err
	}

	query := make([]string, 4)
	query[0] = "country=" + params.Country
	query[1] = "currency=" + params.Currency
	query[2] = "appid=" + params.AppID
	query[3] = "market_hash_name=" + params.MarketHashName

	url.RawQuery = constructQuery(query)

	return c.runQuerry(url.String())
}
