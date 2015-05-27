package client

// Residences fetches a (filtered) list of residences from the API
func (c *Client) Residences(optionalParams ...map[string]string) ([]byte, error) {
	params := map[string]string{}
	if len(optionalParams) > 0 {
		params = optionalParams[0]
	}

	return c.Get("residences", params)
}