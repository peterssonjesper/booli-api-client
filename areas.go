package client

import (
	"fmt"
)

// ListingAreas fetches all areas that a listing is placed in
func (c *Client) ListingAreas(id int, optionalParams ...map[string]string) ([]byte, error) {
	params := map[string]string{}
	if len(optionalParams) > 0 {
		params = optionalParams[0]
	}

	return c.Get(fmt.Sprintf("listings/%d/areas", id), params)
}

// SoldPropertyAreas fetches all areas that a sold property was placed in
func (c *Client) SoldPropertyAreas(id int, optionalParams ...map[string]string) ([]byte, error) {
	params := map[string]string{}
	if len(optionalParams) > 0 {
		params = optionalParams[0]
	}

	return c.Get(fmt.Sprintf("sold/%d/areas", id), params)
}
