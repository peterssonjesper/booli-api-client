package client

import "fmt"

// Project searches for a single project with given project ID
func (c *Client) Project(projectID int, optionalParams ...map[string]string) ([]byte, error) {
	params := map[string]string{}
	if len(optionalParams) > 0 {
		params = optionalParams[0]
	}

	return c.Get(fmt.Sprintf("projects/%d", projectID), params)
}

// Projects searches for projects
func (c *Client) Projects(optionalParams ...map[string]string) ([]byte, error) {
	params := map[string]string{}
	if len(optionalParams) > 0 {
		params = optionalParams[0]
	}

	return c.Get("projects", params)
}
