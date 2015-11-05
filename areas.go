package client

import (
	"fmt"
)

// Area fetches an area by given ID
func (c *Client) Area(areaID int) ([]byte, error) {
	return c.Get(fmt.Sprintf("areas/%d", areaID))
}

// ListingAreas fetches all areas that a listing is placed in
func (c *Client) ListingAreas(id int, optionalParams ...map[string]string) ([]byte, error) {
	params := map[string]string{}
	if len(optionalParams) > 0 {
		params = optionalParams[0]
	}

	return c.Get(fmt.Sprintf("listings/%d/areas", id), params)
}

// SoldPropertyAreas fetches all areas that a sold property was placed in
func (c *Client) SoldPropertyAreas(id int, optionalParams ...map[string]string) ([]byte, error) {
	params := map[string]string{}
	if len(optionalParams) > 0 {
		params = optionalParams[0]
	}

	return c.Get(fmt.Sprintf("sold/%d/areas", id), params)
}

// AddressAreas fetches all areas that an address was placed in
func (c *Client) AddressAreas(addressID int, optionalParams ...map[string]string) ([]byte, error) {
	params := map[string]string{}
	if len(optionalParams) > 0 {
		params = optionalParams[0]
	}
	params["addressId"] = fmt.Sprintf("%d", addressID)

	return c.Get("areas", params)
}
