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

	values := url.Query()

	values.Add("country", params.Country)
	values.Add("currency", params.Currency)
	values.Add("appid", params.AppID)
	values.Add("market_hash_name", params.MarketHashName)

	url.RawQuery = values.Encode()

	return c.runQuerry(url.String())
}
