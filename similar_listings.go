package client

import (
	"fmt"
)

func (this *Client) SimilarListings(booliId int, optionalParams ...map[string]string) ([]byte, error) {
	params := map[string]string{}
	if len(optionalParams) > 0 {
		params = optionalParams[0]
	}

	return this.Get(fmt.Sprintf("listings/%d/similar", booliId), params)
}
