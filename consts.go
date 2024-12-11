package steamcommunityapi

import "fmt"

const BaseWebApiUrl = "https://www.steamcommunity.com"

var (
	CooldownNotPassed         = fmt.Errorf("Cooldown since last request has not passed")
	SteamRateLimitExceeded    = fmt.Errorf("Steam rate limit exceeded")
	UnexpectedSteamStatusCode = fmt.Errorf("Received unexpected status code from Steam")
)
