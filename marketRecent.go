package steamcommunityapi

func (c *controller) MarketRecent(params MarketRecentParams) ([]byte, error) {
	err := params.validate()
	if err != nil {
		return []byte{}, err
	}

	url, err := constructPath("/market/recent")
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

type MarketRecentParams struct {
	Country  country
	Language language
	Currency currency
}

func (m MarketRecentParams) validate() error {
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
