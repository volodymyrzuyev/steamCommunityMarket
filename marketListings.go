package steamcommunityapi

import (
	"fmt"
	"net/url"
	"strconv"
)

func (c *controller) MarketListings(params MarketListingsParams) ([]byte, error) {
	err := params.validate()
	if err != nil {
		return []byte{}, err
	}

	path := fmt.Sprintf("/market/listings/%s/%s/render", params.AppID, params.MarketHashName)
	curUrl, err := constructPath(path)
	if err != nil {
		return []byte{}, err
	}

	query := make([]string, 5)
	query[0] = "currency=" + url.QueryEscape(params.Currency.code)
	query[1] = "coutry=" + url.QueryEscape(params.Country.code)
	query[2] = "start=" + url.QueryEscape(params.Start)
	query[3] = "count=" + url.QueryEscape(params.Count)
	query[4] = "language=" + url.QueryEscape(params.Language.fullname)

	curUrl.RawQuery = constructQuery(query)

	return c.runQuery(curUrl.String())
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
