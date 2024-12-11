package steamcommunityapi

func (c *Controller) MarketRecent() ([]byte, error) {
	path, err := constructPath("/market/recent")
	if err != nil {
		return []byte{}, err
	}

	return c.runQuerry(path.String())
}
