package steamcommunityapi

import "time"

type state interface {
	updateLastRequest(time.Time)
	cooldownPassed() bool
}

type stateStore struct {
	cooldown    time.Duration
	lastRequest time.Time
}

func (s *stateStore) updateLastRequest(newRequest time.Time) {
	s.lastRequest = newRequest
}

func (s *stateStore) cooldownPassed() bool {
	if time.Since(s.lastRequest) > s.cooldown {
		return true
	}
	return false
}
