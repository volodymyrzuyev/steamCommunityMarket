package steamcommunityapi

func (c *controller) MarketPriceOverview(params MarketPriceOverviewParams) ([]byte, error) {
	err := params.validate()
	if err != nil {
		return []byte{}, err
	}

	url, err := constructPath("/market/priceoverview/")
	if err != nil {
		return []byte{}, err
	}

	query := make([]string, 4)
	query[0] = "country=" + params.Country.code
	query[1] = "currency=" + params.Currency.code
	query[2] = "appid=" + params.AppID
	query[3] = "market_hash_name=" + params.MarketHashName

	url.RawQuery = constructQuery(query)

	return c.runQuery(url.String())
}

type MarketPriceOverviewParams struct {
	AppID          string
	MarketHashName string
	Country        country
	Currency       currency
}

func (m MarketPriceOverviewParams) validate() error {
	if m.AppID == "" {
		return InvalidParameter
	}
	if m.MarketHashName == "" {
		return InvalidParameter
	}
	if m.Country == blankCountry {
		return InvalidParameter
	}
	if m.Currency == blankCur {
		return InvalidParameter
	}

	return nil
}
