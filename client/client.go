package client

import (
	"fmt"
	"errors"
	"time"
	"io"
	"encoding/json"
	"crypto/sha1"
	"io/ioutil"
	"math/rand"
	"net/url"
	"net/http"
)

type Client struct {
	hostname, callerId, apiKey string
}

func New(hostname, callerId, apiKey string) *Client {
	return &Client{
		hostname: hostname,
		callerId: callerId,
		apiKey: apiKey,
	}
}

func (this *Client) Get(endpoint string, optionalParams ...map[string]string) ([]byte, error) {
	var params map[string]string
	if len(optionalParams) > 0 {
		params = optionalParams[0]
	} else {
		params = map[string]string {}
	}

	url := this.url(endpoint, params)

	client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/vnd.booli-v2+json")

	response, err := client.Do(req)

	if err != nil {
		return nil, errors.New("Could make GET request to " + url)
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	return body, nil
}

func (this *Client) Listing(booliId string) (*Listing, error) {
	response, err := this.Get("listings/" + booliId)

	if err != nil {
		return nil, err
	}

	var envelope ListingsEnvelope
	err = json.Unmarshal(response, &envelope)

	if err != nil {
		return nil, err
	}

	if len(envelope.Listings) == 0 {
		return nil, errors.New("Listing not found.")
	}

	return &envelope.Listings[0], nil
}

func (this *Client) hash (time, unique string) string {
	hash := sha1.New()
	io.WriteString(hash, this.callerId + time + this.apiKey + unique)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func (this *Client) url (endpoint string, params map[string]string) string {
	time := this.time()
	unique := this.unique()
	hash := this.hash(time, unique)

	p := url.Values{}
	for k, v := range params {
		p.Set(k, v)
	}
	p.Set("callerId", this.callerId)
	p.Set("time", time)
	p.Set("unique", unique)
	p.Set("hash", hash)

	query := p.Encode()

	return "http://" + this.hostname + "/" + endpoint + "?" + query
}

func (this *Client) time() string {
	return fmt.Sprintf("%d", time.Now().Unix())
}

func (this *Client) unique() string {
	random := rand.New(rand.NewSource(time.Now().UnixNano())).Uint32()
	return fmt.Sprintf("%d", random)
}
