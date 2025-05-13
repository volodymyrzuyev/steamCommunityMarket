package steamcommunityapi

import (
	"io"
	"net/http"
	"net/url"
	"time"
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
		query += "&" + url.QueryEscape(v)
	}

	return query
}

func runHttp(url string) (*http.Response, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.82 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
