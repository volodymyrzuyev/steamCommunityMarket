package steamcommunityapi

import (
	"fmt"
	"strconv"
)

func (c *controller) MarketListings(params MarketListingsParams) ([]byte, error) {
	err := params.validate()
	if err != nil {
		return []byte{}, err
	}

	path := fmt.Sprintf("/market/listings/%s/%s/render", params.AppID, params.MarketHashName)
	url, err := constructPath(path)
	if err != nil {
		return []byte{}, err
	}

	query := make([]string, 5)
	query[0] = "currency=" + params.Currency.code
	query[1] = "coutry=" + params.Country.code
	query[2] = "start=" + params.Start
	query[3] = "count=" + params.Count
	query[4] = "language=" + params.Language.fullname

	url.RawQuery = constructQuery(query)

	return c.runQuery(url.String())
}

type MarketListingsParams struct {
	AppID          string
	MarketHashName string
	Country        country
	Currency       currency
	Start          string
	Count          string
	Language       language
}

func (m MarketListingsParams) validate() error {
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
	_, e := strconv.Atoi(m.Start)
	if e != nil {
		return InvalidParameter
	}
	_, e = strconv.Atoi(m.Count)
	if e != nil {
		return InvalidParameter
	}
	if m.Language == blankLang {
		return InvalidParameter
	}

	return nil
}
