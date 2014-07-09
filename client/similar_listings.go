package client

import (
	"fmt"
	"encoding/json"
)

func (this *Client) SimilarListings(booliId int, optionalParams ...map[string]string) ([]Listing, error) {
	params := map[string]string{}
	if len(optionalParams) > 0 {
		params = optionalParams[0]
	}

	response, err := this.Get("listings/" + fmt.Sprintf("%d", booliId) + "/similar", params)

	if err != nil {
		return nil, err
	}

	var envelope ListingsEnvelope
	err = json.Unmarshal(response, &envelope)

	if err != nil {
		return nil, err
	}

	return envelope.Listings, nil
}
