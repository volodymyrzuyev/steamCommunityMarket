package steamcommunityapi

import (
	"net/http"
	"time"
)

type Api interface {
	MarketListings(params MarketListingsParams) ([]byte, error)
	MarketPriceOverview(params MarketPriceOverviewParams) ([]byte, error)
	MarketRecent(params MarketRecentParams) ([]byte, error)
	MarketSearch(params MarketSearchParams) ([]byte, error)
}

type controller struct {
	state      state
	httpRunner func(url string) (*http.Response, error)
}

func NewApiController(cooldownInterwal time.Duration) Api {

	state := stateStore{
		cooldown: cooldownInterwal,
		// so the first request will not have to wait for the cooldown to be over
		lastRequest: time.Now().Add(-cooldownInterwal * 2),
	}

	controller := controller{
		state:      &state,
		httpRunner: runHttp,
	}

	return &controller
}

func (c *controller) runQuery(url string) ([]byte, error) {
	if !c.state.cooldownPassed() {
		return []byte{}, CooldownNotPassed
	}

	resp, err := c.httpRunner(url)
	// update last reqest time to calculate even failed requests
	c.state.updateLastRequest(time.Now())
	if err != nil {
		return []byte{}, err
	}

	return filterOutBasedOnStatusCode(*resp)
}
