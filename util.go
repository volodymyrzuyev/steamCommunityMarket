package steamcommunityapi

import (
	"io"
	"net/http"
	"net/url"
)

func filterOutBasedOnStatusCode(resp http.Response) ([]byte, error) {
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
