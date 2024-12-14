package steamcommunityapi

import "strconv"

func (c *controller) MarketSearch(params MarketSearchParams) ([]byte, error) {
	err := params.validate()
	if err != nil {
		return []byte{}, err
	}

	url, err := constructPath("/market/search/render")
	if err != nil {
		return []byte{}, err
	}

	query := make([]string, 5)
	query[0] = "norender=1"
	query[1] = "start=" + params.Start
	query[2] = "count=" + params.Count
	query[3] = "query=" + params.Query
	query[4] = "appid=" + params.AppID

	url.RawQuery = constructQuery(query)

	return c.runQuery(url.String())
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
