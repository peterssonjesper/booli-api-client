package client

import "fmt"

// Address fetches an address with given ID
func (c *Client) Address(addressID int, optionalParams ...map[string]string) ([]byte, error) {
	params := map[string]string{}
	if len(optionalParams) > 0 {
		params = optionalParams[0]
	}

	return c.Get(fmt.Sprintf("/addresses/%d", addressID), params)
}
