package client

import "encoding/json"

// Residences fetches a (filtered) list of residences from the API
func (c *Client) Residences(optionalParams ...map[string]string) ([]byte, error) {
	params := map[string]string{}
	if len(optionalParams) > 0 {
		params = optionalParams[0]
	}

	return c.Get("residences", params)
}

// Estimate will POST to /estimation-subscription
func (c *Client) SubscribeToEstimation(params map[string]string) ([]byte, error) {
	json, err := json.Marshal(params)

	if err != nil {
		return nil, err
	}

	return c.Post("estimation-subscriptions", json)
}
