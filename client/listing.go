package client

import (
	"fmt"
)

func (this *Client) Listing(booliId int) ([]byte, error) {
	return this.Get("listings/" + fmt.Sprintf("%d", booliId))
}
