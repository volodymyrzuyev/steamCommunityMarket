package steamcommunityapi

import "net/url"

func (c *controller) MarketPriceOverview(params MarketPriceOverviewParams) ([]byte, error) {
	err := params.validate()
	if err != nil {
		return []byte{}, err
	}

	curUrl, err := constructPath("/market/priceoverview/")
	if err != nil {
		return []byte{}, err
	}

	query := make([]string, 4)
	query[0] = "country=" + url.QueryEscape(params.Country.code)
	query[1] = "currency=" + url.QueryEscape(params.Currency.code)
	query[2] = "appid=" + url.QueryEscape(params.AppID)
	query[3] = "market_hash_name=" + url.QueryEscape(params.MarketHashName)

	curUrl.RawQuery = constructQuery(query)

	return c.runQuery(curUrl.String())
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
