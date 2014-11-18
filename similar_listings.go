package client

import (
	"fmt"
)

// SimilarListings fetches listings similar to a listing given its Booli ID
func (c *Client) SimilarListings(booliID int, optionalParams ...map[string]string) ([]byte, error) {
	params := map[string]string{}
	if len(optionalParams) > 0 {
		params = optionalParams[0]
	}

	return c.Get(fmt.Sprintf("listings/%d/similar", booliID), params)
}
