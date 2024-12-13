package steamcommunityapi

type MarketRecentParams struct {
	Country  string
	Language string
	Currency string
}

func (c *controller) MarketRecent(params MarketRecentParams) ([]byte, error) {
	url, err := constructPath("/market/recent")
	if err != nil {
		return []byte{}, err
	}

	query := make([]string, 4)
	query[0] = "country=" + params.Country
	query[1] = "language=" + params.Language
	query[2] = "currency=" + params.Currency
	query[3] = "norender=1"

	url.RawQuery = constructQuery(query)

	return c.runQuery(url.String())
}
