package steamcommunityapi

func (c *controller) MarketRecentCompleted(params MarketRecentCompletedParams) ([]byte, error) {
	err := params.validate()
	if err != nil {
		return []byte{}, err
	}

	url, err := constructPath("/market/recentcompleted")
	if err != nil {
		return []byte{}, err
	}

	query := make([]string, 4)
	query[0] = "country=" + params.Country.code
	query[1] = "language=" + params.Language.fullname
	query[2] = "currency=" + params.Currency.code
	query[3] = "norender=1"

	url.RawQuery = constructQuery(query)

	return c.runQuery(url.String())
}

type MarketRecentCompletedParams struct {
	Country  country
	Language language
	Currency currency
}

func (m MarketRecentCompletedParams) validate() error {
	if m.Country == blankCountry {
		return InvalidParameter
	}
	if m.Language == blankLang {
		return InvalidParameter
	}
	if m.Currency == blankCur {
		return InvalidParameter
	}

	return nil
}
