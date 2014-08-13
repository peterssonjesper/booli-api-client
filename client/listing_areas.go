package client

import (
	"fmt"
)

func (this *Client) ListingAreas(id int, optionalParams ...map[string]string) ([]byte, error) {
	params := map[string]string{}
	if len(optionalParams) > 0 {
		params = optionalParams[0]
	}

	return this.Get(fmt.Sprintf("listings/%d/areas", id), params)
}
