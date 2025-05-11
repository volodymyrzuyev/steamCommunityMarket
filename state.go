package steamcommunityapi

import "time"

type state interface {
	updateLastRequest(time.Time)
	cooldownPassed() bool
}

type stateStore struct {
	cooldown        time.Duration
	waitForCooldown bool
	lastRequest     time.Time
}

func (s *stateStore) updateLastRequest(newRequest time.Time) {
	s.lastRequest = newRequest
}

func (s *stateStore) cooldownPassed() bool {
	if time.Since(s.lastRequest) > s.cooldown {
		return true
	}
	if s.waitForCooldown {
		s.sleepUptilCooldown()
		return true
	}

	return false
}

func (s *stateStore) sleepUptilCooldown() {
	timeToSleep := s.lastRequest.Add(s.cooldown).Sub(time.Now())
	time.Sleep(timeToSleep)
}
