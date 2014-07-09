package client

import (
	"fmt"
	"errors"
	"encoding/json"
)

func (this *Client) Listing(booliId int, optionalParams ...map[string]string) (*Listing, error) {
	params := map[string]string{}
	if len(optionalParams) > 0 {
		params = optionalParams[0]
	}

	response, err := this.Get("listings/" + fmt.Sprintf("%d", booliId), params)

	if err != nil {
		return nil, err
	}

	var envelope ListingsEnvelope
	err = json.Unmarshal(response, &envelope)

	if err != nil {
		return nil, err
	}

	if len(envelope.Listings) == 0 {
		return nil, errors.New("Listing not found")
	}

	return &envelope.Listings[0], nil
}
