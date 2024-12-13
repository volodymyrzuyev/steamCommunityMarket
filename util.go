package steamcommunityapi

import (
	"io"
	"net/http"
	"net/url"
)

func filterOutBasedOnStatusCode(resp *http.Response) ([]byte, error) {
	statusCode := resp.StatusCode

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	switch statusCode {
	case 200:
		return body, nil
	case 429:
		return body, SteamRateLimitExceeded
	}
	// TODO: add a logger to log unexpected responses
	return body, UnexpectedSteamStatusCode
}

func constructPath(path string) (url.URL, error) {
	newUrl, err := url.Parse(BaseWebApiUrl)
	if err != nil {
		return url.URL{}, err
	}

	newUrl.Path = path

	return *newUrl, nil
}

// since URL.Query().Encode() returns the query in a random order, have to create
// a custom function
func constructQuery(params []string) string {
	query := ""
	for i, v := range params {
		if i == 0 {
			query += v
			continue
		}
		query += "&" + v
	}

	return query
}

func runHttp(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
