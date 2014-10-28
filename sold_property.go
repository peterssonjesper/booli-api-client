package client

import (
	"fmt"
)

func (this *Client) SoldProperty(booliId int, optionalParams ...map[string]string) ([]byte, error) {
	params := map[string]string{}
	if len(optionalParams) > 0 {
		params = optionalParams[0]
	}

	return this.Get("sold/" + fmt.Sprintf("%d", booliId), params)
}
