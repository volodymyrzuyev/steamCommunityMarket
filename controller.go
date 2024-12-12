package steamcommunityapi

import (
	"net/http"
	"time"
)

type Controller struct {
	cooldown    time.Duration
	lastRequest time.Time
}

func NewApiController(cooldownInterwal float64) *Controller {
	cooldown := time.Duration(cooldownInterwal) * time.Second
	controller := Controller{
		cooldown: cooldown,
		// so the first request will not have to wait for the cooldown to be over
		lastRequest: time.Now().Add(-cooldown * 2),
	}

	return &controller
}

func (c *Controller) cooldownPassed() bool {
	if time.Since(c.lastRequest) > c.cooldown {
		return true
	}
	return false
}

func (c *Controller) runQuerry(url string) ([]byte, error) {
	if !c.cooldownPassed() {
		return []byte{}, CooldownNotPassed
	}

	resp, err := http.Get(url)
	// update last reqest time to calculate even failed requests
	c.lastRequest = time.Now()
	if err != nil {
		return []byte{}, err
	}

	return filterOutBasedOnStatusCode(*resp)
}
