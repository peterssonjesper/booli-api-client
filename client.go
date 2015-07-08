package client

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

// Interface describes the Booli API Client
type Interface interface {
	Get(endpoint string, params ...map[string]string) ([]byte, error)
	Post(endpoint string, payload []byte, params ...map[string]string) ([]byte, error)
	Listing(id int, params ...map[string]string) ([]byte, error)
	Listings(params map[string]string) ([]byte, error)
	SoldProperty(id int, params ...map[string]string) ([]byte, error)
	SoldProperties(params ...map[string]string) ([]byte, error)
	SimilarListings(id int, params ...map[string]string) ([]byte, error)
	ListingImages(id int, params ...map[string]string) ([]byte, error)
	ListingAreas(id int, params ...map[string]string) ([]byte, error)
	SoldPropertyAreas(id int, params ...map[string]string) ([]byte, error)
	Residences(params ...map[string]string) ([]byte, error)
	SubscribeToEstimation(params map[string]string) ([]byte, error)
}

// Client holds data needed to talk to the API
type Client struct {
	host, callerID, apiKey string
}

// New is a constructor for Client
func New(host, callerID, apiKey string) Interface {
	return &Client{
		host:     host,
		callerID: callerID,
		apiKey:   apiKey,
	}
}

// Get makes a generic GET request to the API
func (c *Client) Get(endpoint string, optionalParams ...map[string]string) ([]byte, error) {
	params := map[string]string{}
	if len(optionalParams) > 0 {
		params = optionalParams[0]
	}

	url := c.url(endpoint, params)

	client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/vnd.booli-v2+json")

	response, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Could not make GET request to %s: %q", url, err.Error())
	}

	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	if response.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Could not get a proper response "+
			"from server. Got response code %d but expected %d from URL %s",
			response.StatusCode,
			http.StatusOK,
			url)
	}

	return body, nil
}

func (c *Client) Post(endpoint string, payload []byte, optionalParams ...map[string]string) ([]byte, error) {
	params := map[string]string{}
	if len(optionalParams) > 0 {
		params = optionalParams[0]
	}

	url := c.url(endpoint, params)

	client := &http.Client{}
	buf := bytes.NewBuffer(payload)
	req, err := http.NewRequest("POST", url, buf)
	req.Header.Set("Accept", "application/vnd.booli-v2+json")

	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Could not make POST request to %s: %q", url, err.Error())
	}

	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK || response.StatusCode == http.StatusCreated {
		return body, nil
	}

	return nil, fmt.Errorf("Could not get a proper response "+
		"from server. Got response code %d but expected %d or %d from URL %s",
		response.StatusCode,
		http.StatusOK,
		http.StatusCreated,
		url)
}

func (c *Client) hash(time, unique string) string {
	hash := sha1.New()
	io.WriteString(hash, c.callerID+time+c.apiKey+unique)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func (c *Client) url(endpoint string, params map[string]string) string {
	time := c.time()
	unique := c.unique()
	hash := c.hash(time, unique)

	p := url.Values{}
	for k, v := range params {
		p.Set(k, v)
	}
	p.Set("callerId", c.callerID)
	p.Set("time", time)
	p.Set("unique", unique)
	p.Set("hash", hash)

	query := p.Encode()

	return c.host + "/" + endpoint + "?" + query
}

func (c *Client) time() string {
	return fmt.Sprintf("%d", time.Now().Unix())
}

func (c *Client) unique() string {
	random := rand.New(rand.NewSource(time.Now().UnixNano())).Uint32()
	return fmt.Sprintf("%d", random)
}
