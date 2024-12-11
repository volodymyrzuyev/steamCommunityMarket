package steamcommunityapi

type MarketSearchParams struct {
	AppID  string
	Start  string
	Count  string
	Querry string
}

func (c *Controller) MarketSearch(params MarketSearchParams) ([]byte, error) {
	url, err := constructPath("/market/search/render")
	if err != nil {
		return []byte{}, err
	}

	values := url.Query()
	values.Add("start", params.Start)
	values.Add("count", params.Count)
	values.Add("querry", params.Querry)
	values.Add("appid", params.AppID)

	url.RawQuery = values.Encode()

	return c.runQuerry(url.String())
}
