package client

func (this *Client) PreviousSales(_ int, optionalParams ...map[string]string) ([]byte, error) {
	params := map[string]string{}
	if len(optionalParams) > 0 {
		params = optionalParams[0]
	}

	return this.Get("sold", params)
}
