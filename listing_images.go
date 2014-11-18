package client

import (
	"fmt"
)

// ListingImages fetches images for a listing
func (c *Client) ListingImages(booliID int, optionalParams ...map[string]string) ([]byte, error) {
	params := map[string]string{}
	if len(optionalParams) > 0 {
		params = optionalParams[0]
	}

	return c.Get(fmt.Sprintf("listings/%d/images", booliID), params)
}
