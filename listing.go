package client

import (
	"fmt"
)

// Listing fetches listing data from the API given a Booli ID of the listing
func (c *Client) Listing(booliID int, optionalParams ...map[string]string) ([]byte, error) {
	params := map[string]string{}
	if len(optionalParams) > 0 {
		params = optionalParams[0]
	}

	return c.Get("listings/"+fmt.Sprintf("%d", booliID), params)
}

// Listings searches for listings given a map of parameters
func (c *Client) Listings(params map[string]string) ([]byte, error) {
	return c.Get("listings", params)
}
