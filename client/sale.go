package client

import (
	"fmt"
)

func (this *Client) Sale(booliId int) ([]byte, error) {
	return this.Get("sold/" + fmt.Sprintf("%d", booliId))
}
