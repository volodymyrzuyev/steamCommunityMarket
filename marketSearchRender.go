package steamcommunityapi

import (
	"net/url"
	"strconv"
)

func (c *controller) MarketSearch(params MarketSearchParams) ([]byte, error) {
	err := params.validate()
	if err != nil {
		return []byte{}, err
	}

	curUrl, err := constructPath("/market/search/render")
	if err != nil {
		return []byte{}, err
	}

	query := make([]string, 5)
	query[0] = "norender=1"
	query[1] = "start=" + url.QueryEscape(params.Start)
	query[2] = "count=" + url.QueryEscape(params.Count)
	query[3] = "query=" + url.QueryEscape(params.Query)
	query[4] = "appid=" + url.QueryEscape(params.AppID)

	curUrl.RawQuery = constructQuery(query)

	return c.runQuery(curUrl.String())
}

type MarketSearchParams struct {
	AppID string
	Start string
	Count string
	Query string
}

func (m MarketSearchParams) validate() error {
	if m.AppID == "" {
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
	if m.Query == "" {
		return InvalidParameter
	}

	return nil
}
