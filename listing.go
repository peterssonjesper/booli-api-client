package client

import (
	"fmt"
)

func (this *Client) Listing(booliId int, optionalParams ...map[string]string) ([]byte, error) {
	params := map[string]string{}
	if len(optionalParams) > 0 {
		params = optionalParams[0]
	}

	return this.Get("listings/" + fmt.Sprintf("%d", booliId), params)
}

func (this *Client) Listings(params map[string]string) ([]byte, error) {
	return this.Get("listings", params)
}
