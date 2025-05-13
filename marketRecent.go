package steamcommunityapi

import "net/url"

func (c *controller) MarketRecent(params MarketRecentParams) ([]byte, error) {
	err := params.validate()
	if err != nil {
		return []byte{}, err
	}

	curUrl, err := constructPath("/market/recent")
	if err != nil {
		return []byte{}, err
	}

	query := make([]string, 4)
	query[0] = "country=" + url.QueryEscape(params.Country.code)
	query[1] = "language=" + url.QueryEscape(params.Language.fullname)
	query[2] = "currency=" + url.QueryEscape(params.Currency.code)
	query[3] = "norender=1"

	curUrl.RawQuery = constructQuery(query)

	return c.runQuery(curUrl.String())
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
