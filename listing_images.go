package client

import (
	"fmt"
)

func (this *Client) ListingImages(booliId int, optionalParams ...map[string]string) ([]byte, error) {
	params := map[string]string{}
	if len(optionalParams) > 0 {
		params = optionalParams[0]
	}

	return this.Get(fmt.Sprintf("listings/%d/images", booliId), params)
}
