package client

import (
	"fmt"
)

// SoldProperty fetches a sold property given a booli ID
func (c *Client) SoldProperty(booliID int, optionalParams ...map[string]string) ([]byte, error) {
	params := map[string]string{}
	if len(optionalParams) > 0 {
		params = optionalParams[0]
	}

	return c.Get("sold/"+fmt.Sprintf("%d", booliID), params)
}

// SoldProperties makes a searches for sold properties given a map of parameters
func (c *Client) SoldProperties(optionalParams ...map[string]string) ([]byte, error) {
	params := map[string]string{}
	if len(optionalParams) > 0 {
		params = optionalParams[0]
	}

	return c.Get("sold", params)
}
