package steamcommunityapi

type MarketSearchParams struct {
	AppID  string
	Start  string
	Count  string
	Querry string
}

func (c *controller) MarketSearch(params MarketSearchParams) ([]byte, error) {
	url, err := constructPath("/market/search/render")
	if err != nil {
		return []byte{}, err
	}

	query := make([]string, 5)
	query[0] = "norender=1"
	query[1] = "start=" + params.Start
	query[2] = "count=" + params.Count
	query[3] = "query=" + params.Querry
	query[4] = "appid=" + params.AppID

	url.RawQuery = constructQuery(query)

	resp, err := c.runQuery(url.String())
	if err != nil {
		return []byte{}, err
	}

	return c.filter(resp)
}
